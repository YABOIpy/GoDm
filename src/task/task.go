package task

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"source/src/modules"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	Mod = modules.Modules{}
	Ws  = modules.Sock{}
)

// should've used pointers. too late to go back now so sadly gotta sleep within the functions
func StartTask(in []modules.Instance, Task func(c modules.Instance)) {
	cfg, _ := Mod.LoadConfig("config.json")

	var routines int
	routines = len(in)
	if cfg.Mode.Configs.CCManager {
		routines = cfg.Mode.Configs.MaxRoutines
	}
	var wg sync.WaitGroup
	wg.Add(routines)
	for i := 0; i < len(in); i++ {
		c := in[i]
		go func(i int, c modules.Instance) {
			defer wg.Done()
			Task(c)
		}(i, c)

		time.Sleep(c.Cfg.Mode.Configs.Interval * time.Second)
	}
	wg.Wait()
}

func MassDmTask(in []modules.Instance, msg string, interval time.Duration) {
	cfg, _ := Mod.LoadConfig("config.json")
	id, _, _ := Mod.ReadFile("data/ids.txt")

	if len(id) == 0 {
		log.Println("No IDS found in tokens.txt..")
		Return(1)
	}
	C := len(in)
	ids := len(id)
	dist := ids / C
	over := ids % C

	group := make(map[string][]string)

	var start, ex int
	for i, token := range in {
		ex = 0
		if i < over {
			ex = 1
		}
		s := start
		e := start + dist + ex
		start = e
		group[token.Token] = id[s:e]
	}

	var routines int
	routines = len(in)
	if cfg.Mode.Configs.CCManager {
		routines = cfg.Mode.Configs.MaxRoutines
	}

	var wg sync.WaitGroup
	wg.Add(routines)
	s := time.Now()

	for i := 0; i < len(group); i++ {
		c := in[i]

		go func(i int, c modules.Instance) {

			defer wg.Done()
			uid := group[c.Token]
			for j := 0; j < len(uid); j++ {

				if c.Eligible(c, uid[j]) {
					data := c.CreateChannel(c, uid[j])
					resp, body := c.Message(c, msg, data.Id, modules.MessageOptions{Loop: false})

					var m modules.MessageResp
					if err := json.Unmarshal(body, &m); err != nil {
						log.Println(err)
					}
					switch resp {
					case 200:
						Mod.Checker.Valid++
					case 429:
						c.TokenProps.RateLimit = data.Retry
						time.Sleep(time.Duration(c.TokenProps.RateLimit+float64(rand.Intn(9)+2)) * time.Second)
					default:
					retry:
						//More nested than a birds nest!
						// TODO: fix nesting
						if strings.Contains(string(body), "captcha_sitekey") {
							if !cfg.Mode.Configs.Solver {
								Mod.Checker.Invalid++
							} else {
								Mod.Checker.Locked++
								captcha := c.Captcha(modules.CapCfg{
									ApiKey:  c.Cfg.Mode.Discord.CapAPI[1],
									SiteKey: m.SiteKey,
									PageURL: "https://discord.com",
								})
								for count := 0; count <= cfg.Mode.Configs.CaptchaRetry; count++ {
									resp, body = c.Message(c, msg, data.Id,
										modules.MessageOptions{
											Loop:    false,
											Captcha: captcha,
										})
									if resp != 200 {
										time.Sleep(interval * time.Second)
										goto retry
									} else {
										break
									}
								}
							}
						} else {
							Mod.Checker.Invalid++
						}
					}
					Mod.Checker.All++
				}
				time.Sleep(interval * time.Second)
			}
		}(i, c)

		time.Sleep(c.Cfg.Mode.Configs.Interval * time.Second)
	}
	wg.Wait()

	fmt.Printf(modules.MassDmFormat, time.Since(s).String()[:4], Mod.Checker.Locked, Mod.Checker.Invalid, Mod.Checker.Valid, Mod.Checker.All)
	Mod.Checker = modules.Modules{}.Checker
}

func ScrapeTask(Token string, in modules.Instance, GID string, CID string) {
	_, _, con := Ws.Connect(Token, in)
	os.Truncate("data/ids.txt", 0)

	var iter, pv int
	for {
		s := time.Now()
		con.ScrapeUsers(GID, CID, iter)
		if con.Break {
			break
		}
		cv := len(con.Members)
		if cv != pv && cv > 0 {
			Mod.StrlogV("Got Online Member Chunk", strconv.Itoa(len(con.Members)), s)
		}
		pv = cv
		iter++
	}
	var ids []string
	for i := 0; i < len(con.Members); i++ {
		ids = append(ids, con.Members[i].User.ID)
	}
	data := Mod.FilterArray(ids)
	Mod.WriteFileArray("data/ids.txt", data)
}

func CheckerTask(in []modules.Instance) {

	f := [3]string{
		"data/valid.txt",
		"data/locked.txt",
		"data/invalid.txt",
	}
	for i := 0; i < len(f); i++ {
		os.Truncate(f[i], 0)
	}

	var delta time.Duration
	var wg sync.WaitGroup
	token := make([]string, 0)
	wg.Add(len(in))

	s := time.Now()
	for i := 0; i < len(in); i++ {
		c := in[i]
		go func(i int) {
			defer wg.Done()
			resp, t := c.Check()
			switch resp {
			case 200:
				Mod.WriteFile("data/valid.txt", c.Token)
				token = append(token, c.Token)
				Mod.Checker.Valid++
			case 403:
				Mod.WriteFile("data/locked.txt", c.Token)
				Mod.Checker.Locked++
			default:
				Mod.WriteFile("data/invalid.txt", c.Token)
				Mod.Checker.Invalid++
			}
			Mod.Checker.All++
			atomic.AddInt64((*int64)(&delta), int64(time.Since(t)))
		}(i)
	}
	wg.Wait()

	fmt.Printf(modules.CheckerFormat, time.Since(s).String()[:4], Mod.Checker.Locked, Mod.Checker.Invalid, Mod.Checker.Valid, Mod.Checker.All)
	fmt.Printf("Delta:\u001B[34;1m %s\n\u001B[0m", delta/time.Duration(len(in)))
	if Mod.Checker.All != Mod.Checker.Valid && Mod.Checker.Valid != 0 {
		if Mod.Input(modules.WriteValidMention) == "y" {
			os.Truncate("tokens.txt", 0)
			var d []string
			for i := 0; i < len(token); i++ {
				pass, mail := in[i].TokenProps.Pass, in[i].TokenProps.Email
				if len(mail) > 0 && len(pass) > 0 {
					d = append(d, fmt.Sprintf(modules.TokenFormat, mail, pass, token[i]))
				} else {
					d = append(d, token[i])
				}
			}
			Mod.WriteFileArray("tokens.txt", d)
		}
	}
	Mod.Checker = modules.Modules{}.Checker
}

func Return(i int) {
	fmt.Println("Going Back To Menu =>..")
	time.Sleep(time.Duration(i) * time.Second)
	Mod.Cls()
	Mod.Menu()
}
