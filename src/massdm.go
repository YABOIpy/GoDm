package massdm

import (
	"bytes"
	"encoding/json"
	"fmt"
	http "github.com/Danny-Dasilva/fhttp"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func (Xc *Config) Dm(ID string, Token string, Msg string, Cookies string) {
	payload := map[string]string{
		"content": Msg,
	}

	req, err := http.NewRequest("POST", "https://discord.com/api/v9/channels/"+ID+"/messages",
		bytes.NewBuffer(Xc.Marsh(payload)),
	)
	Xc.Errs(err)

	Hd.Head_Dm(req, Token)
	resp, err := Client.Do(req)
	Xc.Errs(err)

	if resp.StatusCode == 200 {
		fmt.Println(""+grn+"▏"+r+"("+grn+"+"+r+") Sent Message To:"+clr+"", ID)
	} else {
		fmt.Println(""+red+"▏"+r+"("+red+"-"+r+") Failed Sent Message To:"+clr+"", ID,
			Xc.Errmsg(*resp),
		)
	}
}

func (Xc *Config) CloseDm(ID string, Token string, Cookies string) {
	req, err := http.NewRequest("DELETE",
		"https://discord.com/api/v9/channels/"+ID+"?silent=false",
		nil,
	)
	Xc.Errs(err)

	Hd.Head_CloseDm(req, Token)
	resp, err := Client.Do(req)
	Xc.Errs(err)

	if resp.StatusCode == 200 {
		fmt.Println(""+grn+"▏"+r+"("+grn+"+"+r+") Closed Channel:"+clr+"", ID)
	} else {
		fmt.Println(""+red+"▏"+r+"("+red+"-"+r+") Failed To Close Dm:"+clr+"", ID, "ERR |: ",
			Xc.Errmsg(*resp),
		)
	}

}

func (Xc *Config) React(Token string, link string) {
	payload := map[string]string{}
	req, err := http.NewRequest("PUT",
		link,
		bytes.NewBuffer(
			Xc.Marsh(payload),
		),
	)
	Xc.Errs(err)

	Hd.Head_React(req, Token)
	resp, err := Client.Do(req)
	Xc.Errs(err)

	switch resp.StatusCode {
	case 204:
		fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Added Emoji")
	default:
		fmt.Println(""+red+"▏"+r+"("+red+"-"+r+") Failed To Add Emoji"+clr+"", "ERR |: "+r,
			Xc.Errmsg(*resp),
		)
	}
}

func (Xc *Config) Create(ID int, Token string, Msg string) (string, error) {
	payload := []byte("{\"recipients\":[\"" + strconv.Itoa(ID) + "\"]}")
	req, err := http.NewRequest("POST", "https://discord.com/api/v9/users/@me/channels",
		bytes.NewBuffer(payload),
	)

	Xc.Errs(err)

	Hd.Head_Create(req, Token)
	resp, err := Client.Do(req)

	var flake ChannelID
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := ReadBody(*resp)
		Xc.Errs(err)
		errx := json.Unmarshal(body, &flake)
		Xc.Errs(errx)
		fmt.Println(""+grn+"▏"+r+"("+grn+"+"+r+") Created Channel:"+clr+"", flake.ID)
		Xc.Dm(flake.ID, Token, Msg, Cookies)
		return flake.ID, err

	} else {
		fmt.Println(""+red+"▏"+r+"("+red+"+"+r+") Failed To Create Channel:",
			Xc.Errmsg(*resp),
		)
	}
	return flake.ID, err
}

func (Xc *Config) Block(ID int, Token string, Cookies string) {
	p := map[string]string{"type": "2"}
	req, err := http.NewRequest("PUT", "https://discord.com/api/v9/users/@me/relationships/"+strconv.Itoa(ID)+"",
		bytes.NewBuffer(Xc.Marsh(p)),
	)

	Xc.Errs(err)

	Hd.Head_Block(req, Token, ID)

	resp, err := Client.Do(req)
	Xc.Errs(err)
	if resp.StatusCode == 204 {
		fmt.Println(""+grn+"▏"+r+"("+grn+"+"+r+") Blocked User:"+clr+"", ID)
	} else {
		fmt.Println(""+red+"▏"+r+"("+red+"-"+r+") Failed To Block:"+clr+"", ID, "ERR |: ",
			Xc.Errmsg(*resp),
		)
	}
}

func (Xc *Config) Dm_Spam(ID string, Token string, Msg string) {
	Xc.Dm(ID, Token, Msg, Cookies)
}

func (Xc *Config) Joiner(Token string, invite string) {
	req, err := http.NewRequest("POST", "https://discord.com/api/v9/invites/"+invite+"",
		bytes.NewBuffer(
			Xc.Marsh(
				map[string]string{"": ""},
			),
		),
	)
	Xc.Errs(err)

	Hd.Head_Joiner(req, Token)
	resp, err := Client.Do(req)
	Xc.Errs(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	Xc.Errs(err)

	if resp.StatusCode == 200 {
		fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Joined " + clr + "discord.gg/" + invite)
	} else if resp.StatusCode == 429 {
		fmt.Println(""+red+"▏"+r+"("+red+"+"+r+") Failed To Join "+clr+"discord.gg/"+invite, yel+" RateLimit", r)
	} else if strings.Contains(string(body), "captcha_sitekey") {
		fmt.Println(""+yel+"▏"+r+"("+yel+"+"+r+") Failed To Join "+clr+"discord.gg/"+invite, yel+" Captcha", r)
	} else {

		fmt.Println(""+yel+"▏"+r+"("+red+"+"+r+") Failed To Join "+clr+"discord.gg/"+invite,
			Xc.Errmsg(*resp),
		)
	}

}

func (Xc *Config) Leaver(Token string, ID string) {
	req, err := http.NewRequest("DELETE", "https://discord.com/api/v9/users/@me/guilds/"+ID+"",
		bytes.NewBuffer(
			Xc.Marsh(
				map[string]string{"lurking": "false"},
			),
		),
	)
	Xc.Errs(err)

	Hd.Head_Leaver(req, Token)

	resp, err := Client.Do(req)
	_ = err
	if resp.StatusCode == 204 {
		fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Left Server")
	} else {
		fmt.Println(""+yel+"▏"+r+"("+yel+"+"+r+") Failed To Leave "+clr+ID,
			Xc.Errmsg(*resp),
		)
	}
}

func (Xc *Config) Agree(Token string, invite string, ID string) {
	req, err := http.NewRequest("POST",
		"https://discord.com/api/v9/guilds/"+ID+"/member-verification?with_guild=false&invite_code="+invite+"",
		nil,
	)
	Xc.Errs(err)

	Hd.Head_Agree(req, Token, 1)
	resp, err := Client.Do(req)
	Xc.Errs(err)

	defer resp.Body.Close()
	body, ers := ReadBody(*resp)
	Xc.Errs(ers)

	var data Bypass
	err = json.Unmarshal(body, &data)
	Xc.Errs(err)

	for i := 0; i < len(data.FormFields); i++ {
		data.FormFields[i].Response = true
	}

	payload, _ := json.Marshal(data)
	reqs, er := http.NewRequest("POST", "https://discord.com/api/v9/guilds/"+ID+"/requests/@me",
		bytes.NewBuffer(
			payload,
		),
	)
	Xc.Errs(er)

	Hd.Head_Agree(reqs, Token, 0)
	resps, es := Client.Do(reqs)
	_ = es
	if resps.StatusCode == 201 {
		fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Verified Token")
	} else {
		fmt.Println(""+red+"▏"+r+"("+red+"-"+r+") Failed",
			Xc.Errmsg(*resp),
		)
	}

}

func (Xc *Checker) Check(token string) string {

	req, _ := http.NewRequest("GET", urls, nil)

	req.Header.Set("authorization", token)
	resp, _ := Xc.Client.Do(req)

	var typ = Xc.Resp

	if resp.StatusCode == 200 {
		typ = true
		fmt.Println(""+grn+"▏("+grn+"✓"+r+") ("+grn+"+"+r+"):", token[:50]+"...")
		Xc.Valid++

	} else if resp.StatusCode == 403 {
		typ = false
		fmt.Println(""+yel+"▏("+yel+"/"+r+"):", token[:50]+"...")
		Xc.Locked++
	} else {
		fmt.Println(""+red+"▏("+red+"x"+r+"):", token[:50]+"...")
		Xc.Invalid++
	}

	Xc.All++
	Xc.Resp = typ
	Xc.Token = token
	return Xc.Token
}

func (Xc *Config) Scrape_ID(Ws *Sock, Token string, CID string, GID string, index int) {
	if index == 0 {
		payload := Data{
			GuildId:    GID,
			Typing:     true,
			Threads:    true,
			Activities: true,
			Members:    []Member{},
			Channels: map[string]interface{}{
				CID: []interface{}{[2]int{0, 99}},
			},
		}
		err := Ws.Conn.WriteJSON(&Event{
			Op:   14,
			Data: payload,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	var x []interface{}
	if index == 0 {
		x = []interface{}{[2]int{0, 99}}
	} else if index == 1 {
		x = []interface{}{[2]int{0, 99}, [2]int{100, 199}}
	} else if index == 2 {
		x = []interface{}{[2]int{0, 99}, [2]int{100, 199}, [2]int{200, 299}}
	} else {
		x = []interface{}{[2]int{0, 99}, [2]int{100, 199}, [2]int{index * 100, (index * 100) + 99}}
	}
	payload := Data{
		GuildId: GID,
		Channels: map[string]interface{}{
			CID: x,
		},
	}
	err := Ws.Conn.WriteJSON(&Event{
		Op:   14,
		Data: payload,
	})
	if err != nil {
		log.Fatal(err)
	}

	_, resp, err := Ws.Conn.ReadMessage()

	var data interface{}
	json.Unmarshal(resp, &data)
	fmt.Println(data)

	//fmt.Println(string([]byte(Ws.Conn)))
	var memberids []string
	for _, member := range Ws.Members {
		memberids = append(memberids, member.User.ID)
		fmt.Println(memberids)
	}
}

func (Xc *Config) Raider(Token string, message string, ID string) {
	for true {

		req, err := http.NewRequest("POST", "https://discord.com/api/v9/channels/"+ID+"/messages",
			bytes.NewBuffer(
				Xc.Marsh(map[string]string{
					"content": message,
				}),
			),
		)
		Xc.Errs(err)

		Hd.Head_Raider(req, Token, ID)
		resp, ers := Client.Do(req)
		_ = ers
		if resp.StatusCode == 200 {
			fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Sent Message")
		} else if resp.StatusCode == 429 {
			fmt.Println("" + yel + "▏" + r + "(" + yel + "+" + r + ") RateLimit")
		} else {
			fmt.Println(""+red+"▏"+r+"("+red+"+"+r+") Failed To Send:",
				Xc.Errmsg(*resp),
			)
		}
	}
}

func (Xc *Config) MassPing(Token string, Message string, Amount int, ID string) {
	for true {
		var msg string
		rand.Seed(time.Now().Unix())
		users, _ := Xc.ReadFile("ids.txt")
		for i := 0; i < Amount; i++ {

			ping := users[rand.Intn(len(users))]
			msg += "<@" + ping + ">"
		}

		req, err := http.NewRequest("POST", "https://discord.com/api/v9/channels/"+ID+"/messages",
			bytes.NewBuffer(
				Xc.Marsh(map[string]string{
					"content": msg,
				}),
			),
		)
		Xc.Errs(err)

		Hd.Head_MassPing(req, Token, ID)
		resp, ers := Client.Do(req)
		Xc.Errs(ers)
		if resp.StatusCode == 200 {
			fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Sent Message")
		} else if resp.StatusCode == 429 {
			fmt.Println("" + yel + "▏" + r + "(" + yel + "+" + r + ") RateLimit")
		} else {
			fmt.Println(""+red+"▏"+r+"("+red+"+"+r+") Failed To Send:",
				Xc.Errmsg(*resp),
			)
		}
	}
}

func X() Config {
	x := Config{}
	return x
}

func T() Checker {
	x := Checker{}
	return x
}
