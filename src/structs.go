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

	Settings struct {
		Websock	 bool `json:"Websocket"`
		Block	 bool `json:"Block_Usr"`
		Close    bool `json:"Close_DM"`
		Proxy 	 string `json:"Proxy"`
	} `json:"Config"`
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
	proxy = c.Config().Settings.Proxy
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
)