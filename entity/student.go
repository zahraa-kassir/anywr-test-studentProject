package entity

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
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

// ByPageNum offset and limit
func ByPageNum(page, size string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		//handle and convert page numb to int
		pageNb, _ := strconv.Atoi(page)
		if pageNb == 0 {
			pageNb = 1
		}
		//handle and convert pageSize to int
		pageSize, _ := strconv.Atoi(size)
		//set offset based on pageSize jump x nb
		offset := (pageNb - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}

}

// -----------------------------------------------------------------

func (s Student) String() string {
	return fmt.Sprintf("Student: %v , Email @: %v ", s.Name, s.Email)
}
