package repository

import (
	"gorm.io/gorm"
	"test-pr/anywr-test-studentProject/entity"
)

type ClassRepository struct {
	DB *gorm.DB
}

// GetAll return all data of classes
func (c ClassRepository) GetAll() []entity.Classes {
	var class []entity.Classes
	_ = c.DB.Find(&class)
	return class
}

func (c ClassRepository) GetById(id int) entity.Classes {
	var class entity.Classes
	_ = c.DB.Scopes(byClassesId(id)).Take(&class)
	return class
}

func (c ClassRepository) GetByCode(code string) *entity.Classes {
	var class entity.Classes
	_ = c.DB.Scopes(byClassCode(code)).Take(&class)
	return &class
}

func (c ClassRepository) GetStByClCode(code string) entity.Classes {
	var class entity.Classes

	if obj := c.GetByCode(code); obj.Id != 0 {
		_ = c.DB.
			Scopes(byClassesId(obj.Id)).
			Preload("Student", "class =?", obj.Id).
			Preload("Teacher").
			Joins("inner join teacher_classes as tc on tc.class_id = classes.id and tc.class_id = ? ", obj.Id).
			Take(&class)

	}
	return class
}

func byClassCode(code string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("classes.code = ? ", code)
	}
}
func byClassesId(id int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("classes.id = ? ", id)
	}
}
