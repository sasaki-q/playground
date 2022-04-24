package usercontroller

import (
	usermodel "demo/models/user_model"
	userservice "demo/services/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSelectedUser(ctx *gin.Context) {
	var service userservice.UserService
	var selectedUser usermodel.UserModel = *service.GetSelectedUser()
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "success",
			"body":    selectedUser,
		})
}
