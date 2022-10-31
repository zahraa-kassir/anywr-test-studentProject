package entity

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
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

// -----------------------------------------------------------------
// Scopes
//

func ByTeachEmail(email string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("teachers.email = ? ", email)
	}
}

// ByTeachSort sorting acs des
func ByTeachSort(sort string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		rType := strings.ToLower(sort)

		if rType == "desc" {
			return db.Order("teachers.name DESC")
		}

		return db.Order("teachers.name ASC")
	}
}

// -----------------------------------------------------------------

func (t Teacher) ToString() string {
	return fmt.Sprintf("Teacher: %v , Email@: %v", t.Name, t.Email)
}
