package route

import "github.com/gin-gonic/gin"

// GetServer Return Gin Server
func GetServer() *gin.Engine {
	app := gin.Default()
	auth := app.Group("/api")
	userRoute(auth)
	return app
}
