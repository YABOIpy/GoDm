package massdm

import (
	"bytes"
	"encoding/json"
	"fmt"
	http "github.com/Danny-Dasilva/fhttp"
	"io/ioutil"
	"math/rand"
	shttp "net/http"
	"strconv"
	"strings"
	"time"
)

func (Xc *Config) Dm(ID string, Token string, Msg string, Cookies string) {
	payload := map[string]string{
		"content": Msg,
	}

	req, err := http.NewRequest("POST", "https://discord.com/api/v9/channels/"+ID+"/messages",
		bytes.NewBuffer(Xc.Marsh(payload)),
	)
	Xc.Errs(err)

	Hd.Head_Dm(req, Token)
	resp, err := Client.Do(req)
	Xc.Errs(err)

	if resp.StatusCode == 200 {
		fmt.Println(""+grn+"▏"+r+"("+grn+"+"+r+") Sent Message To:"+clr+"", ID)
	} else {
		fmt.Println(""+red+"▏"+r+"("+red+"-"+r+") Failed Sent Message To:"+clr+"", ID,
			Xc.Errmsg(*resp),
		)
	}
}

func (Xc *Config) CloseDm(ID string, Token string, Cookies string) {
	req, err := http.NewRequest("DELETE",
		"https://discord.com/api/v9/channels/"+ID+"?silent=false",
		nil,
	)
	Xc.Errs(err)

	Hd.Head_CloseDm(req, Token)
	resp, err := Client.Do(req)
	Xc.Errs(err)

	if resp.StatusCode == 200 {
		fmt.Println(""+grn+"▏"+r+"("+grn+"+"+r+") Closed Channel:"+clr+"", ID)
	} else {
		fmt.Println(""+red+"▏"+r+"("+red+"-"+r+") Failed To Close Dm:"+clr+"", ID, "ERR |: ",
			Xc.Errmsg(*resp),
		)
	}

}

func (Xc *Config) React(Token string, link string) {
	payload := map[string]string{}
	req, err := http.NewRequest("PUT",
		link,
		bytes.NewBuffer(
			Xc.Marsh(payload),
		),
	)
	Xc.Errs(err)

	Hd.Head_React(req, Token)
	resp, err := Client.Do(req)
	Xc.Errs(err)

	switch resp.StatusCode {
	case 204:
		fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Added Emoji")
	default:
		fmt.Println(""+red+"▏"+r+"("+red+"-"+r+") Failed To Add Emoji"+clr+"", "ERR |: "+r,
			Xc.Errmsg(*resp),
		)
	}
}

func (Xc *Config) Create(ID int, Token string, Msg string) (string, error) {
	payload := []byte("{\"recipients\":[\"" + strconv.Itoa(ID) + "\"]}")
	req, err := http.NewRequest("POST", "https://discord.com/api/v9/users/@me/channels",
		bytes.NewBuffer(payload),
	)

	Xc.Errs(err)

	Hd.Head_Create(req, Token)
	resp, err := Client.Do(req)

	var flake ChannelID
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := ReadBody(*resp)
		Xc.Errs(err)
		errx := json.Unmarshal(body, &flake)
		Xc.Errs(errx)
		fmt.Println(""+grn+"▏"+r+"("+grn+"+"+r+") Created Channel:"+clr+"", flake.ID)
		Xc.Dm(flake.ID, Token, Msg, Cookies)
		return flake.ID, err

	} else {
		fmt.Println(""+red+"▏"+r+"("+red+"+"+r+") Failed To Create Channel:",
			Xc.Errmsg(*resp),
		)
	}
	return flake.ID, err
}

func (Xc *Config) Block(ID int, Token string, Cookies string) {
	p := map[string]string{"type": "2"}
	req, err := http.NewRequest("PUT", "https://discord.com/api/v9/users/@me/relationships/"+strconv.Itoa(ID)+"",
		bytes.NewBuffer(Xc.Marsh(p)),
	)

	Xc.Errs(err)

	Hd.Head_Block(req, Token, ID)

	resp, err := Client.Do(req)
	Xc.Errs(err)
	if resp.StatusCode == 204 {
		fmt.Println(""+grn+"▏"+r+"("+grn+"+"+r+") Blocked User:"+clr+"", ID)
	} else {
		fmt.Println(""+red+"▏"+r+"("+red+"-"+r+") Failed To Block:"+clr+"", ID, "ERR |: ",
			Xc.Errmsg(*resp),
		)
	}
}

func (Xc *Config) Dm_Spam(ID string, Token string, Msg string) {
	Xc.Dm(ID, Token, Msg, Cookies)
}

func (Xc *Config) Joiner(Token string, invite string, cap string, captoken string) {

	var payload map[string]string
	if len(cap) > 2 {
		payload = map[string]string{
			"captcha_key":     cap,
			"captcha_rqtoken": captoken,
		}
	} else {
		payload = map[string]string{"": ""}
	}

	req, err := http.NewRequest("POST",
		"https://discord.com/api/v9/invites/"+invite+"",
		bytes.NewBuffer(
			Xc.Marsh(
				payload,
			),
		),
	)
	Xc.Errs(err)

	Hd.Head_Joiner(req, Token)
	resp, err := Client.Do(req)
	Xc.Errs(err)

	var data JoinResp
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &data)
	Xc.Errs(err)

	if resp.StatusCode == 200 {
		fmt.Println(""+grn+"▏"+r+"("+grn+"+"+r+") Joined "+clr+"discord.gg/"+invite, r)
	} else if resp.StatusCode == 429 {
		fmt.Println(""+red+"▏"+r+"("+red+"+"+r+") Failed To Join "+clr+"discord.gg/"+invite, yel+" RateLimit", r)
	} else if strings.Contains(string(body), "captcha_sitekey") {
		if Xc.Config().Mode.Configs.Solver {
			fmt.Println(""+yel+"▏"+r+"("+yel+"+"+r+") Solving Captcha... "+clr+"discord.gg/"+invite, r)
			cap := Xc.Captcha(data.SiteKey)
			captoken := data.RqToken
			Xc.Joiner(Token, invite, cap, captoken)
		} else {
			fmt.Println(""+yel+"▏"+r+"("+yel+"+"+r+") Failed To Join "+clr+"discord.gg/"+invite, yel+" Captcha", r)
		}
	} else {

		fmt.Println(""+red+"▏"+r+"("+red+"+"+r+") Failed To Join "+clr+"discord.gg/"+invite,
			Xc.Errmsg(*resp),
		)
	}

}

func (Xc *Config) Leaver(Token string, ID string) {
	req, err := http.NewRequest("DELETE", "https://discord.com/api/v9/users/@me/guilds/"+ID+"",
		bytes.NewBuffer(
			Xc.Marsh(
				map[string]string{"lurking": "false"},
			),
		),
	)
	Xc.Errs(err)

	Hd.Head_Leaver(req, Token)

	resp, err := Client.Do(req)
	_ = err
	if resp.StatusCode == 204 {
		fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Left Server")
	} else {
		fmt.Println(""+yel+"▏"+r+"("+yel+"+"+r+") Failed To Leave "+clr+ID,
			Xc.Errmsg(*resp),
		)
	}
}

func (Xc *Config) Agree(Token string, invite string, ID string) {
	req, err := http.NewRequest("POST",
		"https://discord.com/api/v9/guilds/"+ID+"/member-verification?with_guild=false&invite_code="+invite+"",
		nil,
	)
	Xc.Errs(err)

	Hd.Head_Agree(req, Token, 1)
	resp, err := Client.Do(req)
	Xc.Errs(err)

	defer resp.Body.Close()
	body, ers := ReadBody(*resp)
	Xc.Errs(ers)

	var data Bypass
	err = json.Unmarshal(body, &data)
	Xc.Errs(err)

	for i := 0; i < len(data.FormFields); i++ {
		data.FormFields[i].Response = true
	}

	payload, _ := json.Marshal(data)
	reqs, er := http.NewRequest("POST", "https://discord.com/api/v9/guilds/"+ID+"/requests/@me",
		bytes.NewBuffer(
			payload,
		),
	)
	Xc.Errs(er)

	Hd.Head_Agree(reqs, Token, 0)
	resps, es := Client.Do(reqs)
	_ = es
	if resps.StatusCode == 201 {
		fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Verified Token")
	} else {
		fmt.Println(""+red+"▏"+r+"("+red+"-"+r+") Failed",
			Xc.Errmsg(*resp),
		)
	}

}

func (Xc *Config) Friend(Token string, username string, discrim string) {

	req, err := http.NewRequest("POST",
		"https://discord.com/api/v9/users/@me/relationships",
		bytes.NewBuffer(
			Xc.Marsh(
				map[string]string{
					"username":      username,
					"discriminator": discrim,
				},
			),
		),
	)
	Xc.Errs(err)

	Hd.Head_Friend(req, Token)
	resp, err := Client.Do(req)
	Xc.Errs(err)

	if resp.StatusCode == 204|200 {
		fmt.Println("" + grn + "▏(" + grn + "+" + r + ") Sent Friend Request To: " +
			username + "#" + discrim,
		)
	} else {
		fmt.Println("" + grn + "▏(" + grn + "+" + r + ") Failed To Friend Request: " +
			username + "#" + discrim,
		)
	}

}

func (Xc *Config) Check(token string) int {
	req, err := shttp.NewRequest("GET", urls, nil)
	Xc.Errs(err)

	req.Header.Set("authorization", token)

	sClient := &shttp.Client{}
	resp, err := sClient.Do(req)
	Xc.Errs(err)

	if resp.StatusCode == 200 {
		fmt.Println(""+grn+"▏ "+r+"("+grn+"✓"+r+") ("+grn+"+"+r+"):", token[:50]+"...")
		Xc.Checker.Valid++

	} else if resp.StatusCode == 403 {
		fmt.Println(""+yel+"▏ "+r+"("+yel+"/"+r+"):", token[:50]+"...")
		Xc.Checker.Locked++
	} else {
		fmt.Println(""+red+"▏ "+r+"("+red+"x"+r+"):", token[:50]+"...")
		Xc.Checker.Invalid++
	}

	Xc.Checker.All++
	return resp.StatusCode
}

func (Xc *Config) Buttons(Token string, GID string, CID string, MID string, BotID string, Type int, Comp int, Text string) {

	req, err := http.NewRequest("POST",
		"https://discord.com/api/v9/interactions",
		strings.NewReader(
			string(
				Xc.Marsh_btn(
					map[string]interface{}{
						"application_id": BotID,
						"channel_id":     CID,
						"data": map[string]interface{}{
							"component_type": Comp,
							"custom_id":      Text,
						},
						"guild_id":      GID,
						"message_flags": 0,
						"message_id":    MID,
						"type":          Type,
						"session_id":    Xc.Socket(Token).Data.SessionID,
					},
				),
			),
		),
	)

	Xc.Errs(err)

	Hd.Head_Button(req, Token, GID, CID)
	resp, err := Client.Do(req)
	Xc.Errs(err)

	if resp.StatusCode == 204|200 {
		fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Clicked Button ")
	} else if resp.StatusCode == 429 {
		fmt.Println(""+yel+"▏"+r+"("+yel+"+"+r+") Failed To Click Button "+yel+" RateLimit", r)
	} else {
		fmt.Println(""+red+"▏"+r+"("+red+"+"+r+") Failed To Click Button | Err: ", Xc.Errmsg(*resp))
	}

}

func (Xc *Config) Raider(Token string, message string, ID string) {
	for true {

		req, err := http.NewRequest("POST", "https://discord.com/api/v9/channels/"+ID+"/messages",
			bytes.NewBuffer(
				Xc.Marsh(map[string]string{
					"content": message,
				}),
			),
		)
		Xc.Errs(err)

		Hd.Head_Raider(req, Token, ID)
		resp, ers := Client.Do(req)
		_ = ers
		if resp.StatusCode == 200 {
			fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Sent Message")
		} else if resp.StatusCode == 429 {
			fmt.Println("" + yel + "▏" + r + "(" + yel + "+" + r + ") RateLimit")
		} else {
			fmt.Println(""+red+"▏"+r+"("+red+"+"+r+") Failed To Send:",
				Xc.Errmsg(*resp),
			)
		}
	}
}

func (Xc *Config) MassPing(Token string, Message string, Amount int, ID string) {
	for true {
		var msg string
		rand.Seed(time.Now().Unix())
		users, _ := Xc.ReadFile("ids.txt")
		for i := 0; i < Amount; i++ {

			ping := users[rand.Intn(len(users))]
			msg += "<@" + ping + ">"
		}

		req, err := http.NewRequest("POST", "https://discord.com/api/v9/channels/"+ID+"/messages",
			bytes.NewBuffer(
				Xc.Marsh(map[string]string{
					"content": Message + " " + msg,
				}),
			),
		)
		Xc.Errs(err)

		Hd.Head_MassPing(req, Token, ID)
		resp, ers := Client.Do(req)
		Xc.Errs(ers)
		if resp.StatusCode == 200 {
			fmt.Println("" + grn + "▏" + r + "(" + grn + "+" + r + ") Sent Message")
		} else if resp.StatusCode == 429 {
			fmt.Println("" + yel + "▏" + r + "(" + yel + "+" + r + ") RateLimit")
		} else {
			fmt.Println(""+red+"▏"+r+"("+red+"+"+r+") Failed To Send:",
				Xc.Errmsg(*resp),
			)
		}
	}
}

func X() Config {
	x := Config{}
	return x
}
