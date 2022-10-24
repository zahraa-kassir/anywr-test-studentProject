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

	//main group
	maingrp := e.Group("/anywrUni")

	maingrp.GET("/UniClass", classController.GetAll)
	maingrp.GET("/UniClass/:id", classController.GetById)
	maingrp.GET("/UniClass/:code", classController.GetByCode)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
