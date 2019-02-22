package middleware

import (
	"fmt"

	"github.com/beautiful-you/anniversary/leifengtrend/jwt"
	"github.com/gin-gonic/gin"
)

// Auth 简单的验证是否登陆 uuid use jwt
func Auth(c *gin.Context) {
	// 是否登陆
	_, err := jwt.Chcek(c.ClientIP(), c.GetHeader("token"))
	if err != nil {
		// ... user not login
		x := fmt.Sprintf("%s", err)
		c.JSON(200, gin.H{"errCode": "error", "info": "user not login " + x})
		c.Abort()
	}
	c.Next()
}
