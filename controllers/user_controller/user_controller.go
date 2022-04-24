package usercontroller

import (
	usermodel "demo/models/user_model"
	userservice "demo/services/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	var service userservice.UserService
	var selectedUser []usermodel.UserModel
	selectedUser, err := service.GetUsers()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"result":  "failure",
				"message": err,
			})
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"result": "success",
			"body":   selectedUser,
		})
}
