package modules

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	fhttp "github.com/Danny-Dasilva/fhttp"
	"github.com/spf13/cast"
)

// LoadConfig Loads the json configuration file
func (m *Modules) LoadConfig(filename string) (Config, error) {
	var config Config
	conf, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer conf.Close()

	if err = json.NewDecoder(conf).Decode(&config); err != nil {
		return config, err
	}

	return config, nil
}

// Cookies Gets cookies off discord.com
func (m *Modules) Cookies() (cookie string) {
	resp, err := http.Get("https://discord.com")
	if err != nil {
		return m.Cookies()
	}
	defer resp.Body.Close()

	if resp.Cookies() != nil {
		for _, c := range resp.Cookies() {
			cookie += fmt.Sprintf("%s=%s; ", c.Name, c.Value)
		}
		return cookie
	} else {
		return m.Cookies()
	}
}

// MessageData Get Data from a message within a Channel returns type []MessageResp
func (in *Instance) MessageData(CID, MID string) []MessageResp {
	s := time.Now()
	req, err := fhttp.NewRequest("GET", fmt.Sprintf(
		"https://discord.com/api/v9/channels/%s/messages?limit=1&around=%s", CID, MID),
		nil,
	)
	if err != nil {
		log.Println(err)
		return nil
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZone,
		"x-super-properties": in.Xprop,
	})

	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return nil
	}

	defer resp.Body.Close()
	var data []MessageResp
	body, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
		modules.StrlogE("Failed To Get Message Data", string(body), s)
	}

	return data
}

func (in *Instance) FetchMessages(CID string, limit int) (data []ChannelMessages) {
	req, err := fhttp.NewRequest("GET",
		fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages?limit=%s", CID, cast.ToString(limit)),
		nil,
	)
	if err != nil {
		log.Println(err)
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZone,
		"x-super-properties": in.Xprop,
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}
	return data
}

func (in *Instance) OpenChannels() (data []Friend) {
	req, err := fhttp.NewRequest("GET",
		"https://discord.com/api/v9/users/@me/channels",
		nil,
	)
	if err != nil {
		return nil
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZone,
		"x-super-properties": in.Xprop,
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}

	switch resp.StatusCode {
	case 200:
		return data
	default:
		return nil
	}
}

func (in *Instance) UserInfo(ID string) (data UserInfo) {
	req, err := fhttp.NewRequest("GET",
		"https://discord.com/api/v9/users/"+ID,
		nil,
	)
	if err != nil {
		return data
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZone,
		"x-super-properties": in.Xprop,
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return data
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}
	return data
}

func (in *Instance) Friends() (data []Friend) {
	req, err := fhttp.NewRequest("GET",
		"https://discord.com/api/v9/users/@me/relationships",
		nil,
	)
	if err != nil {
		return nil
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZone,
		"x-super-properties": in.Xprop,
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}

	switch resp.StatusCode {
	case 200:
		return data
	default:
		return nil
	}
}

func (in *Instance) Guilds() []Guilds {
	req, err := fhttp.NewRequest("GET",
		"https://discord.com/api/v9/users/@me/guilds",
		nil,
	)
	if err != nil {
		return nil
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZone,
		"x-super-properties": in.Xprop,
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()
	var data []Guilds
	body, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}

	switch resp.StatusCode {
	case 200:
		return data
	default:
		return nil
	}
}
func (in *Instance) Guild(GID string) (data Server) {
	req, err := fhttp.NewRequest("GET",
		"https://discord.com/api/v9/guilds/"+GID,
		nil,
	)
	if err != nil {
		return
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua-platform": in.Quote(in.BrowserClient.OS),
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": in.TimeZone,
		"x-super-properties": in.Xprop,
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}

	return data
}

func (in *Instance) GuildJoinData(invite string) (data JoinResp) {
	req, err := fhttp.NewRequest("GET", fmt.Sprintf(
		"https://discord.com/api/v9/invites/%s?inputValue=%s&with_counts=true&with_expiration=true", invite, invite),
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua":          in.SecUA(in),
		"referer":            "https://discord.com/channels/@me",
		"x-discord-timezone": in.TimeZone,
		"x-super-properties": in.Xprop,
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}

	return data
}

func (in *Instance) OnboardingData(GID string) (data Onboarding) {
	req, err := fhttp.NewRequest("GET", fmt.Sprintf(
		"https://discord.com/api/v9/guilds/%s/onboarding", GID),
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua":          in.SecUA(in),
		"referer":            "https://discord.com/channels/" + GID + "/onboarding",
		"x-discord-timezone": in.TimeZone,
		"x-super-properties": in.Xprop,
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
	}

	return data
}

// Input Takes user input from the Command line
// and returns it as a string
func (m *Modules) Input(text string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

func (m *Modules) InputBool(text string) bool {
	if m.Input(text+" y/n: ") == "y" {
		return true
	}
	return false
}

// InputInt Takes user input from the Command line
// and returns it as an int
func (m *Modules) InputInt(text string) int {
	var d int

	fmt.Print(text + ": ")
	_, err := fmt.Scanln(&d)
	if err != nil {
		log.Println(err)
	}
	return d
}

func (m *Modules) ReadFile(path string) ([]string, []TokenConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	var (
		tkncfg []TokenConfig
		lines  []string
	)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if strings.Contains(path, "token") {
		tokens := make([]string, 0, len(lines))
		for _, line := range lines {
			if strings.Contains(line, ":") {
				f := strings.Split(line, ":")
				tokens = append(tokens, f[2])
				tkncfg = append(tkncfg, TokenConfig{
					Email: f[0],
					Pass:  f[1],
				})
			}
		}
		if len(tokens) > IntNil {
			lines = tokens
		}
	}

	return lines, tkncfg, scanner.Err()
}

func (m *Modules) ReadDirectory(path, end string) (f []string) {
	Files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range Files {
		if !file.IsDir() {
			name := filepath.Join(path, file.Name())
			if strings.Contains(name, end) {
				Data, er := os.ReadFile(name)
				if er != nil {
					log.Printf("Error reading %s: %s\n", name, err)
					continue
				}

				base64Data := base64.StdEncoding.EncodeToString(Data)
				f = append(f, base64Data)
			}
		}
	}

	return f
}

func (in *Instance) Quote(data string) string {
	return fmt.Sprintf(`"%s"`, data)
}

// TrimZero Trims the first 0 of a string that is found
// 0010 = 10, 1000 = 1000
func (m *Modules) TrimZero(discrim string) string {
	if len(discrim) > 0 && discrim[0] == '0' {
		return strings.TrimLeft(discrim, "0")
	}
	return discrim
}

func (m *Modules) WriteFile(files, item string) {
	f, err := os.OpenFile(files, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	_, ers := f.WriteString(item + "\n")
	if ers != nil {
		log.Println(ers)
		return
	}
}

// WriteFileArray this is for writing large amounts into files fast
// instead of opening the file each time you write using WriteFile
func (m *Modules) WriteFileArray(files string, item []string) {
	f, err := os.OpenFile(files, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	for i := 0; i < len(item); i++ {
		f.WriteString(item[i] + "\n")
	}
}

// FilterArray Filters an array of unwanted stuff like blank lines and duplicates
func (m *Modules) FilterArray(data []string) []string {
	lines := make(map[string]bool)
	var d []string
	for _, line := range data {
		l := strings.TrimSpace(line)
		if l != "" && !lines[l] {
			lines[l] = true
			d = append(d, l)
		}
	}
	return d

}

// FetchInputData read thefiles needed for launching GoDm
func (m *Modules) FetchInputData() ([]string, []TokenConfig, []string) {
	t, d, err := m.ReadFile("tokens.txt")
	if err != nil {
		log.Println("Failed To Load Tokens.txt")
	}
	p, _, err := m.ReadFile("proxies.txt")
	if err != nil {
		log.Println("Failed To Load Proxies.txt")
	}
	return t, d, p

}
func (m *Modules) Sleep(t time.Duration, in *Instance) bool {
	if in.Cfg.Mode.Configs.RateLimit {
		time.Sleep(t * time.Second)
		return true
	}
	return false
}

func (m *Modules) Marsh(v any) *bytes.Buffer {
	data, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		return nil
	}
	return bytes.NewBuffer(data)
}

// Nonce Returns discord nonce value
func (m *Modules) Nonce() string {
	epoch := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1000000
	now := time.Now().UTC().UnixNano() / 1000000
	return strconv.FormatInt((now-epoch)<<22, 10)
}

// HalfToken Splits Tokens using the "." separator
func (m *Modules) HalfToken(T string, v int) string {
	return strings.Split(T, ".")[v]
}

func RandSeed() *CCSeed {
	return &CCSeed{}
}

// GenerateSeed Generates a Random Seed
// using rand.Seed will be deprecated and isn't safe for concurrency
func (c *CCSeed) GenerateSeed() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.seed = time.Now().UnixNano()
	return c.seed
}
func (in *Instance) SecUA(d *Instance) string {
	return fmt.Sprintf(`"Not?A_Brand";v="8", "%s";v="%s"`, d.BrowserClient.Name, d.BrowserClient.Version)
}

// Xprops Builds Discord X-Super-Properties with the given data,
// data will be created using the CreateBrowser function
func (in *Instance) Xprops(data ClientData) string {

	d, err := json.Marshal(XpropData{
		OS:                 data.OS,
		Browser:            data.Name,
		Device:             "",
		SystemLocale:       "en-US",
		BrowserUserAgent:   data.Agent,
		BrowserVersion:     data.Version,
		OSVersion:          data.OSVer,
		Referrer:           "https://www.google.com/",
		ReferringDomain:    "www.google.com",
		SearchEngine:       "google",
		ReferrerCurrent:    "",
		ReferringDomainCur: "",
		ReleaseChannel:     "stable",
		ClientBuildNumber:  cbn,
		ClientEventSource:  nil,
		DesignID:           0,
	})
	if err != nil {
		log.Println(err)
	}
	return base64.StdEncoding.EncodeToString(d)
}

func CreateRange(i int) (v []interface{}) {
	switch i {
	case 0:
		v = []interface{}{[2]int{0, 99}}
	case 1:
		v = []interface{}{[2]int{0, 99}, [2]int{100, 199}}
	default:
		v = []interface{}{[2]int{0, 99}, [2]int{100, 199}, [2]int{i * 100, (i * 100) + 99}}
	}
	return v
}

func RGB(red, green, blue int) int {
	return (red * 256 * 256) + (green * 256) + blue
}

// GetRandomBrowser Returns a random browser from Browser
func (in *Instance) GetRandomBrowser(browsers []BrowserData) BrowserData {
	return browsers[rand.Intn(len(browsers))]
}

func (in *Instance) GetRandomData(data []string) string {
	RSeed.GenerateSeed()
	return data[rand.Intn(len(data))]
}

func (in *Instance) getRandomKey(m map[string][]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	RSeed.GenerateSeed()
	return keys[rand.Intn(len(keys))]
}

func ReturnRandomArray(v []string, num int) []string {
	if num >= len(v) {
		return v
	}
	RSeed.GenerateSeed()
	rand.Shuffle(len(v), func(i, j int) {
		v[i], v[j] = v[j], v[i]
	})

	return v[:num]
}

func (m *Modules) RandString(length int) string {
	RSeed.GenerateSeed()
	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomString[i] = CharSet[rand.Intn(len(CharSet))]
	}
	return string(randomString)
}

func (m *Modules) Contains(data []string, v string) bool {
	for _, k := range data {
		if k == v {
			return true
		}
	}
	return false
}
