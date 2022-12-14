package massdm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	http "github.com/Danny-Dasilva/fhttp"
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
		"content-length":     strconv.Itoa(content(string(Xc.Marsh(payload))).Length),
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

func (Xc *Config) Create(ID int, Token string, Msg string) (string, error) {
	payload := []byte("{\"recipients\":[\"" + strconv.Itoa(ID) + "\"]}")
	Lenght := strconv.Itoa(content(string(payload)).Length)
	req, err := http.NewRequest("POST", "https://discord.com/api/v9/users/@me/channels",
		bytes.NewBuffer(payload),
	)

	Xc.Errs(err)
	for x, o := range map[string]string{
		"accept":               "*/*",
		"accept-encoding":      "gzip, deflate, br",
		"accept-language":      "en-US,en-GB;q=0.9",
		"authorization":        Token,
		"content-length":       Lenght,
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
	Length := strconv.Itoa(content(string(Xc.Marsh(p))).Length)
	req, err := http.NewRequest("PUT", "https://discord.com/api/v9/users/@me/relationships/"+strconv.Itoa(ID)+"",
		bytes.NewBuffer(Xc.Marsh(p)),
	)

	Xc.Errs(err)
	for x, o := range map[string]string{
		"accept":             "*/*",
		"accept-encoding":    "gzip, deflate, br",
		"accept-language":    "en-US,en-GB;q=0.9",
		"authorization":      Token,
		"content-length":     Length,
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

func (Xc *Config) Scrape_ID(Token string, IDs string) {
	// reqs, err := http.NewRequest("GET", "https://discord.com/api/guilds/"+IDs+"/channels",
	// 	nil,
	// )
	// Xc.Errs(err)
	// for x,o := range map[string]string{
	// 	"accept": "*/*",
	// 	"accept-encoding": "gzip, deflate, br",
	// 	"accept-language": "en-US,en-GB;q=0.9",
	// 	"authorization": Token,
	// 	"content-type": "application/json",
	// 	"cookie": Cookies,
	// 	"origin": "https://discord.com",
	// 	"referer": "https://discord.com/channels/",
	// 	"sec-fetch-dest": "empty",
	// 	"sec-fetch-mode": "cors",
	// 	"sec-fetch-site": "same-origin",
	// 	"user-agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9007 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
	// } {
	// 	reqs.Header.Set(x,o)
	// }

	// resps, err := Client.Do(reqs)
	// Xc.Errs(err)

	var ids = []string{}
	// var data1 ChannelData
	// body, err := ReadBody(*resps)
	// Xc.Errs(err)

	// err = json.Unmarshal(body, &data1)
	// Xc.Errs(err)
	// fmt.Println(data1)
	// if reqs.Response.StatusCode == http.StatusOK {
	// 	for _, x := range data1 {
	// 	   fmt.Println(x.ID)
	// 	}
	// }
	req, err := http.NewRequest("GET", "https://discord.com/api/v9/channels/"+IDs+"/messages?limit=100",
		nil,
	)
	Xc.Errs(err)
	for x, o := range map[string]string{
		"accept":          "*/*",
		"accept-encoding": "gzip, deflate, br",
		"accept-language": "en-US,en-GB;q=0.9",
		"authorization":   Token,
		"content-type":    "application/json",
		"cookie":          Cookies,
		"origin":          "https://discord.com",
		"referer":         "https://discord.com/channels/",
		"sec-fetch-dest":  "empty",
		"sec-fetch-mode":  "cors",
		"sec-fetch-site":  "same-origin",
		"user-agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9007 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
	} {
		req.Header.Set(x, o)
	}

	//var ids = []string{}
	resp, ers := Client.Do(req)
	_ = ers

	if resp.StatusCode == 200 {
		defer func(Body io.ReadCloser) {

			err := Body.Close()
			Xc.Errs(err)
		}(resp.Body)

		var data ChannelData
		body, err := ReadBody(*resp)
		Xc.Errs(err)

		err = json.Unmarshal(body, &data)
		Xc.Errs(err)
		for _, x := range data {
			if !contains(ids, x.Author.ID) {
				ids = append(ids, x.Author.ID)

			}
		}
	} else {
		Xc.Decerr(*resp)
	}

	Xc.Write_ID(ids)
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
			continue
		} else if resp.StatusCode == 429 {
			fmt.Println("" + yel + "▏" + r + "(" + yel + "+" + r + ") RateLimit")
			continue
		} else {
			fmt.Println("" + red + "▏" + r + "(" + red + "+" + r + ") Failed To Send")
			continue
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
		Xc.Errs(ers)
		if resp.StatusCode == 200 {
			fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Sent Message")
			continue
		} else if resp.StatusCode == 429 {
			fmt.Println("" + yel + "▏" + r + "(" + yel + "+" + r + ") RateLimit")
			continue
		} else {
			fmt.Println("" + red + "▏" + r + "(" + red + "+" + r + ") Failed To Send")
			continue
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

package massdm

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"encoding/json"
	"fmt"
	http "github.com/Danny-Dasilva/fhttp"
	"github.com/andybalholm/brotli"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	xhttp "net/http"
	"os"
	"os/exec"
	"unicode/utf8"
)

func (Xc *Config) BuildNum() {

}

func ReadBody(resp http.Response) ([]byte, error) {

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipreader, err := zlib.NewReader(bytes.NewReader(body))
		if err != nil {
			return nil, err
		}
		gzipbody, err := ioutil.ReadAll(gzipreader)
		if err != nil {
			return nil, err
		}
		return gzipbody, nil
	}

	if resp.Header.Get("Content-Encoding") == "br" {
		brreader := brotli.NewReader(bytes.NewReader(body))
		brbody, err := ioutil.ReadAll(brreader)
		if err != nil {
			fmt.Println(string(brbody))
			return nil, err
		}

		return brbody, nil
	}
	return body, nil
}

func (Xc *Config) Config() Config {
	var config Config
	conf, err := os.Open("config.json")
	defer conf.Close()
	config.Errs(err)
	xp := json.NewDecoder(conf)
	xp.Decode(&config)
	return config

}

func (Xc *Config) ReadFile(files string) ([]string, error) {
	file, err := os.Open(files)
	Xc.Errs(err)
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func (Xc *Config) WriteFile(files string, item string) {
	f, err := os.OpenFile(files, os.O_RDWR|os.O_APPEND, 0660)
	Xc.Errs(err)
	defer f.Close()
	_, ers := f.WriteString(item + "\n")
	Xc.Errs(ers)
}

func cookies() Config {
	Xc := Config{}
	req, err := http.NewRequest("GET", "https://discord.com", nil)
	Xc.Errs(err)
	resp, er := Client.Do(req)
	Xc.Errs(er)
	defer resp.Body.Close()
	Cookie := Config{}
	if resp.Cookies() != nil {
		for _, cookie := range resp.Cookies() {
			if cookie.Name == "__dcfduid" {
				Cookie.Dcfd = cookie.Value
			}
			if cookie.Name == "__sdcfduid" {
				Cookie.Sdcfd = cookie.Value
			}
		}
	} else {
		cookies()
	}
	return Cookie
}

func (Xc *Config) Decerr(resp http.Response) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	Xc.Errs(err)
	var data interface{}
	json.Unmarshal(body, &data)
	fmt.Println(data)
}

//gecko
func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func (Xc *Config) Marsh(payload map[string]string) []byte {
	x, _ := json.Marshal(payload)
	return x
}

func content(payload string) Config {
	content := Config{}
	amt := utf8.RuneCountInString(payload)
	content.Length = amt
	return content
}

func (Xc *Config) Cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (Xc *Config) Errs(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (Xc *Config) WebSock(token string) {

	dialer := websocket.Dialer{}
	ws, _, err := dialer.Dial("wss://gateway.discord.gg/?encoding=json&v=9&compress=zlib-stream", xhttp.Header{
		"Origin":     []string{"https://discord.com"},
		"User-Agent": []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"},
	})
	Xc.Errs(err)
	_, _, _ = ws.ReadMessage()
	Payload, _ := json.Marshal(&PayloadWsLogin{
		Op: 2,
		D: WsD{
			Token:        token,
			Capabilities: 125,
			Properties: XProperties{
				Os:                     "Windows",
				Browser:                "Chrome",
				Device:                 "",
				SystemLocale:           "en-US",
				BrowserUserAgent:       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36",
				BrowserVersion:         "107.0.0.0",
				OsVersion:              "10",
				Referrer:               "https://www.google.com",
				ReferringDomain:        "www.google.com",
				ReferrerCurrent:        "",
				ReferringDomainCurrent: "",
				ReleaseChannel:         "stable",
				ClientBuildNumber:      158183,
				ClientEventSource:      nil,
			},
		},
		Presence: WsPresence{
			Status:     "online",
			Since:      0,
			Activities: nil,
			Afk:        false,
		},
		Compress: false,
		ClientState: WsClientState{
			GuildHashes:              WsGH{},
			HighestLastMessageID:     "0",
			ReadStateVersion:         0,
			UserGuildSettingsVersion: -1,
			UserSettingsVersion:      -1,
		},
	})
	err = ws.WriteMessage(websocket.TextMessage, Payload)
	Xc.Errs(err)
	_, _, _ = ws.ReadMessage()
	_, _, _ = ws.ReadMessage()
	fmt.Println("" + clr + "▏" + r + "(" + clr + "o" + r + ") Connected to " + clr + "Websocket" + r + "")
	ws.Close()
}

package massdm

import (
	_ "crypto/tls"
	http "github.com/Danny-Dasilva/fhttp"
	goclient "massdm/client"
	"time"
)

type ChannelID struct {
	ID string `json:"id,omitempty"`
}

type Checker struct {
	Client  http.Client
	Invalid int
	Locked  int
	Token   string
	Valid   int
	Resp    bool
	All     int
}

type Config struct {
	Headers map[string]string
	Dcfd    string
	Sdcfd   string
	Length  int
	ID      string

	Settings struct {
		Websock bool `json:"Websocket"`
		Block   bool `json:"Block_Usr"`
		Close   bool `json:"Close_DM"`
	} `json:"Settings"`
	ProxySettings ProxySettings `json:"proxy_settings"`
}
type ProxySettings struct {
	Proxy   string `json:"Proxy"`
	Timeout int    `json:"timeout"`
}

type ChannelData []struct {
	ID               string      `json:"id"`
	LastMessageID    string      `json:"last_message_id,omitempty"`
	LastPinTimestamp time.Time   `json:"last_pin_timestamp,omitempty"`
	Type             int         `json:"type"`
	Name             string      `json:"name"`
	Position         int         `json:"position"`
	ParentID         string      `json:"parent_id"`
	Topic            interface{} `json:"topic,omitempty"`
	GuildID          string      `json:"guild_id"`
	Nsfw             bool        `json:"nsfw"`
	RateLimitPerUser int         `json:"rate_limit_per_user,omitempty"`
	Banner           interface{} `json:"banner,omitempty"`
	Bitrate          int         `json:"bitrate,omitempty"`
	UserLimit        int         `json:"user_limit,omitempty"`
	RtcRegion        interface{} `json:"rtc_region,omitempty"`

	PermissionOverwrites []struct {
		ID       string `json:"id"`
		Type     string `json:"type"`
		Allow    int    `json:"allow"`
		Deny     int    `json:"deny"`
		AllowNew string `json:"allow_new"`
		DenyNew  string `json:"deny_new"`
	} `json:"permission_overwrites"`
	Author struct {
		ID            string `json:"id"`
		Username      string `json:"username"`
		Avatar        string `json:"avatar"`
		Discriminator string `json:"discriminator"`
		PublicFlags   int    `json:"public_flags"`
	} `json:"author"`
}

type FormField struct {
	FieldType   string   `json:"field_type"`
	Label       string   `json:"label"`
	Description string   `json:"description"`
	Required    bool     `json:"required"`
	Values      []string `json:"values"`
	Response    bool     `json:"response"`
}

type Bypass struct {
	Version    string      `json:"version"`
	FormFields []FormField `json:"form_fields"`
}

type PayloadWsLogin struct {
	Op          int           `json:"op"`
	D           WsD           `json:"d"`
	Presence    WsPresence    `json:"presence"`
	Compress    bool          `json:"compress"`
	ClientState WsClientState `json:"client_state"`
}

type WsD struct {
	Token        string      `json:"token"`
	Capabilities int         `json:"capabilities"`
	Properties   XProperties `json:"properties"`
}

type WsPresence struct {
	Status     string        `json:"status"`
	Since      int           `json:"since"`
	Activities []interface{} `json:"activities"`
	Afk        bool          `json:"afk"`
}

type WsGH struct{}

type WsClientState struct {
	GuildHashes              WsGH   `json:"guild_hashes"`
	HighestLastMessageID     string `json:"highest_last_message_id"`
	ReadStateVersion         int    `json:"read_state_version"`
	UserGuildSettingsVersion int    `json:"user_guild_settings_version"`
	UserSettingsVersion      int    `json:"user_settings_version"`
}

type XProperties struct {
	Os                     string      `json:"os"`
	Browser                string      `json:"browser"`
	Device                 string      `json:"device"`
	SystemLocale           string      `json:"system_locale"`
	BrowserUserAgent       string      `json:"browser_user_agent"`
	BrowserVersion         string      `json:"browser_version"`
	OsVersion              string      `json:"os_version"`
	Referrer               string      `json:"referrer"`
	ReferringDomain        string      `json:"referring_domain"`
	ReferrerCurrent        string      `json:"referrer_current"`
	ReferringDomainCurrent string      `json:"referring_domain_current"`
	ReleaseChannel         string      `json:"release_channel"`
	ClientBuildNumber      int         `json:"client_build_number"`
	ClientEventSource      interface{} `json:"client_event_source"`
}

var (
	cfg         = Config{}
	Client, err = goclient.NewClient(goclient.Browser{JA3: "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0", UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0", Cookies: nil}, cfg.ProxySettings.Timeout, false, "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0", "")
	Cookies     = "__dcfduid=" + cookies().Dcfd + "; " + "__sdcfduid=" + cookies().Sdcfd + "; "
	urls        = "https://discord.com/api/v9/users/@me/affinities/guilds"
	grn         = "\033[32m"
	yel         = "\033[33m"
	red         = "\033[31m"
	clr         = "\033[36m"
	r           = "\033[39m"

	Logo = `
	____` + clr + `_____` + r + `__` + clr + `____     ` + r + `____` + clr + `____` + r + `____` + clr + `__  ___
	` + r + `__` + clr + `  ____/` + r + `_  ` + clr + `__ \    ` + r + `___` + clr + `  __ \` + r + `__` + clr + `   |/  /
	` + r + `_ ` + clr + ` / __ ` + r + `_` + clr + `  / / /    ` + r + `__` + clr + `  / / /` + r + `_` + clr + `  /|_/ / 
	` + clr + `/ /_/ / / /_/ /     ` + r + `_  ` + clr + `/_/ /` + r + `_` + clr + `  /  / /  
	\____/  \____/      /_____/ /_/  /_/   
    
	[` + r + `Go Discord Mass Dm` + clr + `]			~` + r + `YABOI` + clr + `
	[` + r + `1` + clr + `]` + r + ` Mass DM ` + clr + `		[` + r + `10` + clr + `]` + r + ` Mass Ping ` + clr + `
	[` + r + `2` + clr + `]` + r + ` Dm Spam ` + clr + `		[` + r + `11` + clr + `]` + r + `	x ` + clr + `
	[` + r + `3` + clr + `]` + r + ` React Verify ` + clr + `	[` + r + `12` + clr + `]` + r + `	x ` + clr + `
	[` + r + `4` + clr + `]` + r + ` Joiner ` + clr + `		[` + r + `13` + clr + `]` + r + `	x ` + clr + `
	[` + r + `5` + clr + `]` + r + ` Leaver ` + clr + `		[` + r + `14` + clr + `]` + r + `	x ` + clr + `
	[` + r + `6` + clr + `]` + r + ` Accept Rules ` + clr + `	[` + r + `15` + clr + `]` + r + `	x ` + clr + `
	[` + r + `7` + clr + `]` + r + ` Raid Channel ` + clr + `	[` + r + `16` + clr + `]` + r + `	x ` + clr + `
	[` + r + `8` + clr + `]` + r + ` Scrape Users ` + clr + `	[` + r + `17` + clr + `]` + r + `	x ` + clr + `
	[` + r + `9` + clr + `]` + r + ` Check Tokens ` + clr + `	[` + r + `18` + clr + `]` + r + `	x

	Choice ` + clr + `>>:` + r + ` `
)
