package main

import (
	"demo/config/db"
	"demo/routes"

	// "os"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	db.Init()
	routes.Routing(engine)
	// engine.Run(":" + os.Getenv("PORTS"))
	engine.Run(":8080")
}
