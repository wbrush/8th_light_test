package question2

import (
	"errors"
	"testing"
)

// func RemoveDuplicates(inputData interface{}) (result interface{}, err error)
func TestQuestion1_RemoveDuplicates_1(t *testing.T) {
	//  NOTE: since the conversion is tested below, we just need to make sure that the interface{} declaration is working properly
	tests := []struct {
		testName       string
		givenInput     []int
		expectedOutput []int
		processError   error
	}{
		{
			testName:       "Happy Path - integer",
			givenInput:     []int{4, 4, 3, 2, 3, 1},
			expectedOutput: []int{4, 3, 2, 1},
			processError:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result, err := RemoveDuplicates(tt.givenInput)
			if (err != nil && tt.processError != nil && tt.processError.Error() != err.Error()) ||
				(err == nil && tt.processError != nil) ||
				(err != nil && tt.processError == nil) {
				t.Error("Expected Error doesn't match resulting error!")
			}
			if len(result.([]int)) != len(tt.expectedOutput) {
				t.Errorf("Expected output slice length (%d) doesn't match processed result length (%d)!", len(tt.expectedOutput), len(result.([]int)))
			}
			for i := range tt.expectedOutput {
				if tt.expectedOutput[i] != result.([]int)[i] {
					t.Errorf("Expected output (%d) doesn't match processed result (%d)", tt.expectedOutput[i], result.([]int)[i])
				}
			}
		})
	}
}

func TestQuestion1_RemoveDuplicates_2(t *testing.T) {
	tests := []struct {
		testName       string
		givenInput     []string
		expectedOutput []string
		processError   error
	}{
		{
			testName:       "Happy Path - strings",
			givenInput:     []string{"a", "b", "c", "a", "b", "d"},
			expectedOutput: []string{"a", "b", "c", "d"},
			processError:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result, err := RemoveDuplicates(tt.givenInput)
			if (err != nil && tt.processError != nil && tt.processError.Error() != err.Error()) ||
				(err == nil && tt.processError != nil) ||
				(err != nil && tt.processError == nil) {
				t.Error("Expected Error doesn't match resulting error!")
			}
			if len(result.([]string)) != len(tt.expectedOutput) {
				t.Errorf("Expected output slice length (%d) doesn't match processed result length (%d)!", len(tt.expectedOutput), len(result.([]string)))
			}
			for i := range tt.expectedOutput {
				if tt.expectedOutput[i] != result.([]string)[i] {
					t.Errorf("Expected output (%s) doesn't match processed result (%s)", tt.expectedOutput[i], result.([]string)[i])
				}
			}
		})
	}
}

func TestQuestion1_RemoveDuplicates_3(t *testing.T) {
	tests := []struct {
		testName       string
		givenInput     []float64
		expectedOutput []float64
		processError   error
	}{
		{
			testName:       "unsupported data type - return array as is without removing dupes",
			givenInput:     []float64{1.2, 1.3, 1.2, 1.5},
			expectedOutput: []float64{1.2, 1.3, 1.2, 1.5},
			processError:   errors.New(unsupportArrayTypeMsg),
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result, err := RemoveDuplicates(tt.givenInput)
			if (err != nil && tt.processError != nil && tt.processError.Error() != err.Error()) ||
				(err == nil && tt.processError != nil) ||
				(err != nil && tt.processError == nil) {
				t.Error("Expected Error doesn't match resulting error!")
			} else if len(result.([]float64)) != len(tt.expectedOutput) {
				t.Errorf("Expected output slice length (%d) doesn't match processed result length (%d)!", len(tt.expectedOutput), len(result.([]float64)))
			}
			for i := range tt.expectedOutput {
				if tt.expectedOutput[i] != result.([]float64)[i] {
					t.Errorf("Expected output (%f) doesn't match processed result (%f)", tt.expectedOutput[i], result.([]float64)[i])
				}
			}
		})
	}
}

// func removeDuplicateStrings(inputData []string) (result []string, err error)
func TestQuestion1_RemoveDuplicateStrings(t *testing.T) {
	tests := []struct {
		testName       string
		givenInput     []string
		expectedOutput []string
		processError   error
	}{
		{
			testName:       "Happy Path - strings",
			givenInput:     []string{"a", "b", "c", "a", "b", "d"},
			expectedOutput: []string{"a", "b", "c", "d"},
			processError:   nil,
		},
		{
			testName:       "Happy Path - descending order",
			givenInput:     []string{"d", "b", "a", "c", "b", "a"},
			expectedOutput: []string{"d", "b", "a", "c"},
			processError:   nil,
		},
		{
			testName:       "Happy Path - larger strings",
			givenInput:     []string{"dont tell", "anybody", "anybody", "cats are cute", "but dogs", "are mans best friend", "but dogs"},
			expectedOutput: []string{"dont tell", "anybody", "cats are cute", "but dogs", "are mans best friend"},
			processError:   nil,
		},
		{
			testName:       "strings - empty slice",
			givenInput:     []string{},
			expectedOutput: []string{},
			processError:   nil,
		},
		{
			testName:       "strings - no duplicates",
			givenInput:     []string{"a", "b", "c", "d"},
			expectedOutput: []string{"a", "b", "c", "d"},
			processError:   nil,
		},
		{
			testName:       "strings - fairly large array",
			givenInput:     []string{"a", "b", "c", "d", "a", "b", "c", "d", "g", "a", "b", "c", "d", "h", "a", "b", "c", "d", "i"},
			expectedOutput: []string{"a", "b", "c", "d", "g", "h", "i"},
			processError:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result, err := removeDuplicateStrings(tt.givenInput)
			if (err != nil && tt.processError != nil && tt.processError.Error() != err.Error()) ||
				(err == nil && tt.processError != nil) ||
				(err != nil && tt.processError == nil) {
				t.Error("Expected Error doesn't match resulting error!")
			}
			if len(result) != len(tt.expectedOutput) {
				t.Errorf("Expected output slice length (%d) doesn't match processed result length (%d)!", len(tt.expectedOutput), len(result))
			}
			for i := range tt.expectedOutput {
				if tt.expectedOutput[i] != result[i] {
					t.Errorf("Expected output[i] (%s) doesn't match processed result[i] (%s)", tt.expectedOutput[i], result[i])
				}
			}
		})
	}
}

// func removeDuplicateIntegers(inputData []int) (result []int, err error)
func TestQuestion2_RemoveDuplicateIntegers(t *testing.T) {
	tests := []struct {
		testName       string
		givenInput     []int
		expectedOutput []int
		processError   error
	}{
		{
			testName:       "Happy Path - ints",
			givenInput:     []int{4, 4, 3, 2, 3, 1},
			expectedOutput: []int{4, 3, 2, 1},
			processError:   nil,
		},
		{
			testName:       "Happy Path - ascending order",
			givenInput:     []int{1, 3, 2, 3, 4, 4},
			expectedOutput: []int{1, 3, 2, 4},
			processError:   nil,
		},
		{
			testName:       "Happy Path - empty slice",
			givenInput:     []int{},
			expectedOutput: []int{},
			processError:   nil,
		},
		{
			testName:       "Happy Path - no duplicates",
			givenInput:     []int{4, 3, 2, 1},
			expectedOutput: []int{4, 3, 2, 1},
			processError:   nil,
		},
		{
			testName:       "Happy Path - fairly large input array",
			givenInput:     []int{100, 99, 98, 99, 95, 90, 85, 80, 80, 101, 99, 97, 99, 96, 92, 84, 80, 80, 4, 3, 2, 1},
			expectedOutput: []int{100, 99, 98, 95, 90, 85, 80, 101, 97, 96, 92, 84, 4, 3, 2, 1},
			processError:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result, err := removeDuplicateIntegers(tt.givenInput)
			if (err != nil && tt.processError != nil && tt.processError.Error() != err.Error()) ||
				(err == nil && tt.processError != nil) ||
				(err != nil && tt.processError == nil) {
				t.Error("Expected Error doesn't match resulting error!")
			}
			if len(result) != len(tt.expectedOutput) {
				t.Errorf("Expected output slice length (%d) doesn't match processed result length (%d)!", len(tt.expectedOutput), len(result))
			}
			for i := range tt.expectedOutput {
				if tt.expectedOutput[i] != result[i] {
					t.Errorf("Expected output[i] (%d) doesn't match processed result[i] (%d)", tt.expectedOutput[i], result[i])
				}
			}
		})
	}
}
