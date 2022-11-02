package payload

// FilterByTeacherAndClass in the filters you are receiving as query params it is query tag not json
type FilterByTeacherAndClass struct {
	TeachEmail string `json:"teach_email"`
	ClassCode  string `json:"class_code"`
	Page       string `json:"page"`
	PageSize   string `json:"page_size"`
	SortBy     SortBy `json:"sort_by"`
	RankBy     string `json:"rank_by"`
}

type SortBy struct {
	StudentName string `json:"student_name"`
	TeacherName string `json:"teacher_name"`
	ClassCode   string `json:"class_code"`
}
