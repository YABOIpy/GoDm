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
	Proxy         string        `json:"Proxy"`
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
	c           = X()
	prox        = c.Config().Proxy
	cfg         = Config{}
	Client, err = goclient.NewClient(goclient.Browser{JA3: "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0", UserAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9006 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36", Cookies: nil}, cfg.ProxySettings.Timeout, false, "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9006 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36", prox)
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
