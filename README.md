# GoDm
<img align="right" width="159px" src="https://github.com/YABOIpy/GoDm/assets/110062350/8ada5a13-f664-470e-9b8e-fe5473ea9d44">

coming soon star for updates!

# Hang in Tight! the new update will be here soon and its faster than ever.

this will release soon the code is just being uploaded here
the functions are still in development
# About

<p align="center" style="text-align: center"> 
  <img src="https://sloc.xyz/github/yaboipy/godm/">
  <img src="https://img.shields.io/github/downloads/yaboipy/godm/total?color=green&label=Release Downloads">
</p>

<p align="center" style="text-align: center"> 
  <img width="900" alt="image" src="https://github.com/YABOIpy/GoDm/assets/110062350/3c0c725d-f002-4b31-be46-2fa7ba148a04">
</p>
# Issues
```md
> Specify which Function
> Show a Screenshot
> note the config u had 
```

# Usage

Go Not installed? 
Download Compiled Version:

https://github.com/YABOIpy/Go-MassDM/releases
```md
Needs & Inputs:

> tokens can be in mail:pass:token format or only token format
> just make sure theyre inside tokens.txt and the separator is ":"
________________________________________________
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
  [13] email:pass:token format
  [14] Nitro Tokens in tokens.txt and server ID
________________________________________________

```

# Config


Proxy Type: Residential rotating


Proxies available at:

- https://iproyal.com
- https://proxiware.com
- https://proxies.gg
- 
Format: user:pass@ip:port

_____________
Recommended Config:
```json
{
  "Modes": {
    "Discord": {
      "Ver": 2.0,
      "CapApi": ["your-captcha-service", "your-captcha-api-key"]
    },
    "Config": {
      "Interval": 0,
      "CCManager": false,
      "MaxRoutines": 300,
      "SolveCaptcha": true,
      "RateLimit": true
    },
    "Net": {
      "JA3": "771,4866-4867-4865-52393-49188-49199-158-49191-49200-49192-107-159-52392-49195-103-49196-49187-255,0-11-10-35-16-22-23-13-43-45-51-21,29-23-30-25-24,0-1-2",
      "Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9006 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
      "Proxy": "your-proxy-url",
      "WebKit": "537.36",
      "Redirect": false,
      "TimeOut": 50
    }
  }
}

```

**Discord**
| Name | Type | Description | 
| ---  | ---  | ---         |
| `CfBm` | HEADER | Discord Clourdflare Cookie 
| `Service` | CAPTCHA | Captcha Service for Api Captcha solver
| `Api_Key` | CAPTCHA | Api Key for the Captcha Service
| `Ver` | CONFIG | GoDm Client Version
| `AppID` | CONFIG | Discord Application ID 
| `Presence` | CONFIG | Discord RPC Client


**Config**
| Name | Type | Description | 
| ---  | ---  | ---         |
| `Interval` | CONFIG | The Time in Seconds before iterating 0 = instant iteration
| `Errors` | ERRORS | Shows you the error from the request
| `SolveCaptcha` | CAPTCHA | Solves The Capthca Using Api
| `ErrorLog` | ERRORS | Logs Errors inside errors.txt

**Client**
| Name | Type | Description | 
| ---  | ---  | ---         |
| `JA3` | SSL/TLS | TLS Fingerprint can be left as is
| `Proxy` | HTTP | Your Proxy address Format: 'username:password@hostname:port
| `Redirect` | HTTP | specifies the policy for handling redirects
| `Agent` | HEADER | UserAgent for the http client
| `TimeOut` | HTTP | Time-[s] after request with no response, allowed 0 = no timeout

**MassDm**
| Name | Type | Description | 
| ---  | ---  | ---         |
| `WebSocket` | ADDON | Connect to discord wss before sending requests
| `Block_Usr` | ADDON | Block the User after sending them a message
| `Close_Dm` | ADDON | Close Dm after sendinf them a message
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
# TODO
```
dynamic Ja3 Fingerprinting
more captcha api
```

# Previews:

![Discord_erO97kHyUK](https://user-images.githubusercontent.com/110062350/224490994-3dee64da-ca1c-4bc8-b563-6c06b2194b40.gif)

this is for educational use ONLY ;)
i am not held responsible for damage caused by this tool
