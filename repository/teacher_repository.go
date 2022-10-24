package repository

import (
	"gorm.io/gorm"
	"test-pr/anywr-test-studentProject/entity"
)

type TeacherRepository struct {
	DB *gorm.DB
}

func (t TeacherRepository) GetAll() []entity.Teacher {
	var teachers []entity.Teacher
	_ = t.DB.Find(&teachers)
	return teachers
}

func (t TeacherRepository) GetByEmail(email string) entity.Teacher {
	var teacher = entity.Teacher{Email: email}
	_ = t.DB.First(&teacher)
	return teacher
}
