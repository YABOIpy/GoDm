package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"source/src/modules"
	"source/src/task"
	"time"

	"github.com/wasilibs/go-re2"
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
			var msg string
			fmt.Println(modules.MassDmMention)
			if interval := Mod.InputInt("CoolDown"); interval != 0 {
				cooldown = time.Duration(interval) + time.Duration(rand.Intn(9)+2)
			}
		r:
			if msg = Mod.Input("Message: "); len(msg) == 0 {
				cfg, _ := Mod.LoadConfig("config.json")
				message := cfg.Mode.Discord.Message
				if message != nil {
					modules.RandSeed().GenerateSeed()
					i := rand.Intn(len(message))
					msg = fmt.Sprintf("%s \n%s \n%s",
						message[i].Title,
						message[i].Body,
						message[i].Link,
					)
				} else {
					log.Println("No Messages Found.")
					time.Sleep(2 * time.Second)
					goto r
				}
			}
			task.MassDmTask(in, msg, cooldown)
		},
		2: func() {
			ID := Mod.Input("UserID: ")
			msg := Mod.Input("Message: ")
			task.StartTask(in, func(c modules.Instance) {
				data := Con.CreateChannel(c, ID)
				if Con.Eligible(c, ID) {
					Con.Message(c, msg, data.Id, modules.MessageOptions{Loop: true})
				}
			})
		},
		3: func() {
			fmt.Println(modules.InServerMention)
			CID := Mod.Input("Channel ID: ")
			MID := Mod.Input("Message ID: ")
			data := Mod.MessageData(in[0], CID, MID)
			for _, v := range data {
				for j, k := range v.Reactions {
					fmt.Printf("\u001B[36m| [\u001B[39m%d\u001B[36m]\u001B[39m %s ", j, k.Emoji.Name)
				}
			}
			fmt.Println()
			emoji := data[0].Reactions[Mod.InputInt("Choice")].Emoji.Name
			task.StartTask(in, func(c modules.Instance) {
				Con.Reaction(c, CID, MID, emoji)
			})
		},
		4: func() {
			inv := Mod.Input("discord.gg/")
			os.Truncate("data/joined.txt", 0)
			task.StartTask(in, func(c modules.Instance) {
				d, _, con := Ws.Connect(c.Token, c)
				defer con.Ws.Close()
				Con.Joiner(c, inv, d.Data.SessionID)
			})
			j, _, _ := Mod.ReadFile("data/joined.txt")
			if len(j) != len(in) && len(j) > 0 {
				if Mod.Input(modules.WriteJoinedMention) == "y" {
					os.Truncate("tokens.txt", 0)
					Mod.WriteFileArray("tokens.txt", j)
					main()
				}
			}
		},
		5: func() {
			ID := Mod.Input("Guild ID: ")
			task.StartTask(in, func(c modules.Instance) {
				Con.Leaver(c, ID)
			})
		},
		6: func() {
			fmt.Println(modules.InServerMention)
			ID := Mod.Input("Guild ID: ")
			inv := Mod.Input("discord.gg/")
			task.StartTask(in, func(c modules.Instance) {
				Con.MemberVerify(c, ID, inv)
			})
		},
		7: func() {
			// Should be clear enough.. fmt.Println(modules.InServerMention)
			msg := Mod.Input("Message: ")
			ID := Mod.Input("Channel ID:")
			task.StartTask(in, func(c modules.Instance) {
				Con.Message(c, msg, ID, modules.MessageOptions{Loop: true})
			})
		},
		8: func() {
			Token := Mod.Input("Token:")
			GID := Mod.Input("Guild ID: ")
			CID := Mod.Input("Channel ID: ")
			task.ScrapeTask(Token, in[0], GID, CID)
		},
		9: func() { task.CheckerTask(in) },
		10: func() {
			msg := Mod.Input("Message: ")
			ID := Mod.Input("Channel ID:")
			ids, _, _ := Mod.ReadFile("data/ids.txt")
			options := modules.MessageOptions{
				Mping:  true,
				Loop:   true,
				IDs:    ids,
				Amount: Mod.InputInt("Ping Per Message"),
			}
			task.StartTask(in, func(c modules.Instance) {
				Con.Message(c, msg, ID, options)
			})
		},
		11: func() {
			//will leave indexing like this. i have yet to see more data.
			link := Mod.Input("Message Link: ")
			ID := re2.MustCompile(`\d+`).FindAllString(link, -1)
			data := Mod.MessageData(in[0], ID[1], ID[2])
			for i, d := range data {
				for j, b := range d.Components[i].Components {
					fmt.Printf("\033[36m| [\033[39m%d\u001B[36m]\u001B[39m %s %s ", j, b.Emoji.Name, b.Label)
				}
			}
			fmt.Println()

			opt := &modules.ButtonOptions{
				Button:  data[0].Components[0].Components[Mod.InputInt("Choice")], // <-
				Type:    3,
				GuildID: ID[0],
			}
			task.StartTask(in, func(c modules.Instance) {
				wsd, _, _ := Ws.Connect(c.Token, c)
				opt.SessionID = wsd.Data.SessionID
				Con.Buttons(c, data[0], *opt)
			})
		},
		12: func() {
			fmt.Println(modules.DiscrimMention)
			data := modules.FriendReq{
				Username: Mod.Input("Username: "),
			}
			disc := Mod.Input(data.Username + "#")
			data.Discrim = nil
			if disc != "" {
				data.Discrim = disc
			}
			task.StartTask(in, func(c modules.Instance) {
				Con.Friend(c, data)
			})
		},
		13: func() {
			choice := Mod.InputInt(modules.TokenOptions)
			switch choice {
			case 1:
				user := Mod.Input("Username: ")
				task.StartTask(in, func(c modules.Instance) {
					Con.DisplayName(c, user)
				})
			case 2:
				bio := Mod.Input("Bio: ")
				task.StartTask(in, func(c modules.Instance) {
					Con.Bio(c, bio)
				})
			case 3:
				fmt.Println(modules.BandWidthMention)
				fmt.Println(modules.ImageFormatMention)
				if !Mod.InputBool("Continue") {
					break
				}
				img := Mod.ReadDirectory("data/pfp", "png")
				task.StartTask(in, func(c modules.Instance) {
					_, _, conn := Ws.Connect(c.Token, c)
					defer conn.Ws.Close()
					Con.Avatar(c, img[rand.Intn(len(img))])
				})
			case 4:
				var data []string
				fmt.Println(modules.TokenFormatMention)
				fmt.Println(modules.PasswordFieldMention)
				password := Mod.Input("Password: ")
				task.StartTask(in, func(c modules.Instance) {
					data = append(data, Con.Password(c, password))
				})
				os.Truncate("tokens.txt", 0)
				Mod.WriteFileArray("tokens.txt", data)
			case 5:
				text := Mod.Input("Pronouns: ")
				task.StartTask(in, func(c modules.Instance) {
					Con.Pronouns(c, text)
				})
			case 6:
				fmt.Println(modules.TokenFormatMention)
				user := Mod.Input("Username: ")
				task.StartTask(in, func(c modules.Instance) {
					Con.Username(c, user)
				})
			case 7:
				// TODO: take combos from txt file
				//user := Mod.Input("Username: ")
				// task.StartTask(in, func(c modules.Instance) {
				// })
				fmt.Println("Coming Soon..")
			}
		},
		14: func() {
			ID := Mod.Input("Guild ID: ")
			task.StartTask(in, func(c modules.Instance) {
				Con.Boost(c, ID)
			})
		},
		15: func() {
			opt := modules.VcOptions{
				GID:  Mod.Input("Guild ID: "),
				CID:  Mod.Input("Channel ID: "),
				Mute: Mod.InputBool("Mute"),
				Deaf: Mod.InputBool("Deafen"),
			}
			task.StartTask(in, func(c modules.Instance) {
				Con.VoiceChat(c, opt)
			})
		},
		16: func() {
			fmt.Println("Coming Soon..")
		},
	}
	for {
		choice := Mod.InputInt("Choice")
		if choice == 0 {
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

// TODO: Options {
// onboarding + captcha support
//}
