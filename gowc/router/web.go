package router

import (
	"github.com/beautiful-you/anniversary/gowc/app/controller"
	"github.com/gin-gonic/gin"
)

var ctr = new(controller.Controller)

// WEB ...
func WEB(app *gin.Engine) {
	wc := app.Group("/wechat/public/account")
	{
		wc.Any("auth_event", ctr.WeChat.AuthEvent)
		wc.Any("message_with_event/:appid", ctr.WeChat.MessageWithEvent)
		wc.Any("auth_call", ctr.WeChat.AuthCall)
		wc.Any("auth_url", ctr.WeChat.AuthURL)
		wc.Any("test", ctr.WeChat.Test)
	}
	app.Run(":80")
}
