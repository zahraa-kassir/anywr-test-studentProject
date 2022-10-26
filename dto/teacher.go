package dto

type Teacher struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Class []Class `json:"class"`
}
