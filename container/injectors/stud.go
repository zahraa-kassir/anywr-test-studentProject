package injectors

import (
	"gorm.io/gorm"
	"test-pr/anywr-test-studentProject/container/providers"
	"test-pr/anywr-test-studentProject/controller"
)

func InitialiseStudCont(db *gorm.DB) controller.StudentController {
	studentRepository := providers.StudRepo(db)
	studentController := providers.StudCont(db, studentRepository)
	return studentController
}
