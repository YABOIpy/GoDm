package massdm

import (
	"net/url"
	"net/http"
	"crypto/tls"
	"github.com/gorilla/websocket"
)


type ChannelID struct {
	ID 		string `json:"id,omitempty"`
}

type Config struct {
	Headers map[string]string
	Dcfd	string
	Sdcfd	string
	Length	int
	ID  	string	

	XConfig struct {
		Proxy 	 string `json:"Proxy"`
	}`json:"Config"`

	Settings struct {
		Websock	 bool `json:"Websocket"`
		Block	 bool `json:"Block_Usr"`
		Close    bool `json:"Close_DM"`
	} `json:"Settings"`
}

type Connect struct {
	Token         string
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

var (
	c = X()
	proxy = c.Config().XConfig.Proxy
	p, _ = url.Parse("http://" + proxy)
	Client = http.Client {
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MaxVersion: tls.VersionTLS13,
			},
			Proxy: http.ProxyURL(p),
		},
	}
	Cookies = "__dcfduid=" + cookies().Dcfd + "; " + "__sdcfduid=" + cookies().Sdcfd + "; "
	grn = "\033[32m"
	red = "\033[31m"
	clr = "\033[36m"
	r   = "\033[39m"

	Logo = `
	____`+clr+`_____`+r+`__`+clr+`____     `+r+`____`+clr+`____`+r+`____`+clr+`__  ___
	`+r+`__`+clr+`  ____/`+r+`_  `+clr+`__ \    `+r+`___`+clr+`  __ \`+r+`__`+clr+`   |/  /
	`+r+`_ `+clr+` / __ `+r+`_`+clr+`  / / /    `+r+`__`+clr+`  / / /`+r+`_`+clr+`  /|_/ / 
	`+clr+`/ /_/ / / /_/ /     `+r+`_  `+clr+`/_/ /`+r+`_`+clr+`  /  / /  
	\____/  \____/      /_____/ /_/  /_/   
    
	[`+r+`Go Discord Mass Dm`+clr+`]		~`+r+`YABOI`+clr+`
	[`+r+`1`+clr+`]`+r+` Mass DM `+clr+`
	[`+r+`2`+clr+`]`+r+` Dm Spam `+clr+`
	[`+r+`3`+clr+`]`+r+` React Verify `+clr+`
	[`+r+`4`+clr+`]`+r+` Joiner `+clr+`
	[`+r+`5`+clr+`]`+r+` Leaver 

	Choice `+clr+`>>:`+r+` `

)
