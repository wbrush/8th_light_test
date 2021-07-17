package question1

import "testing"

// func RemoveNonNumbers(inputString string) (outputString string, err error)
func TestQuestion1_RemoveNonNumbers(t *testing.T) {
	tests := []struct {
		testName       string
		givenInput     string
		expectedOutput string
		processError   error
	}{
		{
			testName:       "Happy Path #1",
			givenInput:     "7-623",
			expectedOutput: "7623",
			processError:   nil,
		},
		{
			testName:       "Happy Path #2",
			givenInput:     "..2965a",
			expectedOutput: "2965",
			processError:   nil,
		},
		{
			testName:       "happy path 3 - string with only numbers",
			givenInput:     "654378",
			expectedOutput: "654378",
			processError:   nil,
		},
		{
			testName:       "edge case 1 - empty string",
			givenInput:     "",
			expectedOutput: "",
			processError:   nil,
		},
		{
			testName:       "edge case 2 - string without numbers",
			givenInput:     "asdfpoi^&*",
			expectedOutput: "",
			processError:   nil,
		},
		{
			testName:       "edge case 3 - string with special characters like CR",
			givenInput:     "123/r45",
			expectedOutput: "12345",
			processError:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result, err := RemoveNonNumbers(tt.givenInput)
			if tt.processError != err {
				t.Error("Expected Error doesn't match resulting error!")
			}
			if tt.expectedOutput != result {
				t.Errorf("Expected output (%s) doesn't match processed result (%s)", tt.expectedOutput, result)
			}
		})
	}
}
