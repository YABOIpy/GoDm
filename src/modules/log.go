package modules

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

func (m *Modules) Menu() {
	opt := map[int]string{
		1:  "Mass DM",
		2:  "Dm Spam",
		3:  "React Verify",
		4:  "Joiner",
		5:  "Leaver",
		6:  "Accept Rules",
		7:  "Raid Channel",
		8:  "Scrape Users",
		9:  "Check Tokens",
		10: "Mass Ping",
		11: "Button Click",
		12: "Friender",
		13: "Token Menu",
		14: "Booster",
		15: "VoiceChat",
		16: "Onboarding",
	}
	tkn, _, _ := modules.ReadFile("tokens.txt")
	//old logo. not going to bother formatting
	fmt.Println(`
	____`+clr+`_____`+r+`__`+clr+`____     `+r+`____`+clr+`____`+r+`____`+clr+`__  ___
	`+r+`__`+clr+`  ____/`+r+`_  `+clr+`__ \    `+r+`___`+clr+`  __ \`+r+`__`+clr+`   |/  /
	`+r+`_ `+clr+` / __ `+r+`_`+clr+`  / / /    `+r+`__`+clr+`  / / /`+r+`_`+clr+`  /|_/ / 
	`+clr+`/ /_/ / / /_/ /     `+r+`_  `+clr+`/_/ /`+r+`_`+clr+`  /  / /  
	\____/  \____/      /_____/ /_/  /_/

  `+clr+`[`+r+`Tokens: `+g+strconv.Itoa(len(tkn))+r+clr+`]_____________________________`, r)
	m.PrintMenu(opt)
}

func (m *Modules) StrlogV(text string, data string, s time.Time) {
	e := time.Since(s)
	fmt.Printf("[%s%sms%s] [%s✓%s]%s: %s%s%s\n", bg, e.String()[:3], rb, g, r, text, gr, data, rb+r)
}

func (m *Modules) StrlogE(text string, data string, s time.Time) {
	e := time.Since(s)
	fmt.Printf("[%s%sms%s] [%sX%s]%s: %s%s%s\n", bg, e.String()[:3], rb, red, r, text, gr, data, rb+r)
}

func (m *Modules) StrlogR(text string, data string, s time.Time) {
	e := time.Since(s)
	fmt.Printf("[%s%sms%s] [%s-%s]%s: %s%s%s\n", bg, e.String()[:3], rb, yellow, r, text, gr, data, rb+r)

}

func (m *Modules) PrintMenu(options map[int]string) {
	var sb strings.Builder
	var count int

	tw := tabwriter.NewWriter(&sb, 0, 0, 2, ' ', 0)
	opts := make([]int, 0, len(options))
	for k := range options {
		opts = append(opts, k)
	}
	sort.Ints(opts)

	numO := len(opts)
	numC := 2
	numR := (numO + numC - 1) / numC

	for row := 0; row < numR; row++ {
		for col := 0; col < numC; col++ {
			i := col*numR + row
			if i < numO {
				count++
				optnum := strconv.Itoa(opts[i])
				if len(optnum) == 1 {
					optnum = "0" + optnum
				}
				fmt.Fprintf(tw, "\t\033[36m[\033[39m%s\033[36m]\033[39m\t%s\t", optnum, options[opts[i]])
			}
		}
		fmt.Fprintln(tw)
	}
	if count%numC != 0 {
		fmt.Fprint(tw, "\t")
	}
	tw.Flush()

	fmt.Println(sb.String())
}

func (m *Modules) Cls() {
	var clearCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		clearCmd = exec.Command("cmd", "/c", "cls")
	} else {
		clearCmd = exec.Command("clear")
	}

	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}
