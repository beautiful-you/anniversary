package wechat

import (
	"errors"
	"fmt"

	"github.com/beautiful-you/anniversary/wechat/platforms"

	"github.com/beautiful-you/anniversary/gowc/config"

	// _ "github.com/beautiful-you/anniversary/gowc/app/controller/wechat"
	"github.com/beautiful-you/anniversary/wechat"
	"github.com/beautiful-you/anniversary/wechat/message"
	"github.com/gin-gonic/gin"
)

// WeChat 控制器
type WeChat struct {
}

// 接口信息
const (
	RedirectURL = "http://am.jyacad.cc/wechat/public/account/auth_call"
)

var lcfg = new(config.Config)

// AuthCall ... 授权后回调地址
func (w *WeChat) AuthCall(c *gin.Context) {

}

// AuthURL ... 授权地址
func (w *WeChat) AuthURL(c *gin.Context) {
	// 获取 component_verify_ticket
	cvt := componentverifyticket()
	if cvt == "" {
		fmt.Println("获取不到缓存中的 component_verify_ticket")
		c.Writer.WriteString("获取不到缓存中的 component_verify_ticket")
		return
	}

	resComponentAccessToken, err := platforms.ComponentAccessToken(lcfg.OW().AppID, lcfg.OW().AppSecret, cvt)
	if err != nil {
		fmt.Println(err)
		c.Writer.WriteString("获取 ComponentAccessToken 出现错误")
		return
	}

	resPreAuthCode, err := platforms.PreAuthCode(lcfg.OW().AppID, resComponentAccessToken.ComponentAccessToken)
	if err != nil {
		fmt.Println(err)
		c.Writer.WriteString("获取 PreAuthCode 出现错误")
		return
	}

	URL := fmt.Sprint(platforms.AuthURL, lcfg.OW().AppID, resPreAuthCode.PreAuthCode, RedirectURL)
	c.Writer.WriteString(URL)
}

// MessageWithEvent ... 消息与事件接收url
func (w *WeChat) MessageWithEvent(c *gin.Context) {
	cfg := new(wechat.Config)
	cfg.AppID = lcfg.OW().AppID
	cfg.AppSecret = lcfg.OW().AppSecret
	cfg.Token = lcfg.OW().Token
	cfg.EncodingAESKey = lcfg.OW().EncodingAESKey
	wc := wechat.NewWechat(cfg)
	server := wc.GetServer(c.Request, c.Writer)

	//设置接收消息的处理方法
	server.SetMessageHandler(messageWithEventHandler)
	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(errors.New("server.Serve() error: "))
		fmt.Println(err)
		return
	}
	//发送回复的消息
	err = server.Send()
	if err != nil {
		fmt.Println(errors.New("server.Send() error: "))
		fmt.Println(err)
		return
	}
}

// messageWithEventHandler  消息与事件处理
func messageWithEventHandler(msg message.MixMessage) *message.Reply {
	return nil
}

// Test 执行测试任务
func (w *WeChat) Test(c *gin.Context) {
	lcfg := new(config.Config)
	ca := lcfg.Cache()
	if c.Request.FormValue("ca") == "set" {
		ca.Set("s", "字符串")
		c.Writer.WriteString(c.Request.FormValue("ca"))
		//return
	}
	fmt.Println(ca.Get("s"))
	str, err := ca.Get("s")
	if err != nil {
		c.Writer.WriteString("error")
		return
	}
	c.Writer.WriteString(str)

}

// AuthEvent 授权事件接收URL
func (w *WeChat) AuthEvent(c *gin.Context) {
	cfg := new(wechat.Config)
	lcfg := new(config.Config)
	cfg.AppID = lcfg.OW().AppID
	cfg.AppSecret = lcfg.OW().AppSecret
	cfg.Token = lcfg.OW().Token
	cfg.EncodingAESKey = lcfg.OW().EncodingAESKey

	wc := wechat.NewWechat(cfg)
	server := wc.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(authEventHandler)
	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(errors.New("server.Serve() error: "))
		fmt.Println(err)
		return
	}
	//发送回复的消息
	err = server.Send()
	if err != nil {
		fmt.Println(errors.New("server.Send() error: "))
		fmt.Println(err)
		return
	}
}

// authEventHandler 授权事件处理
func authEventHandler(msg message.MixMessage) *message.Reply {
	if msg.InfoType == "component_verify_ticket" {
		ca := lcfg.Cache()
		err := ca.Set("ComponentVerifyTicket", msg.ComponentVerifyTicket)
		if err != nil {
			fmt.Println("cache error cvt")
			return nil
		}
		return nil
	}
	return nil
}

// componentverifyticket
func componentverifyticket() string {
	ca := lcfg.Cache()
	str, err := ca.Get("ComponentVerifyTicket")
	if err != nil {
		return str
	}
	fmt.Println("ComponentVerifyTicket no cache")
	return ""
}
