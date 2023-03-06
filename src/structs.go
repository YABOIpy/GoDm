package massdm

import (
	http "github.com/Danny-Dasilva/fhttp"
	"github.com/gorilla/websocket"
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
	Ws      *Sock
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

	Mode struct {
		Network struct {
			Redirect bool   `json:"Redirect"`
			TimeOut  int    `json:"TimeOut"`
			Ja3      string `json:"JA3"`
			Proxy    string `json:"Proxy"`
			Agent    string `json:"Agent"`
		} `json:"Net"`
		Discord struct {
			CfbM    bool   `json:"CfBm"`
			Version int    `json:"Ver"`
			Status  string `json:"Status"`
			AppID   string `json:"AppID"`
			RPC     bool   `json:"Presence"`
		} `json:"Discord"`
	} `json:"Modes"`

	Con struct {
		Solution  string
		Tokenclr  string
		ProxyMode string
		Toks      int
	}
}

type Sock struct {
	Members       []Member
	Token         string
	OfflineScrape chan []byte
	AllMembers    []string
	Messages      chan []byte
	Complete      bool
	Conn          *websocket.Conn
	sessionID     string
	in            chan string
	out           chan []byte
	fatalHandler  func(err error)
	seq           int
	closeChan     chan struct{}
	Reactions     chan []byte
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
	Op int `json:"op"`
	D  WsD `json:"d"`
}

type WsD struct {
	Token        string        `json:"token"`
	Capabilities int           `json:"capabilities"`
	Properties   XProperties   `json:"properties"`
	Presence     WsPresence    `json:"presence"`
	Compress     bool          `json:"compress"`
	ClientState  WsClientState `json:"client_state"`
}

type WsPresence struct {
	Status     string       `json:"status,omitempty"`
	Since      int          `json:"since,omitempty"`
	Activities PresenceData `json:"activities,omitempty"`
	Afk        bool         `json:"afk,omitempty"`
}

type PresenceData struct {
	Name  string `json:"name,omitempty"`
	Type  int    `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Emoji string `json:"emoji,omitempty"`
}

type Event struct {
	Op        int    `json:"op"`
	Data      Data   `json:"d,omitempty"`
	Sequence  int    `json:"s,omitempty"`
	EventName string `json:"t,omitempty"`
}

type Activity struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

type PresenceChange struct {
	Since      int        `json:"since,omitempty"`
	Activities []Activity `json:"activities"`
	Status     string     `json:"status"`
	Afk        bool       `json:"afk"`
}

type ClientState struct {
	HighestLastMessageID     string `json:"highest_last_message_id,omitempty"`
	ReadStateVersion         int    `json:"read_state_version,omitempty"`
	UserGuildSettingsVersion int    `json:"user_guild_settings_version,omitempty"`
}
type Message struct {
	Content   string `json:"content,omitempty"`
	ChannelID string `json:"channel_id,omitempty"`
	Author    User   `json:"author,omitempty"`
	GuildID   string `json:"guild_id,omitempty"`
	MessageId string `json:"id,omitempty"`
	Flags     int    `json:"flags,omitempty"`
}

type Data struct {
	Message
	Identify
	PresenceChange
	ClientState       ClientState            `json:"client_state,omitempty"`
	HeartbeatInterval int                    `json:"heartbeat_interval,omitempty"`
	SessionID         string                 `json:"session_id,omitempty"`
	GuildId           interface{}            `json:"guild_id,omitempty"`
	Channels          map[string]interface{} `json:"channels,omitempty"`
	Ops               []Ops                  `json:"ops,omitempty"`
	ChannelID         string                 `json:"channel_id,omitempty"`
	Members           []Member               `json:"members,omitempty"`
	Typing            bool                   `json:"typing,omitempty"`
	Threads           bool                   `json:"threads,omitempty"`
	Activities        bool                   `json:"activities,omitempty"`
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

type Identify struct {
	Token        string      `json:"token,omitempty"`
	Properties   XProperties `json:"properties,omitempty"`
	Capabilities int         `json:"capabilities,omitempty"`
	Compress     bool        `json:"compress,omitempty"`
	Presence     Presence    `json:"presence,omitempty"`
}

type Presence struct {
	Status     string   `json:"status,omitempty"`
	Since      int      `json:"since,omitempty"`
	Activities []string `json:"activities,omitempty"`
	AFK        bool     `json:"afk,omitempty"`
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
	c       = X()
	cfg     = Config{}
	Cookies = c.GetCookie()
	urls    = "https://discord.com/api/v9/users/@me/affinities/guilds"
	grn     = "\033[32m"
	yel     = "\033[33m"
	red     = "\033[31m"
	clr     = "\033[36m"
	r       = "\033[39m"
)
