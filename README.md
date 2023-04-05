# Go-MassDM
coming soon star for updates!

# https://gomassdm.github.io
this will release soon the code is just being uploaded here
the functions are still in development

# About
<p align="center" style="text-align: center"> 
  <img src="https://img.shields.io/tokei/lines/github/yaboipy/go-massdm">3.5k
  <img src="https://img.shields.io/github/downloads/yaboipy/go-massdm/total?color=blue&label=users">
</p>


```
GoDm Was Made to be Fast undetectable and simple
to use and configure for users
```

![image](https://user-images.githubusercontent.com/110062350/222306213-a9d43e89-1a1e-4315-ae02-5d562664d53e.png)

# Errors
```ruby
ERROR:                                               FIX:
 "Timed out waiting for pipe '\.\pipe\discord-ipc-0" #disable "Presence" in config.json

```

# Issues
```
1. Specify which Function
2. Show a Screenshot
3. note the config u had 
```

# Usage

Go Not installed? 
Download Compiled Version:

https://github.com/YABOIpy/Go-MassDM/releases
```md
Needs & Inputs:
________________________________________________
  [1] Message and Scraped Ids in ids.txt
  [2] victims User ID
  [3] Api Request Link
  [4] Server invite gg/"invite"
  [5] Server ID
  [6] Server Invite and Server ID
  [7] Channel ID and Message
  [8] Server ID and Channel ID
  [9] Tokens in tokens.txt
  [10] Channeld ID And Amount To ping in a single message 
  [11] Server ID Channel ID Message ID and Bot user ID
________________________________________________

```

# Config


Proxy Type: Residential rotating


Proxies available at:


https://iproyal.com


https://proxiware.com


https://proxies.gg

Format: user:pass@ip:port

_____________
Recommended Config:
```json
{
    "Modes": {
        "Discord": {
            "Ver": 1.4,
            "CfBm": false,
            "Service": "",
            "Api_Key": "Capmonster Api Key",
            "Status": "idle",
            "AppID": "1082211264081707068",
            "Presence": true
        },
        "Config": {
            "Interval": 0,
            "Errors": true,
            "SolveCaptcha": false,
            "ErrorLog": true
        },
        "Net": {
            "JA3": "771,4866-4867-4865-49199-49187-52393-49191-107-158-52392-49200-103-49196-49192-159-49188-49195-255,0-11-10-35-16-22-23-13-43-45-51-21,29-23-30-25-24,0-1-2",
            "Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9006 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
            "Proxy": "Your.Proxy.Adress",
            "Redirect": false,
            "TimeOut": 0
        }
    },
    "Settings": {
        "Websocket": true,
        "Block_Usr": false,
        "Close_DM": false
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

# TODO
```
add the other options
dynamic Ja3 Fingerprinting
small nice to have's

```

# Previews:

![Discord_erO97kHyUK](https://user-images.githubusercontent.com/110062350/224490994-3dee64da-ca1c-4bc8-b563-6c06b2194b40.gif)

