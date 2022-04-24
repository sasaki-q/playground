package main

import (
	"demo/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	routes.Routing(engine)
	engine.Run(":" + os.Getenv("PORTS"))
}
