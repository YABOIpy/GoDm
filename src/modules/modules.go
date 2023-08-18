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
	"github.com/wasilibs/go-re2"
)

// LoadConfig Loads the json configuration file
func (m *Modules) LoadConfig(filename string) (Config, error) {
	var config Config
	conf, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer conf.Close()

	err = json.NewDecoder(conf).Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// Cookies Gets cookies off discord.com
func (m *Modules) Cookies() (string, Config) {
	resp, err := http.Get("https://discord.com")
	if err != nil {
		return m.Cookies()
	}
	defer resp.Body.Close()
	cookieData := CookieData{
		Cookies: make(map[string]*http.Cookie),
	}

	if resp.Cookies() != nil {
		for _, cookie := range resp.Cookies() {
			switch cookie.Name {
			case "__dcfduid":
				cookieData.Cookies[cookie.Name] = cookie
				m.Dcfd = cookie.Value
			case "__sdcfduid":
				cookieData.Cookies[cookie.Name] = cookie
				m.Sdcfd = cookie.Value
			case "__cfruid":
				cookieData.Cookies[cookie.Name] = cookie
				m.Cfruid = cookie.Value
			}
		}
	} else {
		return m.Cookies()
	}

	return fmt.Sprintf("__dcfduid=%s; __sdcfduid=%s; __cfruid=%s; locale=us", m.Dcfd, m.Sdcfd, m.Cfruid), Config{
		Dcfd: m.Dcfd, Sdcfd: m.Sdcfd, Cfruid: m.Cfruid, Cookie: cookieData,
	}
}

// MessageData Get Data from a message within a Channel returns type []MessageResp
func (m *Modules) MessageData(in Instance, CID string, MID string) []MessageResp {
	req, err := fhttp.NewRequest("GET", fmt.Sprintf(
		"https://discord.com/api/v9/channels/%s/messages?limit=1&around=%s", CID, MID),
		nil,
	)
	if err != nil {
		log.Println(err)
		return []MessageResp{}
	}
	Hd.Header(req, map[string]string{
		"authorization":      in.Token,
		"cookie":             in.Cookie,
		"user-agent":         in.BrowserClient.Agent,
		"sec-ch-ua":          in.SecUA(in),
		"x-discord-timezone": "Europe/Amsterdam",
		"x-super-properties": in.Xprop,
	})

	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return []MessageResp{}
	}

	defer resp.Body.Close()
	var data []MessageResp
	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
	}
	return data
}

// TODO
func (m *Modules) OnboardingData(in Instance, GID string) {
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
		"x-discord-timezone": "Europe/Amsterdam",
		"x-super-properties": in.Xprop,
	})

	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	//err = json.Unmarshal(body, &data)
	//if err != nil {
	//	log.Println(err)
	//}
	return
}

// Input Takes user input from the Command line
// and returns it as a string
func (m *Modules) Input(text string) string {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print(text)
	reader.Scan()
	return reader.Text()
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
// TODO: turn tokenconfig into an array...
func (m *Modules) ReadFile(path string) ([]string, TokenConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, TokenConfig{}, err
	}
	defer file.Close()
	var (
		mail, pass string
		lines      []string
	)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if strings.Contains(path, "token") {
		tokens := make([]string, 0, len(lines))
		for _, line := range lines {
			if strings.Contains(line, ":") {
				format := strings.Split(line, ":")
				tokens = append(tokens, format[2])
				mail, pass = format[0], format[1]
			}
		}
		if len(tokens) > 0 {
			lines = tokens
		}
	}

	return lines, TokenConfig{Email: mail, Pass: pass}, scanner.Err()
}

func (m *Modules) ReadDirectory(path string, end string) (f []string) {
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

func (m *Modules) WriteFile(files string, item string) {
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
func (m *Modules) FetchInputData() ([]string, TokenConfig, []string) {
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
func (m *Modules) Sleep(t time.Duration, in Instance) bool {
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

// BuildInfo Get discords current build number
func BuildInfo() string {
	buildNumber := make(map[string]string)

	js := re2.MustCompile(`([a-zA-Z0-9]+)\.js`)
	build := re2.MustCompilePOSIX(`Build Number: "\)\.concat\("([0-9]{4,8})"`)

	Client := &http.Client{}

	res, err := Client.Get("https://discord.com/app")
	if err != nil {
		log.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return ""
	}

	rs := js.FindAllString(string(body), -1)
	asset := rs[len(rs)-1]
	if strings.Contains(asset, "invisible") {
		asset = rs[len(rs)-2]
	}

	resp, err := Client.Get("https://discord.com/assets/" + asset)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}

	buildInfos := strings.Split(
		strings.ReplaceAll(build.FindAllString(string(b), -1)[0], " ", ""),
		",",
	)
	bn := strings.Split(buildInfos[0], `("`)
	buildNumber["stable"] = strings.ReplaceAll(bn[len(bn)-1], `"`, ``)

	return buildNumber["stable"]
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
func (in *Instance) SecUA(d Instance) string {
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
	})
	if err != nil {
		log.Println(err)
	}
	return base64.StdEncoding.EncodeToString(d)
}

func Content[T map[string]string | map[string]interface{}](data T) string {
	d, _ := json.Marshal(data)
	return strconv.Itoa(len(d))
}

// GetRandomBrowser Returns a random browser from Browser
func (in *Instance) GetRandomBrowser(browsers []BrowserData) BrowserData {
	return browsers[rand.Intn(len(browsers))]
}

// could've used generics for both
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
	for _, i := range data {
		if i == v {
			return true
		}
	}
	return false
}
