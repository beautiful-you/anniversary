package gowc

import (
	"github.com/beautiful-you/anniversary/gowc/app/middleware"
	"github.com/beautiful-you/anniversary/gowc/router"
	"github.com/gin-gonic/gin"
)

// Start GinApp
func Start() {
	GinApp := gin.Default()
	GinApp.Use(middleware.GOWC)
	router.WEB(GinApp)
}
