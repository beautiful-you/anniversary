package controller

import (
	"github.com/beautiful-you/anniversary/gowc/app/controller/wechat"
)

// Controller 控制器
type Controller struct {
}

// WeChat 执行测试任务
func (ctr *Controller) WeChat() *wechat.WeChat {
	wc := new(wechat.WeChat)
	return wc
}
