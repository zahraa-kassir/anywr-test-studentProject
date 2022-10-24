package repository

import (
	"gorm.io/gorm"
	"test-pr/anywr-test-studentProject/entity"
)

type ClassRepository struct {
	DB *gorm.DB
}

func (c ClassRepository) GetAll() []entity.UniClass {
	var class []entity.UniClass
	_ = c.DB.Find(&class)
	return class
}

func (c ClassRepository) GetById(id int) entity.UniClass {
	var class = entity.UniClass{Id: id}
	_ = c.DB.Take(&class)
	return class
}

func (c ClassRepository) GetByCode(code string) entity.UniClass {
	var class entity.UniClass
	_ = c.DB.Take(&class, "code = ?", code)
	return class
}

func (c ClassRepository) GetDataOfThisStudent(id int) []entity.UniClass {
	var classes []entity.UniClass
	if dbc := c.DB.Table("uni_classes").
		Joins("inner join student_classes sc on uni_classes.id = sc.class_id").
		Find(&classes, "stud_id = ? ", id); dbc.Error != nil {
		return nil
	}
	return classes
}
