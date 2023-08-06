package modules

import (
	crand "crypto/rand"
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	shttp "net/http"
	"time"
)

func (ws *Sock) Connect(Token string, in Instance) (*WsResp, []byte, *Sock) {
	var dialer websocket.Dialer

	conn, _, err := dialer.Dial(
		"wss://gateway.discord.gg/?v=9&encoding=json", shttp.Header{
			"Accept-Language":       {"en-US"},
			"Cache-Control":         {"no-cache"},
			"Host":                  {"gateway.discord.gg"},
			"Origin":                {"https://discord.com"},
			"Pragma":                {"no-cache"},
			"Sec-WebSocket-Version": {"13"},
			"User-Agent":            {in.BrowserClient.Agent},
			//"Sec-WebSocket-Key":     {ws.WsKey()},

		},
	)
	if err != nil {
		log.Println(err)
		return nil, nil, nil
	}

	ws.Ws = conn
	d, b, c := ws.Send(conn, Token, in)

	return d, b, &Sock{Ws: c}
}

func (c *Sock) ScrapeUsers(GID string, CID string, iter int) []Member {

	time.Sleep(1 * time.Second)
	c.GuildConnection(c.Ws, GID, CID, iter)

	_, b, _ := c.Ws.ReadMessage()
	var data WsResp
	json.Unmarshal(b, &data)

	if data.Name == EventGuildMemberListUpdate {
		for i := 0; i < len(data.Data.Ops); i++ {
			if len(data.Data.Ops[i].Items) == 0 && data.Data.Ops[i].Op == EventSync {
				c.Break = true
			}
		}
		for i := 0; i < len(data.Data.Ops); i++ {
			if data.Data.Ops[i].Op == EventSync {
				for j := 0; j < len(data.Data.Ops[i].Items); j++ {
					c.Members = append(c.Members, data.Data.Ops[i].Items[j].Member)
				}
			}
		}
	}

	return c.Members
}

func (s *Sock) GuildConnection(Ws *websocket.Conn, GID string, CID string, iter int) {
	if iter == 0 {
		s.GuildCon(Ws, GID, CID)
	}

	var v []interface{}
	switch iter {
	case 0:
		v = []interface{}{[2]int{0, 99}}
	case 1:
		v = []interface{}{[2]int{0, 99}, [2]int{100, 199}}
	case 2:
		v = []interface{}{[2]int{0, 99}, [2]int{100, 199}, [2]int{200, 299}}
	default:
		v = []interface{}{[2]int{0, 99}, [2]int{100, 199}, [2]int{iter * 100, (iter * 100) + 99}}
	}

	Ws.WriteJSON(WsResp{
		Op: 14,
		Data: WsData{
			GuildId: GID,
			Channels: map[string]interface{}{
				CID: v,
			},
		},
	})
}

func (s *Sock) GuildCon(Ws *websocket.Conn, GID string, CID string) {
	Ws.WriteJSON(WsResp{
		Op: 14,
		Data: WsData{
			GuildId:    GID,
			Typing:     true,
			Threads:    true,
			Activities: true,
			Members:    []Member{},
			Channels: map[string]interface{}{
				CID: []interface{}{[2]int{0, 99}},
			},
		},
	})
}

func (ws *Sock) WsKey() string {
	B := make([]byte, 16)
	_, err := crand.Read(B)
	if err != nil {
		log.Println(err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(B)
}

func (c *Sock) Send(conn *websocket.Conn, Token string, in Instance) (*WsResp, []byte, *websocket.Conn) {
	_, d, _ := conn.ReadMessage()
	var resp WsResp
	json.Unmarshal(d, &resp)
	conn.WriteJSON(map[string]interface{}{
		"op": 2,
		"d": map[string]interface{}{
			"token":        Token,
			"capabilities": 125,
			"properties": map[string]interface{}{
				"os":                       in.BrowserClient.OS,
				"browser":                  in.BrowserClient.Name,
				"system_locale":            "en-US",
				"browser_user_agent":       in.BrowserClient.Agent,
				"browser_version":          in.BrowserClient.Version,
				"os_version":               in.BrowserClient.OSVer,
				"referrer":                 "https://www.google.com",
				"referring_domain":         "www.google.com",
				"referrer_current":         "",
				"referring_domain_current": "",
				"release_channel":          "stable",
				"client_build_number":      cbn,
			},
			"presence": map[string]interface{}{
				"status": in.Cfg.Mode.Discord.Presence[rand.Intn(len(in.Cfg.Mode.Discord.Presence))],
				"game": map[string]interface{}{
					"name": "GoDm - github.com/yaboipy/godm",
					"type": 0,
				},
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

	_, b, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return nil, nil, nil //return c.Send(conn, Token, in)
	}
	var data WsResp
	json.Unmarshal(b, &data)
	data.Data.HeartbeatInterval = resp.Data.HeartbeatInterval

	return &data, b, conn
}
