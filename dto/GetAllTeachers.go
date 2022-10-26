package dto

type GetAllTeachers struct {
	Teacher SimpleTeacherData `json:"teacher"`
	Class   []Class           `json:"class"`
}
