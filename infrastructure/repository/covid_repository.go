package repository

import (
	"encoding/json"
	"io"
	"lmwn-assignment/domain/entities"
	"lmwn-assignment/infrastructure/configuration"
	repo_response "lmwn-assignment/infrastructure/response"

	"net/http"
)

type CovidRepository struct {
}

func NewCovidRepository() *CovidRepository {
	return &CovidRepository{}
}

func (_ *CovidRepository) GetCovidCaseRecord() (*[]entities.Covid, error) {

	response, err := http.Get(configuration.GetCovidUrl())

	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := new(repo_response.CovidCaseResponse)

	if err = json.Unmarshal(responseBody, result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}
