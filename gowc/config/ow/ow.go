package ow

// OpenWeChat config struct
type OpenWeChat struct {
	AppID           string //
	AppSecret       string //
	Token           string // 消息校验Token
	EncodingAESKey  string // 消息加解密KeyL9LSFLlTj8lYLy8Y8Q2D6668d6y982ef6Q0G26jdetb
	AuthRedirectURL string // 授权第三方公众号的 回调链接
}

// AesKey ... JWT | AppID, AppSecret, Token, EncodingAESKey, AuthRedirectURL ... OPWC
const (
	AesKey          = "MY4tM5jSFLlTj8l35cf"
	AppID           = "wx45fc784e7eb235cf"
	AppSecret       = "22afb1d94c9c7ea0e8f403b8c7133305"
	Token           = "HekMkeMY4tM5jeO40jTOaTtSPn5slMEg"
	EncodingAESKey  = "L9LSFLlTj8lYLy8Y8Q2D6668d6y982ef6Q0G26jdetb"
	AuthRedirectURL = "http://am.oovmi.com/wechat/public/account/auth_call"
)
