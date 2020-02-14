package route

import (
	"github.com/Penetration-Platform-Go/User-Service/middleware"
	"github.com/gin-gonic/gin"
)

// GetServer Return Gin Server
func GetServer() *gin.Engine {
	app := gin.Default()
	app.Use(middleware.Cors())
	auth := app.Group("/api")
	userRoute(auth)
	return app
}
