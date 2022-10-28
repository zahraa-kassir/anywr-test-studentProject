package dto

import "test-pr/anywr-test-studentProject/entity"

type UniClassData struct {
	Class    SimpleClassData `json:"class"`
	Students []SimpleData    `json:"students"`
	Teachers []SimpleData    `json:"teachers"`
}

// -----------------------------------------------------------------
// re-construct with dto

func ReconstructUniClassData(class entity.Classes) UniClassData {
	var std []SimpleData
	var teach []SimpleData
	if class.Id != 0 {
		for i, v := range class.Student {
			student := ReconstructSimpleData(i, v.Name, v.Email)
			std = append(std, student)
		}
		for i, v := range class.Teacher {
			th := ReconstructSimpleData(i, v.Name, v.Email)
			teach = append(teach, th)
		}
	}

	uniClassData := UniClassData{
		Class:    ReconstructSimpleClass(class),
		Students: std,
		Teachers: teach,
	}

	return uniClassData
}

// -----------------------------------------------------------------
