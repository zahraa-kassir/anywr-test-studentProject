package entity

import "fmt"

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
func (u Classes) String() string {
	return fmt.Sprintf("code: %v -name: %v -credit: %v", u.Code, u.Name, u.CreditNb)
}
