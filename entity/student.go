package entity

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type Student struct {
	Id       int `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
	Class    int     `gorm:"index"`
	Classes  Classes `gorm:"foreignKey:id;references:class"`
	Notes    int
}

func (s Student) TableName() string {
	//it is better to name db table in singular as each record represents a singular student
	return "students"
}

// -----------------------------------------------------------------
// Scopes
//

// ByStudentEmail where student.email
func ByStudentEmail(email string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		//it is better here to check if email is empty or not before doing any query
		return db.Where("students.email = ?", email)
	}
}

// ByStudentClassId where student.class
func ByStudentClassId(id int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if id == 0 {
			return db
		}
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

// ByRank by rank max or min
func ByRank(db *gorm.DB, rank string) *gorm.DB {
	//handle the rank text , verifie
	rType := strings.ToLower(rank)

	if rType == "min" {
		return db.Select("students.class as class,Min(students.notes) as notes").Group("students.class").Table("students")
	} else if rType == "max" {
		return db.Select("students.class as class,Max(students.notes) as notes").Group("students.class").Table("students")
	}

	return db.Select("")
}

// BySort sorting acs des
func BySort(sort string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		rType := strings.ToLower(sort)

		if rType == "desc" {
			return db.Order("students.name DESC")
		}

		return db.Order("students.name ASC")
	}
}

// -----------------------------------------------------------------

func (s Student) String() string {
	return fmt.Sprintf("Student: %v , Email @: %v ", s.Name, s.Email)
}
