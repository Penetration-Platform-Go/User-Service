package route

import "github.com/gin-gonic/gin"

import "github.com/Penetration-Platform-Go/User-Service/controller"

func userRoute(route *gin.RouterGroup) {
	route.POST("/user", controller.CreateUser)
	route.PUT("/user", controller.UpdateUser)
}
