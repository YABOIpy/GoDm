package main

import (
	"bufio"
	"fmt"
	"massdm/scraper"
	"massdm/src"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/zenthangplus/goccm"
)

var (
	c = massdm.X()
	z = massdm.T()
	s = Scraper.X()
)

func MassDm(message string) {

	var wg goccm.ConcurrencyManager
	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)
	ids, err := c.ReadFile("ids.txt")
	c.Errs(err)

	if len(Token) > 300 {
		wg = goccm.New(len(Token))
	} else {
		wg = goccm.New(300)
	}

	for i := 0; i < len(Token); i++ {
		wg.Wait()
		go func(i int) {
			defer wg.Done()
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])
			}
			for _, UserID := range ids {
				ID, _ := strconv.Atoi(UserID)
				CID, err := c.Create(ID, Token[i], message)
				c.Errs(err)

				if c.Config().Settings.Close == true {
					c.CloseDm(CID, Token[i], massdm.Cookies)
				}
				if c.Config().Settings.Block == true {
					c.Block(ID, Token[i], massdm.Cookies)
				}
			}
		}(i)
	}
	wg.WaitAllDone()
	Return()

}

func Spam_Dm(UserID string, message string) {

	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)

	var wg sync.WaitGroup
	wg.Add(len(Token))

	ID, _ := strconv.Atoi(UserID)
	for i := 0; i < len(Token); i++ {
		go func(i int) {
			defer wg.Done()
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])
			}
			CID, err := c.Create(ID, Token[i], message)
			c.Errs(err)
			for true {
				c.Dm_Spam(CID, Token[i], message)
			}
		}(i)
	}
	wg.Wait()
}

func Join(invite string) {

	var wg goccm.ConcurrencyManager
	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)
	interval := c.Config().Mode.Configs.Interval
	if interval > 0 {
		for i := 0; i < len(Token); i++ {
			time.Sleep(time.Duration(interval) * time.Second)
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])
			}
			c.Joiner(Token[i], invite, "", "")

		}
		Return()
	} else {
		if len(Token) > 300 {
			wg = goccm.New(len(Token))
		} else {
			wg = goccm.New(300)
		}
		for i := 0; i < len(Token); i++ {
			wg.Wait()
			go func(i int) {
				defer wg.Done()
				if c.Config().Settings.Websock == true {
					c.WebSock(Token[i])
				}
				c.Joiner(Token[i], invite, "", "")
			}(i)
		}
		wg.WaitAllDone()
		Return()
	}
}

func Leave(ID string) {

	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)

	var wg sync.WaitGroup
	wg.Add(len(Token))

	for i := 0; i < len(Token); i++ {
		go func(i int) {
			defer wg.Done()
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])
			}
			c.Leaver(Token[i], ID)
		}(i)
	}
	wg.Wait()
	Return()
}

func Check() {
	var wg goccm.ConcurrencyManager
	token, err := c.ReadFile("tokens.txt")
	c.Errs(err)
	f := [3]string{
		"data/valid.txt",
		"data/locked.txt",
		"data/invalid.txt",
	}
	for i := 0; i < len(f); i++ {
		os.Truncate(f[i], 0)
	}

	if len(token) > 300 {
		wg = goccm.New(len(token))
	} else {
		wg = goccm.New(300)
	}

	start := time.Now()
	for i := 0; i < len(token); i++ {
		wg.Wait()
		go func(i int) {
			defer wg.Done()
			Token, Valid := z.Check(token[i])
			if Valid != "" {
				c.WriteFile("data/valid.txt", Valid)
			} else if z.Resp == false {
				c.WriteFile("data/locked.txt", Token)
			} else {
				c.WriteFile("data/invalid.txt", Token)
			}

		}(i)
	}
	wg.WaitAllDone()
	elapsed := time.Since(start)
	fmt.Println("[\033[32m✓\033[39m] (TIME\033[39m):", elapsed.String()[:4]+"Ms", "\033[39m(\033[33mLOCKED\033[39m):", z.Locked, "(\033[31mINVALID\033[39m):", z.Invalid, "(\033[32mVALID\033[39m):", z.Valid, "(\u001b[34;1mTOTAL\033[39m):", z.All)
	Return()
}

func Reac(link string) {
	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)

	var wg sync.WaitGroup
	wg.Add(len(Token))

	for i := 0; i < len(Token); i++ {
		go func(i int) {
			defer wg.Done()
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])
			}
			c.React(Token[i], link)
		}(i)
	}
	wg.Wait()
	Return()
}

func Rules(invite string, ID string) {

	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)

	var wg sync.WaitGroup
	wg.Add(len(Token))

	for i := 0; i < len(Token); i++ {
		go func(i int) {
			defer wg.Done()
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])
			}
			c.Agree(Token[i], invite, ID)
		}(i)
	}
	wg.Wait()
	Return()
}

func Raid(message string, ID string) {

	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)

	var wg sync.WaitGroup
	wg.Add(len(Token))

	for i := 0; i < len(Token); i++ {
		go func(i int) {
			defer wg.Done()
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])
			}
			c.Raider(Token[i], message, ID)
		}(i)
	}
	wg.Wait()
}

func Friend(user string) {

	username := strings.Split(user, "#")
	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)

	var wg sync.WaitGroup
	wg.Add(len(Token))

	for i := 0; i < len(Token); i++ {
		go func(i int) {
			defer wg.Done()
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])
			}
			c.Friend(Token[i], username[0], username[1])
		}(i)
	}
	wg.Wait()
	Return()
}

func Scrape(Token string, GID string, CID string) {
	for {
		data := s.Connect(Token)
		fmt.Println(data)
		//s.Scrape(GID, CID, 0)
		fmt.Println("scraping...")
	}
	Return()
}

func Ping(message string, amount int, ID string) {

	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)

	var wg sync.WaitGroup
	wg.Add(len(Token))

	for i := 0; i < len(Token); i++ {
		go func(i int) {
			defer wg.Done()
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])
			}
			c.MassPing(Token[i], message, amount, ID)
		}(i)
	}
	wg.Wait()
	Return()

}

func Click(GID string, CID string, MID string, BotID string, Type int, Comp int, Text string) {
	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)

	var wg sync.WaitGroup
	wg.Add(len(Token))

	x, _ := strconv.Atoi(strconv.Itoa(Type))
	s, _ := strconv.Atoi(strconv.Itoa(Comp))
	if Type == x {
		Type = 3
	}
	if Comp == s {
		Comp = 2
	}
	for i := 0; i < len(Token); i++ {
		go func(i int) {
			defer wg.Done()
			c.Buttons(Token[i], GID, CID, MID, BotID, Type, Comp, Text)
		}(i)
	}
	wg.Wait()

	Return()
}

func main() {
	c.Cls()
	scn := bufio.NewScanner(os.Stdin)
	c.CheckConfig()
	fmt.Print(c.Logo())
	var choice int
	fmt.Scanln(&choice)
	if choice == 1 {
		fmt.Print("	[Message]>: ")
		scn.Scan()
		msg := scn.Text()
		MassDm(msg)
	} else if choice == 2 {
		var ID string
		fmt.Print("	[UserID]>: ")
		fmt.Scanln(&ID)
		fmt.Print("	[Message]>: ")
		scn.Scan()
		msg := scn.Text()
		Spam_Dm(ID, msg)
	} else if choice == 3 {
		var link string
		fmt.Print("	[Link]>: ")
		fmt.Scanln(&link)
		Reac(link)
	} else if choice == 4 {
		var invite string
		fmt.Print("	discord.gg/")
		fmt.Scanln(&invite)
		Join(invite)
	} else if choice == 5 {
		var ID string
		fmt.Print("	[ID]>: ")
		fmt.Scanln(&ID)
		Leave(ID)

	} else if choice == 6 {
		var invite, ID string
		fmt.Print("	discord.gg/")
		fmt.Scanln(&invite)
		fmt.Print("	[ServerID]>: ")
		fmt.Scanln(&ID)
		Rules(invite, ID)
	} else if choice == 7 {
		var ID string
		fmt.Print("	[Message]>: ")
		scn.Scan()
		msg := scn.Text()
		fmt.Print("	[ChannelID]>: ")
		fmt.Scanln(&ID)
		Raid(msg, ID)
	} else if choice == 8 {
		var token, GID, CID string
		fmt.Print("	[Token]>: ")
		fmt.Scanln(&token)
		fmt.Print("	[ServerID]>: ")
		fmt.Scanln(&GID)
		fmt.Print("	[ChannelID]>: ")
		fmt.Scanln(&CID)
		Scrape(token, GID, CID)
	} else if choice == 9 {
		Check()
	} else if choice == 10 {
		var (
			ID    string
			count int
		)
		fmt.Print("	[Message]>: ")
		scn.Scan()
		msg := scn.Text()
		fmt.Print("	[Ping Amount]>: ")
		fmt.Scanln(&count)
		fmt.Print("	[ChannelID]>: ")
		fmt.Scanln(&ID)
		Ping(msg, count, ID)
	} else if choice == 11 {
		var GID, CID, MID, BotID, Text string
		var Type, Comp int
		fmt.Print("	[ServerID]>: ")
		fmt.Scanln(&GID)
		fmt.Print("	[ChannelID]>: ")
		fmt.Scanln(&CID)
		fmt.Print("	[MessageID]>: ")
		fmt.Scanln(&MID)
		fmt.Print("	[BOT UserID]>: ")
		fmt.Scanln(&BotID)
		fmt.Print("	[Button Name]>: ")
		fmt.Scanln(&Text)
		fmt.Println("	\u001B[36mLeave Empty For Defualts\u001B[39m")
		fmt.Print("	[Button type INT]>: ")
		fmt.Scanln(&Type)
		fmt.Print("	[Component type INT]>: ")
		fmt.Scanln(&Comp)
		Click(GID, CID, MID, BotID, Type, Comp, Text)

	} else if choice == 12 {
		var user string
		fmt.Print("	[username#0000]>: ")
		fmt.Scanln(&user)
		Friend(user)
	} else {
		fmt.Println("[\u001B[31m~\u001B[39m]	Wrong Input")
		time.Sleep(1 * time.Second)
		main()
	}
}

func Return() {
	fmt.Println("\u001B[39mGoing Back to menu...")
	time.Sleep(3 * time.Second)
	main()
}
