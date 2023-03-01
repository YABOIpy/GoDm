package massdm

import goclient "massdm/client"

var (
	Client, _ = goclient.NewClient(goclient.Browser{
		JA3:       c.Config().Mode.Network.Ja3,
		UserAgent: c.Config().Mode.Network.Agent,
		Cookies:   nil,
	},
		c.Config().Mode.Network.TimeOut,
		c.Config().Mode.Network.Redirect,
		c.Config().Mode.Network.Agent,
		"http://"+c.Config().Mode.Network.Proxy,
	)
)
