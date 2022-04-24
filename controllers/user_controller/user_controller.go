package usercontroller

import (
	usermodel "demo/models/user_model"
	userservice "demo/services/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	var (
		service userservice.UserService
		users   []usermodel.UserModel
		err     error
	)
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
