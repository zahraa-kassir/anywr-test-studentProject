package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"test-pr/anywr-test-studentProject/repository"
)

var (
	ErrUserNotFound = errors.New("user not found, no data")
)

type ClassController struct {
	ClassRepository   repository.ClassRepository
	StudentRepository repository.StudentRepository
}

func (r ClassController) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, r.ClassRepository.GetAll())
}

func (r ClassController) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, r.ClassRepository.GetById(id))
}

func (r ClassController) GetByCode(c echo.Context) error {
	return c.JSON(http.StatusOK, r.ClassRepository.GetByCode(c.Param("code")))
}
func (r ClassController) GetByStudentMail(c echo.Context) error {

	ownerOfEmail := r.StudentRepository.GetByEmail(c.Param("email"))
	if ownerOfEmail == nil {

		return c.JSON(http.StatusBadRequest, ErrUserNotFound)

	}
	return c.JSON(http.StatusOK, r.ClassRepository.GetDataOfThisStudent(ownerOfEmail.Id))

}
