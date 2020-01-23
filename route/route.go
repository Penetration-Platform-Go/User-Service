package route

import "github.com/gin-gonic/gin"


// GetServer Return Gin Server
func GetServer() *gin.Engine {
	app := gin.Default()
	return app
}
