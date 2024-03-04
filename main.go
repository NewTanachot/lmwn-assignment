package main

import (
	"fmt"
	"lmwn-assignment/api/route"
	"lmwn-assignment/infrastructure/configuration"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	route.MapControllerRoute(app)

	app.SetTrustedProxies([]string{})

	portString := configuration.GetEnvPort()
	fmt.Printf("server runing on http://localhost%s\n", portString)
	app.Run(portString)
}
