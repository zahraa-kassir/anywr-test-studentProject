package dto

import "test-pr/anywr-test-studentProject/entity"

type ClassFoFilter struct {
	Code     string
	Name     string
	CreditNb int
	Teacher  []SimpleData
}

// -----------------------------------------------------------------
// re-construct with dto

func ReconstructClassFoFilter(data entity.Classes) ClassFoFilter {
	var teachers []SimpleData
	for i, v := range data.Teacher {
		t := SimpleData{
			Id:    i,
			Name:  v.Name,
			Email: v.Email,
		}
		teachers = append(teachers, t)
	}

	return ClassFoFilter{
		Code:     data.Code,
		Name:     data.Name,
		CreditNb: data.CreditNb,
		Teacher:  teachers,
	}
}

// -----------------------------------------------------------------
