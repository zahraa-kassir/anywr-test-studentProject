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
	if dbc := s.DB.Preload("Classes").Scopes(byStudentEmail(email)).Find(&student); dbc.Error != nil {
		return nil
	}
	return &student
}

func (s StudentRepository) GetByCode(id int) []entity.Student {
	var students []entity.Student
	if dbc := s.DB.Scopes(byStudentClassId(id)).Preload("Classes").Find(&students); dbc.Error != nil {
		return nil
	}
	return students

}

func (s StudentRepository) GetByFilter(code int) []entity.Student {

	var students []entity.Student

	if dbc := s.DB.
		//Where("students.class = ?", code).
		Scopes(byStudentClassId(code)).
		Preload("Classes", "id=?", code).
		Find(&students); dbc.Error != nil {
		return nil
	}

	return students

}

func byStudentEmail(email string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("students.email = ?", email)
	}
}

func byStudentClassId(id int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("students.class = ? ", id)
	}

}
