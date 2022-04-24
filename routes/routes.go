package routes

import (
	usercontroller "demo/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func Routing(router *gin.Engine) {
	baseGroup := router.Group("/apis")
	{
		userGroup := baseGroup.Group("/user")
		{
			userGroup.GET("/", usercontroller.GetSelectedUser)
		}
	}
}
