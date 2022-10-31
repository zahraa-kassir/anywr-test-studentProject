package dto

import (
	"test-pr/anywr-test-studentProject/entity"
	"test-pr/anywr-test-studentProject/payload"
)

type FilterData struct {
	PageData PaginationData      `json:"page_data"`
	Student  []SimpleStudentData `json:"student"`
}

// -----------------------------------------------------------------
// re-construct with dto

func ReconstructFilterData(stud []entity.Student, filterData payload.FilterByTeacherAndClass, nb int64, text string) FilterData {

	var students []SimpleStudentData
	for i, v := range stud {
		n := ReconstructSimpleStudentData(i, v)
		students = append(students, n)
	}

	return FilterData{
		PageData: ReconstructPaginationData(filterData, nb, text),
		Student:  students,
	}
}

// -----------------------------------------------------------------
