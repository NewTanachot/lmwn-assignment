package controller

import (
	"lmwn-assignment/application/usecase"
	"lmwn-assignment/infrastructure/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCovidSummary(context *gin.Context) {
	covidRepo := repository.NewCovidRepository()
	covidUseCase := usecase.NewCovidUseCase(covidRepo)

	response, err := covidUseCase.GetCovidSummary()

	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}

	context.JSON(http.StatusOK, response)
}
