package dto

type FilterData struct {
	Teacher SimpleTeacherData   `json:"teacher"`
	Class   SimpleClassData     `json:"class"`
	Student []SimpleStudentData `json:"student"`
}
