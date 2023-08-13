package modules

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
	
	http "github.com/Danny-Dasilva/fhttp"
	shttp "net/http"
)

func (in *Instance) Browser() (browsers []BrowserData) {
	browsers = []BrowserData{
		{
			Name:     "Chromium",
			Versions: []string{"116", "115", "114", "113", "112", "111", "110", "109", "108", "107", "106", "105", "104", "103", "102", "101", "100", "99", "98", "97", "96", "95", "94", "93", "92", "91", "90"},
			OSver: map[string][]string{
				"Windows": {"10", "8.1", "8", "7"},
				"Mac":     {"11", "10.15", "10.14", "10.13"},
				"Linux":   {"Ubuntu/20", "Debian/10", "Fedora/34"},
			},
			UserAgent: map[string]string{
				"Windows": "Mozilla/5.0 (Windows NT %s; Win64; x64) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s",
				"Mac":     "Mozilla/5.0 (Macintosh; Intel Mac OS X %s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s",
				"Linux":   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s %s",
			},
		},
		{
			Name:     "Mozilla Firefox",
			Versions: []string{"115.0", "114.0", "113.0", "112.0", "111.0", "110.0", "109.0", "108.0", "107.0", "106.0", "105.0", "104.0", "103.0", "102.0", "101.0", "100.0", "99.0", "98.0", "97.0", "96.0", "95.0", "94.0", "93.0", "92.0", "91.0", "90.0"},

			OSver: map[string][]string{
				"Windows": {"10", "8.1", "8", "7"},
				"Mac":     {"11", "10.15", "10.14", "10.13"},
				"Linux":   {"Ubuntu/20", "Debian 10", "Fedora/34"},
			},
			UserAgent: map[string]string{
				"Windows": "Mozilla/5.0 (Windows NT %s) Gecko/20100101 Firefox/%s",
				"Mac":     "Mozilla/5.0 (Macintosh; Intel Mac OS X %s) Gecko/20100101 Firefox/%s",
				"Linux":   "Mozilla/5.0 (X11; Linux x86_64) Gecko/20100101 Firefox/%s %s",
			},
		},
		{
			Name:     "Safari",
			Versions: []string{"15", "14", "13", "12", "11"},
			OSver: map[string][]string{
				"Mac": {"11", "10.15", "10.14", "10.13"},
			},
			UserAgent: map[string]string{
				"Mac": "Mozilla/5.0 (Macintosh; Intel Mac OS X %s) AppleWebKit/%s (KHTML, like Gecko) Version/%s Safari/%s",
			},
		},
		{
			Name:     "Microsoft Edge",
			Versions: []string{"88", "88", "89", "89", "90", "90", "91", "91", "92", "92", "93", "93", "94", "94", "95", "95", "96", "96", "97", "97", "98", "98", "99", "99", "100", "100", "101", "102", "102", "103", "103", "104", "104", "105", "105", "106", "106", "107", "107", "108", "108", "109", "109", "110", "110", "111", "111", "112", "112", "113", "113", "114", "114", "115", "115", "116", "117", "118"},
			OSver: map[string][]string{
				"Windows": {"10", "8.1", "8", "7"},
				"Mac":     {"11", "10.15", "10.14", "10.13"},
			},
			UserAgent: map[string]string{
				"Windows": "Mozilla/5.0 (Windows NT %s; Win64; x64) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s Edg/%s",
				"Mac":     "Mozilla/5.0 (Macintosh; Intel Mac OS X %s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s Edg/%s",
			},
		},
		{
			Name:     "Opera",
			Versions: []string{"76.0", "75.0", "74.0", "73.0", "72.0", "71.0", "70.0", "69.0", "68.0", "67.0", "66.0", "65.0"},
			OSver: map[string][]string{
				"Windows": {"10", "8.1", "8", "7"},
				"Mac":     {"11", "10.15", "10.14", "10.13"},
				"Linux":   {"Ubuntu/20", "Debian/10", "Fedora/34"},
			},
			UserAgent: map[string]string{
				"Windows": "Mozilla/5.0 (Windows NT %s; Win64; x64) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s OPR/%s",
				"Mac":     "Mozilla/5.0 (Macintosh; Intel Mac OS X %s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s OPR/%s",
				"Linux":   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s OPR/%s %s",
			},
		},
	}

	return browsers
}

// was thinking of adding specific timezones depended on the ip. but most will use residential Proxies...
func (in *Instance) TimeZones() string {
	zones := []string{
		"Pacific/Midway",
		"Pacific/Pago_Pago",
		"Pacific/Honolulu",
		"America/Anchorage",
		"America/Los_Angeles",
		"America/Denver",
		"America/Chicago",
		"America/New_York",
		"America/Sao_Paulo",
		"Atlantic/South_Georgia",
		"Atlantic/Azores",
		"Europe/London",
		"Europe/Paris",
		"Europe/Istanbul",
		"Africa/Cairo",
		"Africa/Johannesburg",
		"Asia/Damascus",
		"Asia/Jerusalem",
		"Asia/Riyadh",
		"Asia/Tehran",
		"Asia/Dubai",
		"Asia/Yekaterinburg",
		"Asia/Kolkata",
		"Asia/Kathmandu",
		"Asia/Dhaka",
		"Asia/Bangkok",
		"Asia/Hong_Kong",
		"Asia/Tokyo",
		"Australia/Sydney",
		"Pacific/Noumea",
		"Pacific/Auckland",
		"Pacific/Tongatapu",
	}
	return in.GetRandomData(zones)
}

func (in *Instance) UserAgent(Browser BrowserData, data ClientData) Agents {

	cfg, _ := modules.LoadConfig("config.json")
	WebKit := cfg.Mode.Network.WebKit
	os := data.OS

	switch Browser.Name {
	case "Chromium", "Safari":
		agent := fmt.Sprintf(Browser.UserAgent[os], in.GetRandomData(Browser.OSver[os]), WebKit, in.GetRandomData(Browser.Versions), WebKit)
		if Browser.Name == "Chromium" {
			return Agents{
				Windows: agent,
				Linux:   fmt.Sprintf(Browser.UserAgent[os], WebKit, in.GetRandomData(Browser.Versions), WebKit, in.GetRandomData(Browser.OSver[os])),
				Mac:     agent,
			}
		}
		return Agents{Mac: agent}

	case "Microsoft Edge", "Opera":
		d := in.Browser()
		agent := fmt.Sprintf(Browser.UserAgent[os], in.GetRandomData(Browser.OSver[os]), WebKit, in.GetRandomData(d[0].Versions), WebKit, in.GetRandomData(Browser.Versions))
		if Browser.Name == "Opera" {
			return Agents{
				Windows: agent,
				Linux:   fmt.Sprintf(Browser.UserAgent[os], WebKit, in.GetRandomData(d[0].Versions), WebKit, in.GetRandomData(Browser.Versions), in.GetRandomData(Browser.OSver[os])),
				Mac:     agent,
			}
		}
		return Agents{
			Mac:     agent,
			Windows: agent,
		}

	case "Mozilla Firefox":
		agent := fmt.Sprintf(Browser.UserAgent[os], in.GetRandomData(Browser.OSver[os]), in.GetRandomData(Browser.Versions))
		return Agents{
			Windows: agent,
			Linux:   fmt.Sprintf(Browser.UserAgent[os], in.GetRandomData(Browser.Versions), in.GetRandomData(Browser.OSver[os])),
			Mac:     agent,
		}
	}
	return Agents{Mac: cfg.Mode.Network.Agent, Linux: cfg.Mode.Network.Agent, Windows: cfg.Mode.Network.Agent}
}

func (in *Instance) CreateBrowser() ClientData {
	RSeed.GenerateSeed()

	Browser := in.GetRandomBrowser(in.Browser())
	os := in.getRandomKey(Browser.OSver)
	ua := in.UserAgent(Browser, ClientData{
		Name:    Browser.Name,
		OS:      os,
		OSVer:   in.GetRandomData(Browser.OSver[os]),
		Version: in.GetRandomData(Browser.Versions),
	})

	var agent string
	switch os {
	case "Windows":
		agent = ua.Windows
	case "Mac":
		agent = ua.Mac
	case "Linux":
		agent = ua.Linux
	}
	return ClientData{
		Name:    Browser.Name,
		OS:      os,
		OSVer:   in.GetRandomData(Browser.OSver[os]),
		Version: in.GetRandomData(Browser.Versions),
		Agent:   agent,
	}
}

func (in *Instance) Captcha(data CapCfg) string {

	switch strings.ToLower(in.Cfg.Mode.Discord.CapAPI[0]) {
	case "capmonster":
		return in.SolveCapMonster(data)
	case "capsolver":
		return in.SolveCapSolver(data)
	}
	return ""
}

// i havent checked this yet all captcha docs are very badly written. so is their api wrapper on github
func (in *Instance) SolveCapSolver(cfg CapCfg) string {
	s := time.Now()
	req, err := http.NewRequest("POST",
		"https://api.capsolver.com/createTask",
		modules.Marsh(map[string]string{
			"ApiKey": cfg.ApiKey,
		}),
	)
	if err != nil {
		log.Println(err)
		return ""
	}
	for h, o := range map[string]string{
		"Content-Type": "application/json",
	} {
		req.Header.Set(h, o)
	}
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()

	var data struct {
		ErrorId          int    `json:"errorId"`
		ErrorCode        string `json:"errorCode"`
		ErrorDescription string `json:"errorDescription"`
		Status           string `json:"status"`
		Solution         struct {
			Text string `json:"text"`
		} `json:"solution"`
		TaskId string `json:"taskId"`
	}

	body, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
		return ""
	}

	switch resp.StatusCode {
	case 200:
		return data.TaskId
	default:
		modules.StrlogE("Failed To Get TaskID", data.ErrorDescription, s)
	}

	return ""
}

func (in *Instance) SolveCapMonster(cfg CapCfg) string {
	req, err := shttp.NewRequest("POST",
		"https://api.capmonster.cloud/createTask",
		modules.Marsh(map[string]string{
			"type":       "HCaptchaTaskProxyless",
			"userAgent":  in.BrowserClient.Agent,
			"websiteKey": cfg.SiteKey,
			"websiteURL": cfg.PageURL,
		}),
	)
	if err != nil {
		log.Println(err)
		return ""
	}
	req.Header.Set("content-type", "application/json")
	Client := &shttp.Client{}

	resp, err := Client.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var data struct {
		Captcha  interface{}
		solution interface{}
		TaskID   int `json:"taskId"`
	}
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}

	if resp.StatusCode == 200 {
		for true {
			re, er := shttp.NewRequest("POST",
				"https://api.capmonster.cloud/getTaskResult",
				modules.Marsh(map[string]interface{}{
					"clientKey": cfg.ApiKey,
					"taskId":    data.TaskID,
				}),
			)
			if er != nil {
				log.Println(err)
				return ""
			}
			res, er := Client.Do(re)
			if er != nil {
				log.Println(err)
				return ""
			}

			defer res.Body.Close()
			bod := make(map[string]interface{})
			json.NewDecoder(res.Body).Decode(&bod)

			switch bod["status"] {
			case "ready":
				return bod["solution"].(map[string]interface{})["gRecaptchaResponse"].(string)
			case "processing":
				continue
			default:
				log.Println(bod)
			}
		}
	}

	return ""
}
