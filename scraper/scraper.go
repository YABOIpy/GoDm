package Scraper

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

func Subscribe(ws *Sock, guildid, Channel string) error {
	payload := Data{
		GuildId:    guildid,
		Typing:     true,
		Threads:    true,
		Activities: true,
		Members:    []Member{},
		Channels: map[string]interface{}{
			Channel: []interface{}{[2]int{0, 99}},
		},
	}

	err := ws.Ws.WriteJSON(
		map[string]interface{}{
			"Op": 14,
			"d":  payload,
		})
	if err != nil {
		return err
	}
	return nil
}

func (Ws *Sock) Chann(i int, GID string, CID string) []interface{} {
	var x []interface{}
	switch i {
	default:
		{
			x = []interface{}{[2]int{0, 99}, [2]int{100, 199}, [2]int{i * 100, (i * 100) + 99}}
		}
	case 0:
		{
			err := Subscribe(Ws, GID, CID)
			if err != nil {
				log.Fatal(err)
			}
			x = []interface{}{[2]int{0, 99}}
		}
	case 1:
		{
			x = []interface{}{[2]int{0, 99}, [2]int{100, 199}}
		}
	case 2:
		{
			x = []interface{}{[2]int{0, 99}, [2]int{100, 199}, [2]int{200, 299}}
		}
	}
	return x
}

func (Ws *Sock) Scrape(GID string, CID string, i int) {
	x := Ws.Chann(i, GID, CID)
	err := Ws.Ws.WriteJSON(map[string]interface{}{
		"op": 14,
		"d": map[string]interface{}{
			"guild_id":   GID,
			"typing":     true,
			"activities": true,
			"threads":    true,
			"channels": map[string]interface{}{
				CID: x,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	go Ws.ReadMsg()
}

func (Ws *Sock) ReadMsg() {
	for {
		_, body, err := Ws.Ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("WORKSS!")
		var data WsResp
		if err := json.Unmarshal(body, &data); err != nil {
			continue
		}
		if data.EventName == "GUILD_MEMBER_LIST_UPDATE" {
			for i := 0; i < len(data.Data.Ops); i++ {
				if len(data.Data.Ops[i].Items) == 0 && data.Data.Ops[i].Op == "SYNC" {
					fmt.Println("Scraped Users!")
				}
			}
			for i := 0; i < len(data.Data.Ops); i++ {
				if data.Data.Ops[i].Op == "SYNC" {
					for j := 0; j < len(data.Data.Ops[i].Items); j++ {
						data.Data.Members = append(data.Data.Members, data.Data.Ops[i].Items[j].Member)
					}
				}
			}
		}
		fmt.Println(data.Data.Members)
	}
}

func (Ws *Sock) Connect(Token string) *WsResp {
	var dailer websocket.Dialer
	ws, _, err := dailer.Dial("wss://gateway.discord.gg/?v=10&encoding=json", http.Header{
		"Accept-Encoding":          []string{"gzip, deflate, br"},
		"Accept-Language":          []string{"en-US,en;q=0.9"},
		"Cache-Control":            []string{"no-cache"},
		"Pragma":                   []string{"no-cache"},
		"Sec-WebSocket-Extensions": []string{"permessage-deflate; client_max_window_bits"},
		"User-Agent":               []string{"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9006 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36"},
	})
	if err != nil {
		log.Fatal(err)
	}
	Ws.Ws = ws
	interval, err := Ws.ReadHello()
	if err != nil {
		Ws.Ws.Close()
		log.Fatal(err)
	}

	err = Ws.Ws.WriteJSON(map[string]interface{}{
		"op": 2,
		"d": map[string]interface{}{
			"token":        Token,
			"capabilities": 125,
			"properties": map[string]interface{}{
				"os":                       "Windows",
				"browser":                  "Vivaldi",
				"system_locale":            "en-US",
				"browser_user_agent":       "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9006 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
				"browser_version":          "94.0",
				"os_version":               "10",
				"referrer":                 "",
				"referring_domain":         "",
				"referrer_current":         "",
				"referring_domain_current": "",
				"release_channel":          "stable",
				"client_build_number":      103981,
			},
			"presence": map[string]interface{}{
				"status":     "online",
				"since":      0,
				"activities": []string{},
				"afk":        false,
			},
			"compress": false,
			"client_state": map[string]interface{}{
				"highest_last_message_id":     "0",
				"read_state_version":          0,
				"user_guild_settings_version": -1,
				"user_settings_version":       -1,
			},
		},
	})
	var data WsResp
	if err != nil {
		log.Fatal(err)
	} else {
		_, b, err := Ws.Ws.ReadMessage()

		json.Unmarshal(b, &data)
		if err != nil {
			log.Fatal(err)

		}

		go Ws.Ping(time.Duration(interval) * time.Millisecond)

	}
	return &data
}

//function from V4nsh4j
func (Ws *Sock) ReadHello() (int, error) {
	_, message, err := Ws.Ws.ReadMessage()
	if err != nil {
		return 0, err
	}
	var body WsResp
	if err := json.Unmarshal(message, &body); err != nil {
		return 0, fmt.Errorf("error while Unmarshalling incoming hello websocket message: %v", err)
	}

	if body.Data.HeartbeatInterval <= 0 {
		return 0, fmt.Errorf("heartbeat interval is not valid")
	}

	return body.Data.HeartbeatInterval, nil

}

func (Ws *Sock) Ping(interval time.Duration) {
	go func() {
		t := time.NewTicker(interval)
		defer t.Stop()
		for {
			_ = Ws.Ws.WriteJSON(&WsResp{
				Op: 10,
			})
		}
	}()
}

func X() Sock {
	x := Sock{}
	return x
}
