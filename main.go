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
	db.Connect("postgres://postgres:admin@localhost:5432/anywr-student-project")

	//Routes
	classController := controller.ClassController{
		ClassRepository: repository.ClassRepository{
			DB: db.Instance,
		},
	}

	StudentController := controller.StudentController{
		StudentRepository: repository.StudentRepository{
			DB: db.Instance,
		},
	}

	TeacherController := controller.TeacherContoller{
		TeacherRepository: repository.TeacherRepository{
			DB: db.Instance,
		},
	}

	//main group
	maingrp := e.Group("/anywrUni")

	//class option
	classgrp := maingrp.Group("/UniClass")
	classgrp.GET("", classController.GetAll)
	classgrp.GET("/:id", classController.GetById)
	classgrp.GET("/:code", classController.GetByCode)

	//student option
	stdgrp := maingrp.Group("/st")
	stdgrp.GET("", StudentController.GetAll)
	stdgrp.GET("/:email", StudentController.GetByEmail)

	//teacher option
	teachgrp := maingrp.Group("/tea")
	teachgrp.GET("", TeacherController.GetAll)
	teachgrp.GET("/:email", TeacherController.GetByEmail)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
