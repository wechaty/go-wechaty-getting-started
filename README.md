# go-wechaty-getting-started

Go Wechaty Starter Project Template that Works Out-of-the-Box

![Go Version](https://img.shields.io/github/go-mod/go-version/wechaty/go-wechaty)
[![Go](https://github.com/wechaty/go-wechaty/workflows/Go/badge.svg)](https://github.com/wechaty/go-wechaty/actions?query=workflow%3AGo)

## Connecting Chatbots

[![Powered by Wechaty](https://img.shields.io/badge/Powered%20By-Wechaty-brightgreen.svg)](https://github.com/Wechaty/wechaty)

Wechaty is a RPA SDK for Wechat **Individual** Account that can help you create a chatbot in 6 lines of Go.

## The World's Shortest Go ChatBot: 7 lines of Code

```go
package main

import (
  "fmt"

  "github.com/wechaty/go-wechaty/wechaty"
)

func main() {
  _ = wechaty.NewWechaty().
    OnScan(func(qrCode, status string) {
      fmt.Printf("Scan QR Code to login: %s\nhttps://api.qrserver.com/v1/create-qr-code/?data=%s\n", status, qrCode)
    }).
    OnLogin(func(user string) { fmt.Printf("User %s logined\n", user) }).
    OnMessage(func(message string) { fmt.Printf("Message: %s\n", message) }).
    Start()
}
```

## Requirements

1. Go 1.14+

## Install

```shell
make install
```

## Run

```sh
make bot
```

## Copyright & License

- Code & Docs © 2020-now Wechaty <https://github.com/wechaty>
- Code released under the Apache-2.0 License
- Docs released under Creative Commons
