package modules

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
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
		16: "SoundBoard",
		17: "OnBoarding",
		18: "Server Info",
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

	num := (len(opts) + 2 - 1) / 2

	for row := 0; row < num; row++ {
		for col := 0; col < 2; col++ {
			i := col*num + row
			if i < len(opts) {
				optnum := strconv.Itoa(opts[i])
				if len(optnum) == 1 {
					optnum = "0" + optnum
				}
				fmt.Fprintf(tw, "\t\033[36m[\033[39m%s\033[36m]\033[39m\t%s\t", optnum, options[opts[i]])
				count++
			}
		}
		fmt.Fprintln(tw)
	}
	if count%2 != 0 {
		fmt.Fprint(tw, "\t")
	}
	tw.Flush()

	fmt.Println(sb.String())
}

func PrintStruct(data interface{}) {
	v := reflect.ValueOf(data)

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fmt.Printf("\u001B[36m[\u001B[39m%s\u001B[36m]\u001B[39m: ",
			reflect.TypeOf(data).Field(i).Name,
		)
		switch f.Kind() {
		case reflect.Struct:
			PrintStruct(f.Interface())
		default:
			fmt.Println(f.Interface())
		}
	}
}

func (m *Modules) Cls() {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
