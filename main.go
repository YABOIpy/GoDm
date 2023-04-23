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
	c  = massdm.X()
	s  = Scraper.X()
	wg = goccm.New(300)
)

func MassDm(message string) {

	var wg goccm.ConcurrencyManager
	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)
	ids, err := c.ReadFile("ids.txt")
	c.Errs(err)

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
					c.CloseDm(CID, Token[i], massdm.Cookie)
				}
				if c.Config().Settings.Block == true {
					c.Block(ID, Token[i], massdm.Cookie)
				}
			}
		}(i)
	}
	wg.WaitAllDone()
	Return()

}

func Spam_Dm(UserID string, message string) {

	var wg goccm.ConcurrencyManager
	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)

	ID, _ := strconv.Atoi(UserID)
	for i := 0; i < len(Token); i++ {
		wg.Wait()
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
	wg.WaitAllDone()
}

func Join(invite string) {

	var wg goccm.ConcurrencyManager
	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)
	interval := c.Config().Mode.Interval.Intjoiner
	if interval > 0 {
		for i := 0; i < len(Token); i++ {
			time.Sleep(time.Duration(interval) * time.Second)
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])
			}
			c.Joiner(Token[i], invite, "", "", 0)

		}
		Return()
	} else {

		for i := 0; i < len(Token); i++ {
			wg.Wait()
			go func(i int) {
				defer wg.Done()
				if c.Config().Settings.Websock == true {
					c.WebSock(Token[i])
				}
				c.Joiner(Token[i], invite, "", "", 0)
			}(i)
		}
		wg.WaitAllDone()
		Return()
	}
}

func Leave(ID string) {

	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)

	var wg goccm.ConcurrencyManager

	for i := 0; i < len(Token); i++ {
		go func(i int) {
			defer wg.Done()
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])
			}
			c.Leaver(Token[i], ID)
		}(i)
	}
	wg.WaitAllDone()
	Return()
}

func Check() {

	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)
	f := [3]string{
		"data/valid.txt",
		"data/locked.txt",
		"data/invalid.txt",
	}
	for i := 0; i < len(f); i++ {
		os.Truncate(f[i], 0)
	}
	var (
		grn = "\033[32m"
		yel = "\033[33m"
		red = "\033[31m"
		r   = "\033[39m"
	)

	var wg sync.WaitGroup
	wg.Add(len(Token))

	start := time.Now()
	for i := 0; i < len(Token); i++ {
		go func(i int) {
			defer wg.Done()
			r := c.Check(Token[i])
			switch r {
			case 200:
				c.WriteFile("data/valid.txt", Token[i])
			case 403:
				c.WriteFile("data/locked.txt", Token[i])
			default:
				c.WriteFile("data/invalid.txt", Token[i])
			}

		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("["+grn+"✓"+r+"] (TIME):", elapsed.String()[:4]+"Ms", "("+yel+"LOCKED"+r+"):", c.Checker.Locked, "("+red+"INVALID"+r+"):", c.Checker.Invalid, "("+grn+"VALID"+r+"):", c.Checker.Valid, "(TOTAL):", c.Checker.All)
	c.Checker.All = 0
	c.Checker.Locked = 0
	c.Checker.Valid = 0
	c.Checker.Invalid = 0
	Return()
}

func Reac(link string) {
	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)

	var wg goccm.ConcurrencyManager

	for i := 0; i < len(Token); i++ {
		wg.Wait()
		go func(i int) {
			defer wg.Done()
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])
			}
			c.React(Token[i], link)
		}(i)
	}
	wg.WaitAllDone()
	Return()
}

func Rules(invite string, ID string) {

	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)

	var wg goccm.ConcurrencyManager

	for i := 0; i < len(Token); i++ {
		wg.Wait()
		go func(i int) {
			defer wg.Done()
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])
			}
			c.Agree(Token[i], invite, ID)
		}(i)
	}
	wg.WaitAllDone()
	Return()
}

func Raid(message string, ID string) {

	var wg goccm.ConcurrencyManager
	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)

	for i := 0; i < len(Token); i++ {
		wg.Wait()
		go func(i int) {
			defer wg.Done()
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])
			}
			c.Raider(Token[i], message, ID)
		}(i)
	}
	wg.WaitAllDone()
}

func Friend(user string) {

	username := strings.Split(user, "#")
	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)

	var wg goccm.ConcurrencyManager

	for i := 0; i < len(Token); i++ {
		wg.Wait()
		go func(i int) {
			defer wg.Done()
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])
			}
			c.Friend(Token[i], username[0], username[1])
		}(i)
	}
	wg.WaitAllDone()
	Return()
}

func Scrape(Token string, GID string, CID string) {
	for {
		resp, rep := s.Connect(Token)
		fmt.Println(resp.Data, string(rep))
		//s.Scrape(resp, GID, CID, 0)
		//fmt.Println(data)
		//s.Scrape(GID, CID, 0)
		fmt.Println("scraping...")
	}
	Return()
}

func Ping(message string, amount int, ID string) {

	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)

	var wg goccm.ConcurrencyManager

	for i := 0; i < len(Token); i++ {
		wg.Wait()
		go func(i int) {
			defer wg.Done()
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])
			}
			c.MassPing(Token[i], message, amount, ID)
		}(i)
	}
	wg.WaitAllDone()
	Return()

}

func Click(GID string, CID string, MID string, BotID string, Type int, Comp int, Text string) {
	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)

	var wg goccm.ConcurrencyManager

	x, _ := strconv.Atoi(strconv.Itoa(Type))
	s, _ := strconv.Atoi(strconv.Itoa(Comp))
	if Type == x {
		Type = 3
	}
	if Comp == s {
		Comp = 2
	}
	for i := 0; i < len(Token); i++ {
		wg.Wait()
		go func(i int) {
			defer wg.Done()
			c.Buttons(Token[i], GID, CID, MID, BotID, Type, Comp, Text)
		}(i)
	}
	wg.WaitAllDone()

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
