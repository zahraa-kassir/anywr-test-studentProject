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
	_ = c.DB.First(&class)
	return class
}

func (c ClassRepository) GetByCode(code string) entity.UniClass {
	var class = entity.UniClass{Code: code}
	_ = c.DB.First(&class)
	return class
}
