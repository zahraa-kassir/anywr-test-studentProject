package dto

type UniClassData struct {
	Class    Class               `json:"class"`
	Students []SimpleStudentData `json:"students"`
	Teachers []SimpleTeacherData `json:"teachers"`
}
