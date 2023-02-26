package massdm

import (
	http "github.com/Danny-Dasilva/fhttp"
	"github.com/gorilla/websocket"
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
type Instance struct {
	Token string
	Ws    *Socket
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
	Ja3           string        `json:"JA3"`
	Proxy         string        `json:"proxy"`
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
	Sequence          int                    `json:"seq,omitempty"` // For sending only
	GuildId           interface{}            `json:"guild_id,omitempty"`
	Channels          map[string]interface{} `json:"channels,omitempty"`
	Ops               []Ops                  `json:"ops,omitempty"`
	ChannelID         string                 `json:"channel_id,omitempty"`
	Members           []Member               `json:"members,omitempty"`
	Typing            bool                   `json:"typing,omitempty"`
	Threads           bool                   `json:"threads,omitempty"`
	Activities        bool                   `json:"activities,omitempty"`
	ThreadMemberLists interface{}            `json:"thread_member_lists,omitempty"`
	// Emoji React
	UserID    string `json:"user_id,omitempty"`
	MessageID string `json:"message_id,omitempty"`
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

type Socket struct {
	Members      []Member
	Token        string
	AllMembers   []string
	Messages     chan []byte
	Conn         *websocket.Conn
	sessionID    string
	in           chan string
	out          chan []byte
	fatalHandler func(err error)
	seq          int
	closeChan    chan struct{}
	Reactions    chan []byte
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
	c         = X()
	proxi     = c.Config().Proxy
	cfg       = Config{}
	Client, _ = goclient.NewClient(goclient.Browser{
		JA3:       c.Config().Ja3,
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9006 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
		Cookies:   nil,
	},
		cfg.ProxySettings.Timeout,
		false,
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9008 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
		proxi,
	)
	Cookies = "__dcfduid=" + cookies().Dcfd + "; " + "__sdcfduid=" + cookies().Sdcfd + "; "
	urls    = "https://discord.com/api/v9/users/@me/affinities/guilds"
	grn     = "\033[32m"
	yel     = "\033[33m"
	red     = "\033[31m"
	clr     = "\033[36m"
	r       = "\033[39m"

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

func Z() Socket {
	x := Socket{}
	return x
}
