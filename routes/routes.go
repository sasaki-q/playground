package routes

import (
	simplecontroller "demo/controllers/simple_controller"
	usercontroller "demo/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func Routing(router *gin.Engine) {
	baseGroup := router.Group("/apis")
	{
		simple := baseGroup.Group("/simple")
		{
			simple.GET("/", simplecontroller.ReturnString)
		}
		user := baseGroup.Group("/user")
		{
			user.GET("/", usercontroller.GetUsers)
			user.POST("/", usercontroller.Create)
		}
	}
}
