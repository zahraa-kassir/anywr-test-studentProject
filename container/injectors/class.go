package injectors

import (
	"gorm.io/gorm"
	"test-pr/anywr-test-studentProject/container/providers"
	"test-pr/anywr-test-studentProject/controller"
)

func InitialiseClassCont(db *gorm.DB) controller.ClassController {
	classRepository := providers.ClassRepo(db)
	classController := providers.ClassCont(db, classRepository)
	return classController
}
