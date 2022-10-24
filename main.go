package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"test-pr/anywr-test-studentProject/db"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connect to the "bank" database
	db.Connect("postgres://postgres:admin@localhost:5432/anywr-student-project")

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
