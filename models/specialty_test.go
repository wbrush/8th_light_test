package models

import "testing"

// func (sp Specialty) Validate() (bool)
func TestQuestion1_Validate(t *testing.T) {
	tests := []struct {
		testName       string
		givenInput     Specialty
		expectedOutput bool
	}{
		{
			testName: "Happy Path #1",
			givenInput: Specialty{
				Id:   1,
				Name: "name1",
			},
			expectedOutput: true,
		},
		{
			testName: "Negative ID",
			givenInput: Specialty{
				Id:   -13,
				Name: "name1",
			},
			expectedOutput: false,
		},
		{
			testName: "Zero ID",
			givenInput: Specialty{
				Id:   0,
				Name: "name1",
			},
			expectedOutput: false,
		},
		{
			testName: "Zero Name",
			givenInput: Specialty{
				Id:   5,
				Name: "",
			},
			expectedOutput: false,
		},
		{
			testName: "Large values",
			givenInput: Specialty{
				Id:   999999999999999999,
				Name: "the brown fox jumps over the lazy dog on Sunday when everyone is supposed to be resting!",
			},
			expectedOutput: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			record := tt.givenInput
			result := record.Validate()
			if result != tt.expectedOutput {
				t.Errorf("Expected result doesn't match expected result for id %d!", tt.givenInput.Id)
			}
		})
	}
}
