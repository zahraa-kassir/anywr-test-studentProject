package dto

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Class Class  `json:"class"`
}
