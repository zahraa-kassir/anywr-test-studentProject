package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"test-pr/anywr-test-studentProject/controller"
	"test-pr/anywr-test-studentProject/db"
	"test-pr/anywr-test-studentProject/repository"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connect to the "bank" database
	db.Connect("postgres://postgres:admin@localhost:5432/anywrstudentproject")

	//Routes
	classController := controller.ClassController{
		ClassRepository: repository.ClassRepository{
			DB: db.Instance,
		},
		StudentRepository: repository.StudentRepository{
			DB: db.Instance,
		},
		TeacherRepository: repository.TeacherRepository{
			DB: db.Instance,
		},
	}

	StudentController := controller.StudentController{
		StudentRepository: repository.StudentRepository{
			DB: db.Instance,
		},
		ClassRepository: repository.ClassRepository{
			DB: db.Instance,
		},
		TeacherRepository: repository.TeacherRepository{
			DB: db.Instance,
		},
	}

	TeacherController := controller.TeacherContoller{
		TeacherRepository: repository.TeacherRepository{
			DB: db.Instance,
		},
	}

	//main group
	mainGrp := e.Group("/anywrUni")

	//class option
	classGrp := mainGrp.Group("/UniClass")
	classGrp.GET("", classController.GetAll)
	classGrp.GET("/id/:id", classController.GetById)
	classGrp.GET("/code/:code", classController.GetByCode)
	classGrp.GET("/dataCode", classController.GetStByClCode)

	//student option
	stdGrp := mainGrp.Group("/st")
	stdGrp.GET("", StudentController.GetAll)
	stdGrp.GET("/:email", StudentController.GetByEmail)
	stdGrp.GET("/code/:code", StudentController.GetByCode)
	stdGrp.GET("/filter", StudentController.GetByThAndClass)

	//teacher option
	teachGrp := mainGrp.Group("/tea")
	teachGrp.GET("", TeacherController.GetAll)
	teachGrp.GET("/:email", TeacherController.GetByEmail)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
