package massdm

import (
	"encoding/json"
	"unicode/utf8"
	"io/ioutil"
	"os/exec"
	"bytes"
	"fmt"
	"os"
	"net/http"
	"bufio"
	"log"
	"compress/zlib"
	"github.com/andybalholm/brotli"
	"github.com/gorilla/websocket"
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
	resp, err := Client.Do(req)
	Xc.Errs(err)
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

func (Xc * Config) Marsh(payload map[string]string) []byte {
	x,_ := json.Marshal(payload)
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
	ws, _, err := dialer.Dial("wss://gateway.discord.gg/?encoding=json&v=9&compress=zlib-stream", http.Header{
		"Origin": []string{"https://discord.com"},
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
	fmt.Println(""+clr+"▏"+r+"("+clr+"o"+r+") Connected to "+clr+"Websocket"+r+"")
	ws.Close()
}


