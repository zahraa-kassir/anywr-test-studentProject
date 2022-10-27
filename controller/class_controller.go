package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"test-pr/anywr-test-studentProject/dto"
	"test-pr/anywr-test-studentProject/payload"
	"test-pr/anywr-test-studentProject/repository"
)

var (
	ErrUserNotFound = errors.New("user not found, no data")
	ErrDataOption   = errors.New("data option not found")
)

type ClassController struct {
	ClassRepository   repository.ClassRepository
	StudentRepository repository.StudentRepository
	TeacherRepository repository.TeacherRepository
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

func (r ClassController) GetStByClCode(c echo.Context) error {
	//empty StudentForClass data
	stClassData := &payload.StudentForClass{}

	//bind data from client request ro StudentForClass
	if err := c.Bind(stClassData); err != nil {
		return err
	}

	/*class := r.ClassRepository.GetByCode(stClassData.MCode)
	cl := dto.Class{
		Code:   class.Code,
		Name:   class.Name,
		Credit: class.CreditNb,
	}

	students := r.StudentRepository.GetByCode(class.Id)
	var std []dto.SimpleStudentData
	for i, v := range students {
		student := dto.SimpleStudentData{
			Id:    i,
			Name:  v.Name,
			Email: v.Email,
		}
		std = append(std, student)
	}
	teachers := r.TeacherRepository.GetByCode(class.Id)
	var teach []dto.SimpleTeacherData
	for i, v := range teachers {
		th := dto.SimpleTeacherData{
			Id:    i,
			Name:  v.Name,
			Email: v.Email,
		}
		teach = append(teach, th)
	}

	uniClassData := &dto.UniClassData{
		Class:    cl,
		Students: std,
		Teachers: teach,
	}*/

	class := r.ClassRepository.GetStByClCode(stClassData.MCode)
	var std []dto.SimpleStudentData
	var teach []dto.SimpleTeacherData
	if class.Id != 0 {
		for i, v := range class.Student {
			student := dto.SimpleStudentData{
				Id:    i,
				Name:  v.Name,
				Email: v.Email,
			}
			std = append(std, student)
		}
		for i, v := range class.Teacher {
			th := dto.SimpleTeacherData{
				Id:    i,
				Name:  v.Name,
				Email: v.Email,
			}
			teach = append(teach, th)
		}
	}

	uniClassData := &dto.UniClassData{
		Class: dto.Class{
			Code:   class.Code,
			Name:   class.Name,
			Credit: class.CreditNb,
		},
		Students: std,
		Teachers: teach,
	}

	return c.JSON(http.StatusOK, uniClassData)
}
