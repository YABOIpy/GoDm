package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"github.com/wasilibs/go-re2"
	"log"
	"math/rand"
	"os"
	"runtime"
	"source/src/modules"
	"source/src/task"
	"strings"
	"time"
)

var (
	Mod = modules.Modules{}
	Con = modules.Instance{}
	Ws  = modules.Sock{}
)

type FuncMap map[int]func()

func main() {
	in := initialize()
	task.Return(0)
	LoadChoice(in)
}

func initialize() []modules.Instance {
	log.Println(modules.Initializing)
	modules.RSeed.GenerateSeed()
	instances, err := Con.Configuration()
	if err != nil {
		log.Println(err)
	}
	return instances
}

func LoadChoice(in []modules.Instance) {

	opt := FuncMap{
		1: func() {
			var cooldown time.Duration

			switch Mod.InputInt("1: Mass DM \n2: Mass Friend \nChoice") {
			case 1:

				var msg string
				fmt.Println(modules.MassDmMention)
				if interval := Mod.InputInt("CoolDown"); interval != modules.IntNil {
					cooldown = time.Duration(interval) + time.Duration(rand.Intn(9)+2)
				}
			r:
				if msg = Mod.Input(modules.MessageInput); len(msg) == modules.IntNil {
					cfg, _ := Mod.LoadConfig("config.json")
					message := cfg.Mode.Discord.Message
					if message != nil {
						modules.RandSeed().GenerateSeed()
						i := rand.Intn(len(message))
						msg = fmt.Sprintf("%s\n%s\n%s",
							message[i].Title,
							message[i].Body,
							message[i].Link,
						)
						time.Sleep(time.Hour)
					} else {
						log.Println("No Messages Found.")
						time.Sleep(2 * time.Second)
						goto r
					}
				}
				task.MassDmTask(in, msg, cooldown)
			case 2:
				fmt.Println(modules.MassFriendOptionMention)
				if interval := Mod.InputInt("CoolDown"); interval != modules.IntNil {
					cooldown = time.Duration(interval) + time.Duration(rand.Intn(9)+2)
				}
				if user := Mod.Input("Username: "); user != modules.StringNil {
					task.StartTask(in, func(c modules.Instance) {
						c.DisplayName(user)
					})
				}
				if bio := Mod.Input("Bio: "); bio != modules.StringNil {
					task.StartTask(in, func(c modules.Instance) {
						c.Bio(bio)
					})
				}
				task.MassFriendTask(in, cooldown)
			}
		},
		2: func() {
			ID := Mod.Input("UserID: ")
			msg := Mod.Input(modules.MessageInput)
			task.StartTask(in, func(c modules.Instance) {
				data := c.CreateChannel(ID)
				if c.Eligible(ID) {
					c.Message(msg, data.Id, modules.MessageOptions{Loop: true})
				}
			})
		},
		3: func() {
			fmt.Println(modules.InServerMention)
			CID := Mod.Input(modules.ChannelInput)
			MID := Mod.Input("Message ID: ")
			data := in[0].MessageData(CID, MID)
			for _, v := range data {
				for j, k := range v.Reactions {
					fmt.Printf("\u001B[36m| [\u001B[39m%d\u001B[36m]\u001B[39m %s ", j, k.Emoji.Name)
				}
			}
			fmt.Println()
			emoji := data[0].Reactions[Mod.InputInt("Choice")].Emoji.Name
			task.StartTask(in, func(c modules.Instance) {
				c.Reaction(CID, MID, emoji)
			})
		},
		4: func() {
			typ := Mod.InputInt("1: Direct API\n2: Add Server API\nChoice")
			inv := Mod.Input(modules.InviteInput)
			os.Truncate("data/joined.txt", 0)
			task.StartTask(in, func(c modules.Instance) {
				d, con, err := Ws.Connect(c.Token, &c)
				if err != nil {
					return
				}
				defer con.Ws.Close()
				c.Joiner(inv, d.Data.SessionID, typ)
			})
			j, _, _ := Mod.ReadFile("data/joined.txt")
			if len(j) != len(in) && len(j) > modules.IntNil {
				if Mod.InputBool(modules.WriteJoinedMention) {
					os.Truncate("tokens.txt", 0)
					Mod.WriteFileArray("tokens.txt", j)
					main()
				}
			}
		},
		5: func() {
			ID := Mod.Input(modules.GuildInput)
			task.StartTask(in, func(c modules.Instance) {
				c.Leaver(ID)
			})
		},
		6: func() {
			fmt.Println(modules.InServerMention)
			inv := Mod.Input(modules.InviteInput)
			ID := in[0].GuildJoinData(inv).GuildId
			task.StartTask(in, func(c modules.Instance) {
				c.MemberVerify(ID, inv)
			})
		},
		7: func() {
			msg := Mod.Input(modules.MessageInput)
			ID := Mod.Input(modules.ChannelInput)
			task.StartTask(in, func(c modules.Instance) {
				c.Message(msg, ID, modules.MessageOptions{Loop: true})
			})
		},
		8: func() {
			task.ScrapeTask(in[0],
				Mod.Input(modules.GuildInput),
				Mod.Input(modules.ChannelInput),
			)
		},
		9: func() { task.CheckerTask(in) },
		10: func() {
			msg := Mod.Input(modules.MessageInput)
			ID := Mod.Input(modules.ChannelInput)
			ids, _, _ := Mod.ReadFile("data/ids.txt")
			options := modules.MessageOptions{
				Mping:  true,
				Loop:   true,
				IDs:    ids,
				Amount: Mod.InputInt("Ping Per Message"),
			}
			task.StartTask(in, func(c modules.Instance) {
				c.Message(msg, ID, options)
			})
		},
		11: func() {
			//will leave indexing like this. i have yet to see more data.
			ID := re2.MustCompile(`\d+`).FindAllString(Mod.Input("Message Link: "), -1)
			data := in[0].MessageData(ID[1], ID[2])
			for i, d := range data {
				for j, b := range d.Components[i].Components {
					fmt.Printf("\033[36m| [\033[39m%d\u001B[36m]\u001B[39m %s %s ", j, b.Emoji.Name, b.Label)
				}
			}
			fmt.Println()

			opt := &modules.ButtonOptions{
				Button:  data[0].Components[0].Components[Mod.InputInt("Choice")], // <-
				Type:    Mod.InputInt("Button Type"),
				GuildID: ID[0],
			}
			task.StartTask(in, func(c modules.Instance) {
				wsd, _, _ := Ws.Connect(c.Token, &c)
				opt.SessionID = wsd.Data.SessionID
				c.Buttons(data[0], *opt)
			})
		},
		12: func() {
			fmt.Println(modules.DiscrimMention)
			data := modules.FriendReq{
				Username: Mod.Input("Username: "),
			}
			disc := Mod.Input(data.Username + "#")
			data.Discrim = nil
			if disc != modules.StringNil {
				data.Discrim = disc
			}
			task.StartTask(in, func(c modules.Instance) {
				c.Friend(data)
			})
		},
		13: func() {
			choice := Mod.InputInt(modules.TokenOptions)
			switch choice {
			case 1:
				user := Mod.Input("Username: ")
				task.StartTask(in, func(c modules.Instance) {
					c.DisplayName(user)
				})
			case 2:
				bio := Mod.Input("Bio: ")
				task.StartTask(in, func(c modules.Instance) {
					c.Bio(bio)
				})
			case 3:
				fmt.Println(modules.BandWidthMention)
				fmt.Println(modules.ImageFormatMention)
				if !Mod.InputBool("Continue") {
					break
				}
				img := Mod.ReadDirectory("data/pfp", "png")
				task.StartTask(in, func(c modules.Instance) {
					_, con, err := Ws.Connect(c.Token, &c)
					if err != nil {
						return
					}
					defer con.Ws.Close()
					c.Avatar(img[rand.Intn(len(img))])
				})
			case 4:
				var data []string
				fmt.Println(modules.TokenFormatMention)
				fmt.Println(modules.PasswordFieldMention)
				password := Mod.Input("Password: ")
				task.StartTask(in, func(c modules.Instance) {
					data = append(data, c.Password(password))
				})
				os.Truncate("tokens.txt", 0)
				Mod.WriteFileArray("tokens.txt", data)
			case 5:
				text := Mod.Input("Pronouns: ")
				task.StartTask(in, func(c modules.Instance) {
					c.Pronouns(text)
				})
			case 6:
				fmt.Println(modules.TokenFormatMention)
				user := Mod.Input("Username: ")
				task.StartTask(in, func(c modules.Instance) {
					c.Username(user)
				})
			case 7:
				// TODO: take combos from txt file
				//user := Mod.Input("Username: ")
				// task.StartTask(in, func(c modules.Instance) {
				// })
				fmt.Println("Coming Soon..")
			case 8:
				fmt.Println(modules.RGBMention)
				clr := strings.Split(fmt.Sprint(Mod.Input("Input RGB: ")), ",")
				task.StartTask(in, func(c modules.Instance) {
					c.ChangeBanner(modules.RGB(
						cast.ToInt(clr[0]), cast.ToInt(clr[1]), cast.ToInt(clr[2])),
					)
				})
			case 9:
				task.StartTask(in, func(c modules.Instance) {
					for _, d := range c.OpenChannels() {
						c.CloseDM(d.Id)
					}
					for _, d := range c.Friends() {
						c.RemoveFriend(d)
					}
					for _, d := range c.Guilds() {
						time.Sleep(850 * time.Millisecond)
						c.Leaver(d.Id)
					}
				})
			case 10:
				for {
					c := in[0]
					_, ws, _ := Ws.Connect(in[0].Token, &c)
					var data modules.WsResp

					//for _, d := range Mod.Guilds(c) {
					ws.Ws.WriteJSON(map[string]interface{}{
						"op": 8,
						"d": map[string]interface{}{
							"guild_id": []string{
								"125440014904590336",
							},
							"presences": false,
						}})
					_, b, _ := ws.Ws.ReadMessage()
					json.Unmarshal(b, &data)
					fmt.Println(data.Name)
					if data.Name == modules.EventMessageCreate {
						fmt.Println(data.Data.Message.Content, data.Data.Message.MessageId)
						fmt.Println(data.Data.Message)
					}
					//Mod.FetchMessages(Mod.Guild(d.Id).Id, 100)
					//}
				}
			}
		},
		14: func() {
			ID := Mod.Input(modules.GuildInput)
			task.StartTask(in, func(c modules.Instance) {
				c.Boost(ID)
			})
		},
		15: func() {
			opt := modules.VcOptions{
				GID:  Mod.Input(modules.GuildInput),
				CID:  Mod.Input(modules.ChannelInput),
				Mute: Mod.InputBool("Mute"),
				Deaf: Mod.InputBool("Deafen"),
			}
			task.StartTask(in, func(c modules.Instance) {
				c.VoiceChat(opt)
			})
		},
		16: func() {
			CID := Mod.Input(modules.ChannelInput)
			opt := map[int]modules.SoundBoardOptions{
				0: {"1", "🦆"}, 1: {"2", "🔊"},
				2: {"3", "🦗"}, 3: {"4", "👏"},
				4: {"5", "🎺"}, 5: {"6", "🥁"},
			}
			for j, k := range []int{0, 1, 2, 3, 4, 5} { // i could use PrintMenu but i like the look of this more.
				if v, ok := opt[k]; ok {
					fmt.Printf("\u001B[36m| [\u001B[39m%d\u001B[36m]\u001B[39m %s ", j, v.Emoji)
				}
			}
			sound := opt[Mod.InputInt("\nChoice")]
			ok := Mod.InputBool("Loop")
			task.StartTask(in, func(c modules.Instance) {
			l:
				c.SoundBoard(CID, sound)
				if ok {
					goto l
				}
			})
		},
		17: func() {
			var opt []string
			var verify bool

			inv := Mod.Input(modules.InviteInput)
			guild := in[0].GuildJoinData(inv)
			data := in[0].OnboardingData(guild.GuildId)

			if Mod.Contains(guild.Guild.Features, modules.MemberVerificationGateEnabled) {
				verify = Mod.InputBool("Server Has Member Verification. Verify?")
			}
			if !Mod.Contains(guild.Guild.Features, modules.GuildOnboarding) {
				fmt.Println("Server Doesn't Have an OnBoarding Prompt")
				return
			}
			for _, d := range data.Prompts {
				if d.Required {
					fmt.Printf("\u001B[36m[\u001B[39m%s\u001B[36m]\u001B[39m:\n", d.Title)
					for i, o := range d.Options {
						fmt.Printf("%d: (%s)=%s\n", i, o.Title, o.Description)
					}
					opt = append(opt, d.Options[Mod.InputInt("Choice")].Id)
				}
			}
			task.StartTask(in, func(c modules.Instance) {
				c.OnBoard(guild.GuildId, opt)
				if verify {
					c.MemberVerify(guild.GuildId, inv)
				}
			})
		},
		18: func() {
			switch Mod.InputInt("1: Server Info \n2: In Guild Checker \nOption") {
			case 1:
				s := time.Now()
				data := in[0].GuildJoinData(Mod.Input(modules.InviteInput))
				if data.Message != modules.StringNil {
					Mod.StrlogE("Failed To Fetch Data", data.Message, s)
					return
				}
				modules.PrintStruct(data)
				Mod.Input("Press Enter To Continue")
			case 2:
				var i []string
				GID := Mod.Input(modules.GuildInput)
				task.StartTask(in, func(c modules.Instance) {
					s := time.Now()
					data := c.Guild(GID)
					switch len(data.Id) {
					case 0:
						Mod.StrlogE(fmt.Sprintf("\u001B[31m[\u001B[39m%s\u001B[31m]\u001B[39m", Mod.HalfToken(c.Token, 0)), "Not In Server: "+GID, s)
					default:
						Mod.StrlogV(fmt.Sprintf("\u001B[32m[\u001B[39m%s\u001B[32m]\u001B[39m", Mod.HalfToken(c.Token, 0)), "In Server: "+GID, s)
						i = append(i, c.Token)
					}
				})
				if len(i) != len(in) && len(i) > modules.IntNil {
					if Mod.InputBool(modules.WriteInServerMention) {
						os.Truncate("tokens.txt", 0)
						Mod.WriteFileArray("tokens.txt", i)
						main()
					}
				}
			default:
				return
			}
		},
	}
	for {
		choice := Mod.InputInt("Choice")
		if choice == modules.IntNil {
			//restart the client
			runtime.GC()
			Mod.Cls()
			main()
		}
		if function, v := opt[choice]; v {
			function()
			task.Return(3)
		} else {
			fmt.Println("Invalid Choice..")
			task.Return(1)
		}
	}
}
