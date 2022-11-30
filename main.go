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
				ID,_ := strconv.Atoi(UserID)
				c.Create(ID, strings.Join(Token, ""), message)
				if c.Config().Settings.Close == true {
					//									
				} 
				if c.Config().Settings.Block == true {
					//
				}
			}
		}(i)
	}
	wg.Wait()
}




func main() {
	logo := `
	
	`
	fmt.Println(logo)
	MassDm()	
}