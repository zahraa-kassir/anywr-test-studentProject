package dto

type StudentByClass struct {
	Code     string              `json:"code"`
	Name     string              `json:"name"`
	Students []SimpleStudentData `json:"students"`
}
