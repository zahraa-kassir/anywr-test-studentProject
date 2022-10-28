package dto

type GetAllTeachers struct {
	Teacher SimpleData        `json:"teacher"`
	Class   []SimpleClassData `json:"class"`
}
