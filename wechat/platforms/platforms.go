// Package platforms 参考 https://github.com/beautiful-you/anniversary/wechat/tree/master/util
package platforms

import (
	"fmt"
	"net/http"

	"github.com/beautiful-you/anniversary/wechat"
	"github.com/beautiful-you/anniversary/wechat/message"
)

// SetComponentVerifyTicket 缓存该票据
func SetComponentVerifyTicket(config *wechat.Config, rw http.ResponseWriter, req *http.Request) {
	wc := wechat.NewWechat(config)
	server := wc.GetServer(req, rw)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		fmt.Println(msg.ComponentVerifyTicket)
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

// GetComponentVerifyTicket 获取该票据
func GetComponentVerifyTicket() {

}
