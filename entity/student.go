package entity

import (
	"fmt"
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

func (s Student) String() string {
	return fmt.Sprintf("Student: %v , Email @: %v ", s.Name, s.Email)
}
