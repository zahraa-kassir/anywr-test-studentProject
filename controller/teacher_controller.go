package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"test-pr/anywr-test-studentProject/repository"
)

type TeacherContoller struct {
	TeacherRepository repository.TeacherRepository
}

func (t TeacherContoller) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, t.TeacherRepository.GetAll())
}

func (t TeacherContoller) GetByEmail(c echo.Context) error {
	email := c.Param("email")
	return c.JSON(http.StatusOK, t.TeacherRepository.GetByEmail(email))
}
