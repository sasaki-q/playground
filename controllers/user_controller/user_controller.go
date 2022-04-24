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
		return
	}
	ctx.IndentedJSON(
		http.StatusOK,
		users,
	)
	return
}

func Create(ctx *gin.Context) {
	var (
		body usermodel.CreateModel
		err  error
	)
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"result":  "failure",
				"message": err.Error(),
			})
		return
	}
	ctx.IndentedJSON(
		http.StatusOK,
		body,
	)
	return
}
