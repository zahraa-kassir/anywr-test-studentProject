package dto

import "test-pr/anywr-test-studentProject/entity"

type Student struct {
	Id    int             `json:"id"`
	Name  string          `json:"name"`
	Email string          `json:"email"`
	Class SimpleClassData `json:"class"`
}

// -----------------------------------------------------------------
// re-construct with dto

func ReconstructStudent(data entity.Student) Student {
	std := Student{
		Id:    data.Id,
		Name:  data.Name,
		Email: data.Email,
		Class: ReconstructSimpleClass(data.Classes),
	}
	return std
}

// -----------------------------------------------------------------
