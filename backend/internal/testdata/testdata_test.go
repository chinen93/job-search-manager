package testdata

import (
	"job-search-manager/internal/dao"
	"reflect"
	"testing"
)

func Test_DataCreation(t *testing.T) {

	dao.Init()
	defer dao.Close()
	InitTest()

	t.Run("Get All Companies", func(t *testing.T) {
		result, err := companyController.GetAll()

		if err != nil {
			t.Fatal("Companies not found", err)
		}

		for index := 0; index < len(result); index++ {
			element := *result[index]
			expected := Companies[index]

			if !reflect.DeepEqual(element, expected) {
				t.Error("Companies does not match", element, expected)
			}
		}
	})

	t.Run("Get All Jobs", func(t *testing.T) {
		result, err := jobController.GetAll()

		if err != nil {
			t.Fatal("Jobs not found", err)
		}

		for index := 0; index < len(result); index++ {
			element := *result[index]
			expected := Jobs[index]

			if !reflect.DeepEqual(element, expected) {
				t.Error("Jobs does not match", element, expected)
			}
		}
	})
}
