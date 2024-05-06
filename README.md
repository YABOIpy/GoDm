<img align="right" width="159px" src="https://github.com/YABOIpy/GoDm/assets/110062350/8ada5a13-f664-470e-9b8e-fe5473ea9d44">


# OverView
- [Introduction](#Usage)
- [How](#How)
- [Configuration](#Config)
- [Donate](#Donations)

## being recoded. probably wont be maintained after.
# Join the Chat!: https://t.me/+ZsFhZWu8ZJMyZmNk

# About


<p align="center" style="text-align: center"> 
  <img src="https://sloc.xyz/github/yaboipy/godm/">
</p>

<p align="center" style="text-align: center"> 
  <img width="1200" alt="image" src="https://github.com/YABOIpy/GoDm/assets/110062350/e5d31ea6-7f1c-40bc-b5a8-2ec34528ce39">
 
</p>


# How
What Happens Behind GoDm
![image](https://github.com/YABOIpy/GoDm/assets/110062350/7b096cf5-9b88-46eb-8b1f-fd785c8cb2a7)


# Usage
golang can be installed at: https://go.dev
version 1.20+

Go Not installed? 
Download Compiled Version:

https://github.com/YABOIpy/GoDm/releases
```md
Inputs:
________________________________________________
  [0] Refreshes GoDm
  [1] Message & Scraped Ids in ids.txt
  [2] victims User ID
  [3] Channel ID & Message ID
  [4] Server invite
  [5] Server ID
  [6] Server Invite and Server ID
  [7] Channel ID & Message
  [8] Server ID & Channel ID
  [9] Tokens in tokens.txt
  [10] Channeld ID & Amount To ping in a single message 
  [11] Message Link & Button Type https://discord.com/developers/docs/interactions/message-components
  [12] Full Username / USER#0000
  [13] email:pass:token format / 10 options
  [14] Nitro Tokens in tokens.txt & server ID
  [15] Server ID & CHannel ID
  [16] SoundBoard Option & Channel ID
  [17] Server Invite & Options
  [18] Server Invite
________________________________________________

time logging: the logger will return in ms so 500ms, but if its 5.23ms than its 5s & 230ms
before scraping, have a atleast a single token within the list

> tokens can be in mail:pass:token format or only token format
> just make sure theyre inside tokens.txt and the separator is ":"
> also do not mix the formats together


Recommended:
> use 1 - 3000 tokens for stability
> above 20 mbps netspeed dependent on amount of tokens
________________________________________________
> godm creates a random context for each token, its not advised to restart godm after different functions.
> as it will automatically return to the menu


GoDm was made to be fast so even small amounts are enough for significant speeds
```

# Config


Proxy Type: Residential rotating / static list


Proxies available at:

- https://iproyal.com
- https://proxiware.com
- https://proxies.gg


Format: user:pass@ip:port

if using residential proxies
just put it in the config.json

if using proxies from proxies.txt
use the same or more amount of proxies as tokens 
within the file
_____________
Recommended Config:
```json
{
  "Modes": {
    "Config": {
      "Interval": 0,
      "CCManager": false,
      "MaxRoutines": 300,
      "SolveCaptcha": false,
      "CaptchaRetry": 2,
      "RateLimit": true
    },
    "Net": {
      "JA3": "772,4865-4866-49195-49199-49196-49200-52393-49171-49172-156-157-47-53,35-51-13-17513-45-10-65281-0-11-27-16-23-18-5-43,29-23-24,0",
      "Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.366",
      "Proxy": "your-proxy-address",
      "WebKit": "537.36",
      "Redirect": false,
      "TimeOut": 50
    },
    "Discord": {
      "Ver": 2.0,
      "CapApi": ["captcha-service", "service-api-key"], <- no captcha support until recode.
      "Presence": ["online", "dnd", "idle"], 
      "Message": [
        {
          "Title": "# Hey!",
          "Body": "> This Message Was sent using",
          "Link": "> https://github.com/yaboipy"
        },
        {
          "Title": "# Hello!",
          "Body": "> mind joining my server!",
          "Link": "> discord.gg/"
        },
      ]
    }
  }
}

```

**Config**
| Name | Type | Description | 
| ---  | ---  | ---         |
| `Interval` | CONFIG | Global Time-[sec] before iterating 0 = instant iteration
| `SolveCaptcha` | CAPTCHA | Solves The Capthca Using Api
| `CCManager` | PROCESS | False will have no limit to MaxRountines
| `MaxRoutines` | PROCESS | The Max ammount of Concurrent WaitGroups allowed to run
| `CaptchaRetry` | CONFIG | The Max ammount of times a captcha is re-solved 
| `RateLimit` | ANTI | Will Safely stop for if it encounnters a ratelimit and go after its over

**Client**
| Name | Type | Description | 
| ---  | ---  | ---         |
| `JA3` | TLS | TLS Fingerprint can be left as is
| `Proxy` | HTTP | Your Proxy address Format: username:password@hostname:port
| `Redirect` | HTTP | specifies the policy for handling redirects
| `Agent` | HEADER | UserAgent To fall back on
| `WebKit` | HEADER | The webkit Used for the Useragents
| `TimeOut` | HTTP | Time-[sec] after request with no response, allowed 0 = no timeout

**Discord**
| Name | Type | Description | 
| ---  | ---  | ---         |
| `Ver` | CONFIG | GoDm Client Version
| `CapApi` | CAPTCHA | Supported Captcha Service & Api Captcha solver
| `Presence` | CONFIG | Array of presence's the tokens have.
| `Message` | MASSDM | Array of messages that will be randomly chosen and sent

</p>

# Donations

```md
ETH: 0xf96F04521F59dDcEc404d3A90Bdf91715D202a06
BTC: bc1q5lunjjahjemql8mfyszpjfcp4cwu5u4vgu69jf
BTC: bc1qk9pdv82jd9zletczuw08jhg9ly4gn7y9g55dfq
LTC: LLMauaJr69njn1wug4E169oqFsFTDKSirq
XMR: 46TwyYsGQCqUREXJa2jgKTb6awTaftkSmPvh3aoz5se1MokrH38UvN7Co1doJ4uhLc3MeEbTEe5evMu6z5oTMbra4Hzjgc6
PYP: https://paypal.me/yaboipy  
```
Thanks Alot to the users who use GoDm

i appreciate it and enjoy making these projects for you to use
________________________

# Issues
```md
> Specify which Function
> Show a Screenshot
> note the inputs & config u had 
```

# TODO
```
Captcha Support
fix file structure
Clean Some code / nesting
```
# Sources
```
https://github.com/uber-go/guide
https://github.com/Danny-Dasilva/CycleTLS
https://github.com/V4NSH4J/discord-mass-DM-GO
```
# Previews:
soon.


this is for educational use ONLY 😉
