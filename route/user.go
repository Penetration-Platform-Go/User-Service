package route

import (
	"github.com/Penetration-Platform-Go/User-Service/controller"
	"github.com/Penetration-Platform-Go/User-Service/middleware"
	"github.com/gin-gonic/gin"
)

func userRoute(route *gin.RouterGroup) {
	route.POST("/user", controller.CreateUser)
	route.PUT("/user", middleware.Auth(), controller.UpdateUser)
	route.POST("/avatar", middleware.Auth(), controller.UploadAvatar)
}
