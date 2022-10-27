package entity

import (
	"fmt"
	"gorm.io/gorm"
)

type Teacher struct {
	Id       int `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
	Classes  []Classes `gorm:"many2many:teacher_classes;joinForeignKey:teach_id;joinReferences:class_id"`
}

func (t Teacher) TableName() string {
	return "teachers"
}

func ByTeachEmail(email string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("teachers.email = ? ", email)
	}
}

func (t Teacher) ToString() string {
	return fmt.Sprintf("Teacher: %v , Email@: %v", t.Name, t.Email)
}
