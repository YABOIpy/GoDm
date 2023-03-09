package Scraper

import (
	"github.com/gorilla/websocket"
	"time"
)

type Sock struct {
	Ws *websocket.Conn
}

type WsResp struct {
	Op        int    `json:"op"`
	Data      Data   `json:"d,omitempty"`
	Sequence  int    `json:"s,omitempty"`
	EventName string `json:"t,omitempty"`
}

type Data struct {
	Content           string                 `json:"content,omitempty"`
	Author            User                   `json:"author,omitempty"`
	GuildID           string                 `json:"guild_id,omitempty"`
	GuildId           interface{}            `json:"guild_id,omitempty"`
	MessageId         string                 `json:"id,omitempty"`
	Flags             int                    `json:"flags,omitempty"`
	Token             string                 `json:"token,omitempty"`
	Capabilities      int                    `json:"capabilities,omitempty"`
	Compress          bool                   `json:"compress,omitempty"`
	Since             int                    `json:"since,omitempty"`
	Status            string                 `json:"status"`
	Afk               bool                   `json:"afk"`
	HeartbeatInterval int                    `json:"heartbeat_interval,omitempty"`
	SessionID         string                 `json:"session_id,omitempty"`
	Channels          map[string]interface{} `json:"channels,omitempty"`
	Ops               []Ops                  `json:"ops,omitempty"`
	Members           []Member               `json:"members,omitempty"`
	Typing            bool                   `json:"typing,omitempty"`
	Threads           bool                   `json:"threads,omitempty"`
	Activities        bool                   `json:"activities,omitempty"`
	ThreadMemberLists interface{}            `json:"thread_member_lists,omitempty"`
	UserID            string                 `json:"user_id,omitempty"`
	MessageID         string                 `json:"message_id,omitempty"`
}

type Ops struct {
	Items []Userinfo  `json:"items,omitempty"`
	Range interface{} `json:"range,omitempty"`
	Op    string      `json:"op,omitempty"`
}

type Userinfo struct {
	Member Member `json:"member,omitempty"`
}

type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Bot           bool   `json:"bot"`
}

type Member struct {
	User                       User        `json:"user,omitempty"`
	Roles                      []string    `json:"roles"`
	PremiumSince               interface{} `json:"premium_since"`
	Pending                    bool        `json:"pending"`
	Nick                       string      `json:"nick"`
	Mute                       bool        `json:"mute"`
	JoinedAt                   time.Time   `json:"joined_at"`
	Flags                      int         `json:"flags"`
	Deaf                       bool        `json:"deaf"`
	CommunicationDisabledUntil interface{} `json:"communication_disabled_until"`
	Avatar                     interface{} `json:"avatar"`
}

type WsPayload struct {
}
