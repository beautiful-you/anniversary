package main

// OpenWeChat config struct
type OpenWeChat struct {
	APPID           string //
	Appsecret       string //
	Token           string // 消息校验Token
	MsgKey          string // 消息加解密KeyL9LSFLlTj8lYLy8Y8Q2D6668d6y982ef6Q0G26jdetb
	AuthRedirectURL string // 授权第三方公众号的 回调链接
}

// AesKey ...
const AesKey = "MY4tM5jSFLlTj8l35cf"

// OpenWeChatConfig export
func OpenWeChatConfig() *OpenWeChat {
	ow := new(OpenWeChat)
	ow.APPID = "wx45fc784e7eb235cf"
	ow.Appsecret = "22afb1d94c9c7ea0e8f403b8c7133305"
	ow.Token = "HekMkeMY4tM5jeO40jTOaTtSPn5slMEg"
	ow.MsgKey = "L9LSFLlTj8lYLy8Y8Q2D6668d6y982ef6Q0G26jdetb"
	ow.AuthRedirectURL = "http://am.oovmi.com/wechat/public/account/auth_call"
	return ow
}
