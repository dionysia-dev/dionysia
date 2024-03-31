package api

import "github.com/gin-gonic/gin"

func addRoutes(e *gin.Engine) {
	e.POST("/ready", onReady)
}
