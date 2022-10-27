package dto

import "test-pr/anywr-test-studentProject/entity"

type Class struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	CreditNb    int    `json:"credit_nb"`
	Description string `json:"description"`
}

// -----------------------------------------------------------------
// re-construct with dto

func ReconstructClass(data *entity.Classes) Class {
	cl := Class{
		Code:        data.Code,
		Name:        data.Name,
		CreditNb:    data.CreditNb,
		Description: data.Description,
	}
	return cl
}

// -----------------------------------------------------------------
