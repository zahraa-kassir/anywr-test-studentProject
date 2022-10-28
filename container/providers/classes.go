package providers

import (
	"gorm.io/gorm"
	"test-pr/anywr-test-studentProject/controller"
	"test-pr/anywr-test-studentProject/repository"
)

func ClassRepo(db *gorm.DB) repository.ClassRepository {
	return repository.ClassRepository{
		DB: db,
	}
}

func ClassCont(db *gorm.DB, rep repository.ClassRepository) controller.ClassController {
	return controller.ClassController{
		ClassRepository:   rep,
		StudentRepository: StudRepo(db),
		TeacherRepository: TeachRepo(db),
	}
}
