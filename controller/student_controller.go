package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"test-pr/anywr-test-studentProject/repository"
)

type StudentController struct {
	StudentRepository repository.StudentRepository
}

func (s StudentController) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, s.StudentRepository.GetAll())
}

func (s StudentController) GetByEmail(c echo.Context) error {
	email := c.Param("email")
	return c.JSON(http.StatusOK, s.StudentRepository.GetByEmail(email))
}
