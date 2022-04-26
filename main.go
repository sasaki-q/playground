package main

import (
	"demo/routes"
	"demo/util"

	// "os"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// db.Init()
	routes.Routing(engine)
	// engine.Run(":" + os.Getenv("PORTS"))
	engine.Run(":8080")
	util.Debug("start runnning")
}
