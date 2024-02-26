package unittest

import (
	"fmt"
	"lmwn-assignment/application/model"
	usecase_response "lmwn-assignment/application/response"
	"lmwn-assignment/application/usecase"
	"testing"
)

func TestCovidUseCase(test *testing.T) {

	test.Run("Match Expected - Expect Success", func(t *testing.T) {
		mockRepo := NewMockCovidRepository()
		useCase := usecase.NewCovidUseCase(mockRepo)
		result, err := useCase.GetCovidSummary()

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result == nil {
			t.Errorf("Result should not be nil")
		}

		expectResult := usecase_response.NewCovidSummaryResponse()

		expectResult.ProvinceGroupCount[Province1] = 1
		expectResult.ProvinceGroupCount[Province2] = 1
		expectResult.ProvinceGroupCount[Province3] = 1
		expectResult.ProvinceGroupCount[Province4] = 1
		expectResult.ProvinceGroupCount[model.Na] = 1

		expectResult.AgeGroupCount[model.Na] = 1
		expectResult.AgeGroupCount[model.LessThan30] = 2
		expectResult.AgeGroupCount[model.MoreThan30LessThan60] = 1
		expectResult.AgeGroupCount[model.MoreThan60] = 1

		if !compareCovidSummaryResponses(result, expectResult) {
			t.Errorf("Result doesn't match expected")
		}
	})

	test.Run("Covid Record is Empty - Expect Success", func(t *testing.T) {
		mockRepo := NewMockCovidRepositoryWithEmptyCovidRecord()
		useCase := usecase.NewCovidUseCase(mockRepo)
		result, err := useCase.GetCovidSummary()

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result == nil {
			t.Errorf("Result should not be nil")
		}

		expectResult := usecase_response.NewCovidSummaryResponse()

		fmt.Println(result)
		fmt.Println(expectResult)

		if !compareCovidSummaryResponses(result, expectResult) {
			t.Errorf("Result doesn't match expected")
		}
	})
}

func compareCovidSummaryResponses(result, expectResult *usecase_response.CovidSummaryResponse) bool {
	if len(result.ProvinceGroupCount) != len(expectResult.ProvinceGroupCount) || len(result.AgeGroupCount) != len(expectResult.AgeGroupCount) {
		return false
	}

	for province, count := range result.ProvinceGroupCount {
		if expectCount, exist := expectResult.ProvinceGroupCount[province]; !exist || expectCount != count {
			return false
		}
	}

	for age, count := range result.AgeGroupCount {
		if expectCount, exist := expectResult.AgeGroupCount[age]; !exist || expectCount != count {
			return false
		}
	}

	return true
}
