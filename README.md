# GoDm
<img align="right" width="159px" src="https://github.com/YABOIpy/GoDm/assets/110062350/8ada5a13-f664-470e-9b8e-fe5473ea9d44">

coming soon star for updates!
# Hang in Tight! the new update will be here soon and its faster than ever.  1/08/2023

# OverView
- [Introduction](#Usage)
- [How](#How)
- [Configuration](#Config)
- [Donate](#Donations)

# About


<p align="center" style="text-align: center"> 
  <img src="https://sloc.xyz/github/yaboipy/godm/">
  <img src="https://img.shields.io/github/downloads/yaboipy/godm/total?color=green&label=Release Downloads">
</p>

<p align="center" style="text-align: center"> 
  <img width="900" alt="image" src="https://github.com/YABOIpy/GoDm/assets/110062350/3c0c725d-f002-4b31-be46-2fa7ba148a04">
</p>


# How
What Happens Behind GoDm
![image](https://github.com/YABOIpy/GoDm/assets/110062350/7b096cf5-9b88-46eb-8b1f-fd785c8cb2a7)


# Usage
golang can be installed at: https://go.dev
version 1.20+

Go Not installed? 
Download Compiled Version:

https://github.com/YABOIpy/Go-MassDM/releases
```md
Needs & Inputs:

> tokens can be in mail:pass:token format or only token format
> just make sure theyre inside tokens.txt and the separator is ":"
________________________________________________
  [0] Refreshes GoDm
  [1] Message and Scraped Ids in ids.txt
  [2] victims User ID
  [3] Channel ID & Message ID
  [4] Server invite
  [5] Server ID
  [6] Server Invite and Server ID
  [7] Channel ID and Message
  [8] Server ID and Channel ID
  [9] Tokens in tokens.txt
  [10] Channeld ID And Amount To ping in a single message 
  [11] Server ID Channel ID Message ID and Bot user ID
  [12] Full Username / USER#0000
  [13] email:pass:token format / 6 options
  [14] Nitro Tokens in tokens.txt and server ID
________________________________________________

Recommended:
> use 0 - 3000 tokens for stability
> above 20 mbps netspeed dependent on amount of tokens

GoDm was made to be fast so even small amounts are enough for significant speeds
```

# Config


Proxy Type: Residential rotating


Proxies available at:

- https://iproyal.com
- https://proxiware.com
- https://proxies.gg


Format: user:pass@ip:port

if using residential proxies
just put it in the config.json

if using proxies from proxies.txt
use the same amount of tokens as proxies
within the txts
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
      "Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9006 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
      "Proxy": "your-proxy-address",
      "WebKit": "537.36",
      "Redirect": false,
      "TimeOut": 50
    },
    "Discord": {
      "Ver": 2.0,
      "CapApi": ["captcha-service", "service-api-key"],
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

**Discord**
| Name | Type | Description | 
| ---  | ---  | ---         |
| `CapApi` | CAPTCHA | Supported Captcha Service & Api Captcha solver
| `Ver` | CONFIG | GoDm Client Version
| `Message` | MASSDM | Array of messages that will be randomly chosen and sent


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

</p>

# Donations

```md
ETH: 0xf96F04521F59dDcEc404d3A90Bdf91715D202a06
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
> note the config u had 
```

# TODO
```
dynamic Ja3 Fingerprinting
more captcha api
```

# Previews:



this is for educational use ONLY 😉
