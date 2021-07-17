package question1

import (
	"regexp"
)

//  since the regexp.Compile command can return an error, I allowed for passing this error back to the caller. Prefer
//  to do this since sometimes we can hit the error and by not checking it can create issues. Especially if this routine
//  were modified in the future to make the removed characters changeable (i.e. have flag to remove only special characters,
//	remove only upper case, etc). Normally I would probably make the regexp an input parameter.
func RemoveNonNumbers(inputString string) (outputString string, err error) {

	//  quick search on google shows that we can use regexp to remove the unwanted characters; Note that Go allows us
	//  to cycle through the string characters if we needed or wanted to so we could have used a manual metod to do this.
	regExpPtr, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return
	}

	outputString = regExpPtr.ReplaceAllString(inputString, "")

	return
}
