package usercontroller

import (
	usermodel "demo/models/user_model"
	userservice "demo/services/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	service userservice.UserService
	users   []usermodel.UserModel
	err     error
)

func GetUsers(ctx *gin.Context) {
	users, err = service.GetUsers()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"result":  "failure",
				"message": err.Error(),
			})
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"result": "success",
			"users":  users,
		})
}
