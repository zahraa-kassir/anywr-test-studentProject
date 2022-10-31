package entity

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type Classes struct {
	Id          int `gorm:"primaryKey"`
	Code        string
	Name        string
	CreditNb    int
	Description string
	Student     []Student `gorm:"ForeignKey:class;"`                                                         // one to many
	Teacher     []Teacher `gorm:"many2many:teacher_classes;joinForeignKey:class_id;joinReferences:teach_id"` // many 2 many
}

func (u Classes) TableName() string {
	return "classes"
}

// -----------------------------------------------------------------
// Scopes
//

// ByClassCode where classes.code
func ByClassCode(code string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("classes.code = ? ", code)
	}
}

// ByClassesId where classes.id
func ByClassesId(id int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("classes.id = ? ", id)
	}
}

// ByClassSort sorting acs des
func ByClassSort(sort string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		rType := strings.ToLower(sort)

		if rType == "desc" {
			return db.Order("classes.code DESC")
		}

		return db.Order("classes.code ASC")
	}
}

// -----------------------------------------------------------------

func (u Classes) String() string {
	return fmt.Sprintf("code: %v -name: %v -credit: %v", u.Code, u.Name, u.CreditNb)
}
