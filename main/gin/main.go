package main

import (
	"fmt"
	"time"

	"github.com/beautiful-you/anniversary/wechat/cache"

	"github.com/beautiful-you/anniversary/wechat"
	"github.com/beautiful-you/anniversary/wechat/message"

	"github.com/gin-gonic/gin"
)

const (
	componentTokenURL = "https://api.weixin.qq.com/cgi-bin/component/api_component_token"
)

func main() {
	str := fmt.Sprintf("updateRemarkURL%s", "accessToken")
	fmt.Println(str)
	/*
		app := gin.New()
		app.Any("wechat/public/account/auth_event", authEvent)
		app.Run(":80")
	*/
}
func authEvent(c *gin.Context) {
	config := new(wechat.Config)
	wc := wechat.NewWechat(config)
	server := wc.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(cachecvt)

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}

// cachecvt 缓存微信票据
func cachecvt(msg message.MixMessage) *message.Reply {
	if len(msg.ComponentVerifyTicket) > 0 {
		ca := cache.NewMemory()
		ca.Set("ComponentVerifyTicket", msg.ComponentVerifyTicket, time.Minute*15)
		return nil
	}
	return nil
}
