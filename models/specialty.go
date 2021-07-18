package models

type Specialty struct {
	Id   int
	Name string
}

//  this is a function that I include with most models. Usually it is very simple but I include it anyway
func (sp Specialty) Validate() bool {
	if sp.Id <= 0 {
		return false
	}
	if len(sp.Name) <= 0 {
		return false
	}
	return true
}
