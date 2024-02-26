package usecase_response

import "sync"

var mutex sync.Mutex

type CovidSummaryResponse struct {
	ProvinceGroupCount map[string]int `json:"Province"`
	AgeGroupCount      map[string]int `json:"AgeGroup"`
}

func NewCovidSummaryResponse() *CovidSummaryResponse {
	return &CovidSummaryResponse{
		ProvinceGroupCount: make(map[string]int),
		AgeGroupCount:      make(map[string]int),
	}
}

func (response *CovidSummaryResponse) IncreaseProviceCount(key string) {
	mutex.Lock()
	defer mutex.Unlock()
	response.ProvinceGroupCount[key]++
}

func (response *CovidSummaryResponse) IncreaseAgeCount(key string) {
	mutex.Lock()
	defer mutex.Unlock()
	response.AgeGroupCount[key]++
}
