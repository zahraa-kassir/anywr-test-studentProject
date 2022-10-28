//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"test-pr/anywr-test-studentProject/controller"
	"test-pr/anywr-test-studentProject/repository"
)

func InitialiseTeachCont(db *gorm.DB) *controller.TeacherController {
	wire.Build(controller.TeachCont, repository.TeachRepo)
	return &controller.TeacherController{}
}

func InitialiseStudCont(db *gorm.DB) *controller.StudentController {
	wire.Build(controller.StudCont, repository.StudRepo, repository.ClassRepo, repository.TeachRepo)
	return &controller.StudentController{}
}
func InitialiseClassCont(db *gorm.DB) *controller.ClassController {
	wire.Build(controller.ClassCont, repository.ClassRepo, repository.StudRepo, repository.TeachRepo)
	return &controller.ClassController{}
}
