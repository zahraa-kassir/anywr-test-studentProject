package dto

import "test-pr/anywr-test-studentProject/entity"

type SimpleClassData struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Credit int    `json:"credit"`
}

// -----------------------------------------------------------------
// re-construct with dto

func ReconstructSimpleClass(data entity.Classes) SimpleClassData {
	cls := SimpleClassData{
		Code:   data.Code,
		Name:   data.Name,
		Credit: data.CreditNb,
	}
	return cls
}

// -----------------------------------------------------------------
