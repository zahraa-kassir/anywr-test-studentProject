package dto

type UniClassData struct {
	Class    SimpleClassData     `json:"class"`
	Students []SimpleStudentData `json:"students"`
	Teachers []SimpleTeacherData `json:"teachers"`
}
