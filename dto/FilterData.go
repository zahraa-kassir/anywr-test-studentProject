package dto

type FilterData struct {
	Teacher SimpleTeacherData   `json:"teacher"`
	Class   *Class              `json:"class"`
	Student []SimpleStudentData `json:"student"`
}
