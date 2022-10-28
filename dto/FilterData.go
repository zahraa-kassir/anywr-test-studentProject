package dto

import "test-pr/anywr-test-studentProject/entity"

type FilterData struct {
	Teacher SimpleData      `json:"teacher"`
	Class   SimpleClassData `json:"class"`
	Student []SimpleData    `json:"student"`
}

// -----------------------------------------------------------------
// re-construct with dto

func ReconstructFilterData(teach entity.Teacher, class entity.Classes, stud []entity.Student) FilterData {

	var students []SimpleData
	for i, v := range stud {
		n := ReconstructSimpleData(i, v.Name, v.Email)
		students = append(students, n)
	}
	fl := FilterData{

		Teacher: ReconstructSimpleData(teach.Id, teach.Name, teach.Email),
		Class:   ReconstructSimpleClass(class),
		Student: students,
	}

	return fl
}

// -----------------------------------------------------------------
