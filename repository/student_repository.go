package repository

import (
	"gorm.io/gorm"
	"test-pr/anywr-test-studentProject/entity"
	"test-pr/anywr-test-studentProject/payload"
)

type StudentRepository struct {
	DB *gorm.DB
}

func StudRepo(db *gorm.DB) StudentRepository {
	return StudentRepository{
		DB: db,
	}
}

func (s StudentRepository) GetAll() []entity.Student {
	var students []entity.Student
	_ = s.DB.Preload("Classes").Find(&students)
	return students
}
func (s StudentRepository) GetByEmail(email string) *entity.Student {
	var student entity.Student
	if dbc := s.DB.Preload("Classes").Scopes(entity.ByStudentEmail(email)).Find(&student); dbc.Error != nil {
		return nil
	}
	return &student
}

func (s StudentRepository) GetByCode(id int, sorStudent string, sortTeach string) ([]entity.Student, int64) {
	var students []entity.Student
	dbc := s.DB.Scopes(entity.ByStudentClassId(id), entity.BySort(sorStudent)).Preload("Classes.Teacher", entity.ByTeachSort(sortTeach)).Find(&students)
	rowNum := dbc.RowsAffected
	if dbc.Error != nil {
		return nil, 0
	}
	return students, rowNum

}

func (s StudentRepository) GetByFilterTeacherAndClass(code int, data payload.FilterByTeacherAndClass, teacher entity.Teacher) ([]entity.Student, int64) {
	var students []entity.Student
	query := s.DB.
		Scopes(entity.ByStudentClassId(code),
			entity.ByPageNum(data.Page, data.PageSize),
			entity.BySort(data.SortBy.StudentName)).
		Preload("Classes.Teacher", entity.ByTeachSort(data.SortBy.TeacherName), "teachers.id = ?", teacher.Id).
		Find(&students)
	rowNum := query.RowsAffected
	if query.Error != nil {
		return nil, 0
	}

	return students, rowNum

}

func (s StudentRepository) GetByRank(classId int, data payload.FilterByTeacherAndClass) ([]entity.Student, int64) {
	var students []entity.Student

	query := s.DB.
		Scopes(entity.ByStudentClassId(classId),
			entity.ByPageNum(data.Page, data.PageSize),
			entity.BySort(data.SortBy.StudentName)).
		Preload("Classes.Teacher", entity.ByTeachSort(data.SortBy.TeacherName)).
		Joins("inner join (?) q on q.class = students.class and q.notes = students.notes", entity.ByRank(s.DB, data.RankBy)).
		Find(&students)
	rowNum := query.RowsAffected
	if query.Error != nil {
		return nil, 0
	}
	return students, rowNum
}
