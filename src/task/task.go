package task

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"github.com/wasilibs/go-re2"
	"log"
	"math/rand"
	"os"
	"source/src/modules"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/zenthangplus/goccm"
)

var (
	Mod = modules.Modules{}
	Ws  = modules.Sock{}
)

func StartTask(in []modules.Instance, Task func(c modules.Instance)) {
	cfg, _ := Mod.LoadConfig("config.json")

	routines := len(in)
	if cfg.Mode.Configs.CCManager {
		routines = cfg.Mode.Configs.MaxRoutines
	}
	wg := goccm.New(routines)

	for i := 0; i < len(in); i++ {
		c := in[i]
		wg.Wait()
		go func(i int, c modules.Instance) {
			defer wg.Done()
			Task(c)

			Mod.Sleep(time.Duration(c.TokenProps.RateLimit), &c)
		}(i, c)

		time.Sleep(c.Cfg.Mode.Configs.Interval * time.Second)
	}
	wg.WaitAllDone()
}

func MassDmTask(in []modules.Instance, msg string, interval time.Duration) {
	cfg, _ := Mod.LoadConfig("config.json")
	id, _, _ := Mod.ReadFile("data/ids.txt")

	if len(id) == 0 {
		log.Println("No IDS found in tokens.txt..")
		Return(1)
	}
	dist := len(id) / len(in)
	over := len(id) % len(in)

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

	routines := len(in)
	if cfg.Mode.Configs.CCManager {
		routines = cfg.Mode.Configs.MaxRoutines
	}

	wg := goccm.New(routines)
	s := time.Now()

	for i := 0; i < len(group); i++ {
		c := in[i]
		wg.Wait()
		go func(i int, c modules.Instance) {

			defer wg.Done()
			uid := group[c.Token]
			for j := 0; j < len(uid); j++ {

				if c.Eligible(uid[j]) {

					cmd := re2.MustCompile(`<([^>]+)>`).FindAllStringSubmatch(msg, -1)
					if len(cmd) > modules.IntNil && Mod.Contains(cmd[0], "user") {
						msg = strings.ReplaceAll(msg, "<user>", uid[j])
					}

					data := c.CreateChannel(uid[j])
					resp, body := c.Message(msg, data.Id, modules.MessageOptions{Loop: false})

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
						// More nested than a birds nest!
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
								for count := 0; count < cfg.Mode.Configs.CaptchaRetry; {
									count++
									resp, body = c.Message(msg, data.Id,
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
	wg.WaitAllDone()

	fmt.Printf(modules.MassDmFormat, time.Since(s).String()[:4], Mod.Checker.Locked, Mod.Checker.Invalid, Mod.Checker.Valid, Mod.Checker.All)
	Mod.Checker = modules.Modules{}.Checker
}

func ScrapeTask(in modules.Instance, GID, CID string) {
	_, con, _ := Ws.Connect(in.Token, &in)
	os.Truncate("data/ids.txt", 0)

	var ids []string
	var iter, pv int

	for {
		s := time.Now()
		con.ScrapeUsers(GID, CID, iter)
		cv := len(con.Members)
		if cv != pv && cv > modules.IntNil {
			Mod.StrlogV("Got Online Member Chunk", strconv.Itoa(len(con.Members)), s)
		}
		pv = cv
		iter++
		if con.Break {
			break
		}
	}
	for _, c := range con.Members {
		ids = append(ids, c.User.ID)
	}
	data := Mod.FilterArray(ids)
	Mod.WriteFileArray("data/ids.txt", data)
}

func CheckerTask(in []modules.Instance) {
	cfg, _ := Mod.LoadConfig("config.json")

	f := [3]string{
		"data/valid.txt",
		"data/locked.txt",
		"data/invalid.txt",
	}
	for i := 0; i < len(f); i++ {
		os.Truncate(f[i], 0)
	}
	var delta time.Duration

	routines := len(in)
	if cfg.Mode.Configs.CCManager {
		routines = cfg.Mode.Configs.MaxRoutines
	}

	wg := goccm.New(routines)
	var token []string

	s := time.Now()
	for i := 0; i < len(in); i++ {
		c := in[i]
		wg.Wait()
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
	wg.WaitAllDone()

	fmt.Printf(modules.CheckerFormat, time.Since(s).String()[:4], Mod.Checker.Locked, Mod.Checker.Invalid, Mod.Checker.Valid, Mod.Checker.All)
	fmt.Printf("Delta:\u001B[34;1m %s\n\u001B[0m", delta/time.Duration(len(in)))

	if Mod.Checker.All != Mod.Checker.Valid && Mod.Checker.Valid != modules.IntNil {
		if Mod.InputBool(modules.WriteValidMention) {

			os.Truncate("tokens.txt", 0)
			var d []string

			for i := 0; i < len(token); i++ {
				pass, mail := in[i].TokenProps.Pass, in[i].TokenProps.Email
				if len(mail) > modules.IntNil && len(pass) > modules.IntNil {
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

func MassFriendTask(in []modules.Instance, interval time.Duration) {
	cfg, _ := Mod.LoadConfig("config.json")
	id, _, _ := Mod.ReadFile("data/ids.txt")

	if len(id) == 0 {
		log.Println("No IDS found in tokens.txt..")
		Return(1)
		return
	}
	dist := len(id) / len(in)
	over := len(id) % len(in)

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

	routines := len(in)
	if cfg.Mode.Configs.CCManager {
		routines = cfg.Mode.Configs.MaxRoutines
	}

	wg := goccm.New(routines)
	s := time.Now()

	for i := 0; i < len(group); i++ {
		c := in[i]
		wg.Wait()

		go func(i int, c modules.Instance) {

			defer wg.Done()
			uid := group[c.Token]
			for j := 0; j < len(uid); j++ {

				var d modules.UserInfo
				if d = c.UserInfo(uid[j]); d == (modules.UserInfo{}) {
					Mod.StrlogE("No User Found With ID", uid[j], s)
					return
				}

				data := modules.FriendReq{
					Username: d.Username,
				}
				data.Discrim = nil
				if d.Discriminator != "0" {
					data.Discrim = cast.ToInt(d.Discriminator)
				}

				resp, body := c.Friend(data)
				switch resp {
				case 200:
					Mod.Checker.Valid++
				case 429:
					Mod.Checker.Invalid++
				default:
					if strings.Contains(string(body), "captcha_sitekey") {
						Mod.Checker.Locked++
						//solve cap
					} else {
						Mod.Checker.Invalid++
					}
				}
				time.Sleep(interval * time.Second)
			}
		}(i, c)

		time.Sleep(c.Cfg.Mode.Configs.Interval * time.Second)
	}
	wg.WaitAllDone()

	fmt.Printf(modules.MassFriendFormat, time.Since(s).String()[:4], Mod.Checker.Locked, Mod.Checker.Invalid, Mod.Checker.Valid, Mod.Checker.All)
	Mod.Checker = modules.Modules{}.Checker
}
