package repository

import (
	"gorm.io/gorm"
	"test-pr/anywr-test-studentProject/entity"
)

type ClassRepository struct {
	DB *gorm.DB
}

func (c ClassRepository) GetAll() []entity.Classes {
	var class []entity.Classes
	_ = c.DB.Find(&class)
	return class
}

func (c ClassRepository) GetById(id int) entity.Classes {
	var class = entity.Classes{Id: id}
	_ = c.DB.Take(&class)
	return class
}

func (c ClassRepository) GetByCode(code string) *entity.Classes {
	var class entity.Classes
	_ = c.DB.Take(&class, "code = ?", code)
	return &class
}
