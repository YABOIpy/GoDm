package modules

import (
	"github.com/Danny-Dasilva/fhttp"
	"github.com/gorilla/websocket"
	shttp "net/http"
	"strconv"
	"sync"
	"time"
)

const (
	DiscrimMention       = "\u001B[30;1m[!] Note: press enter if the user has no discrim\u001B[0m"
	WriteValidMention    = "\u001B[30;1m[!] Write Unlocked Tokens To Tokens.txt? y/n:\u001B[0m"
	WriteJoinedMention   = "\u001B[30;1m[!] Write Joined Tokens To Tokens.txt? y/n:\u001B[0m"
	InServerMention      = "\u001B[30;1m[!] Tokens Need to be in the Server\u001B[0m"
	TokenFormatMention   = "\u001B[30;1m[!] Tokens Need to be in Email:Pass:Token Format\u001B[0m"
	PasswordFieldMention = "\u001B[30;1m[!] No Input = Random Password\u001B[0m"
	MassDmMention        = "\u001B[30;1m[!] if Custom Message in config, Leave Message input empty\u001B[0m"
	BandWidthMention     = "\u001B[30;1m[!] Warning, Will use alot of proxy bandwidth.\u001B[0m"
	ImageFormatMention   = "\u001B[30;1m[!] PNG's inside data/pfp Must be 128x128\u001B[0m"
)

const (
	TokenOptions = "1: Change DisplayName \n2: Change Bio \n3: Change pfp \n4: Change Password \n5: Change Pronounce \n6: Change Username\n7: Change To Unique Username\nChoice"
	CacheLoading = "Loading Cache: \u001B[36m[\u001B[39m%s\u001B[36m]\u001B[39m | \033[36m[\u001B[39m%s\u001B[36m]\u001B[39m"
	Initializing = "\033[36mInitializing...\033[39m "
	CharSet      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	CheckerFormat = "[\u001B[32m✓\033[39m] (TIME): %sMs (\033[33mLOCKED\033[39m): %v (\033[31mINVALID\033[39m): %v (\033[32mVALID\033[39m): %v (TOTAL): %v \n"
	MassDmFormat  = "[\u001B[32m✓\033[39m] (TIME): %sMs (\033[33mCAPTCHA\033[39m): %v (\033[31mFAILED\033[39m): %v (\033[32mSENT\033[39m): %v (TOTAL DM'S): %v \n"
	TokenFormat   = "%s:%s:%s"
)

const (
	EventSync                       = "SYNC"
	EventReady                      = "READY"
	EventResumed                    = "RESUMED"
	EventReadySupplemental          = "READY_SUPPLEMENTAL"
	EventChannelCreate              = "CHANNEL_CREATE"
	EventChannelUpdate              = "CHANNEL_UPDATE"
	EventChannelDelete              = "CHANNEL_DELETE"
	EventChannelPinsUpdate          = "CHANNEL_PINS_UPDATE"
	EventGuildCreate                = "GUILD_CREATE"
	EventGuildUpdate                = "GUILD_UPDATE"
	EventGuildDelete                = "GUILD_DELETE"
	EventGuildBanAdd                = "GUILD_BAN_ADD"
	EventGuildBanRemove             = "GUILD_BAN_REMOVE"
	EventGuildEmojisUpdate          = "GUILD_EMOJIS_UPDATE"
	EventGuildIntegrationsUpdate    = "GUILD_INTEGRATIONS_UPDATE"
	EventGuildMemberAdd             = "GUILD_MEMBER_ADD"
	EventGuildMemberRemove          = "GUILD_MEMBER_REMOVE"
	EventGuildMemberUpdate          = "GUILD_MEMBER_UPDATE"
	EventGuildMembersChunk          = "GUILD_MEMBERS_CHUNK"
	EventGuildMemberListUpdate      = "GUILD_MEMBER_LIST_UPDATE"
	EventGuildRoleCreate            = "GUILD_ROLE_CREATE"
	EventGuildRoleUpdate            = "GUILD_ROLE_UPDATE"
	EventGuildRoleDelete            = "GUILD_ROLE_DELETE"
	EventMessageCreate              = "MESSAGE_CREATE"
	EventMessageUpdate              = "MESSAGE_UPDATE"
	EventMessageDelete              = "MESSAGE_DELETE"
	EventMessageDeleteBulk          = "MESSAGE_DELETE_BULK"
	EventMessageReactionAdd         = "MESSAGE_REACTION_ADD"
	EventMessageReactionRemove      = "MESSAGE_REACTION_REMOVE"
	EventMessageReactionRemoveAll   = "MESSAGE_REACTION_REMOVE_ALL"
	EventMessageReactionRemoveEmoji = "MESSAGE_REACTION_REMOVE_EMOJI"
	EventPresenceUpdate             = "PRESENCE_UPDATE"
	EventTypingStart                = "TYPING_START"
	EventUserUpdate                 = "USER_UPDATE"
	EventVoiceStateUpdate           = "VOICE_STATE_UPDATE"
	EventVoiceServerUpdate          = "VOICE_SERVER_UPDATE"
	EventWebhooksUpdate             = "WEBHOOKS_UPDATE"
	EventSessionReplace             = "SESSIONS_REPLACE"
)

var (
	modules = Modules{}
	websock = Sock{}
	RSeed   = RandSeed()
	cbn, _  = strconv.Atoi(BuildInfo())

	x, r, g, bg, rb, gr, u, clr, yellow, red, prp = "\u001b[30;1m", "\033[39m", "\033[32m", "\u001b[34;1m", "\u001b[0m", "\u001b[30;1m", "\u001b[4m", "\033[36m", "\033[33m", "\u001B[31m", "\033[35m"
)

type Header struct{}

type Modules struct {
	Dcfd       string
	Sdcfd      string
	Cfruid     string
	CookieData CookieData
	Checker    struct {
		Client  http.Client
		Invalid int
		Locked  int
		Valid   int
		All     int
	}
}

type CookieData struct {
	Cookies map[string]*shttp.Cookie
}

type CCSeed struct {
	mu   sync.Mutex
	seed int64
}

type Sock struct {
	Ws      *websocket.Conn
	Break   bool
	Members []Member
}
type WsResp struct {
	Op   int    `json:"op"`
	Data WsData `json:"d,omitempty"`
	Name string `json:"t,omitempty"`
}
type Ops struct {
	Items []Userinfo  `json:"items,omitempty"`
	Range interface{} `json:"range,omitempty"`
	Op    string      `json:"op,omitempty"`
}

type Userinfo struct {
	Member Member `json:"member,omitempty"`
}

type BoostPayload struct {
	UserPremiumGuildSubscriptionSlotIds []string `json:"user_premium_guild_subscription_slot_ids"`
}
type MeResp struct {
	Id               string        `json:"id"`
	Username         string        `json:"username"`
	GlobalName       interface{}   `json:"global_name"`
	Avatar           string        `json:"avatar"`
	Discriminator    string        `json:"discriminator"`
	PublicFlags      int           `json:"public_flags"`
	Flags            int           `json:"flags"`
	Banner           interface{}   `json:"banner"`
	BannerColor      interface{}   `json:"banner_color"`
	AccentColor      interface{}   `json:"accent_color"`
	Bio              string        `json:"bio"`
	Token            string        `json:"token"`
	Locale           string        `json:"locale"`
	NsfwAllowed      bool          `json:"nsfw_allowed"`
	MfaEnabled       bool          `json:"mfa_enabled"`
	PremiumType      int           `json:"premium_type"`
	LinkedUsers      []interface{} `json:"linked_users"`
	AvatarDecoration interface{}   `json:"avatar_decoration"`
	Email            string        `json:"email"`
	Verified         bool          `json:"verified"`
	Phone            interface{}   `json:"phone"`
	Retry            float64       `json:"retry_after,omitempty"`
}

type Option struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Emoji       Emoji    `json:"emoji"`
	RoleIds     []string `json:"role_ids"`
	ChannelIds  []string `json:"channel_ids"`
}

type Prompt struct {
	Id           string   `json:"id"`
	Title        string   `json:"title"`
	Options      []Option `json:"options"`
	SingleSelect bool     `json:"single_select"`
	Required     bool     `json:"required"`
	InOnboarding bool     `json:"in_onboarding"`
	Type         int      `json:"type"`
}

type Onboarding struct {
	GuildId                 string        `json:"guild_id"`
	Prompts                 []Prompt      `json:"prompts"`
	DefaultChannelIds       []string      `json:"default_channel_ids"`
	Enabled                 bool          `json:"enabled"`
	Mode                    int           `json:"mode"`
	BelowRequirements       bool          `json:"below_requirements"`
	Responses               []interface{} `json:"responses"`
	OnboardingPromptsSeen   struct{}      `json:"onboarding_prompts_seen"`
	OnboardingResponsesSeen struct{}      `json:"onboarding_responses_seen"`
}

type BoostResp struct {
	Id             string `json:"id"`
	Type           int    `json:"type"`
	Invalid        bool   `json:"invalid"`
	Flags          int    `json:"flags"`
	Brand          string `json:"brand"`
	Last4          string `json:"last_4"`
	ExpiresMonth   int    `json:"expires_month"`
	ExpiresYear    int    `json:"expires_year"`
	Country        string `json:"country"`
	PaymentGateway int    `json:"payment_gateway"`
	Default        bool   `json:"default"`
}

type MessageResp struct {
	ID              string            `json:"id"`
	Type            int               `json:"type"`
	Content         string            `json:"content"`
	ChannelID       string            `json:"channel_id"`
	Author          Author            `json:"author"`
	Attachments     []interface{}     `json:"attachments"`
	Embeds          []Embed           `json:"embeds"`
	Mentions        []interface{}     `json:"mentions"`
	MentionRoles    []interface{}     `json:"mention_roles"`
	Pinned          bool              `json:"pinned,omitempty"`
	MentionEveryone bool              `json:"mention_everyone,omitempty"`
	TTS             bool              `json:"tts"`
	Timestamp       time.Time         `json:"timestamp"`
	EditedTimestamp interface{}       `json:"edited_timestamp"`
	Flags           int               `json:"flags"`
	Reactions       []Reaction        `json:"reactions,omitempty"`
	Components      []NestedComponent `json:"components,omitempty"`
	SiteKey         string            `json:"captcha_sitekey,omitempty"`
	Retry           float64           `json:"retry_after,omitempty"`
	ApplicationID   string            `json:"application_id"`
	WebhookId       string            `json:"webhook_id"`
}

type Embed struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Color       int    `json:"color"`
	Author      struct {
		Name string `json:"name"`
	} `json:"author"`
	Footer struct {
		Text         string `json:"text"`
		IconUrl      string `json:"icon_url"`
		ProxyIconUrl string `json:"proxy_icon_url"`
	} `json:"footer"`
}

type NestedComponent struct {
	Type       int         `json:"type"`
	Components []Component `json:"components"`
}

type Component struct {
	Type     int    `json:"type"`
	CustomId string `json:"custom_id"`
	Style    int    `json:"style"`
	Label    string `json:"label"`
	Emoji    Emoji  `json:"emoji"`
}

type DmChannel struct {
	Id            string      `json:"id"`
	Type          int         `json:"type"`
	LastMessageId interface{} `json:"last_message_id"`
	Retry         float64     `json:"retry_after,omitempty"`
	Flags         int         `json:"flags"`
	Recipients    []struct {
		Id               string      `json:"id"`
		Username         string      `json:"username"`
		Avatar           string      `json:"avatar"`
		Discriminator    string      `json:"discriminator"`
		PublicFlags      int         `json:"public_flags"`
		Flags            int         `json:"flags"`
		Banner           interface{} `json:"banner"`
		AccentColor      interface{} `json:"accent_color"`
		GlobalName       interface{} `json:"global_name"`
		AvatarDecoration interface{} `json:"avatar_decoration"`
		DisplayName      interface{} `json:"display_name"`
		BannerColor      interface{} `json:"banner_color"`
	} `json:"recipients"`
}

type Author struct {
	ID               string      `json:"id"`
	Username         string      `json:"username"`
	GlobalName       string      `json:"global_name"`
	Avatar           string      `json:"avatar"`
	Discriminator    string      `json:"discriminator"`
	PublicFlags      int         `json:"public_flags"`
	AvatarDecoration interface{} `json:"avatar_decoration"`
}

type Reaction struct {
	Emoji        Emoji         `json:"emoji"`
	Count        int           `json:"count"`
	CountDetails CountDetail   `json:"count_details"`
	BurstColors  []interface{} `json:"burst_colors"`
	MeBurst      bool          `json:"me_burst"`
	Me           bool          `json:"me"`
	BurstCount   int           `json:"burst_count"`
}

type Emoji struct {
	ID       *string `json:"id"`
	Name     string  `json:"name"`
	Animated bool    `json:"animated"`
}

type CountDetail struct {
	Burst  int `json:"burst"`
	Normal int `json:"normal"`
}

type Agents struct {
	Mac     string
	Windows string
	Linux   string
}

type Instance struct {
	BrowserClient ClientData
	TokenProps    TokenConfig
	Client        *http.Client
	SClient       *http.Client
	Cookie        string
	Xprop         string
	Token         string
	Cfg           Config
}

type BrowserData struct {
	Name           string
	OS             []string `json:"os"`
	OSver          map[string][]string
	Versions       []string
	CipherSuites   []uint16
	UserAgent      map[string]string
	CipherList     map[int][]uint16
	EllipticCurves map[int][]uint16
}

type ClientData struct {
	OS      string
	OSVer   string
	Name    string
	Agent   string
	Version string
}

type TokenConfig struct {
	RateLimit float64
	Email     string
	Pass      string
}

type JoinResp struct {
	SiteKey string  `json:"captcha_sitekey,omitempty"`
	RqToken string  `json:"captcha_rqtoken,omitempty"`
	Retry   float64 `json:"retry_after,omitempty"`
	Message string  `json:"message,omitempty"`
	Code    int     `json:"code"`
}

type ChannelID struct {
	ID string `json:"id,omitempty"`
}

type FriendReq struct {
	Username string
	Discrim  interface{}
}

type JoinReq struct {
	Code      string      `json:"code"`
	Type      int         `json:"type"`
	ExpiresAt interface{} `json:"expires_at"`
	Guild     struct {
		Id                       string      `json:"id"`
		Name                     string      `json:"name"`
		Splash                   string      `json:"splash"`
		Banner                   string      `json:"banner"`
		Description              interface{} `json:"description"`
		Icon                     string      `json:"icon"`
		Features                 []string    `json:"features"`
		VerificationLevel        int         `json:"verification_level"`
		VanityUrlCode            string      `json:"vanity_url_code"`
		NsfwLevel                int         `json:"nsfw_level"`
		Nsfw                     bool        `json:"nsfw"`
		PremiumSubscriptionCount int         `json:"premium_subscription_count"`
	} `json:"guild"`
	Channel struct {
		Id   string `json:"id"`
		Type int    `json:"type"`
		Name string `json:"name"`
	} `json:"channel"`
	ApproximateMemberCount   int `json:"approximate_member_count"`
	ApproximatePresenceCount int `json:"approximate_presence_count"`
}

type JoinContext struct {
	Location            string `json:"location"`
	LocationGuildId     string `json:"location_guild_id"`
	LocationChannelId   string `json:"location_channel_id"`
	LocationChannelType int    `json:"location_channel_type"`
}

type XpropData struct {
	OS                 string      `json:"os"`
	Browser            string      `json:"browser"`
	Device             string      `json:"device"`
	SystemLocale       string      `json:"system_locale"`
	BrowserUserAgent   string      `json:"browser_user_agent"`
	BrowserVersion     string      `json:"browser_version"`
	OSVersion          string      `json:"os_version"`
	Referrer           string      `json:"referrer"`
	ReferringDomain    string      `json:"referring_domain"`
	SearchEngine       string      `json:"search_engine"`
	ReferrerCurrent    string      `json:"referrer_current"`
	ReferringDomainCur string      `json:"referring_domain_current"`
	ReleaseChannel     string      `json:"release_channel"`
	ClientBuildNumber  int         `json:"client_build_number"`
	ClientEventSource  interface{} `json:"client_event_source"`
}

type WsData struct {
	Message struct {
		Content   string `json:"content,omitempty"`
		ChannelID string `json:"channel_id,omitempty"`
		GuildID   string `json:"guild_id,omitempty"`
		MessageId string `json:"id,omitempty"`
		Flags     int    `json:"flags,omitempty"`
	}
	Identify struct {
		Token        string    `json:"token,omitempty"`
		Properties   XpropData `json:"properties,omitempty"`
		Capabilities int       `json:"capabilities,omitempty"`
		Compress     bool      `json:"compress,omitempty"`
	}
	ClientState struct {
		HighestLastMessageID     string `json:"highest_last_message_id,omitempty"`
		ReadStateVersion         int    `json:"read_state_version,omitempty"`
		UserGuildSettingsVersion int    `json:"user_guild_settings_version,omitempty"`
	} `json:"client_state,omitempty"`
	Ops               []Ops                  `json:"ops,omitempty"`
	Members           []Member               `json:"members,omitempty"`
	HeartbeatInterval int                    `json:"heartbeat_interval,omitempty"`
	SessionID         string                 `json:"session_id,omitempty"`
	GuildId           interface{}            `json:"guild_id,omitempty"`
	Channels          map[string]interface{} `json:"channels,omitempty"`
	ChannelID         string                 `json:"channel_id,omitempty"`
	Typing            bool                   `json:"typing,omitempty"`
	Threads           bool                   `json:"threads,omitempty"`
	Activities        bool                   `json:"activities,omitempty"`
	ThreadMemberLists interface{}            `json:"thread_member_lists,omitempty"`
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

type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Bot           bool   `json:"bot"`
}

type Guild struct {
	Version     time.Time   `json:"version"`
	FormFields  []GuildForm `json:"form_fields"`
	Description string      `json:"description"`
}

type GuildForm struct {
	FieldType   string      `json:"field_type"`
	Label       string      `json:"label"`
	Description interface{} `json:"description"`
	Automations interface{} `json:"automations"`
	Required    bool        `json:"required"`
	Values      []string    `json:"values"`
	Response    bool        `json:"response,omitempty"`
}

type VcOptions struct {
	CID  string
	GID  string
	Mute bool
	Deaf bool
}

type ButtonOptions struct {
	Button    Component
	SessionID string
	Type      int
	GuildID   string
}

type Button struct {
	AppID   string     `json:"application_id"`
	CID     string     `json:"channel_id"`
	GID     string     `json:"guild_id"`
	MID     string     `json:"message_id"`
	Flags   int        `json:"message_flags"`
	Type    int        `json:"type"`
	Nonce   string     `json:"nonce"`
	Session string     `json:"session_id"`
	Data    ButtonData `json:"data"`
}

type ButtonData struct {
	Type int    `json:"component_type"`
	ID   string `json:"custom_id"`
}

type MessageOptions struct {
	Captcha string
	Mping   bool
	Loop    bool
	IDs     []string
	Amount  int
}

type CapCfg struct {
	SiteKey string
	ApiKey  string
	PageURL string
}

type Captcha2api struct {
	Status  int    `json:"status"`
	Request string `json:"request"`
}

type Message struct {
	Title string `json:"Title"`
	Body  string `json:"Body"`
	Link  string `json:"Link"`
}

type Config struct {
	Dcfd   string
	Sdcfd  string
	Cfruid string

	Cookie CookieData
	Con    struct {
		ProxyMode string
		Errors    bool
	}

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
			WebKit   string `json:"WebKit"`
		} `json:"Net"`
		Discord struct {
			Version  float64   `json:"Ver"`
			CapAPI   []string  `json:"CapAPI"`
			Message  []Message `json:"Message"`
			Presence []string  `json:"Presence"`
		} `json:"Discord"`
		Configs struct {
			CCManager    bool          `json:"CCManager"`
			MaxRoutines  int           `json:"MaxRoutines"`
			Solver       bool          `json:"SolveCaptcha"`
			Interval     time.Duration `json:"Interval"`
			RateLimit    bool          `json:"RateLimit"`
			WSProxy      bool          `json:"WSProxy"`
			CaptchaRetry int           `json:"CaptchaRetry"`
		} `json:"Config"`
	} `json:"Modes"`
}
