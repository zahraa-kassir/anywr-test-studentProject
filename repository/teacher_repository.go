package repository

import (
	"gorm.io/gorm"
	"test-pr/anywr-test-studentProject/entity"
)

type TeacherRepository struct {
	DB *gorm.DB
}

func TeachRepo(db *gorm.DB) TeacherRepository {
	return TeacherRepository{
		DB: db,
	}
}

func (t TeacherRepository) GetAll() []entity.Teacher {
	var teachers []entity.Teacher
	_ = t.DB.Preload("Classes").
		Find(&teachers)
	return teachers
}

func (t TeacherRepository) GetByEmail(email string) *entity.Teacher {
	var teacher entity.Teacher
	if dbc := t.DB.Preload("Classes").Scopes(entity.ByTeachEmail(email)).First(&teacher); dbc.Error != nil {
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
