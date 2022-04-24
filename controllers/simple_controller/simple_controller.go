package simplecontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnString(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "Hello from ECS !",
		})
}
