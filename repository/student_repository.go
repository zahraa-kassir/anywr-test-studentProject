package repository

import (
	"gorm.io/gorm"
	"test-pr/anywr-test-studentProject/entity"
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

func (s StudentRepository) GetByCode(id int) []entity.Student {
	var students []entity.Student
	if dbc := s.DB.Scopes(entity.ByStudentClassId(id)).Preload("Classes").Find(&students); dbc.Error != nil {
		return nil
	}
	return students

}

func (s StudentRepository) GetByFilter(code int, page string, pageSize string, email string) []entity.Student {
	var students []entity.Student
	if dbc := s.DB.
		Scopes(entity.ByStudentClassId(code), entity.ByPageNum(page, pageSize)).
		Preload("Classes", "id=?", code).
		Joins("inner join classes on classes.id = students.class").
		Joins("inner join teacher_classes on teacher_classes.class_id = classes.id").
		Joins("inner join teachers on teachers.email = ?", email).
		Find(&students); dbc.Error != nil {
		return nil
	}

	return students

}
