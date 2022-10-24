package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"test-pr/anywr-test-studentProject/repository"
)

type ClassController struct {
	ClassRepository repository.ClassRepository
}

func (r ClassController) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, r.ClassRepository.GetAll())
}

func (r ClassController) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, r.ClassRepository.GetById(id))
}

func (r ClassController) GetByCode(c echo.Context) error {
	code, _ := strconv.Atoi(c.Param("code"))
	return c.JSON(http.StatusOK, r.ClassRepository.GetById(code))
}
