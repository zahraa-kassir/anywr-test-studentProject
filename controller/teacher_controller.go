package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"test-pr/anywr-test-studentProject/dto"
	"test-pr/anywr-test-studentProject/repository"
)

type TeacherContoller struct {
	TeacherRepository repository.TeacherRepository
}

var (
	ErrNotFounds = errors.New("data not found")
)

func (t TeacherContoller) GetAll(c echo.Context) error {
	//create a handle for teachers data
	tData := t.TeacherRepository.GetAll()
	if tData == nil {
		return c.JSON(http.StatusBadRequest, ErrNotFounds)
	}
	//create response data form
	var teachers []dto.GetAllTeachers
	for i, v := range tData {
		tea := dto.GetAllTeachers{
			Teacher: dto.SimpleTeacherData{
				Id:    i,
				Name:  v.Name,
				Email: v.Email,
			},
		}

		var cls []dto.Class
		for _, b := range v.Classes {
			cl := dto.Class{
				Code:   b.Code,
				Name:   b.Name,
				Credit: b.CreditNb,
			}

			cls = append(cls, cl)

		}

		tea.Class = cls

		teachers = append(teachers, tea)

	}

	return c.JSON(http.StatusOK, teachers)
}

func (t TeacherContoller) GetByEmail(c echo.Context) error {
	email := c.Param("email")
	data := t.TeacherRepository.GetByEmail(email)
	if data == nil {
		return c.JSON(http.StatusBadRequest, ErrNotFound)
	}
	teach := dto.Teacher{
		Id:    data.Id,
		Name:  data.Name,
		Email: data.Email,
	}
	var cl []dto.Class
	for _, v := range data.Classes {
		cls := dto.Class{
			Code:   v.Code,
			Name:   v.Name,
			Credit: v.CreditNb,
		}

		cl = append(cl, cls)
	}
	teach.Class = cl

	return c.JSON(http.StatusOK, teach)
}
