package modules

import (
	"github.com/Danny-Dasilva/fhttp"
)

var (
	Hd = Header{}
)

func (Hd *Header) Header(req *http.Request, headers map[string]string) {
	for h, o := range map[string]string{
		"accept":           "*/*",
		"accept-language":  "en-US,en-GB;q=0.9",
		"content-type":     "application/json",
		"origin":           "https://discord.com",
		"sec-ch-ua-mobile": "?0",
		"sec-fetch-dest":   "empty",
		"sec-fetch-mode":   "cors",
		"sec-fetch-site":   "same-origin",
		"x-debug-options":  "bugReporterEnabled",
		"x-discord-locale": "en-US",
	} {
		req.Header.Set(h, o)
	}
	for h, o := range headers {
		req.Header.Set(h, o)

	}
}
