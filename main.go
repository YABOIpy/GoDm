package main


import (
	"fmt"
	"sync"
	"time"
	"strconv"
	"massdm/src"
)

var (
	c = massdm.X()
)


func MassDm(message string) {

	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)
	ids, err := c.ReadFile("ids.txt")
	c.Errs(err)

	var wg sync.WaitGroup
	wg.Add(len(ids))
	//start := time.Now()
	for i := 0; i < len(Token); i++ {
		go func(i int) {
			defer wg.Done()
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])							
			}
			for _, UserID := range ids {
				ID,_ := strconv.Atoi(UserID)
				CID, err := c.Create(ID, Token[i], message)
				c.Errs(err)

				if c.Config().Settings.Close == true {
					c.CloseDm(CID, Token[i], massdm.Cookies)							
				} else if c.Config().Settings.Close == false {}
				if c.Config().Settings.Block == true {
					c.Block(ID, Token[i], massdm.Cookies)
				}
			}
		}(i)
	}
	wg.Wait()
	
}



func Spam_Dm(UserID string, message string) {
	
	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)

	var wg sync.WaitGroup
	wg.Add(len(Token))
	
	ID,_ := strconv.Atoi(UserID)
	for i := 0; i < len(Token); i++ {
		go func(i int) {
			defer wg.Done()
			if c.Config().Settings.Websock == true {
				c.WebSock(Token[i])							
			}
			CID,err := c.Create(ID, Token[i], message)
			c.Errs(err)
			for true {
				c.Dm_Spam(CID, Token[i], message)
			}
		}(i)
	}
	wg.Wait()
}



func Join(invite string) {

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
			c.Joiner(Token[i], invite)
		}(i)
	}
	wg.Wait()
}




func main() {
	c.Cls()
	fmt.Print(massdm.Logo)
	var choice int
	fmt.Scanln(&choice)
	if choice == 1 {
		var msg string
		fmt.Print("[Message]>: ")
		MassDm(msg)
		fmt.Scanln(&choice)
	} else if choice == 2 {
		var msg, ID string
		fmt.Print("	[UserID]>: ")
		fmt.Scanln(&ID)
		fmt.Print("	[Message]>: ")
		fmt.Scanln(&msg)
		Spam_Dm(ID, msg)
	} else if choice == 3 {

	} else if choice == 4 {
		var invite string
		fmt.Print("	discord.gg/")
		fmt.Scanln(&invite)
		Join(invite)
	} else if choice == 5 {

	} else if choice == 6 {

	} else {
		fmt.Println("	Wrong Input")
		time.Sleep(1 *time.Second)
		main()
	}		
}
