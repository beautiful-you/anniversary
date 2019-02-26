package cache

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

var ca = cache.New(time.Minute*120, time.Minute*10)

// Set 设置缓存
func Set(key string, x interface{}, t time.Duration) {
	ca.Set(key, x, t)
}

// Get 获取缓存
func Get(key string) (string, bool) {
	s, b := ca.Get(key)
	if b {
		return s.(string), b
	}
	return "error", false
}
