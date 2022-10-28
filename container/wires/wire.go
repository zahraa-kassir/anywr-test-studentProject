package wires

import (
	"gorm.io/gorm"
	"test-pr/anywr-test-studentProject/controller"
	"test-pr/anywr-test-studentProject/repository"
)

func InitialiseTeachCont(db *gorm.DB) controller.TeacherController {
	teacherRepository := repository.TeachRepo(db)
	teacherController := controller.TeachCont(teacherRepository)
	return teacherController
}

func InitialiseStudCont(db *gorm.DB) controller.StudentController {
	studentRepository := repository.StudRepo(db)
	classRepository := repository.ClassRepo(db)
	teacherRepository := repository.TeachRepo(db)
	studentController := controller.StudCont(studentRepository, classRepository, teacherRepository)
	return studentController
}

func InitialiseClassCont(db *gorm.DB) controller.ClassController {
	classRepository := repository.ClassRepo(db)
	studentRepository := repository.StudRepo(db)
	teacherRepository := repository.TeachRepo(db)
	classController := controller.ClassCont(classRepository, studentRepository, teacherRepository)
	return classController
}
