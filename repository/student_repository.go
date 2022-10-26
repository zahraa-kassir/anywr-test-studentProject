package repository

import (
	"gorm.io/gorm"
	"test-pr/anywr-test-studentProject/entity"
)

type StudentRepository struct {
	DB *gorm.DB
}

func (s StudentRepository) GetAll() []entity.Student {
	var students []entity.Student
	_ = s.DB.Preload("Classes").Find(&students)
	return students
}
func (s StudentRepository) GetByEmail(email string) *entity.Student {
	var student entity.Student
	if dbc := s.DB.Preload("Classes").Where("students.email = ?", email).Find(&student); dbc.Error != nil {
		return nil
	}
	return &student
}

func (s StudentRepository) GetByCode(code int) []entity.Student {
	var students []entity.Student
	if dbc := s.DB.Preload("Classes").Find(&students, "class = ?", code); dbc.Error != nil {
		return nil
	}
	return students

}

func (s StudentRepository) GetByFilter(code int) []entity.Student {

	var students []entity.Student

	if dbc := s.DB.
		Where("students.class = ?", code).
		Preload("Classes", "id = ? ", code).
		Find(&students); dbc.Error != nil {
		return nil
	}

	return students

}
