// Package platforms 参考 https://github.com/beautiful-you/anniversary/wechat/tree/master/util
package platforms

import (
	"net/http"
	"sync"

	"github.com/beautiful-you/anniversary/wechat/cache"
	"github.com/beautiful-you/anniversary/wechat/context"
	"github.com/beautiful-you/anniversary/wechat/server"
)

// Wechat struct
type Wechat struct {
	Context *context.Context
}

// Config for user
type Config struct {
	AppID          string // platforms appid
	AppSecret      string // platforms AppSecret
	Token          string // platforms Token
	EncodingAESKey string // platforms EncodingAESKey
	PayMchID       string // 支付 - 商户 ID
	PayNotifyURL   string // 支付 - 接受微信支付结果通知的接口地址
	PayKey         string // 支付 - 商户后台设置的支付 key
	Cache          cache.Cache
}

// NewPlatform init
func NewPlatform(cfg *Config) *Wechat {
	context := new(context.Context)
	copyConfigToContext(cfg, context)
	return &Wechat{context}
}
func copyConfigToContext(cfg *Config, context *context.Context) {
	context.AppID = cfg.AppID
	context.AppSecret = cfg.AppSecret
	context.Token = cfg.Token
	context.EncodingAESKey = cfg.EncodingAESKey
	context.Cache = cfg.Cache
	context.SetAccessTokenLock(new(sync.RWMutex))
	context.SetJsAPITicketLock(new(sync.RWMutex))
}

// GetServer 消息管理
func (wc *Wechat) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return server.NewServer(wc.Context)
}
