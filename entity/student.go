package entity

import (
	"fmt"
	"gorm.io/gorm"
)

type Student struct {
	Id       int `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
	Class    int     `gorm:"index"`
	Classes  Classes `gorm:"foreignKey:id;references:class"`
}

func (s Student) TableName() string {
	return "students"
}

// -----------------------------------------------------------------
// Scopes
//

// ByStudentEmail where student.email
func ByStudentEmail(email string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("students.email = ?", email)
	}
}

// ByStudentClassId where student.class
func ByStudentClassId(id int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("students.class = ? ", id)
	}

}

// -----------------------------------------------------------------

func (s Student) String() string {
	return fmt.Sprintf("Student: %v , Email @: %v ", s.Name, s.Email)
}
