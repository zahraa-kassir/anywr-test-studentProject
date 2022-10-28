package dto

type FilterData struct {
	Teacher SimpleData      `json:"teacher"`
	Class   SimpleClassData `json:"class"`
	Student []SimpleData    `json:"student"`
}
