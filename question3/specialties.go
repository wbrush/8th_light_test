package question3

import (
	"errors"
	"strconv"

	"github.com/wbrush/8th_light_test/models"
	"github.com/wbrush/8th_light_test/question1"
	"github.com/wbrush/8th_light_test/question2"
)

const specialtyMissingIdError = "Specialties list doesn't contain id #"

func GetSpecialty(ids []string, specialties []models.Specialty) (list []string, err error) {
	if list == nil {
		list = make([]string, 0)
	}

	//  convert specialties list to map
	specialtiesMap := make(map[int]string)
	for i := range specialties {
		specialtiesMap[specialties[i].Id] = specialties[i].Name
	}

	//  remove non-numeric characters from the string
	cleanIds := make([]string, 0)
	for i := range ids {
		cleanId, rerr := question1.RemoveNonNumbers(ids[i])
		if rerr != nil {
			err = rerr
			return
		}
		cleanIds = append(cleanIds, cleanId)
	}

	//  remove duplicate numbers
	validIds, err := question2.RemoveDuplicates(cleanIds)
	if err != nil {
		return
	}

	//  for each id
	processedList := make([]string, 0)
	for _, item := range validIds.([]string) {

		//  convert to integer
		id, rerr := strconv.Atoi(item)
		if rerr != nil {
			err = rerr
			return
		}

		//  get specialty for id
		value, present := specialtiesMap[id]
		if !present {
			err = errors.New(specialtyMissingIdError + item)
			return
		}

		processedList = append(processedList, value)
	}

	list = processedList
	return
}
