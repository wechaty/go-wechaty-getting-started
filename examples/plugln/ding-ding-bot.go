/**
 *   Wechaty - https://github.com/wechaty/wechaty
 *
 *   @copyright 2020-now Wechaty
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 *
 */
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/wechaty/go-wechaty/wechaty"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"github.com/wechaty/go-wechaty/wechaty/user"
)

var err error

func main() {
	var bot = wechaty.NewWechaty()

	bot.OnScan(func(context *wechaty.Context, qrCode string, status schemas.ScanStatus, data string) {
		fmt.Printf("Scan QR Code to login: %v\nhttps://wechaty.js.org/qrcode/%s\n", status, qrCode)
	}).OnLogin(func(context *wechaty.Context, user *user.ContactSelf) {
		fmt.Printf("User %s logined\n", user.Name())
	}).OnLogout(func(ctx *wechaty.Context, user *user.ContactSelf, reason string) {
		fmt.Printf("User %s logouted: %s\n", user, reason)
	})

	bot.OnMessage(onMessage)
	// The First Plug-in
	bot.Use(PlugOne())
	// The Second Plug-in
	bot.Use(PlugTwo("Hello, PlugTwo"))

	bot.DaemonStart()
}

// The first onMessage processing logic
func onMessage(context *wechaty.Context, message *user.Message) {
	log.Println(message)

	if message.Self() {
		log.Println("Message discarded because its outgoing")
	}

	if message.Age() > 2*60*time.Second {
		log.Println("Message discarded because its TOO OLD(than 2 minutes)")
	}

	if message.Type() != schemas.MessageTypeText || message.Text() != "ding" {
		log.Println("Message discarded because it does not match 'ding'")

		// Set needReply Status (`return`)
		context.SetData("needReply", false)
		return
	}

	// 1. reply 'dong'
	if _, err = message.Say("dong"); err != nil {
		log.Println(err)
		return
	}
	log.Println("REPLY: dong")

	// Set needReply Status
	context.SetData("needReply", true)

	// 2. reply image(qrcode image)
	//fileBox, _ := file_box.FromUrl("https://wechaty.github.io/wechaty/images/bot-qr-code.png", "", nil)
	//_, err = message.Say(fileBox)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Printf("REPLY: %s\n", fileBox)
}

// PlugOne (The First Plug-in)
func PlugOne() *wechaty.Plugin {
	newPlug := wechaty.NewPlugin()

	// The Second Onmessage Processing Logic
	newPlug.OnMessage(func(context *wechaty.Context, message *user.Message) {

		// Get Value From Wechaty Context
		needReply, ok := context.GetData("needReply").(bool)

		// Determine If Parsing Error (`!`)
		if !ok {
			log.Println("context GetData needReply Error")
			return
		}

		// Determine If You Have Responded (`!`)
		if !needReply {
			return
		}

		if _, err = message.Say("PlugOne OnMessage Here!"); err != nil {
			log.Println(err)
			return
		}
		log.Println("PlugOne OnMessage Here!")
	})
	return newPlug
}

// PlugTwo (The Second Plug-in)
func PlugTwo(args string) *wechaty.Plugin {
	newPlug := wechaty.NewPlugin()

	// The Third Onmessage Processing Logic
	newPlug.OnMessage(func(context *wechaty.Context, message *user.Message) {

		// Get Value From Wechaty Context
		needReply, ok := context.GetData("needReply").(bool)

		// Determine If Parsing Error (`!`)
		if !ok {
			log.Println("context GetData needReply Error")
			return
		}

		log.Println(args)

		// Determine If You Have Responded (`!`)
		if !needReply {
			return
		}

		if _, err = message.Say("PlugTwo OnMessage Here!"); err != nil {
			log.Println(err)
			return
		}
		log.Println("PlugTwo OnMessage Here!")
	})
	return newPlug
}
