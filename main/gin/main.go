package main

import (
	"fmt"

	"github.com/beautiful-you/anniversary/wechat"
	"github.com/beautiful-you/anniversary/wechat/message"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.Any("wechat/public/account/auth_event", authEvent)
	app.Run(":80")
}
func authEvent(c *gin.Context) {
	config := new(wechat.Config)
	config.AppID = OpenWeChatConfig().APPID
	config.AppSecret = OpenWeChatConfig().Appsecret
	config.EncodingAESKey = OpenWeChatConfig().MsgKey
	config.Token = OpenWeChatConfig().Token
	wc := wechat.NewWechat(config)
	server := wc.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(mh)

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}
func mh(msg message.MixMessage) *message.Reply {
	fmt.Println("事件类型1=", msg.Event, "事件内容1=", msg.ComponentVerifyTicket)
	fmt.Println("消息内容1=", msg.Event, "消息内容1=", msg.Content)
	return nil
}
