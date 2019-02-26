package config

import (
	"github.com/beautiful-you/anniversary/gowc/config/cache"

	"github.com/beautiful-you/anniversary/gowc/config/ow"
)

// Config 配置相关
type Config struct {
}

// owc 微信配置相关
var owc = new(ow.OpenWeChat)

// OW 微信配置相关
func (cfg *Config) OW() *ow.OpenWeChat {
	owc.AppID = ow.AppID
	owc.AppSecret = ow.AppSecret
	owc.Token = ow.Token
	owc.EncodingAESKey = ow.EncodingAESKey
	owc.AuthRedirectURL = ow.AuthRedirectURL
	return owc
}

// ca Cache配置相关
var ca = cache.New()

// Cache Cache配置相关
func (cfg *Config) Cache() *cache.Cache {
	return ca
}
