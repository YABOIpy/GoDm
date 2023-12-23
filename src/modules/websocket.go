package modules

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"time"
	"math/rand"
	shttp "net/http"
)

func (ws *Sock) Connect(Token string, in *Instance) (*WsResp, *Sock, error) {
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
		},
	)
	if err != nil {
		log.Println(err)
		if len(Token) != IntNil {
			return websock.Connect(Token, in)
		}
	}
	conn.ReadMessage()

	d, c, err := send(conn, Token, in)
	if err != nil {
		return nil, nil, err
	}

	return d, &Sock{Ws: c}, nil
}

func (ws *Sock) ScrapeUsers(GID, CID string, iter int) []Member {

	time.Sleep(time.Millisecond * 500)
	ws.GuildConnection(GID, CID, iter)

	_, b, _ := ws.Ws.ReadMessage()
	var data WsResp
	json.Unmarshal(b, &data)

	if data.Name == EventGuildMemberListUpdate {
		for i := 0; i < len(data.Data.Ops); i++ {
			if len(data.Data.Ops[i].Items) == IntNil && data.Data.Ops[i].Op == EventSync {
				ws.Break = true
			}
		}
		for i := 0; i < len(data.Data.Ops); i++ {
			if data.Data.Ops[i].Op == EventSync {
				for j := 0; j < len(data.Data.Ops[i].Items); j++ {
					ws.Members = append(ws.Members, data.Data.Ops[i].Items[j].Member)
				}
			}
		}
	}
	return ws.Members
}

func (ws *Sock) GuildConnection(GID, CID string, iter int) {
	if iter == IntNil {
		ws.GuildCon(GID, CID)
	}

	if err := ws.Ws.WriteJSON(WsResp{
		Op: 14,
		Data: WsData{
			GuildId: GID,
			Channels: map[string]interface{}{
				CID: CreateRange(iter),
			},
		},
	}); err != nil {
		log.Println(err)
	}
}

func (ws *Sock) GuildCon(GID, CID string) {
	if err := ws.Ws.WriteJSON(WsResp{
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
	}); err != nil {
		log.Println(err)
	}
}

func send(conn *websocket.Conn, Token string, in *Instance) (*WsResp, *websocket.Conn, error) {
	err := conn.WriteJSON(WsResp{
		Op: 2,
		Data: WsData{
			Token:        Token,
			Capabilities: capabilities,
			Compress:     false,
			Properties: XpropData{
				OS:                 in.BrowserClient.OS,
				Browser:            in.BrowserClient.Name,
				SystemLocale:       "en-US",
				BrowserUserAgent:   in.BrowserClient.Agent,
				OSVersion:          in.BrowserClient.OSVer,
				Referrer:           "https://www.google.com",
				ReferringDomain:    "www.google.com",
				ReferrerCurrent:    "",
				ReferringDomainCur: "",
				Device:             "",
				ReleaseChannel:     "stable",
				ClientBuildNumber:  cbn,
				ClientEventSource:  nil,
			},
			Presence: Presence{
				Status:     in.Cfg.Mode.Discord.Presence[rand.Intn(len(in.Cfg.Mode.Discord.Presence))],
				Activities: []string{},
				Afk:        false,
				Game: Game{
					Name: "GoDm - github.com/yaboipy/godm",
					Type: 0,
				},
			},
			ClientState: ClientState{
				GuildVersions:            struct{}{},
				HighestLastMessageID:     "0",
				PrivateChannelsVersion:   "0",
				ReadStateVersion:         0,
				UserGuildSettingsVersion: -1,
				UserSettingsVersion:      -1,
				ApiCodeVersion:           0,
			},
		},
	})
	if err != nil {
		log.Println(err)
	}

	_, b, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	var data WsResp
	if err = json.Unmarshal(b, &data); err != nil {
		log.Println(err)
	}

	return &data, conn, nil
}
