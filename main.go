package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"test-pr/anywr-test-studentProject/container/db"
	"test-pr/anywr-test-studentProject/container/wires"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connect to the "bank" database
	//This one should be inside container/database.go
	db.Connect("postgres://postgres:admin@localhost:5432/anywrstudentproject")

	//Routes (can be moved to a route file to keep main clean)
	TeacherController := wires.InitialiseTeachCont(db.Instance)
	StudentController := wires.InitialiseStudCont(db.Instance)
	classController := wires.InitialiseClassCont(db.Instance)

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
