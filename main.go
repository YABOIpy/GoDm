package main


import (
	"fmt"
	"sync"
	"strconv"
	"strings"
	"massdm/src"
)

var (
	c = massdm.X()
)


func MassDm() {
	message := "nicker"
	
	Token, err := c.ReadFile("tokens.txt")
	c.Errs(err)
	ids, err := c.ReadFile("ids.txt")
	c.Errs(err)

	var wg sync.WaitGroup
	wg.Add(len(ids))
	//start := time.Now()
	for i := 0; i < len(ids); i++ {
		go func(i int) {
			defer wg.Done()
			for _, UserID := range ids {
				if c.Config().Settings.Websock == true {
					c.WebSock(strings.Join(Token, ""))							
				}
				ID,_ := strconv.Atoi(UserID)
				CID, err := c.Create(ID, strings.Join(Token, ""), message)
				c.Errs(err)

				if c.Config().Settings.Close == true {
					c.CloseDm(CID, strings.Join(Token, ""), massdm.Cookies)							
				} else if c.Config().Settings.Close == false {}
				if c.Config().Settings.Block == true {
					c.Block(ID, strings.Join(Token, ""), massdm.Cookies)
				}
			}
		}(i)
	}
	wg.Wait()
}




func main() {
	logo := `
	_______________     ______________  ___
	__  ____/_  __ \    ___  __ \__   |/  /
	_  / __ _  / / /    __  / / /_  /|_/ / 
	/ /_/ / / /_/ /     _  /_/ /_  /  / /  
	\____/  \____/      /_____/ /_/  /_/   
    
	[Go Discord Mass Dm]		~YABOI
	[1]
	[2]
	[3]
	[4]
	`
	fmt.Println(logo)
	MassDm()	
}
