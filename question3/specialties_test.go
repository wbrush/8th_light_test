package question3

import (
	"errors"
	"testing"

	"github.com/wbrush/8th_light_test/models"
)

// func GetSpecialty(ids []string, specialties []models.Specialty) (list []string, err error)
func TestQuestion3_RemoveNonNumbers(t *testing.T) {
	type givenData struct {
		ids         []string
		specialties []models.Specialty
	}
	type resultData struct {
		list []string
		err  error
	}
	tests := []struct {
		testName       string
		givenInput     givenData
		expectedOutput resultData
	}{
		{
			testName: "Happy Path #1",
			givenInput: givenData{
				ids: []string{"7-623", "8235", "8-235"},
				specialties: []models.Specialty{
					{Id: 1381, Name: "front-end web development"},
					{Id: 8235, Name: "data engineering"},
					{Id: 3434, Name: "API design"},
					{Id: 7623, Name: "security"},
					{Id: 9153, Name: "UX"},
				},
			},
			expectedOutput: resultData{
				list: []string{"security", "data engineering"},
				err:  nil,
			},
		},
		{
			testName: "No ids given",
			givenInput: givenData{
				ids: []string{},
				specialties: []models.Specialty{
					{Id: 1381, Name: "front-end web development"},
					{Id: 8235, Name: "data engineering"},
					{Id: 3434, Name: "API design"},
					{Id: 7623, Name: "security"},
					{Id: 9153, Name: "UX"},
				},
			},
			expectedOutput: resultData{
				list: []string{},
				err:  nil,
			},
		},
		{
			testName: "No specialties given",
			givenInput: givenData{
				ids:         []string{"7-623", "8235", "8-235"},
				specialties: []models.Specialty{},
			},
			expectedOutput: resultData{
				list: []string{},
				err:  errors.New(specialtyMissingIdError + "7623"),
			},
		},
		{
			testName: "id had no number when non-numeric characters are removed",
			givenInput: givenData{
				ids: []string{"-=(*%^$!@", "8235", "8-235"},
				specialties: []models.Specialty{
					{Id: 1381, Name: "front-end web development"},
					{Id: 8235, Name: "data engineering"},
					{Id: 3434, Name: "API design"},
					{Id: 7623, Name: "security"},
					{Id: 9153, Name: "UX"},
				},
			},
			expectedOutput: resultData{
				list: []string{},
				err:  errors.New("strconv.Atoi: parsing \"\": invalid syntax"),
			},
		},
		{
			testName: "have id but there is no matching specialty",
			givenInput: givenData{
				ids: []string{"7-623", "5", "8235", "8-235"},
				specialties: []models.Specialty{
					{Id: 1381, Name: "front-end web development"},
					{Id: 8235, Name: "data engineering"},
					{Id: 3434, Name: "API design"},
					{Id: 7623, Name: "security"},
					{Id: 9153, Name: "UX"},
				},
			},
			expectedOutput: resultData{
				list: []string{},
				err:  errors.New(specialtyMissingIdError + "5"),
			},
		},
		{ //  Note: requirements state that there will be NO duplicates in the specialties list. However, since this may be changed in the future, include a test so that it will fail if changed
			testName: "have id but there is no matching specialty",
			givenInput: givenData{
				ids: []string{"7-623", "8235", "8-235"},
				specialties: []models.Specialty{
					{Id: 1381, Name: "front-end web development"},
					{Id: 8235, Name: "data engineering"},
					{Id: 3434, Name: "API design"},
					{Id: 7623, Name: "security"},
					{Id: 8235, Name: "data engineering"}, //  note that this shouldn't affect the current operation
					{Id: 9153, Name: "UX"},
				},
			},
			expectedOutput: resultData{
				list: []string{"security", "data engineering"},
				err:  nil,
			},
		},

		{ //  Note: requirements state that the same id will not have multiple specialties. However, since this may be changed in the future, include a test so that it will fail if changed
			testName: "have id but there is no matching specialty",
			givenInput: givenData{
				ids: []string{"7-623", "8235", "8-235"},
				specialties: []models.Specialty{
					{Id: 1381, Name: "front-end web development"},
					{Id: 8235, Name: "data engineering"},
					{Id: 3434, Name: "API design"},
					{Id: 7623, Name: "security"},
					{Id: 8235, Name: "cloud engineering"},
					{Id: 9153, Name: "UX"},
				},
			},
			expectedOutput: resultData{
				list: []string{"security", "cloud engineering"}, //  we will guess and take the last one in the list
				err:  nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result, err := GetSpecialty(tt.givenInput.ids, tt.givenInput.specialties)
			if (err != nil && tt.expectedOutput.err != nil && tt.expectedOutput.err.Error() != err.Error()) ||
				(err == nil && tt.expectedOutput.err != nil) ||
				(err != nil && tt.expectedOutput.err == nil) {
				if err != nil {
					t.Error("Expected Error doesn't match resulting error! err = " + err.Error())
				} else {
					t.Error("Expected Error doesn't match resulting error!")
				}
			}
			if len(result) != len(tt.expectedOutput.list) {
				if result != nil {
					t.Errorf("Expected output slice length (%d) doesn't match processed result length (%d)!", len(tt.expectedOutput.list), len(result))
				} else {
					t.Errorf("Result length is nil when we expected a result array!")
				}
			} else {
				for i, item := range tt.expectedOutput.list {
					if item != result[i] {
						t.Errorf("Expected output (%s) doesn't match processed result (%s)", item, result[i])
					}
				}
			}
		})
	}
}
