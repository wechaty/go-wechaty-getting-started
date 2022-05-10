package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/mdp/qrterminal/v3"
	"github.com/spf13/viper"
	. "github.com/wechaty/go-wechaty-getting-started/examples/plug-demo/DingTools"
	. "github.com/wechaty/go-wechaty-getting-started/examples/plug-demo/StructTools" // 导入结构体
	"github.com/wechaty/go-wechaty/wechaty"
	puppet "github.com/wechaty/go-wechaty/wechaty-puppet"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"github.com/wechaty/go-wechaty/wechaty/user"
)

var (
	err error
)

/*
	初始化插件
*/
func init() {
	// 初始化配置文件
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath(abPath)
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Read config failed: %v", err)
	}
	if viper.GetString("WECHATY.TOKEN") == "" {
		fmt.Printf("Please Set WECHATY.TOKEN in config.yaml")
		os.Exit(1)
	}
}

/*
	主程序
*/
func main() {
	var bot = wechaty.NewWechaty(wechaty.WithPuppetOption(puppet.Option{
		Token:    viper.GetString("WECHATY.TOKEN"),
		Endpoint: viper.GetString("WECHATY.ENDPOINT"),
	}))
	bot.
		OnScan(func(context *wechaty.Context, qrCode string, status schemas.ScanStatus, data string) {
			fmt.Printf("\n\n")
			if status.String() == "ScanStatusWaiting" {
				qrterminal.GenerateWithConfig(qrCode, qrterminal.Config{
					Level:     qrterminal.L,
					Writer:    os.Stdout,
					BlackChar: " \u2005",
					WhiteChar: "💵",
					QuietZone: 1,
				})
				fmt.Printf("\n\n")
				fmt.Printf("%v[Scan] https://wechaty.js.org/qrcode/%v %v", viper.GetString("info"), qrCode, data)
			} else if status.String() == "ScanStatusScanned" {
				fmt.Printf("%v[Scan] Status: %v %v\n", viper.GetString("info"), status.String(), data)
			} else {
				fmt.Printf("%v[Scan] Status: %v %v\n", viper.GetString("info"), status.String(), data)
			}
		}).
		OnLogin(func(context *wechaty.Context, user *user.ContactSelf) {
			fmt.Printf("\n微信机器人: [%v] 已经登录成功了。\n", user.Name())
			// 可 别小看 context
			context.SetData("botName", user.Name())
		}).
		OnMessage(func(context *wechaty.Context, message *user.Message) {
			var m MessageInfo

			// 群聊消息 || 好友消息
			// 这样最简单，最快捷吧
			if message.Room() != nil {
				m.Status = true
				// 如果 @ 我了
				if message.MentionSelf() {
					m.AtMe = true
					// 如果是@所有人
					if strings.Contains(message.Text(), "所有人") {
						m.Pass = true
						m.PassResult = "全员消息"
					}
				} else {
					m.AtMe = false
				}
			} else {
				m.Status = false
				m.AtMe = true
			}

			// 测试是否运作
			if message.Type() == schemas.MessageTypeText {
				if message.Age() > 2*60*time.Second {
					if message.Text() == "ding" {
						if _, err := message.Say("dong"); err != nil {
							fmt.Printf("Send Message Error, Error: [%v]", err)
							// 这里不能 return ， 因为 下面还有一个 数据没导出，我自己是单独写了个模块处理这些数据
						}
						// name回复过了，就设置下回复状态吧，不然别人问你一句，你会别人十句
						m.Reply = true
						m.ReplyResult = "ding"
					}
				}
			}

			// 保存 结构体，当前消息所在的任何函数都可调用
			context.SetData("MessageInfo", m)
			// 执行完毕 ，按照 顺序，下一个执行的就是 DingTools
		}).

		// wechaty 执行是按照先后顺序的
		// OnScan 和 Use 插件 内部的 Onscan 也是按顺序执行

		Use(DingTools()).

		// 启动守护程序
		DaemonStart()
}
