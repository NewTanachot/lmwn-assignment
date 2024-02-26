package main

import (
	"lmwn-assignment/api/route"
	"lmwn-assignment/infrastructure/configuration"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	route.MapControllerRoute(app)
	app.Run(configuration.GetEnvPort())
}
