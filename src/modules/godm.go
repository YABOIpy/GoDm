package modules

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/Danny-Dasilva/fhttp"
)

func (*Instance) Joiner(in Instance, invite, session string, typ int) {
	s := time.Now()
	var (
		captcha, rqtoken string
		payload map[string]string
		ContextData []byte
		Count int
	)
retry:
	payload = map[string]string{"session_id": session}
	if len(captcha) > 0 {
		payload = map[string]string{
			"captcha_key": captcha,
			"captcha_rqtoken": rqtoken,
			"session_id": session,
		}
	} 
	if typ != 1{
		req, err := http.NewRequest("GET", fmt.Sprintf(
			"https://discord.com/api/v9/invites/%s?inputValue=%s&with_counts=true&with_expiration=true", invite, invite),
			nil,
		)
		if err != nil {
			log.Println(err)
		}
	
		Hd.Header(req, map[string]string{
			"authorization":      in.Token,
			"cookie":             in.Cookie,
			"user-agent":         in.BrowserClient.Agent,
			"referer":            "https://discord.com/channels/@me",
			"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
			"sec-ch-ua":          in.SecUA(in),
			"x-discord-timezone": in.TimeZones(),
			"x-super-properties": in.Xprop,
		})
		resp, err := in.Client.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
	
		var data JoinReq
		defer resp.Body.Close()
	
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		json.Unmarshal(body, &data)
	
		ContextData, err = json.Marshal(JoinContext{
			Location:            "Join Guild",
			LocationGuildId:     data.Guild.Id,
			LocationChannelId:   data.Channel.Id,
			LocationChannelType: data.Channel.Type,
		})
	}

	req2, err := http.NewRequest("POST",
		"https://discord.com/api/v9/invites/"+invite,
		modules.Marsh(
			payload,
		),
	)
	if err != nil {
		log.Println(err)
		return
	}

	Hd.Header(req2, map[string]string{
		"authorization":        in.Token,
		"cookie":               in.Cookie,
		"user-agent":           in.BrowserClient.Agent,
		"referer":              "https://discord.com/",
		"sec-ch-ua-platform":   in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":            in.SecUA(in),
		"x-discord-timezone":   in.TimeZones(),
		"x-super-properties":   in.Xprop,
	})
	if typ != 1{
		req2.Header.Set("x-context-properties", base64.StdEncoding.EncodeToString(ContextData))
		req2.Header.Set("referer", "https://discord.com/channels/@me",)
	}
	resp2, err := in.Client.Do(req2)
	if err != nil {
		log.Println(err)
		return
	}

	var dat JoinResp
	defer resp2.Body.Close()

	bod, err := io.ReadAll(resp2.Body)
	if err = json.Unmarshal(bod, &dat); err != nil {
		//log.Println(err)
	}

	switch resp2.StatusCode {
	case 200:
		modules.StrlogV("Joined", "gg/"+invite, s)
		pass, mail := in.TokenProps.Pass, in.TokenProps.Email
		if len(mail) > 0 && len(pass) > 0 {
			modules.WriteFile("data/joined.txt", fmt.Sprintf(TokenFormat, mail, pass, in.Token))
		} else {
			modules.WriteFile("data/joined.txt", in.Token)
		}

	case 429:
		modules.StrlogR("Failed", "RateLimit", s)
		in.TokenProps.RateLimit = dat.Retry
		modules.Sleep(time.Duration(dat.Retry), in)

	default:
		if strings.Contains(string(bod), "captcha_sitekey") {
			if in.Cfg.Mode.Configs.Solver {

				modules.StrlogR(fmt.Sprintf("%s[%s%d%s]%s %s", clr, bg, Count, rb, clr, "Captcha"), "Solving..", s)
				captcha = in.Captcha(CapCfg{
					ApiKey:  in.Cfg.Mode.Discord.CapAPI[1],
					SiteKey: dat.SiteKey,
					PageURL: "https://discord.com",
				})
				rqtoken = dat.RqToken
				Count++

				goto retry
			} else {
				modules.StrlogR("Captcha", "Ignoring", s)
			}
		} else {
			if strings.Contains(dat.Message, "verify") {
				modules.StrlogE("Locked", modules.HalfToken(in.Token, 0), s)
			} else {
				modules.StrlogE("Failed", string(bod), s)
			}
		}
	}
}

func (*Instance) Leaver(in Instance, ID string) {
	s := time.Now()
	req, err := http.NewRequest("DELETE",
		"https://discord.com/api/v9/users/@me/guilds/"+ID,
		modules.Marsh(map[string]interface{}{
			"lurking": false,
		}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZones(),
		"x-super-properties": in.Xprop,
		"referer":            "https://discord.com/channels/@me/" + ID,
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 204:
		modules.StrlogV("Left Server", ID, s)
	default:
		body, _ := io.ReadAll(resp.Body)
		modules.StrlogE("Failed To Leave", string(body), s)
	}
}

func (*Instance) Message(in Instance, msg, ID string, opt MessageOptions) (int, []byte) {
	var message, ping string
	message = msg
	for {
		if opt.Mping {
			ping = ""
			users := ReturnRandomArray(opt.IDs, opt.Amount)
			for _, user := range users {
				ping += fmt.Sprintf("<@%s>", user)
			}
			message = msg + " | " + ping
		}

		payload := map[string]interface{}{
			"content": message,
			"flags":   0,
			"nonce":   modules.Nonce(),
			"tts":     false,
		}
		if len(opt.Captcha) != 0 {
			payload["captcha"] = opt.Captcha
		}

		s := time.Now()
		req, err := http.NewRequest("POST", fmt.Sprintf(
			"https://discord.com/api/v9/channels/%s/messages", ID),
			modules.Marsh(payload),
		)
		if err != nil {
			log.Println(err)
			return 0, nil
		}

		Hd.Header(req, map[string]string{
			"authorization":      in.Token,
			"cookie":             in.Cookie,
			"user-agent":         in.BrowserClient.Agent,
			"referer":            "https://discord.com/channels/" + ID,
			"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
			"sec-ch-ua":          in.SecUA(in),
			"x-discord-timezone": in.TimeZones(),
			"x-super-properties": in.Xprop,
		})

		resp, err := in.Client.Do(req)
		if err != nil {
			log.Println(err)
			return 0, nil
		}

		var data MessageResp
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		switch resp.StatusCode {
		case 200:
			modules.StrlogV("Sent Message", strings.ReplaceAll(msg, "\n", ""), s)
		case 429:
			in.TokenProps.RateLimit = data.Retry //discord doesnt seem to return a value on this endpoint
			if opt.Loop {
				if !modules.Sleep(2, in) {
					modules.StrlogR("RateLimit", strconv.Itoa(2), s)
				}
			}
		default:
			if strings.Contains(string(body), "captcha_sitekey") {
				modules.StrlogE("Captcha", resp.Status, s)
			} else {
				modules.StrlogE("Failed", string(body), s)
			}

		}
		if !opt.Loop {
			return resp.StatusCode, body
		}
	}
}

func (in *Instance) Check() (int, time.Time) {
	s := time.Now()
	req, err := http.NewRequest("GET",
		"https://discord.com/api/v9/users/@me/affinities/guilds",
		nil,
	)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("authorization", in.Token)

	resp, err := in.SClient.Do(req)
	if err != nil {
		return in.Check()
	}
	defer resp.Body.Close()

	half := modules.HalfToken(in.Token, 0)
	if resp.StatusCode == 200 {
		modules.StrlogV("", half, s)
	} else if resp.StatusCode == 403 {
		modules.StrlogR("", half, s)
	} else {
		modules.StrlogE("", half, s)
	}

	return resp.StatusCode, s
}

func (*Instance) Friend(in Instance, data FriendReq) {
	s := time.Now()

	var Discrim interface{}
	if data.Discrim != nil {
		Discrim = modules.TrimZero(data.Discrim.(string))
	}

	req, err := http.NewRequest("POST",
		"https://discord.com/api/v9/users/@me/relationships",
		modules.Marsh(
			map[string]interface{}{
				"username":      data.Username,
				"discriminator": Discrim,
			},
		),
	)
	if err != nil {
		log.Println(err)
	}

	Hd.Header(req, map[string]string{
		"authorization":        in.Token,
		"cookie":               in.Cookie,
		"user-agent":           in.BrowserClient.Agent,
		"sec-ch-ua-platform":   in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":            in.SecUA(in),
		"x-discord-timezone":   in.TimeZones(),
		"referer":              "https://discord.com/channels/@me",
		"x-context-properties": "eyJsb2NhdGlvbiI6IkFkZCBGcmllbmQifQ==", //{"location":"Add Friend"}
		"x-super-properties":   in.Xprop,
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	var dat struct {
		Retry float64 `json:"retry_after,omitempty"`
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	switch resp.StatusCode {
	case 204, 200:
		modules.StrlogV("Sent Friend Request To", data.Username, s)
	case 429:
		json.Unmarshal(body, &dat)
		in.TokenProps.RateLimit = dat.Retry
		modules.Sleep(time.Duration(dat.Retry), in)
	default:
		modules.StrlogE("Failed To Send", string(body), s)
	}
}

func (*Instance) MemberVerify(in Instance, ID, invite string) {
	s := time.Now()
	req, err := http.NewRequest("GET", fmt.Sprintf(
		"https://discord.com/api/v9/guilds/%s/member-verification?with_guild=false&invite_code=%s", ID, invite),
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZones(),
		"x-super-properties": in.Xprop,
		"referer":            "https://discord.com/channels/%s" + ID,
	})

	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	var data Guild
	body, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}
	switch resp.StatusCode {
	case 200:
		reqs, ers := http.NewRequest("PUT",
			"https://discord.com/api/v9/guilds/"+ID+"/requests/@me",
			modules.Marsh(Guild{
				Version:     data.Version,
				FormFields:  data.FormFields,
				Description: data.Description,
			}),
		)
		if ers != nil {
			log.Println(ers)
			return
		}

		Hd.Header(reqs, map[string]string{
			"authorization":      in.Token,
			"cookie":             in.Cookie,
			"user-agent":         in.BrowserClient.Agent,
			"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
			"sec-ch-ua":          in.SecUA(in),
			"x-discord-timezone": in.TimeZones(),
			"x-super-properties": in.Xprop,
			"referer":            "https://discord.com/channels/" + ID,
		})

		res, ers := in.Client.Do(reqs)
		if ers != nil {
			log.Println(ers)
			return
		}
		defer res.Body.Close()
		bod, e := io.ReadAll(res.Body)
		if e != nil {
			log.Println(e)
		}

		switch res.StatusCode {
		case 201:
			modules.StrlogV("Verified To Guild", modules.HalfToken(in.Token, 0)+" | Guild:"+ID, s)
		default:
			modules.StrlogE("Failed To Verify", string(bod), s)
		}
	default:
		modules.StrlogE("Failed To Get Guild Data", string(body), s)
	}

}

func (*Instance) Reaction(in Instance, CID, MID, emoji string) {
	s := time.Now()
	req, err := http.NewRequest("PUT", fmt.Sprintf(
		"https://discord.com/api/v9/channels/%s/messages/%s/reactions/%s/", CID, MID, emoji)+"%40me?location=Message&type=0",
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}

	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZones(),
		"x-super-properties": in.Xprop,
		"referer":            "https://discord.com/channels/",
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 204:
		modules.StrlogV("Added Emoji", emoji, s)
	default:
		body, _ := io.ReadAll(resp.Body)
		modules.StrlogE("Failed To Add Emoji", string(body), s)
	}
}

func (*Instance) DisplayName(in Instance, username string) {
	s := time.Now()
	req, err := http.NewRequest("PATCH",
		"https://discord.com/api/v9/users/@me",
		modules.Marsh(map[string]string{
			"global_name": username,
		}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZones(),
		"x-super-properties": in.Xprop,
		"referer":            "https://discord.com/channels/@me/",
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		modules.StrlogV(modules.HalfToken(in.Token, 0)+" Changed Username To", username, s)
	default:
		body, _ := io.ReadAll(resp.Body)
		modules.StrlogE(modules.HalfToken(in.Token, 0)+"Failed To Change Username", string(body), s)
	}
}

func (*Instance) Username(in Instance, username string) {
	s := time.Now()
	req, err := http.NewRequest("PATCH",
		"https://discord.com/api/v9/users/@me",
		modules.Marsh(map[string]string{
			"password": in.TokenProps.Pass,
			"username": username,
		}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZones(),
		"x-super-properties": in.Xprop,
		"Referer":            "https://discord.com/channels/@me",
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		modules.StrlogV(modules.HalfToken(in.Token, 0)+" Changed Username To", username, s)
	default:
		body, _ := io.ReadAll(resp.Body)
		modules.StrlogE(modules.HalfToken(in.Token, 0)+"Failed To Change Username", string(body), s)
	}
}

func (*Instance) Pronouns(in Instance, pronoun string) {
	s := time.Now()
	req, err := http.NewRequest("PATCH",
		"https://discord.com/api/v9/users/%40me/profile",
		modules.Marsh(map[string]string{
			"pronouns": pronoun,
		}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZones(),
		"x-super-properties": in.Xprop,
		"referer":            "https://discord.com/channels/@me/",
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	switch resp.StatusCode {
	case 200:
		modules.StrlogV("Changed Pronouns", pronoun, s)
	default:
		modules.StrlogE("Failed To Change Pronouns", string(body), s)
	}
}

func (*Instance) Bio(in Instance, bio string) {
	s := time.Now()
	req, err := http.NewRequest("PATCH",
		"https://discord.com/api/v9/users/%40me/profile",
		modules.Marsh(map[string]string{
			"bio": bio,
		}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZones(),
		"x-super-properties": in.Xprop,
		"referer":            "https://discord.com/channels/@me/",
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	switch resp.StatusCode {
	case 200:
		modules.StrlogV("Changed Bio", bio, s)
	default:
		modules.StrlogE("Failed To Change Bio", string(body), s)
	}
}

func (*Instance) Avatar(in Instance, pfp string) {
	s := time.Now()
	req, err := http.NewRequest("PATCH",
		"https://discord.com/api/v9/users/@me",
		modules.Marsh(map[string]interface{}{
			"avatar": "data:image/png;base64," + pfp,
		}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	Hd.Header(req, map[string]string{
		"authorization":        in.Token,
		"cookie":               in.Cookie,
		"user-agent":           in.BrowserClient.Agent,
		"sec-ch-ua-platform":   in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":            in.SecUA(in),
		"x-discord-timezone":   in.TimeZones(),
		"x-context-properties": "eyJsb2NhdGlvbiI6IlF1aWNrIE1lc3NhZ2UgSW5wdXQifQ==", //{"location":"Quick Message Input"}
		"x-super-properties":   in.Xprop,
		"referer":              "https://discord.com/channels/@me/",
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	switch resp.StatusCode {
	case 200:
		modules.StrlogV("Changed Avatar", "", s)
	default:
		modules.StrlogE("Failed To Change Avatar", string(body), s)
	}
}

func (*Instance) Password(in Instance, pass string) string {
	s := time.Now()

	if len(pass) == 0 {
		pass = modules.RandString(rand.Intn(25-15+1) + 15)
	}
	req, err := http.NewRequest("PATCH",
		"https://discord.com/api/v9/users/@me",
		modules.Marsh(map[string]string{
			"new_password": pass,
			"password":     in.TokenProps.Pass,
		}),
	)
	if err != nil {
		log.Println(err)
	}

	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZones(),
		"x-super-properties": in.Xprop,
		"referer":            "https://discord.com/channels/@me/",
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	var data MeResp
	body, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}
	switch resp.StatusCode {
	case 200, 201, 204:
		modules.StrlogV(fmt.Sprintf("Changed Password %s[%s]%s", g, r+modules.HalfToken(in.Token, 0)+g, r), in.TokenProps.Pass+" | "+pass, s)
		return fmt.Sprintf(TokenFormat, data.Email, pass, data.Token)
	case 429:
		in.TokenProps.RateLimit = data.Retry
		modules.Sleep(time.Duration(data.Retry), in)
	default:
		modules.StrlogE("Failed To Change Password", string(body), s)
		return fmt.Sprintf(TokenFormat, in.TokenProps.Email, in.TokenProps.Pass, in.Token)
	}
	return fmt.Sprintf(TokenFormat, in.TokenProps.Email, in.TokenProps.Pass, in.Token)
}

func (*Instance) CreateChannel(in Instance, ID string) DmChannel {
	s := time.Now()
	req, err := http.NewRequest("POST",
		"https://discord.com/api/v9/users/@me/channels",
		modules.Marsh(map[string]interface{}{
			"recipients": []string{ID},
		}),
	)
	if err != nil {
		log.Println(err)
		return DmChannel{}
	}
	Hd.Header(req, map[string]string{
		"authorization":        in.Token,
		"cookie":               in.Cookie,
		"user-agent":           in.BrowserClient.Agent,
		"sec-ch-ua-platform":   in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":            in.SecUA(in),
		"x-discord-timezone":   in.TimeZones(),
		"x-context-properties": "eyJsb2NhdGlvbiI6IlF1aWNrIE1lc3NhZ2UgSW5wdXQifQ==", //{"location":"Quick Message Input"}
		"x-super-properties":   in.Xprop,
		"referer":              "https://discord.com/channels/",
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return DmChannel{}
	}

	defer resp.Body.Close()
	var data DmChannel
	body, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}

	switch resp.StatusCode {
	case 200:
		return data
	case 429:
		in.TokenProps.RateLimit = data.Retry
		modules.Sleep(time.Duration(data.Retry), in)
	default:
		modules.StrlogE("Failed To Get Channel", string(body), s)
	}

	return DmChannel{}
}

func (*Instance) Eligible(in Instance, ID string) bool {
	req, err := http.NewRequest("POST", fmt.Sprintf(
		"https://discord.com/api/v9/users/@me/referrals/%s/preview", ID),
		nil,
	)
	if err != nil {
		log.Println(err)
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZones(),
		"x-super-properties": in.Xprop,
		"referer":            "https://discord.com/channels/@me/",
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	var data struct {
		IsEligible bool `json:"is_eligible"`
	}
	body, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}

	switch resp.StatusCode {
	case 200, 204, 201:
		if data.IsEligible {
			return true
		}
	default:
		return false
	}
	return true
}

func (*Instance) Buttons(in Instance, data MessageResp, opt ButtonOptions) {
	s := time.Now()
	req, err := http.NewRequest("POST",
		"https://discord.com/api/v9/interactions",
		modules.Marsh(Button{
			AppID: data.ApplicationID,
			CID:   data.ChannelID,
			Data: ButtonData{
				Type: opt.Button.Type,
				ID:   opt.Button.CustomId,
			},
			GID:     opt.GuildID,
			Flags:   data.Flags,
			MID:     data.ID,
			Nonce:   modules.Nonce(),
			Session: opt.SessionID,
			Type:    opt.Type,
		}),
	)
	if err != nil {
		log.Println(err)
		return
	}

	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZones(),
		"x-super-properties": in.Xprop,
		"referer":            fmt.Sprintf("https://discord.com/channels/%s/%s", opt.GuildID, data.ChannelID),
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	switch resp.StatusCode {
	case 200, 204:
		modules.StrlogV("Clicked Button", opt.Button.Label, s)
	case 429:
		in.TokenProps.RateLimit = data.Retry
		modules.Sleep(time.Duration(data.Retry), in)
	default:
		modules.StrlogE("Failed To Click Button", string(body), s)
	}
}

func (*Instance) Boost(in Instance, ID string) {
	s := time.Now()
	req, err := http.NewRequest("GET",
		"https://discord.com/api/v9/users/@me/guilds/premium/subscription-slots",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZones(),
		"x-super-properties": in.Xprop,
		"referer":            "https://discord.com/channels/@me",
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var data []BoostResp
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}

	tkn := modules.HalfToken(in.Token, 0)
	switch resp.StatusCode {
	default:
		modules.StrlogE(fmt.Sprintf("%s[%s]%s Failed To Boost%s", red, tkn, red, r), ID, s)

	case 200:
		for _, v := range data {

			var slotID []string
			slotID = append(slotID, v.Id)
			re, er := http.NewRequest("PUT",
				"https://discord.com/api/v9/guilds/"+ID+"/premium/subscriptions",
				modules.Marsh(
					BoostPayload{
						UserPremiumGuildSubscriptionSlotIds: slotID,
					},
				),
			)
			if er != nil {
				log.Println(er)
			}

			Hd.Header(req, map[string]string{
				"authorization":      in.Token,
				"cookie":             in.Cookie,
				"user-agent":         in.BrowserClient.Agent,
				"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
				"sec-ch-ua":          in.SecUA(in),
				"x-discord-timezone": in.TimeZones(),
				"x-super-properties": in.Xprop,
				"referer":            "https://discord.com/channels/@me",
			})
			res, er := in.Client.Do(re)
			if er != nil {
				log.Fatal(er)
			}
			defer res.Body.Close()

			switch res.StatusCode {
			case 201:
				modules.StrlogV(fmt.Sprintf("%s[%s]%s Boosted%s", g, tkn, g, r), ID, s)
			default:
				modules.StrlogE(fmt.Sprintf("%s[%s]%s Failed To Boost%s", red, tkn, red, r), ID, s)
			}
		}
	}
}

func (*Instance) VoiceChat(in Instance, opt VcOptions) {
	Ws := Sock{}
	s := time.Now()
	_, _, conn := Ws.Connect(in.Token, in)
	for {
		conn.Ws.WriteJSON(map[string]interface{}{
			"op": 4,
			"d": map[string]interface{}{
				"guild_id":   opt.GID,
				"channel_id": opt.CID,
				"self_mute":  opt.Mute,
				"self_deaf":  opt.Deaf,
			},
		})
		_, b, _ := conn.Ws.ReadMessage()

		var data WsResp
		json.Unmarshal(b, &data)

		Events := []string{
			EventReadySupplemental,
			EventSessionReplace,
			EventVoiceServerUpdate,
			EventPresenceUpdate,
			EventVoiceStateUpdate,
		}

		if modules.Contains(Events, data.Name) {
			modules.StrlogV(fmt.Sprintf("Connected %s[%s]%s", g, r+modules.HalfToken(in.Token, 0)+g, r), data.Name, s)
			time.Sleep(30 * time.Second)
			continue
		} else {
			modules.StrlogE(fmt.Sprintf("Failed %s[%s]%s", g, r+modules.HalfToken(in.Token, 0)+g, r), data.Name, s)
			time.Sleep(3 * time.Second)
			continue
		}
		//modules.StrlogV(fmt.Sprintf("Connected %s[%s]%s", g, r+modules.HalfToken(in.Token, 0)+g, r), opt.CID, s)
		//time.Sleep(30 * time.Millisecond)
	}
}

func (*Instance) OnBoard(in Instance, GID string) {
	// TODO: make this!
	modules.OnboardingData(in, GID)
	req, err := http.NewRequest("POST", fmt.Sprintf(
		"https://discord.com/api/v9/guilds/%s/onboarding-responses", GID),
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZones(),
		"x-super-properties": in.Xprop,
		"referer":            "https://discord.com/channels/" + GID + "/onboarding",
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	var data struct{}
	body, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}

	switch resp.StatusCode {
	case 200:

	default:

	}

}
