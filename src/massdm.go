package massdm

import (
	"bytes"
	"encoding/json"
	"fmt"
	http "github.com/Danny-Dasilva/fhttp"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
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
	for x, o := range map[string]string{
		"accept":             "*/*",
		"accept-encoding":    "gzip, deflate, br",
		"accept-language":    "en-US,en-GB;q=0.9",
		"authorization":      Token,
		"content-type":       "application/json",
		"cookie":             Cookies,
		"origin":             "https://discord.com",
		"referer":            "https://discord.com/channels/",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9007 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
		"x-debug-options":    "bugReporterEnabled",
		"x-discord-locale":   "en-US",
		"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDA3Iiwib3NfdmVyc2lvbiI6IjEwLjAuMjIwMDAiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTYwNjQ1LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==",
	} {
		req.Header.Set(x, o)
	}
	resp, err := Client.Do(req)
	Xc.Errs(err)
	if resp.StatusCode == 200 {
		fmt.Println(""+grn+"▏"+r+"("+grn+"+"+r+") Sent Message To:"+clr+"", ID)
	} else {
		fmt.Println(""+red+"▏"+r+"("+red+"-"+r+") Failed Sent Message To:"+clr+"", ID)
	}
}

func (Xc *Config) CloseDm(ID string, Token string, Cookies string) {
	req, err := http.NewRequest("DELETE",
		"https://discord.com/api/v9/channels/"+ID+"?silent=false",
		nil,
	)
	Xc.Errs(err)
	for x, o := range map[string]string{
		"accept":             "*/*",
		"accept-encoding":    "gzip, deflate, br",
		"accept-language":    "en-US,en-GB;q=0.9",
		"authorization":      Token,
		"content-type":       "application/json",
		"cookie":             Cookies,
		"origin":             "https://discord.com",
		"referer":            "https://discord.com/channels/",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9007 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
		"x-debug-options":    "bugReporterEnabled",
		"x-discord-locale":   "en-US",
		"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDA3Iiwib3NfdmVyc2lvbiI6IjEwLjAuMjIwMDAiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTYwNjQ1LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==",
	} {
		req.Header.Set(x, o)
	}
	resp, err := Client.Do(req)
	Xc.Errs(err)
	if resp.StatusCode == 200 {
		fmt.Println(""+grn+"▏"+r+"("+grn+"+"+r+") Closed Channel:"+clr+"", ID)
	} else {
		fmt.Println(""+red+"▏"+r+"("+red+"-"+r+") Failed To Close Dm:"+clr+"", ID, "ERR |: ")
		Xc.Decerr(*resp)
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
	for x, o := range map[string]string{
		"accept":             "*/*",
		"accept-encoding":    "gzip, deflate, br",
		"accept-language":    "en-US,en-GB;q=0.9",
		"authorization":      Token,
		"content-type":       "application/json",
		"cookie":             Cookies,
		"origin":             "https://discord.com",
		"referer":            "https://discord.com/channels/",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9007 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
		"x-debug-options":    "bugReporterEnabled",
		"x-discord-locale":   "en-US",
		"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDA3Iiwib3NfdmVyc2lvbiI6IjEwLjAuMjIwMDAiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTYwNjQ1LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==",
	} {
		req.Header.Set(x, o)
	}
	resp, err := Client.Do(req)
	Xc.Errs(err)

	switch resp.StatusCode {
	case 204:
		fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Added Emoji")
	default:
		fmt.Println(""+red+"▏"+r+"("+red+"-"+r+") Failed To Add Emoji"+clr+"", "ERR |: "+r)
		Xc.Decerr(*resp)

	}
}


func (Xc *Config) Create(ID int, Token string, Msg string) (string, error) {
	payload := []byte("{\"recipients\":[\"" + strconv.Itoa(ID) + "\"]}")
	req, err := http.NewRequest("POST", "https://discord.com/api/v9/users/@me/channels",
		bytes.NewBuffer(payload),
	)

	Xc.Errs(err)
	for x, o := range map[string]string{
		"accept":               "*/*",
		"accept-encoding":      "gzip, deflate, br",
		"accept-language":      "en-US,en-GB;q=0.9",
		"authorization":        Token,
		"content-type":         "application/json",
		"cookie":               Cookies,
		"origin":               "https://discord.com",
		"referer":              "https://discord.com/channels/",
		"sec-fetch-dest":       "empty",
		"sec-fetch-mode":       "cors",
		"sec-fetch-site":       "same-origin",
		"user-agent":           "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9007 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
		"x-context-properties": "e30=",
		"x-debug-options":      "bugReporterEnabled",
		"x-discord-locale":     "en-US",
		"x-super-properties":   "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDA3Iiwib3NfdmVyc2lvbiI6IjEwLjAuMjIwMDAiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTYwNjQ1LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==",
	} {
		req.Header.Set(x, o)
	}
	var flake ChannelID
	resp, err := Client.Do(req)
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
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		Xc.Errs(err)
		var data interface{}
		json.Unmarshal(body, &data)
		fmt.Print(data)
	}
	return flake.ID, err
}

func (Xc *Config) Block(ID int, Token string, Cookies string) {
	p := map[string]string{"type": "2"}
	req, err := http.NewRequest("PUT", "https://discord.com/api/v9/users/@me/relationships/"+strconv.Itoa(ID)+"",
		bytes.NewBuffer(Xc.Marsh(p)),
	)

	Xc.Errs(err)
	for x, o := range map[string]string{
		"accept":             "*/*",
		"accept-encoding":    "gzip, deflate, br",
		"accept-language":    "en-US,en-GB;q=0.9",
		"authorization":      Token,
		"content-type":       "application/json",
		"cookie":             Cookies,
		"origin":             "https://discord.com",
		"referer":            "https://discord.com/channels/@me/" + strconv.Itoa(ID) + "",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9007 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
		"x-debug-options":    "bugReporterEnabled",
		"x-discord-locale":   "en-US",
		"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDA3Iiwib3NfdmVyc2lvbiI6IjEwLjAuMjIwMDAiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTYwNjQ1LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==",
	} {
		req.Header.Set(x, o)
	}
	resp, err := Client.Do(req)
	Xc.Errs(err)
	if resp.StatusCode == 204 {
		fmt.Println(""+grn+"▏"+r+"("+grn+"+"+r+") Blocked User:"+clr+"", ID)
	} else {
		fmt.Println(""+red+"▏"+r+"("+red+"-"+r+") Failed To Block:"+clr+"", ID, "ERR |: ")
		Xc.Decerr(*resp)
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
	_ = err
	for x, o := range map[string]string{
		"accept":             "*/*",
		"accept-encoding":    "gzip, deflate, br",
		"accept-language":    "en-US,en-GB;q=0.9",
		"authorization":      Token,
		"content-type":       "application/json",
		"cookie":             Cookies,
		"origin":             "https://discord.com",
		"referer":            "https://discord.com/channels/",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9007 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
		"x-debug-options":    "bugReporterEnabled",
		"x-discord-locale":   "en-US",
		"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDA3Iiwib3NfdmVyc2lvbiI6IjEwLjAuMjIwMDAiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTYyNjg2LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==",
	} {
		req.Header.Set(x, o)
	}
	resp, err := Client.Do(req)
	_ = err
	if resp.StatusCode == 200 {
		fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Joined " + clr + "discord.gg/" + invite)
	} else {
		Xc.Decerr(*resp)
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
	_ = err
	for x, o := range map[string]string{
		"accept":             "*/*",
		"accept-encoding":    "gzip, deflate, br",
		"accept-language":    "en-US,en-GB;q=0.9",
		"authorization":      Token,
		"content-type":       "application/json",
		"cookie":             Cookies,
		"origin":             "https://discord.com",
		"referer":            "https://discord.com/channels/",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9007 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
		"x-debug-options":    "bugReporterEnabled",
		"x-discord-locale":   "en-US",
		"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDA3Iiwib3NfdmVyc2lvbiI6IjEwLjAuMjIwMDAiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTYyNjg2LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==",
	} {
		req.Header.Set(x, o)
	}
	resp, err := Client.Do(req)
	_ = err
	if resp.StatusCode == 204 {
		fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Left Server")
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		Xc.Errs(err)
		fmt.Println(string(body))
		Xc.Decerr(*resp)
	}
}

func (Xc *Config) Agree(Token string, invite string, ID string) {
	req, err := http.NewRequest("POST",
		"https://discord.com/api/v9/guilds/"+ID+"/member-verification?with_guild=false&invite_code="+invite+"",
		nil,
	)
	Xc.Errs(err)
	for x, o := range map[string]string{
		"accept":             "*/*",
		"accept-encoding":    "gzip, deflate, br",
		"accept-language":    "en-US,en-GB;q=0.9",
		"authorization":      Token,
		"content-type":       "application/json",
		"cookie":             Cookies,
		"origin":             "https://discord.com",
		"referer":            "https://discord.com/channels/",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9007 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
		"x-debug-options":    "bugReporterEnabled",
		"x-discord-locale":   "en-US",
		"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDA3Iiwib3NfdmVyc2lvbiI6IjEwLjAuMjIwMDAiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTYyNzg4LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==",
	} {
		req.Header.Set(x, o)
	}

	resp, err := http.DefaultClient.Do(req)
	Xc.Errs(err)

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
	for x, o := range map[string]string{
		"accept":             "*/*",
		"accept-encoding":    "gzip, deflate, br",
		"accept-language":    "en-US,en-GB;q=0.9",
		"authorization":      Token,
		"content-type":       "application/json",
		"cookie":             Cookies,
		"origin":             "https://discord.com",
		"referer":            "https://discord.com/channels/",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9007 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
		"x-debug-options":    "bugReporterEnabled",
		"x-discord-locale":   "en-US",
		"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDA3Iiwib3NfdmVyc2lvbiI6IjEwLjAuMjIwMDAiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTYyNzg4LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==",
	} {
		reqs.Header.Set(x, o)
	}
	resps, es := http.DefaultClient.Do(reqs)
	_ = es
	if resps.StatusCode == 201 {
		fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Verified Token")
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

func (Xc *Config) Scrape_ID(Ws *Socket, Token string, CID string, GID string, index int) {
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
		err := Ws.WriteJSONe(&Event{
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
	err := Ws.WriteJSONe(&Event{
		Op:   14,
		Data: payload,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Ws)
	var memberids []string
	for _, member := range Ws.Members {
		memberids = append(memberids, member.User.ID)
		fmt.Println(memberids)
	}
}
func (c *Socket) WriteJSONe(e *Event) error {
	return c.Conn.WriteJSON(e)
}

func (Xc *Config) Write_ID(ids []string) {
	for _, v := range ids {
		fmt.Println(""+grn+"▏"+r+"("+grn+"+"+r+") ID:", v)
		f, err := os.OpenFile("ids.txt", os.O_RDWR|os.O_APPEND, 0660)
		Xc.Errs(err)
		defer f.Close()
		_, ers := f.WriteString(v + "\n")
		Xc.Errs(ers)
	}
}

func (Xc *Config) Raider(Token string, message string, ID string) {
	for true {

		payload := map[string]string{
			"content": message,
		}
		xp, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", "https://discord.com/api/v9/channels/"+ID+"/messages", bytes.NewBuffer(xp))
		_ = err
		for x, o := range map[string]string{
			"accept":             "*/*",
			"accept-encoding":    "gzip, deflate, br",
			"accept-language":    "en-US,en-NL;q=0.9,en-GB;q=0.8",
			"authorization":      Token,
			"content-type":       "application/json",
			"cookie":             Cookies,
			"origin":             "https://discord.com",
			"referer":            "https://discord.com/channels/@me/" + ID + "",
			"sec-fetch-dest":     "empty",
			"sec-fetch-mode":     "cors",
			"sec-fetch-site":     "same-origin",
			"user-agent":         "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9006 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
			"x-debug-options":    "bugReporterEnabled",
			"x-discord-locale":   "en-US",
			"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDA3Iiwib3NfdmVyc2lvbiI6IjEwLjAuMjIwMDAiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTYyNzg4LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==",
		} {
			req.Header.Set(x, o)
		}
		resp, ers := Client.Do(req)
		_ = ers
		if resp.StatusCode == 200 {
			fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Sent Message")
		} else if resp.StatusCode == 429 {
			fmt.Println("" + yel + "▏" + r + "(" + yel + "+" + r + ") RateLimit")
		} else {
			fmt.Println("" + red + "▏" + r + "(" + red + "+" + r + ") Failed To Send")
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

		Message += msg
		req, err := http.NewRequest("POST", "https://discord.com/api/v9/channels/"+ID+"/messages",
			bytes.NewBuffer(
				Xc.Marsh(map[string]string{
					"content": Message,
				}),
			),
		)
		Xc.Errs(err)
		for x, o := range map[string]string{
			"accept":             "*/*",
			"accept-encoding":    "gzip, deflate, br",
			"accept-language":    "en-US,en-NL;q=0.9,en-GB;q=0.8",
			"authorization":      Token,
			"content-type":       "application/json",
			"cookie":             Cookies,
			"origin":             "https://discord.com",
			"referer":            "https://discord.com/channels/@me/" + ID + "",
			"sec-fetch-dest":     "empty",
			"sec-fetch-mode":     "cors",
			"sec-fetch-site":     "same-origin",
			"user-agent":         "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9006 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
			"x-debug-options":    "bugReporterEnabled",
			"x-discord-locale":   "en-US",
			"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDA3Iiwib3NfdmVyc2lvbiI6IjEwLjAuMjIwMDAiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTYyNzg4LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==",
		} {
			req.Header.Set(x, o)
		}
		fmt.Println(Token)

		resp, ers := Client.Do(req)
		Xc.Errs(ers)
		if resp.StatusCode == 200 {
			fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Sent Message")
		} else if resp.StatusCode == 429 {
			fmt.Println("" + yel + "▏" + r + "(" + yel + "+" + r + ") RateLimit")
		} else {
			Xc.Decerr(*resp)
			fmt.Println("" + red + "▏" + r + "(" + red + "+" + r + ") Failed To Send")
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
