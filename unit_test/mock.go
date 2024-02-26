package unittest

import (
	"lmwn-assignment/domain/entities"
)

var (
	Age1      = 10
	Age2      = 40
	Age3      = 80
	Age4      = 0
	Province1 = "Phrae"
	Province2 = "Suphan Buri"
	Province3 = "Roi Et"
	Province4 = "Bangkok"
)

type MockCovidRepository struct{}
type MockCovidRepositoryWithEmptyCovidRecord struct{}

func NewMockCovidRepository() *MockCovidRepository {
	return &MockCovidRepository{}
}

func NewMockCovidRepositoryWithEmptyCovidRecord() *MockCovidRepositoryWithEmptyCovidRecord {
	return &MockCovidRepositoryWithEmptyCovidRecord{}
}

func (m *MockCovidRepository) GetCovidCaseRecord() (*[]entities.Covid, error) {
	mock := getCovidEntityMock()
	return &mock, nil
}

func (m *MockCovidRepositoryWithEmptyCovidRecord) GetCovidCaseRecord() (*[]entities.Covid, error) {
	return &[]entities.Covid{}, nil
}

func getCovidEntityMock() []entities.Covid {
	return []entities.Covid{
		{
			ConfirmDate:    "2021-05-04",
			No:             0,
			Age:            &Age1,
			Gender:         "หญิง",
			GenderEn:       "Female",
			Nation:         "",
			NationEn:       "China",
			Province:       Province1,
			ProvinceId:     46,
			District:       "",
			ProvinceEn:     Province1,
			StatQuarantine: 5,
		},
		{
			ConfirmDate:    "2021-05-04",
			No:             0,
			Age:            &Age2,
			Gender:         "หญิง",
			GenderEn:       "Female",
			Nation:         "",
			NationEn:       "China",
			Province:       Province2,
			ProvinceId:     46,
			District:       "",
			ProvinceEn:     Province2,
			StatQuarantine: 5,
		},
		{
			ConfirmDate:    "2021-05-04",
			No:             0,
			Age:            &Age3,
			Gender:         "หญิง",
			GenderEn:       "Female",
			Nation:         "",
			NationEn:       "China",
			Province:       "",
			ProvinceId:     46,
			District:       "",
			ProvinceEn:     "",
			StatQuarantine: 5,
		},
		{
			ConfirmDate:    "2021-05-04",
			No:             0,
			Age:            &Age4,
			Gender:         "หญิง",
			GenderEn:       "Female",
			Nation:         "",
			NationEn:       "China",
			Province:       Province3,
			ProvinceId:     46,
			District:       "",
			ProvinceEn:     Province3,
			StatQuarantine: 5,
		},
		{
			ConfirmDate:    "2021-05-04",
			No:             0,
			Age:            nil,
			Gender:         "หญิง",
			GenderEn:       "Female",
			Nation:         "",
			NationEn:       "China",
			Province:       Province4,
			ProvinceId:     46,
			District:       "",
			ProvinceEn:     Province4,
			StatQuarantine: 5,
		},
	}
}
