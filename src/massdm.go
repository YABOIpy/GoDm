package massdm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)


func (Xc *Config) Dm(ID string, Token string, Msg string, Cookies string) {
	payload := map[string]string{
		"content": Msg,
	}
	req, err := http.NewRequest("POST", "https://discord.com/api/v9/channels/"+ID+"/messages",
		bytes.NewBuffer(Xc.Marsh(payload)),
	)
	Xc.Errs(err)
	for x,o := range map[string]string{
		"accept": "*/*",
		"accept-encoding": "gzip, deflate, br",
		"accept-language": "en-US,en-GB;q=0.9",
		"authorization": Token,
		"content-length": strconv.Itoa(content(string(Xc.Marsh(payload))).Length),
		"content-type": "application/json",
		"cookie": Cookies,
		"origin": "https://discord.com",
		"referer": "https://discord.com/channels/",
		"sec-fetch-dest": "empty",
		"sec-fetch-mode": "cors",
		"sec-fetch-site": "same-origin",
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9007 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
		"x-debug-options": "bugReporterEnabled",
		"x-discord-locale": "en-US",
		"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDA3Iiwib3NfdmVyc2lvbiI6IjEwLjAuMjIwMDAiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTYwNjQ1LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==",
	} { 
		req.Header.Set(x,o)
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
	for x,o := range map[string]string{
		"accept": "*/*",
		"accept-encoding": "gzip, deflate, br",
		"accept-language": "en-US,en-GB;q=0.9",
		"authorization": Token,
		"content-type": "application/json",
		"cookie": Cookies,
		"origin": "https://discord.com",
		"referer": "https://discord.com/channels/",
		"sec-fetch-dest": "empty",
		"sec-fetch-mode": "cors",
		"sec-fetch-site": "same-origin",
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9007 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
		"x-debug-options": "bugReporterEnabled",
		"x-discord-locale": "en-US",
		"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDA3Iiwib3NfdmVyc2lvbiI6IjEwLjAuMjIwMDAiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTYwNjQ1LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==",
	} {
		req.Header.Set(x,o)
	}
	resp, err := Client.Do(req)
	Xc.Errs(err)
	if resp.StatusCode == 200 {
		fmt.Println(""+grn+"▏"+r+"("+grn+"+"+r+") Closed Channel:"+clr+"", ID)
	} else {
		fmt.Println(""+red+"▏"+r+"("+red+"-"+r+") Failed To Close Dm:"+clr+"", ID, "ERR |: ",)
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
	for x,o := range map[string]string{
		"accept": "*/*",
		"accept-encoding": "gzip, deflate, br",
		"accept-language": "en-US,en-GB;q=0.9",
		"authorization": Token,
		"content-length": Lenght,
		"content-type": "application/json",
		"cookie": Cookies,
		"origin": "https://discord.com",
		"referer": "https://discord.com/channels/",
		"sec-fetch-dest": "empty",
		"sec-fetch-mode": "cors",
		"sec-fetch-site": "same-origin",
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9007 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
		"x-context-properties": "e30=",
		"x-debug-options": "bugReporterEnabled",
		"x-discord-locale": "en-US",
		"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDA3Iiwib3NfdmVyc2lvbiI6IjEwLjAuMjIwMDAiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTYwNjQ1LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==",
	} { 
		req.Header.Set(x,o)
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
	for x,o := range map[string]string{
		"accept": "*/*",
		"accept-encoding": "gzip, deflate, br",
		"accept-language": "en-US,en-GB;q=0.9",
		"authorization": Token,
		"content-length": Length,
		"content-type": "application/json",
		"cookie": Cookies,
		"origin": "https://discord.com",
		"referer": "https://discord.com/channels/@me/"+strconv.Itoa(ID)+"",
		"sec-fetch-dest": "empty",
		"sec-fetch-mode": "cors",
		"sec-fetch-site": "same-origin",
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9007 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
		"x-debug-options": "bugReporterEnabled",
		"x-discord-locale": "en-US",
		"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDA3Iiwib3NfdmVyc2lvbiI6IjEwLjAuMjIwMDAiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTYwNjQ1LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==",
	} {
		req.Header.Set(x,o)
	}
	resp, err := Client.Do(req)
	Xc.Errs(err)
	if resp.StatusCode == 204 {
		fmt.Println(""+grn+"▏"+r+"("+grn+"+"+r+") Blocked User:"+clr+"", ID)
	} else {
		fmt.Println(""+red+"▏"+r+"("+red+"-"+r+") Failed To Block:"+clr+"", ID, "ERR |: ",)
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
				map[string]string{"":""},
			),
		),
	)
	Xc.Errs(err)
	for x,o := range map[string]string{
		"accept": "*/*",
		"accept-encoding": "gzip, deflate, br",
		"accept-language": "en-US,en-GB;q=0.9",
		"authorization": Token,
		"content-type": "application/json",
		"cookie": Cookies,
		"origin": "https://discord.com",
		"referer": "https://discord.com/channels/",
		"sec-fetch-dest": "empty",
		"sec-fetch-mode": "cors",
		"sec-fetch-site": "same-origin",
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9007 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
		"x-debug-options": "bugReporterEnabled",
		"x-discord-locale": "en-US",
		"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDA3Iiwib3NfdmVyc2lvbiI6IjEwLjAuMjIwMDAiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTYyNjg2LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==",
	} {
		req.Header.Set(x,o)
	}
	resp, err := Client.Do(req)
	if resp.StatusCode == 200 {
		fmt.Println(""+grn+"▏"+r+"("+grn+"+"+r+") Joined "+clr+"discord.gg/"+invite)
	} else {
		Xc.Decerr(*resp)
	}
} 




func X() Config {
	x := Config{}
	return x
}
