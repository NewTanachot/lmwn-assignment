package usecase

import usecase_response "lmwn-assignment/application/response"

type ICovidUseCase interface {
	GetCovidSummary() (*usecase_response.CovidSummaryResponse, error)
}
