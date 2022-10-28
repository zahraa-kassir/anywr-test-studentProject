package payload

type FilterByTeacherAndClass struct {
	TeachEmail string `json:"teach_email"`
	ClassCode  string `json:"class_code"`
	Page       string `json:"page"`
	PageSize   string `json:"page_size"`
}
