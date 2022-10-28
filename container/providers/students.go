package providers

import (
	"gorm.io/gorm"
	"test-pr/anywr-test-studentProject/controller"
	"test-pr/anywr-test-studentProject/repository"
)

func StudRepo(db *gorm.DB) repository.StudentRepository {
	return repository.StudentRepository{
		DB: db,
	}
}

func StudCont(db *gorm.DB, rep repository.StudentRepository) controller.StudentController {
	return controller.StudentController{
		StudentRepository: rep,
		ClassRepository:   ClassRepo(db),
		TeacherRepository: TeachRepo(db),
	}
}
