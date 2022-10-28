package injectors

import (
	"gorm.io/gorm"
	"test-pr/anywr-test-studentProject/container/providers"
	"test-pr/anywr-test-studentProject/controller"
)

func InitialiseTeachCont(db *gorm.DB) controller.TeacherController {
	teacherRepository := providers.TeachRepo(db)
	teacherController := providers.TeachCont(teacherRepository)
	return teacherController
}
