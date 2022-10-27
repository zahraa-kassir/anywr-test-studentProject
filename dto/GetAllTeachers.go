package dto

type GetAllTeachers struct {
	Teacher SimpleTeacherData `json:"teacher"`
	Class   []SimpleClassData `json:"class"`
}
