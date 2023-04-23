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
	"github.com/hugolgst/rich-go/client"
	"io/ioutil"
	"massdm/scraper"
	"math/rand"
	xhttp "net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

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

func (Xc *Config) Captcha(WebKey string) (cap string) {

	service := Xc.Config().Mode.Discord.CapAPI[0]

	if service == "capmonster" || service == "Capmonster" || service == "CAPMONSTER" {
		cap = Xc.Capmonster_Solve(WebKey)
	} else if service == "2captcha" || service == "2Captcha" || service == "2CAPTCHA" {
		cap = Xc.TwoCaptcha_Solve(WebKey)
	}

	return cap
}

func (Xc *Config) TwoCaptcha_Solve(WebKey string) (cap string) {
	url := "http://2captcha.com/in.php?key=" + Xc.Config().Mode.Discord.CapAPI[1] + "&method=hcaptcha&sitekey=" + WebKey + "&pageurl=https://discord.com"
	req, err := http.NewRequest("GET", url, nil)
	Xc.Errs(err)

	resp, err := Client.Do(req)
	Xc.Errs(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	Xc.Errs(err)

	if resp.StatusCode == 200 {
		fmt.Println(resp, string(body))
	} else {
		fmt.Println("Somethign Whent wrong: ", resp)
	}
	return cap
}

func (Xc *Config) Capmonster_Solve(WebKey string) (cap string) {
	Key := Xc.Config().Mode.Discord.CapAPI[1]
	payload := map[string]interface{}{
		"clientKey": Key,
		"task": map[string]interface{}{
			"type":       "HCaptchaTaskProxyless",
			"userAgent":  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36 Edg/92.0.902.73",
			"websiteKey": WebKey,
			"websiteURL": "https://discord.com/",
		},
	}
	req, err := xhttp.NewRequest("POST",
		"https://api.capmonster.cloud/createTask",
		bytes.NewBuffer(
			Xc.Marsh_btn(payload),
		),
	)
	Xc.Errs(err)

	resp, err := xhttp.DefaultClient.Do(req)
	Xc.Errs(err)

	var data CapmonsterResp
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)
	Xc.Errs(err)

	if resp.StatusCode == 200 {
		task := data.TaskID

		for {
			req, err := xhttp.NewRequest("POST",
				"https://api.capmonster.cloud/getTaskResult",
				bytes.NewBuffer(
					Xc.Marsh(
						map[string]string{
							"clientKey": Key,
							"taskId":    strconv.Itoa(task),
						},
					),
				),
			)
			Xc.Errs(err)

			resp, err := xhttp.DefaultClient.Do(req)
			Xc.Errs(err)

			responseBody := make(map[string]interface{})
			json.NewDecoder(resp.Body).Decode(&responseBody)
			status := responseBody["status"]
			Xc.Errs(err)
			if status == "ready" {
				cap = responseBody["solution"].(map[string]interface{})["gRecaptchaResponse"].(string)
				break
			} else if status == "processing" {
				continue
			} else {
				fmt.Println("[ERR] | ", string(body))
			}
		}
	} else {
		Xc.Errs(err)
	}

	return cap
}

func (Xc *Config) Errmsg(response http.Response) (errmsg string) {
	body, _ := ReadBody(response)

	if cfg.Con.Errors {
		errmsg = string(body)
	} else {
		return ""
	}

	return errmsg

}

func (Xc *Config) CheckConfig() {

	T, err := Xc.ReadFile("tokens.txt")
	Xc.Errs(err)

	if strings.Count(Xc.Config().Mode.Network.Proxy, "") > 1 {
		cfg.Con.ProxyMode = grn + "True"
	} else {
		cfg.Con.ProxyMode = red + "False"
	}

	if Xc.Config().Mode.Discord.RPC {
		go func() {
			Xc.Presence(len(T))
		}()
	}

	if Xc.Config().Mode.Configs.Errormsg {
		cfg.Con.Errors = true
	}

	if len(T) >= 1 && len(T) < 100 {
		cfg.Con.Solution = yel + strconv.Itoa(len(T))

	} else if len(T) == 0 {
		cfg.Con.Solution = red + strconv.Itoa(len(T))

	} else if len(T) >= 100 {
		cfg.Con.Tokenclr = grn
		cfg.Con.Solution = cfg.Con.Tokenclr + strconv.Itoa(len(T))

	} else {
		cfg.Con.Tokenclr = grn
		cfg.Con.Solution = cfg.Con.Tokenclr + strconv.Itoa(c.Con.Toks)
	}

}

func (Xc *Config) GetJa3() (ja3 string) {
	req, err := http.NewRequest("GET", "https://tls.peet.ws/api/clean", nil)

	Xc.Errs(err)

	resp, err := Client.Do(req)
	Xc.Errs(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	Xc.Errs(err)

	var data Ja3resp
	err = json.Unmarshal(body, &data)
	Xc.Errs(err)

	switch resp.StatusCode {
	case 200:
		fmt.Println(string(body))
		ja3 = string(data.Ja3)
	default:
		//ja3 = Xc.Config().Mode.Network.Ja3
	}

	return ja3
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


func (Xc *Config) GetCookie() string {
	var cook string
	go func() { cook = cookies() }()
	return cook
}

func (Xc *Config) ReadFile(files string) ([]string, error) {

	file, err := os.Open(files)
	Xc.Errs(err)
	defer file.Close()
	var value bool
	var lines, tokens []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if strings.Contains(files, "token") {
		for i := 0; i < len(lines); i++ {
			if strings.Contains(lines[i], ":") {
				format := strings.Split(lines[i], ":")
				tokens = append(tokens, format[2])
				value = true
			}
		}
		if value {
			lines = nil
			lines = tokens
		}
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

func cookies() (Cookies string) {
	Xc := Config{}

	cookie := []*http.Cookie{}
	req, err := http.NewRequest("GET", "https://discord.com", nil)
	Xc.Errs(err)

	resp, er := Client.Do(req)
	Xc.Errs(er)
	defer resp.Body.Close()

	if resp.Cookies() == nil {
		cookies()
	}

	cookie = append(cookie, resp.Cookies()...)
	for i := 0; i < len(cookie); i++ {
		if i == len(cookie)-1 {
			Cookies += fmt.Sprintf(`%s=%s`,
				cookie[i].Name,
				cookie[i].Value,
			)
		} else {
			Cookies += fmt.Sprintf(`%s=%s; `,
				cookie[i].Name,
				cookie[i].Value,
			)
		}
	}
	if !strings.Contains(Cookies, "locale=en-US; ") {
		Cookies += "; locale=en-US "
	}

	return Cookies
}

func (Xc *Config) CfBm() (string, error) {
	r, _ := http.Post("https://discord.com/register", "application/json", nil)
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	site := fmt.Sprintf(`https://discord.com/cdn-cgi/bm/cv/result?req_id=%s`, r)
	d := strings.Split(string(body), ",m:'")[1]
	data := strings.Split(d, "',s:")[0]

	payload := fmt.Sprintf(
		`
        {
            "m":"`+data+`",      
            "results":["`+c.Dcfd+`","`+c.Sdcfd+`"],
            "timing":`+string(((time.Now().Unix()*1000)-1420070400000)*4194304)+`,
            "fp":
                {
                    "id":3,
                    "e":{"r":[1920,1080],
                    "ar":[1054,1920],
                    "pr":1,
                    "cd":24,
                    "wb":true,
                    "wp":false,
                    "wn":false,
                    "ch":false,
                    "ws":false,
                    "wd":false
                }
            }
        }
        `, data, 60+rand.Intn(60),
	)
	req, err := http.NewRequest("POST", site, strings.NewReader(payload))
	Xc.Errs(err)

	resp, err := Client.Do(req)
	Xc.Errs(err)
	defer resp.Body.Close()
	if resp.Cookies() == nil {
		fmt.Println("Failed To Get Cookies: NIL")
	}
	if len(resp.Cookies()) == 0 {
		fmt.Println("Failed To Get Cookies: NIL")
	}
	var cookies string
	for _, cookie := range resp.Cookies() {
		cookies = cookies + cookie.Name + "=" + cookie.Value
	}
	return cookies, nil
}

func (Xc *Config) Decerr(resp http.Response) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	Xc.Errs(err)
	var data interface{}
	json.Unmarshal(body, &data)
	fmt.Println(data)
}

func (Xc *Config) Format_Tokens() {

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

func (Xc *Config) Marsh_btn(data map[string]interface{}) []byte {
	x, _ := json.Marshal(data)
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
		if Xc.Config().Mode.Configs.ErrorLog {
			Xc.WriteFile("errors.txt", err.Error())
		}
		fmt.Errorf(red+"▏"+r+"("+red+"+"+r+") Error"+red+":"+r, err)
	}
}

func (Xc *Config) Presence(Count int) {
	now := time.Now()
	err := client.Login(Xc.Config().Mode.Discord.AppID)
	Xc.Errs(err)
	if Count == 0 {
		Count += 1
	}
	client.SetActivity(client.Activity{
		State:      "Bots Loaded:",
		Details:    "Go Mass DM | github.com/YABOIpy",
		LargeImage: "b51b78ecc9e5711274931774e433b5e6",
		LargeText:  "https://github.com/yaboipy",
		SmallImage: "go-logo",
		SmallText:  "Ver " + fmt.Sprint(Xc.Config().Mode.Discord.Version),
		Buttons: []*client.Button{
			&client.Button{
				Label: "Go To Page",
				Url:   "https://github.com/yaboipy/go-massdm",
			},
		},
		Timestamps: &client.Timestamps{
			Start: &now,
		},
		Party: &client.Party{
			ID:         "-1",
			Players:    Count,
			MaxPlayers: Count,
		},
	})
}

func (Xc *Config) Socket(Token string) *Scraper.WsResp {
	c := Scraper.X()
	x, _ := c.Connect(Token)

	return x
}

func (Xc *Config) WebSock(token string) {

	dialer := websocket.Dialer{}
	ws, _, err := dialer.Dial("wss://gateway.discord.gg/?encoding=json&v=9&compress=zlib-stream", xhttp.Header{
		"Origin":     []string{"https://discord.com"},
		"User-Agent": []string{Xc.Config().Mode.Network.Agent},
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
			Presence: WsPresence{
				Status: Xc.Config().Mode.Discord.Status,
				Since:  0,
				Activities: PresenceData{
					Game: Gamedata{
						Name: "deezer",
						Type: 2,
					},
					Name:  "GoDm | https://github.com/yaboipy",
					Type:  4,
					State: "GoDm V" + fmt.Sprint(Xc.Config().Mode.Discord.Version),
					Emoji: "🌐",
				},
				Afk: false,
			},
			Compress: false,
			ClientState: WsClientState{
				GuildHashes:              WsGH{},
				HighestLastMessageID:     "0",
				ReadStateVersion:         0,
				UserGuildSettingsVersion: -1,
				UserSettingsVersion:      -1,
			},
		},
	})

	err = ws.WriteMessage(websocket.TextMessage, Payload)
	Xc.Errs(err)
	_, _, _ = ws.ReadMessage()
	_, _, _ = ws.ReadMessage()

	fmt.Println("" + clr + "▏" + r + "(" + clr + "o" + r + ") Connected to " + clr + "Websocket" + r + "")
	ws.Close()
}

func (Xc *Config) Logo() string {
	text := `
	____` + clr + `_____` + r + `__` + clr + `____     ` + r + `____` + clr + `____` + r + `____` + clr + `__  ___
	` + r + `__` + clr + `  ____/` + r + `_  ` + clr + `__ \    ` + r + `___` + clr + `  __ \` + r + `__` + clr + `   |/  /
	` + r + `_ ` + clr + ` / __ ` + r + `_` + clr + `  / / /    ` + r + `__` + clr + `  / / /` + r + `_` + clr + `  /|_/ / 
	` + clr + `/ /_/ / / /_/ /     ` + r + `_  ` + clr + `/_/ /` + r + `_` + clr + `  /  / /  
	\____/  \____/      /_____/ /_/  /_/   
    
	[` + r + `Proxy: ` + cfg.Con.ProxyMode + clr + `]   	[` + r + `Tokens: ` + cfg.Con.Solution + clr + `]	~` + r + `YABOI` + clr + `
	[` + r + `1` + clr + `]` + r + ` Mass DM ` + clr + `		[` + r + `10` + clr + `]` + r + ` Mass Ping ` + clr + `
	[` + r + `2` + clr + `]` + r + ` Dm Spam ` + clr + `		[` + r + `11` + clr + `]` + r + ` Button Click ` + clr + `
	[` + r + `3` + clr + `]` + r + ` React Verify ` + clr + `	[` + r + `12` + clr + `]` + r + ` Friender ` + clr + `
	[` + r + `4` + clr + `]` + r + ` Joiner ` + clr + `		[` + r + `13` + clr + `]` + r + `	x ` + clr + `
	[` + r + `5` + clr + `]` + r + ` Leaver ` + clr + `		[` + r + `14` + clr + `]` + r + `	x ` + clr + `
	[` + r + `6` + clr + `]` + r + ` Accept Rules ` + clr + `	[` + r + `15` + clr + `]` + r + `	x ` + clr + `
	[` + r + `7` + clr + `]` + r + ` Raid Channel ` + clr + `	[` + r + `16` + clr + `]` + r + `	x ` + clr + `
	[` + r + `8` + clr + `]` + r + ` Scrape Users ` + clr + `	[` + r + `17` + clr + `]` + r + `	x ` + clr + `
	[` + r + `9` + clr + `]` + r + ` Check Tokens ` + clr + `	[` + r + `18` + clr + `]` + r + `	x

	Choice ` + clr + `>>:` + r + ` `
	return text
}
