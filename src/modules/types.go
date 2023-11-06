package modules

import (
	http "github.com/Danny-Dasilva/fhttp"
	"github.com/gorilla/websocket"
	shttp "net/http"
	"strconv"
	"sync"
	"time"
)

const (
	RGBMention              = "\u001B[30;1m[!] Note: Write 8-bit color (input: 255,255,255)\u001B[0m"
	DiscrimMention          = "\u001B[30;1m[!] Note: press enter if the user has no discrim\u001B[0m"
	WriteValidMention       = "\u001B[30;1m[!] Write Unlocked Tokens To Tokens.txt?\u001B[0m"
	WriteJoinedMention      = "\u001B[30;1m[!] Write Joined Tokens To Tokens.txt?\u001B[0m"
	WriteInServerMention    = "\u001B[30;1m[!] Write In Server Tokens To Tokens.txt?\u001B[0m"
	InServerMention         = "\u001B[30;1m[!] Tokens Need to be in the Server\u001B[0m"
	TokenFormatMention      = "\u001B[30;1m[!] Tokens Need to be in Email:Pass:Token Format\u001B[0m"
	PasswordFieldMention    = "\u001B[30;1m[!] No Input = Random Password\u001B[0m"
	MassDmMention           = "\u001B[30;1m[!] if Custom Message in config, Leave Message input empty\u001B[0m"
	BandWidthMention        = "\u001B[30;1m[!] Warning, Will use alot of proxy bandwidth.\u001B[0m"
	ImageFormatMention      = "\u001B[30;1m[!] PNG's inside data/pfp Must be 128x128\u001B[0m"
	MassFriendOptionMention = "\u001B[30;1m[!] Leave User Inputs Empty for no Changes\u001B[0m"
)

const (
	TokenOptions = "1: Change DisplayName \n2: Change Bio \n3: Change pfp \n4: Change Password \n5: Change Pronounce \n6: Change Username \n7: Change To Unique Username \n8: Change Banner \n9: Token Cleaner \nChoice"
	CacheLoading = "Loading Cache: \u001B[36m[\u001B[39m%s\u001B[36m]\u001B[39m | \033[36m[\u001B[39m%s\u001B[36m]\u001B[39m"
	Initializing = "\033[36mInitializing...\033[39m "
	CharSet      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	CheckerFormat    = "[\u001B[32m✓\033[39m] (TIME): %sMs (\033[33mLOCKED\033[39m): %v (\033[31mINVALID\033[39m): %v (\033[32mVALID\033[39m): %v (TOTAL): %v \n"
	MassDmFormat     = "[\u001B[32m✓\033[39m] (TIME): %sMs (\033[33mCAPTCHA\033[39m): %v (\033[31mFAILED\033[39m): %v (\033[32mSENT\033[39m): %v (TOTAL DM'S): %v \n"
	MassFriendFormat = "[\u001B[32m✓\033[39m] (TIME): %sMs (\033[33mCAPTCHA\033[39m): %v (\033[31mFAILED\033[39m): %v (\033[32mSENT\033[39m): %v (TOTAL REQUESTS): %v \n"
	TokenFormat      = "%s:%s:%s"

	GuildInput   = "Guild ID: "
	InviteInput  = "discord.gg/"
	ChannelInput = "Channel ID: "
	MessageInput = "Message: "

	StringNil = ""
	IntNil    = 0
)

const (
	DiscordBuildAsset = "cec3c372f71b56bc3d44.js"
	DiscordDataAsset  = "675ae39865fd3f811445.js"
)

const (
	GuildOnboardingEverEnabled    = "GUILD_ONBOARDING_EVER_ENABLED"
	MemberVerificationGateEnabled = "MEMBER_VERIFICATION_GATE_ENABLED"
	GuildWebPageVanityURL         = "GUILD_WEB_PAGE_VANITY_URL"
	GuildServerGuide              = "GUILD_SERVER_GUIDE"
	ChannelIconEmojisGenerated    = "CHANNEL_ICON_EMOJIS_GENERATED"
	GuildOnboarding               = "GUILD_ONBOARDING"
	News                          = "NEWS"
	PreviewEnabled                = "PREVIEW_ENABLED"
	Community                     = "COMMUNITY"
	SevenDayThreadArchive         = "SEVEN_DAY_THREAD_ARCHIVE"
	AnimatedBanner                = "ANIMATED_BANNER"
	PrivateThreads                = "PRIVATE_THREADS"
	AutoModeration                = "AUTO_MODERATION"
	InviteSplash                  = "INVITE_SPLASH"
	GuildOnboardingHasPrompts     = "GUILD_ONBOARDING_HAS_PROMPTS"
	Discoverable                  = "DISCOVERABLE"
	Soundboard                    = "SOUNDBOARD"
	ThreeDayThreadArchive         = "THREE_DAY_THREAD_ARCHIVE"
	VanityURL                     = "VANITY_URL"
	EnabledDiscoverableBefore     = "ENABLED_DISCOVERABLE_BEFORE"
	RoleIcons                     = "ROLE_ICONS"
	AnimatedIcon                  = "ANIMATED_ICON"
	MemberProfiles                = "MEMBER_PROFILES"
	Banner                        = "BANNER"
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
	modules      = Modules{}
	websock      = Sock{}
	RSeed        = RandSeed()
	cbn, _       = strconv.Atoi(BuildInfo())
	capabilities = Capabilities()
	fetchClient  = &http.Client{}
	// dont mind this.. i just keep copying the color codes off old projects :p
	x, r, g, bg, rb, gr, u, clr, yellow, red, prp = "\u001b[30;1m", "\033[39m", "\033[32m", "\u001b[34;1m", "\u001b[0m", "\u001b[30;1m", "\u001b[4m", "\033[36m", "\033[33m", "\u001B[31m", "\033[35m"
)

type Header struct{}

type Modules struct {
	Dcfd       string
	Sdcfd      string
	Cfuvid     string
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

type OnBoardPayload struct {
	Responses   []string         `json:"onboarding_responses"`
	PromptsSeen map[string]int64 `json:"onboarding_prompts_seen"`
}

type Onboarding struct {
	GuildId                 string           `json:"guild_id"`
	Prompts                 []Prompt         `json:"prompts"`
	DefaultChannelIds       []string         `json:"default_channel_ids"`
	Enabled                 bool             `json:"enabled"`
	Mode                    int              `json:"mode"`
	BelowRequirements       bool             `json:"below_requirements"`
	Responses               []string         `json:"responses"`
	OnboardingPromptsSeen   map[string]int64 `json:"onboarding_prompts_seen"`
	OnboardingResponsesSeen map[string]int64 `json:"onboarding_responses_seen"`
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
type ChannelMessages struct {
	Id        string `json:"id"`
	Type      int    `json:"type"`
	Content   string `json:"content"`
	ChannelId string `json:"channel_id"`
	Author    struct {
		Id                   string      `json:"id"`
		Username             string      `json:"username"`
		Avatar               *string     `json:"avatar"`
		Discriminator        string      `json:"discriminator"`
		PublicFlags          int         `json:"public_flags"`
		Flags                int         `json:"flags"`
		Banner               interface{} `json:"banner"`
		AccentColor          interface{} `json:"accent_color"`
		GlobalName           string      `json:"global_name"`
		AvatarDecorationData interface{} `json:"avatar_decoration_data"`
		BannerColor          interface{} `json:"banner_color"`
	} `json:"author"`
	Attachments []struct {
		Id                 string `json:"id"`
		Filename           string `json:"filename"`
		Size               int    `json:"size"`
		Url                string `json:"url"`
		ProxyUrl           string `json:"proxy_url"`
		Width              int    `json:"width"`
		Height             int    `json:"height"`
		ContentType        string `json:"content_type"`
		Placeholder        string `json:"placeholder"`
		PlaceholderVersion int    `json:"placeholder_version"`
	} `json:"attachments"`
	Embeds []struct {
		Type     string `json:"type"`
		Url      string `json:"url"`
		Provider struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"provider,omitempty"`
		Thumbnail struct {
			Url      string `json:"url"`
			ProxyUrl string `json:"proxy_url"`
			Width    int    `json:"width"`
			Height   int    `json:"height"`
		} `json:"thumbnail"`
		Video struct {
			Url      string `json:"url"`
			ProxyUrl string `json:"proxy_url"`
			Width    int    `json:"width"`
			Height   int    `json:"height"`
		} `json:"video,omitempty"`
	} `json:"embeds"`
	Mentions []struct {
		Id                   string      `json:"id"`
		Username             string      `json:"username"`
		Avatar               *string     `json:"avatar"`
		Discriminator        string      `json:"discriminator"`
		PublicFlags          int         `json:"public_flags"`
		Flags                int         `json:"flags"`
		Banner               interface{} `json:"banner"`
		AccentColor          interface{} `json:"accent_color"`
		GlobalName           string      `json:"global_name"`
		AvatarDecorationData interface{} `json:"avatar_decoration_data"`
		BannerColor          interface{} `json:"banner_color"`
	} `json:"mentions"`
	MentionRoles     []interface{} `json:"mention_roles"`
	Pinned           bool          `json:"pinned"`
	MentionEveryone  bool          `json:"mention_everyone"`
	Tts              bool          `json:"tts"`
	Timestamp        time.Time     `json:"timestamp"`
	EditedTimestamp  interface{}   `json:"edited_timestamp"`
	Flags            int           `json:"flags"`
	Components       []interface{} `json:"components"`
	MessageReference struct {
		ChannelId string `json:"channel_id"`
		MessageId string `json:"message_id"`
		GuildId   string `json:"guild_id"`
	} `json:"message_reference,omitempty"`
	ReferencedMessage *struct {
		Id        string `json:"id"`
		Type      int    `json:"type"`
		Content   string `json:"content"`
		ChannelId string `json:"channel_id"`
		Author    struct {
			Id                   string      `json:"id"`
			Username             string      `json:"username"`
			Avatar               string      `json:"avatar"`
			Discriminator        string      `json:"discriminator"`
			PublicFlags          int         `json:"public_flags"`
			Flags                int         `json:"flags"`
			Banner               interface{} `json:"banner"`
			AccentColor          interface{} `json:"accent_color"`
			GlobalName           string      `json:"global_name"`
			AvatarDecorationData interface{} `json:"avatar_decoration_data"`
			BannerColor          interface{} `json:"banner_color"`
		} `json:"author"`
		Attachments []struct {
			Id                 string `json:"id"`
			Filename           string `json:"filename"`
			Size               int    `json:"size"`
			Url                string `json:"url"`
			ProxyUrl           string `json:"proxy_url"`
			Width              int    `json:"width"`
			Height             int    `json:"height"`
			ContentType        string `json:"content_type"`
			Placeholder        string `json:"placeholder"`
			PlaceholderVersion int    `json:"placeholder_version"`
		} `json:"attachments"`
		Embeds   []interface{} `json:"embeds"`
		Mentions []struct {
			Id                   string      `json:"id"`
			Username             string      `json:"username"`
			Avatar               string      `json:"avatar"`
			Discriminator        string      `json:"discriminator"`
			PublicFlags          int         `json:"public_flags"`
			Flags                int         `json:"flags"`
			Banner               interface{} `json:"banner"`
			AccentColor          interface{} `json:"accent_color"`
			GlobalName           string      `json:"global_name"`
			AvatarDecorationData interface{} `json:"avatar_decoration_data"`
			BannerColor          interface{} `json:"banner_color"`
		} `json:"mentions"`
		MentionRoles     []interface{} `json:"mention_roles"`
		Pinned           bool          `json:"pinned"`
		MentionEveryone  bool          `json:"mention_everyone"`
		Tts              bool          `json:"tts"`
		Timestamp        time.Time     `json:"timestamp"`
		EditedTimestamp  interface{}   `json:"edited_timestamp"`
		Flags            int           `json:"flags"`
		Components       []interface{} `json:"components"`
		MessageReference struct {
			ChannelId string `json:"channel_id"`
			MessageId string `json:"message_id"`
			GuildId   string `json:"guild_id"`
		} `json:"message_reference,omitempty"`
	} `json:"referenced_message,omitempty"`
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
	TimeZone      string
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
type Guilds struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Icon        *string  `json:"icon"`
	Owner       bool     `json:"owner"`
	Permissions string   `json:"permissions"`
	Features    []string `json:"features"`
}

type JoinResp struct {
	Type      int         `json:"type,omitempty"`
	ExpiresAt interface{} `json:"expires_at,omitempty"`
	Flags     int         `json:"flags,omitempty"`
	Guild     struct {
		Id                       string   `json:"id,omitempty"`
		Name                     string   `json:"name,omitempty"`
		Splash                   string   `json:"splash,omitempty"`
		Banner                   string   `json:"banner,omitempty"`
		Description              string   `json:"description,omitempty"`
		Icon                     string   `json:"icon,omitempty"`
		Features                 []string `json:"features,omitempty"`
		VerificationLevel        int      `json:"verification_level,omitempty"`
		VanityUrlCode            string   `json:"vanity_url_code,omitempty"`
		NsfwLevel                int      `json:"nsfw_level,omitempty"`
		Nsfw                     bool     `json:"nsfw,omitempty"`
		PremiumSubscriptionCount int      `json:"premium_subscription_count,omitempty"`
	} `json:"guild,omitempty"`
	GuildId string `json:"guild_id,omitempty"`
	Channel struct {
		Id   string `json:"id"`
		Type int    `json:"type"`
		Name string `json:"name"`
	} `json:"channel"`
	ApproximateMemberCount   int     `json:"approximate_member_count"`
	ApproximatePresenceCount int     `json:"approximate_presence_count"`
	SiteKey                  string  `json:"captcha_sitekey,omitempty"`
	RqToken                  string  `json:"captcha_rqtoken,omitempty"`
	Retry                    float64 `json:"retry_after,omitempty"`
	Message                  string  `json:"message,omitempty"`
}

type Server struct {
	Id                 string      `json:"id"`
	Name               string      `json:"name"`
	Icon               string      `json:"icon"`
	Description        string      `json:"description"`
	HomeHeader         interface{} `json:"home_header"`
	Splash             string      `json:"splash"`
	DiscoverySplash    string      `json:"discovery_splash"`
	Features           []string    `json:"features"`
	Banner             string      `json:"banner"`
	OwnerId            string      `json:"owner_id"`
	ApplicationId      interface{} `json:"application_id"`
	Region             string      `json:"region"`
	AfkChannelId       string      `json:"afk_channel_id"`
	AfkTimeout         int         `json:"afk_timeout"`
	SystemChannelId    string      `json:"system_channel_id"`
	SystemChannelFlags int         `json:"system_channel_flags"`
	WidgetEnabled      bool        `json:"widget_enabled"`
	WidgetChannelId    interface{} `json:"widget_channel_id"`
	VerificationLevel  int         `json:"verification_level"`
	Roles              []struct {
		Id           string      `json:"id"`
		Name         string      `json:"name"`
		Description  interface{} `json:"description"`
		Permissions  string      `json:"permissions"`
		Position     int         `json:"position"`
		Color        int         `json:"color"`
		Hoist        bool        `json:"hoist"`
		Managed      bool        `json:"managed"`
		Mentionable  bool        `json:"mentionable"`
		Icon         *string     `json:"icon"`
		UnicodeEmoji interface{} `json:"unicode_emoji"`
		Flags        int         `json:"flags"`
		Tags         struct {
			BotId             string      `json:"bot_id,omitempty"`
			PremiumSubscriber interface{} `json:"premium_subscriber"`
		} `json:"tags,omitempty"`
	} `json:"roles"`
	DefaultMessageNotifications int         `json:"default_message_notifications"`
	MfaLevel                    int         `json:"mfa_level"`
	ExplicitContentFilter       int         `json:"explicit_content_filter"`
	MaxPresences                interface{} `json:"max_presences"`
	MaxMembers                  int         `json:"max_members"`
	MaxStageVideoChannelUsers   int         `json:"max_stage_video_channel_users"`
	MaxVideoChannelUsers        int         `json:"max_video_channel_users"`
	VanityUrlCode               string      `json:"vanity_url_code"`
	PremiumTier                 int         `json:"premium_tier"`
	PremiumSubscriptionCount    int         `json:"premium_subscription_count"`
	PreferredLocale             string      `json:"preferred_locale"`
	RulesChannelId              string      `json:"rules_channel_id"`
	SafetyAlertsChannelId       string      `json:"safety_alerts_channel_id"`
	PublicUpdatesChannelId      string      `json:"public_updates_channel_id"`
	HubType                     interface{} `json:"hub_type"`
	PremiumProgressBarEnabled   bool        `json:"premium_progress_bar_enabled"`
	LatestOnboardingQuestionId  string      `json:"latest_onboarding_question_id"`
	Nsfw                        bool        `json:"nsfw"`
	NsfwLevel                   int         `json:"nsfw_level"`
	Emojis                      []struct {
		Id            string        `json:"id"`
		Name          string        `json:"name"`
		Roles         []interface{} `json:"roles"`
		RequireColons bool          `json:"require_colons"`
		Managed       bool          `json:"managed"`
		Animated      bool          `json:"animated"`
		Available     bool          `json:"available"`
	} `json:"emojis"`
	Stickers []struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Tags        string `json:"tags"`
		Type        int    `json:"type"`
		FormatType  int    `json:"format_type"`
		Description string `json:"description"`
		Asset       string `json:"asset"`
		Available   bool   `json:"available"`
		GuildId     string `json:"guild_id"`
	} `json:"stickers"`
	IncidentsData     interface{} `json:"incidents_data"`
	InventorySettings interface{} `json:"inventory_settings"`
	EmbedEnabled      bool        `json:"embed_enabled"`
	EmbedChannelId    string      `json:"embed_channel_id"`
}

type ChannelID struct {
	ID string `json:"id,omitempty"`
}

type FriendReq struct {
	Username string
	Discrim  interface{}
}
type Friend struct {
	Id       string      `json:"id"`
	Type     int         `json:"type"`
	Nickname interface{} `json:"nickname"`
	User     struct {
		Id                   string      `json:"id"`
		Username             string      `json:"username"`
		GlobalName           string      `json:"global_name"`
		Avatar               string      `json:"avatar"`
		AvatarDecorationData interface{} `json:"avatar_decoration_data"`
		Discriminator        string      `json:"discriminator"`
		PublicFlags          int         `json:"public_flags"`
		Bot                  bool        `json:"bot,omitempty"`
	} `json:"user"`
}
type UserInfo struct {
	Id                   string      `json:"id"`
	Username             string      `json:"username"`
	Avatar               interface{} `json:"avatar"`
	Discriminator        string      `json:"discriminator"`
	PublicFlags          int         `json:"public_flags"`
	PremiumType          int         `json:"premium_type"`
	Flags                int         `json:"flags"`
	Banner               interface{} `json:"banner"`
	AccentColor          interface{} `json:"accent_color"`
	GlobalName           string      `json:"global_name"`
	AvatarDecorationData interface{} `json:"avatar_decoration_data"`
	BannerColor          interface{} `json:"banner_color"`
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
	DesignID           int         `json:"design_id"`
}

type WsData struct {
	UserIds           []string               `json:"user_ids"`
	Token             string                 `json:"token,omitempty"`
	Capabilities      int                    `json:"capabilities,omitempty"`
	Compress          bool                   `json:"compress,omitempty"`
	Ops               []Ops                  `json:"ops,omitempty"`
	Members           []Member               `json:"members,omitempty"`
	Properties        XpropData              `json:"properties,omitempty"`
	ClientState       ClientState            `json:"client_state,omitempty"`
	Presence          Presence               `json:"presence"`
	HeartbeatInterval int                    `json:"heartbeat_interval,omitempty"`
	SessionID         string                 `json:"session_id,omitempty"`
	GuildId           interface{}            `json:"guild_id,omitempty"`
	Channels          map[string]interface{} `json:"channels,omitempty"`
	ChannelID         string                 `json:"channel_id,omitempty"`
	Typing            bool                   `json:"typing,omitempty"`
	Threads           bool                   `json:"threads,omitempty"`
	Activities        bool                   `json:"activities,omitempty"`
	ThreadMemberLists interface{}            `json:"thread_member_lists,omitempty"`
	Sequence          int                    `json:"seq"`
	Message           struct {
		Content   string `json:"content,omitempty"`
		ChannelID string `json:"channel_id,omitempty"`
		GuildID   string `json:"guild_id,omitempty"`
		MessageId string `json:"id,omitempty"`
		Flags     int    `json:"flags,omitempty"`
	}
}

type ClientState struct {
	GuildVersions            struct{} `json:"guild_versions"`
	PrivateChannelsVersion   string   `json:"private_channels_version"`
	HighestLastMessageID     string   `json:"highest_last_message_id,omitempty"`
	ReadStateVersion         int      `json:"read_state_version,omitempty"`
	UserSettingsVersion      int      `json:"user_settings_version"`
	ApiCodeVersion           int      `json:"api_code_version"`
	UserGuildSettingsVersion int      `json:"user_guild_settings_version,omitempty"`
}

type Identify struct {
	Token        string `json:"token,omitempty"`
	Capabilities int    `json:"capabilities,omitempty"`
	Compress     bool   `json:"compress,omitempty"`
}

type Presence struct {
	Status     string   `json:"status"`
	Game       Game     `json:"game"`
	Since      int      `json:"since"`
	Activities []string `json:"activities"`
	Afk        bool     `json:"afk"`
}

type Member struct {
	ID                         string
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
type Game struct {
	Name       string   `json:"name"`
	Status     string   `json:"status"`
	Type       int      `json:"type"`
	Since      int      `json:"since"`
	Activities []string `json:"activities"`
	Afk        bool     `json:"afk"`
}

type VcOptions struct {
	CID  string
	GID  string
	Mute bool
	Deaf bool
}

type SoundBoardOptions struct {
	ID    string
	Emoji string
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
type Hcoptcha struct {
	Error        bool `json:"error"`
	TaskId       int  `json:"task_id"`
	TaskTypeInfo struct {
		Price      int    `json:"price"`
		TaskName   string `json:"task_name"`
		TaskTypeId int    `json:"task_type_id"`
	} `json:"task_type_info"`
}

type HcoptchaResponse struct {
	Error bool `json:"error"`
	Task  struct {
		CaptchaKey string `json:"captcha_key"`
		Refunded   bool   `json:"refunded"`
		State      string `json:"state"`
	} `json:"task"`
}

type CapSolver struct {
	ErrorId          int    `json:"errorId"`
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
	Status           string `json:"status"`
	Solution         struct {
		Text string `json:"text"`
	} `json:"solution"`
	TaskId string `json:"taskId"`
}

type Message struct {
	Title string `json:"Title"`
	Body  string `json:"Body"`
	Link  string `json:"Link"`
}

type Config struct {
	Dcfd   string
	Sdcfd  string
	Cfuvid string
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
