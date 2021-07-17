package question2

import (
	"errors"

	"github.com/sirupsen/logrus"
)

const unsupportArrayTypeMsg = "Unsupported array type!"

//  Note that this is capitalized so that it can be accessed by other packages. I split this up to make it easier to
//  modify in the future as well as add more supported data types if needed. It also allowed me to create the unit
//  tests more easily.
func RemoveDuplicates(inputData interface{}) (result interface{}, err error) {

	switch inputData.(type) {
	case []int:
		result, err = removeDuplicateIntegers(inputData.([]int))
		break
	case []string:
		result, err = removeDuplicateStrings(inputData.([]string))
		break
	default:
		result = inputData
		err = errors.New(unsupportArrayTypeMsg)
	}

	return
}

//  this funtion is local only
func removeDuplicateStrings(inputData []string) (result []string, err error) {
	logrus.Debug("[]string")

	stringMap := make(map[string]bool)
	if result == nil {
		result = make([]string, 0)
	}

	for _, key := range inputData {
		_, ok := stringMap[key]
		if ok {
			continue
		}
		result = append(result, key)
		stringMap[key] = true
	}

	return
}

//  this funtion is local only
func removeDuplicateIntegers(inputData []int) (result []int, err error) {
	logrus.Debug("[]int1")

	intMap := make(map[int]bool)
	if result == nil {
		result = make([]int, 0)
	}

	for _, key := range inputData {
		_, ok := intMap[key]
		if ok {
			continue
		}
		result = append(result, key)
		intMap[key] = true
	}

	return
}
