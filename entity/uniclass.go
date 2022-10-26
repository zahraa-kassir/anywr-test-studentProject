package entity

import "fmt"

type Classes struct {
	Id          int `gorm:"primaryKey"`
	Code        string
	Name        string
	CreditNb    int
	Description string
	//[] student one to many
	//[]teachers many 2 many
}

func (u Classes) TableName() string {
	return "classes"
}
func (u Classes) String() string {
	return fmt.Sprintf("code: %v -name: %v -credit: %v", u.Code, u.Name, u.CreditNb)
}
