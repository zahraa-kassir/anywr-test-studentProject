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
	_ = s.DB.Find(&students)
	return students
}

func (s StudentRepository) GetByEmail(email string) entity.Student {
	var student = entity.Student{Email: email}
	_ = s.DB.First(&student)
	return student
}
