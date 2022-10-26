package dto

type StudentData struct {
	Student Student `json:"student"`
	Classes Class   `json:"class"`
}
