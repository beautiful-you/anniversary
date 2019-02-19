package main

import (
	"fmt"

	"github.com/beautiful-you/anniversary/wechat/message"
	"github.com/beautiful-you/anniversary/wechat/platforms"
	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.New()
	app.Any("wechat/public/account/auth_event", authEvent)
}
func authEvent(c *gin.Context) {
	config := new(platforms.Config)
	config.AppID = OpenWeChatConfig().APPID
	config.AppSecret = OpenWeChatConfig().Appsecret
	config.EncodingAESKey = OpenWeChatConfig().MsgKey
	config.Token = OpenWeChatConfig().Token
	wc := platforms.NewPlatform(config)
	server := wc.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		switch msg.MsgType {

		case message.MsgTypeText:
			// 文本消息 (关键词处理)
			text := message.NewText(msg.Content)
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		case message.MsgTypeImage:
			// 图片消息

		case message.MsgTypeVoice:
			// 语音消息

		case message.MsgTypeVideo:
			// 视频消息

		case message.MsgTypeShortVideo:
			// 小视频消息

		case message.MsgTypeLocation:
			// 地理位置消息

		case message.MsgTypeLink:
			// 链接消息

		case message.MsgTypeEvent:
			// 事件消息
			fmt.Println("事件类型=", msg.Event, "事件内容=", msg.ComponentVerifyTicket)
		}
		return nil
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}
