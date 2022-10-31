package dto

import "test-pr/anywr-test-studentProject/entity"

type SimpleStudentData struct {
	Id    int           `json:"id"`
	Name  string        `json:"name"`
	Email string        `json:"email"`
	Note  int           `json:"note"`
	Class ClassFoFilter `json:"class"`
}

// -----------------------------------------------------------------
// re-construct with dto

func ReconstructSimpleStudentData(count int, data entity.Student) SimpleStudentData {

	return SimpleStudentData{
		Id:    count,
		Name:  data.Name,
		Email: data.Email,
		Note:  data.Notes,
		Class: ReconstructClassFoFilter(data.Classes),
	}
}

// -----------------------------------------------------------------
