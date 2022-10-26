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
	_ = t.DB.Preload("Classes").
		Find(&teachers)
	return teachers
}

func (t TeacherRepository) GetByEmail(email string) *entity.Teacher {
	var teacher entity.Teacher
	if dbc := t.DB.Preload("Classes").First(&teacher, "email = ? ", email); dbc.Error != nil {
		return nil
	}
	return &teacher
}

func (t TeacherRepository) GetByCode(code int) []*entity.Teacher {
	var teachers []*entity.Teacher
	if dbc := t.DB.Preload("Classes", "id = ? ", code).
		Joins("inner join teacher_classes as tc on tc.teach_id = teachers.id and tc.class_id = ? ", code).
		Find(&teachers); dbc.Error != nil {
		return nil
	}

	return teachers

}

func (t TeacherRepository) VerifieIfExist(code int) bool {
	if teachers := t.GetByCode(code); len(teachers) != 0 {
		return true
	}
	return false
}
