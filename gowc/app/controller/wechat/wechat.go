package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/beautiful-you/anniversary/wechat/util"

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
	ComponentTokenURL = "https://api.weixin.qq.com/cgi-bin/component/api_component_token"
	PreAuthCodeURL    = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=%s"
	AuthURL           = "https://mp.weixin.qq.com/safe/bindcomponent?action=bindcomponent&auth_type=3&no_scan=1&component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=xxx&biz_appid=xxxx#wechat_redirect"
	RedirectURL       = ""
)

var lcfg = new(config.Config)

// AuthCall ... 授权后回调地址
func (w *WeChat) AuthCall(c *gin.Context) {

}

// ResComponentAccessToken ComponentAccessToken
type ResComponentAccessToken struct {
	util.CommonError
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int64  `json:"expires_in"`
}

// ResPreAuthCode 预授权码
type ResPreAuthCode struct {
	util.CommonError
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int64  `json:"expires_in"`
}

// AuthURL ... 授权地址
func (w *WeChat) AuthURL(c *gin.Context) {
	// 获取 component_verify_ticket
	cvt := componentverifyticket()
	if cvt == "" {
		c.Writer.WriteString("获取不到缓存中的 component_verify_ticket")
		return
	}
	// 获取第三方平台 component_access_token
	body, err := util.PostJSON(ComponentTokenURL, map[string]string{"component_appid": lcfg.OW().AppID, "component_appsecret": lcfg.OW().AppSecret, "component_verify_ticket": cvt})
	if err != nil {
		return
	}
	resComponentAccessToken := new(ResComponentAccessToken)
	err = json.Unmarshal(body, &resComponentAccessToken)
	if err != nil {
		return
	}
	if resComponentAccessToken.ErrMsg != "" {
		err = fmt.Errorf("get access_token error : errcode=%v , errormsg=%v", resComponentAccessToken.ErrCode, resComponentAccessToken.ErrMsg)
		return
	}
	// 获取预授权码 pre_auth_code
	body, err = util.PostJSON(ComponentTokenURL, map[string]string{"component_appid": lcfg.OW().AppID, "component_appsecret": lcfg.OW().AppSecret, "component_verify_ticket": cvt})
	if err != nil {
		return
	}
	resPreAuthCode := new(ResPreAuthCode)
	err = json.Unmarshal(body, &resPreAuthCode)
	if err != nil {
		return
	}
	if resPreAuthCode.ErrMsg != "" {
		err = fmt.Errorf("get auth_code error : errcode=%v , errormsg=%v", resComponentAccessToken.ErrCode, resComponentAccessToken.ErrMsg)
		return
	}
	URL := fmt.Sprint(AuthURL, lcfg.OW().AppID, resPreAuthCode.PreAuthCode, RedirectURL)
	c.Writer.WriteString(URL)
	// cat := resComponentAccessToken.ComponentAccessToken
	// PreAuthCode := "pre_auth_code"

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
		ca.Set("s", "字符串", time.Minute*15)
		c.Writer.WriteString(c.Request.FormValue("ca"))
		//return
	}
	fmt.Println(ca.Get("s"))
	str, bool := ca.Get("s")
	if bool {
		c.Writer.WriteString(str.(string))
		//return
	}
	c.Writer.WriteString("error")

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
		ca.Set("ComponentVerifyTicket", msg.ComponentVerifyTicket, time.Minute*15)
		return nil
	}
	return nil
}

// componentverifyticket
func componentverifyticket() string {
	ca := lcfg.Cache()
	str, bool := ca.Get("ComponentVerifyTicket")
	if bool {
		return str.(string)
	}
	fmt.Println("ComponentVerifyTicket no cache")
	return ""
}
