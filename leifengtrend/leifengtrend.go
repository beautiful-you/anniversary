package leifengtrend

import (
	"github.com/beautiful-you/anniversary/leifengtrend/controller"
	"github.com/beautiful-you/anniversary/leifengtrend/middleware"
	"github.com/gin-gonic/gin"
)

// Controller ... ctr
type Controller struct {
}

// Middleware ... mid
type Middleware struct {
}

// Register ...
func (ctr *Controller) Register(c *gin.Context) {
	controller.Register(c)
}

// Auth ...
func (mid *Middleware) Auth(c *gin.Context) {
	middleware.Auth(c)
}
