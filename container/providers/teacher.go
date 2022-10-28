package providers

import (
	"gorm.io/gorm"
	"test-pr/anywr-test-studentProject/controller"
	"test-pr/anywr-test-studentProject/repository"
)

func TeachRepo(db *gorm.DB) repository.TeacherRepository {
	return repository.TeacherRepository{
		DB: db,
	}
}

func TeachCont(rep repository.TeacherRepository) controller.TeacherController {
	return controller.TeacherController{
		TeacherRepository: rep,
	}
}
