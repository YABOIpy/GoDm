package modules

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Danny-Dasilva/fhttp"
	"github.com/zenthangplus/goccm"

	HttpClient "source/src/client"
)

func (in *Instance) Configuration() ([]Instance, error) {
	cfg, err := modules.LoadConfig("config.json")
	if err != nil {
		log.Println("Failed To Load Config")
		time.Sleep(3 * time.Second)
		return nil, err
	}

	var (
		mutex      sync.Mutex
		Instances  []Instance
		tokenprops TokenConfig
	)

	Tokens, TokenData, Proxies := modules.FetchInputData()
	if Tokens == nil {
		return nil, errors.New("no tokens found in tokens.txt")
	}
	routines := len(Tokens)
	if cfg.Mode.Configs.CCManager {
		routines = cfg.Mode.Configs.MaxRoutines
	}
	wg := goccm.New(routines)

	fmt.Print("\033[s")

	for i := 0; i < len(Tokens); i++ {
		wg.Wait()
		go func(i int) {
			defer wg.Done()
			var proxy string

			cookie := modules.Cookies()
			Browser := in.CreateBrowser()

			fmt.Print("\033[u\033[K")
			fmt.Printf(CacheLoading, cookie[:20], strings.Split(Tokens[i], ".")[0])

			if len(cfg.Mode.Network.Proxy) != IntNil {
				proxy = "http://" + cfg.Mode.Network.Proxy
			} else if len(Proxies) != IntNil {
				proxy = "http://" + Proxies[i]
			}

			if TokenData != nil {
				tokenprops = TokenConfig{
					Email: TokenData[i].Email,
					Pass:  TokenData[i].Pass,
				}
			}
			Client, _ := HttpClient.NewClient(HttpClient.Browser{
				JA3:       cfg.Mode.Network.Ja3,
				UserAgent: Browser.Agent,
				Cookies:   nil,
			},
				cfg.Mode.Network.TimeOut,
				cfg.Mode.Network.Redirect,
				Browser.Agent,
				proxy,
			)

			mutex.Lock()
			Instances = append(Instances, Instance{
				Token:         Tokens[i],
				TokenProps:    tokenprops,
				TimeZone:      in.TimeZones(),
				Client:        Client,
				SClient:       &http.Client{},
				Xprop:         in.Xprops(Browser),
				BrowserClient: Browser,
				Cookie:        cookie,
				Cfg:           cfg,
			})
			mutex.Unlock()
		}(i)
	}
	wg.WaitAllDone()

	return Instances, err
}
