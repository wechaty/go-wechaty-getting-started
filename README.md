# go-wechaty-getting-started ![Go Version](https://img.shields.io/github/go-mod/go-version/wechaty/go-wechaty) [![Go](https://github.com/wechaty/go-wechaty-getting-started/workflows/Go/badge.svg)](https://github.com/wechaty/go-wechaty-getting-started/actions?query=workflow%3AGo)

[![Go Wechaty](https://wechaty.github.io/go-wechaty/images/go-wechaty.png)](https://github.com/wechaty/go-wechaty-getting-started)

[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/wechaty/go-wechaty-getting-started) 

Go Wechaty Starter Project Template that Works Out-of-the-Box

[Go Playground](https://play.golang.org/p/JGw_2TrVQeg): <https://play.golang.org/p/JGw_2TrVQeg>

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
    OnScan(func(qrCode string, status schemas.ScanStatus, data string) {
      fmt.Printf("Scan QR Code to login: %s\nhttps://api.qrserver.com/v1/create-qr-code/?data=%s\n", status, qrCode)
    }).
    OnLogin(user *user.ContactSelf) { fmt.Printf("User %s logined\n", user) }).
    OnMessage(message *user.Message) { fmt.Printf("Message: %s\n", message) }).
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

Get a Token for your Bot first. Learn more from our [Wechaty Developers Program](https://github.com/wechaty/wechaty/wiki/Wechaty-Developer-Program)

```sh
make bot
```

## Wechaty Getting Started in Multiple Languages

- [TypeScript Wechaty Getting Started](https://github.com/wechaty/wechaty-getting-started)
- [Python Wechaty Getting Started](https://github.com/wechaty/python-wechaty-getting-started)
- [Java Wechaty Getting Started](https://github.com/wechaty/java-wechaty-getting-started)
- [Go Wechaty Getting Started](https://github.com/wechaty/go-wechaty-getting-started)

## Maintainers

[@wechaty/go](https://github.com/orgs/wechaty/teams/go/members)

## Copyright & License

- Code & Docs © 2020 Wechaty Contributors <https://github.com/wechaty>
- Code released under the Apache-2.0 License
- Docs released under Creative Commons
