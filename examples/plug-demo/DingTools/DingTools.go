package DingTools

import (
	"fmt"
	"os"
	"time"

	"github.com/blinkbean/dingtalk"
	"github.com/spf13/viper"
	. "github.com/wechaty/go-wechaty-getting-started/examples/plug-demo/StructTools" // 导入结构体
	"github.com/wechaty/go-wechaty/wechaty"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"github.com/wechaty/go-wechaty/wechaty/user"
)

/*
	初始化插件
*/
func DingTools() *wechaty.Plugin {
	// 初始化插件
	plug := wechaty.NewPlugin()
	plug.
		OnScan(onScan).
		OnLogout(onLogout).
		OnMessage(onMessage)

	return plug
}

/*
	处理登录事件，无需传参
*/
func onScan(context *wechaty.Context, qrCode string, status schemas.ScanStatus, data string) {
	if status.String() == "ScanStatusWaiting" {
		for i := 0; i < 6; i++ {
			DingSend(fmt.Sprintf("账号未登录请扫码!\n\n---\n\n[qrCode](https://wechaty.js.org/qrcode/%v)", qrCode))
			time.Sleep(120 * time.Second)
			if i == 5 {
				os.Exit(1)
			}
		}
	} else if status.String() == "ScanStatusScanned" {
		fmt.Printf("%v[Scan] Status: %v %v\n", viper.GetString("info"), status.String(), data)
	} else {
		fmt.Printf("%v[Scan] Status: %v %v\n", viper.GetString("info"), status.String(), data)
	}
}

/*
	处理账号登出事件，无需传参
*/
func onLogout(context *wechaty.Context, user *user.ContactSelf, reason string) {
	DingSend(fmt.Sprintf("%s已登出!\n\n---\n\n**reason**\n\n%v", user.Name(), reason))
}

/*
	处理消息事件，由 DingTools 传入消息对象
*/
func onMessage(context *wechaty.Context, message *user.Message) {
	// 导入我们之前设置的 结构体数据
	m, ok := (context.GetData("MessageInfo")).(MessageInfo)
	if !ok {
		fmt.Printf("Conversion Failed")
		return
	}

	// 判断是否是机器人自己发送的消息
	if message.Self() {
		fmt.Printf("Self")
	}

	// 判断是否超过 2 分钟
	if message.Age() > 2*60*time.Second {
		fmt.Printf("TimeOut")
	}

	// 判断是否回复消息， 尴尬下面还要再写一次, 这里不写的话会有一堆的消息
	if !m.Reply {
		return
	}

	// 设置 DingDing 回复消息的内容
	msg := fmt.Sprintf("%v@我了\n\n---\n\n### 用户属性\n\n用户名: [%v]\n\n用户ID: [%v]", message.Talker().Name(), message.Talker().Name(), message.Talker().ID())

	// 判读是否是群聊
	if m.Status {
		msg += fmt.Sprintf("\n\n---\n\n### 群聊属性\n\n群聊名称: [%v]\n\n群聊ID: [%v]", message.Room().Topic(), message.Room().ID())
	}

	// 继续追加内容
	msg += fmt.Sprintf("\n\n---\n\n**内容**: [%v]", message.Text())

	// 判断是否有别的原因跳过了操作
	if m.Pass {
		msg += fmt.Sprintf("\n\n**Pass**: [%v]", m.PassResult)
	}

	// 判断是否回复内容
	if m.Reply {
		msg += fmt.Sprintf("\n\n**回复**: [%v]", m.ReplyResult)
	}

	// 到这里的时候基本设置好了一些默认的值了
	DingSend(msg)
}

/*
	发送消息到 DingDing
	外部调用:
	DingSend(msg string)
*/
func DingSend(msg string) {
	cli := dingtalk.InitDingTalkWithSecret(viper.GetString("Ding.TOKEN"), viper.GetString("Ding.SECRET"))
	if err := cli.SendMarkDownMessage(msg, msg); err != nil {
		fmt.Printf("DingMessage Error: [%v]", err)
		return
	}
	fmt.Printf("DingTalk Success!")
}
