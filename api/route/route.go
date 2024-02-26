package route

import (
	"lmwn-assignment/api/controller"

	"github.com/gin-gonic/gin"
)

func MapControllerRoute(app *gin.Engine) {
	covidController := app.Group("/covid")
	covidController.GET("/summary", controller.GetCovidSummary)
}
