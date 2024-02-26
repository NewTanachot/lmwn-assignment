package usecase

import (
	"lmwn-assignment/application/model"
	irepository "lmwn-assignment/application/repository"
	usecase_response "lmwn-assignment/application/response"
	"sync"
)

var waitGroup sync.WaitGroup

type CovidUseCase struct {
	CovidRepository irepository.ICovidRepository
}

func NewCovidUseCase(repo irepository.ICovidRepository) *CovidUseCase {
	return &CovidUseCase{CovidRepository: repo}
}

func (useCase *CovidUseCase) GetCovidSummary() (*usecase_response.CovidSummaryResponse, error) {

	covidRecords, err := useCase.CovidRepository.GetCovidCaseRecord()

	if err != nil {
		return nil, err
	}

	covidSummaryResponse := usecase_response.NewCovidSummaryResponse()

	for _, record := range *covidRecords {

		waitGroup.Add(1)
		go checkProvince(covidSummaryResponse, record.Province)

		waitGroup.Add(1)
		go checkAge(covidSummaryResponse, record.Age)
	}

	waitGroup.Wait()
	return covidSummaryResponse, nil
}

func checkProvince(covidSummaryResponse *usecase_response.CovidSummaryResponse, key string) {
	defer waitGroup.Done()
	if key == "" {
		covidSummaryResponse.IncreaseProviceCount("N/A")
	} else {
		covidSummaryResponse.IncreaseProviceCount(key)
	}
}

func checkAge(covidSummaryResponse *usecase_response.CovidSummaryResponse, age *int) {
	defer waitGroup.Done()
	if age == nil {
		covidSummaryResponse.IncreaseAgeCount(model.Na)
	} else if 0 <= *age && *age <= 30 {
		covidSummaryResponse.IncreaseAgeCount(model.LessThan30)
	} else if 31 <= *age && *age <= 60 {
		covidSummaryResponse.IncreaseAgeCount(model.MoreThan30LessThan60)
	} else if 61 <= *age {
		covidSummaryResponse.IncreaseAgeCount(model.MoreThan60)
	}
}
