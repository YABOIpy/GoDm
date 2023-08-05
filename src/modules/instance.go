package modules

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Danny-Dasilva/fhttp"
	HttpClient "source/src/client"
)

func (in *Instance) Configuration() ([]Instance, error) {
	Conf, err := modules.LoadConfig("config.json")
	if err != nil {
		log.Println("Failed To Load Config")
		time.Sleep(3 * time.Second)
		return nil, err
	}

	var (
		wg        sync.WaitGroup
		mutex     sync.Mutex
		Instances []Instance
		Client    *http.Client
	)
	var (
		routines int
		proxy    string
		cfg      = Conf
	)

	Tokens, TokenData, Proxies := modules.FetchInputData()

	routines = len(Tokens)
	if cfg.Mode.Configs.CCManager {
		routines = cfg.Mode.Configs.MaxRoutines
	}
	wg.Add(routines)

	fmt.Print("\033[s")

	for i := 0; i < len(Tokens); i++ {
		go func(i int) {
			defer wg.Done()

			cookie, Cstruct := modules.Cookies()
			Browser := in.CreateBrowser()
			xprop := in.Xprops(ClientData{
				Name:    Browser.Name,
				OS:      Browser.OS,
				OSVer:   Browser.OSVer,
				Version: Browser.Version,
				Agent:   Browser.Agent,
			})

			fmt.Print("\033[u\033[K")
			fmt.Printf(CacheLoading, cookie[:20], strings.Split(Tokens[i], ".")[0])

			if len(cfg.Mode.Network.Proxy) != 0 {
				proxy = "http://" + cfg.Mode.Network.Proxy
			} else {
				proxy = ""
				if len(Proxies) > 0 {
					proxy = "http://" + Proxies[i]
				}
			}
			CookieSettings := Cstruct.Cookie.Cookies
			Client, _ = HttpClient.NewClient(HttpClient.Browser{
				JA3:       cfg.Mode.Network.Ja3,
				UserAgent: Browser.Agent,
				Cookies: []HttpClient.Cookie{
					{Name: "__dcfduid",
						Value:  Cstruct.Dcfd,
						Domain: CookieSettings["__dcfduid"].Domain,
						Secure: CookieSettings["__dcfduid"].Secure,
						MaxAge: CookieSettings["__dcfduid"].MaxAge,
					},
					{Name: "__sdcfduid",
						Value:  Cstruct.Sdcfd,
						Domain: CookieSettings["__sdcfduid"].Domain,
						Secure: CookieSettings["__sdcfduid"].Secure,
						MaxAge: CookieSettings["__sdcfduid"].MaxAge,
					},
					{Name: "__cfruid",
						Value:  Cstruct.Cfruid,
						Domain: CookieSettings["__cfruid"].Domain,
						Secure: CookieSettings["__cfruid"].Secure,
						MaxAge: CookieSettings["__cfruid"].MaxAge,
					},
				},
			},
				cfg.Mode.Network.TimeOut,
				cfg.Mode.Network.Redirect,
				Browser.Agent,
				proxy,
			)

			mutex.Lock()
			Instances = append(Instances, Instance{
				TokenProps: TokenConfig{
					Email: TokenData.Email,
					Pass:  TokenData.Pass,
				},
				Client:        Client,
				SClient:       &http.Client{},
				Xprop:         xprop,
				BrowserClient: Browser,
				Cookie:        cookie,
				Token:         Tokens[i],
				Cfg:           cfg,
			})
			mutex.Unlock()

		}(i)
	}
	wg.Wait()

	return Instances, err
}
