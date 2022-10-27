package dto

type StudentData struct {
	Student Student         `json:"student"`
	Classes SimpleClassData `json:"class"`
}
