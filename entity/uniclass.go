package entity

type UniClass struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	CreditNb    int    `json:"credit_nb"`
	Description string `json:"description"`
}
