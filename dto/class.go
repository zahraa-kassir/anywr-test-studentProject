package dto

type Class struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	CreditNb    int    `json:"credit_nb"`
	Description string `json:"description"`
}
