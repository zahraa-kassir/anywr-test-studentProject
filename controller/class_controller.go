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
	ErrData         = errors.New("data not found")
)

type ClassController struct {
	ClassRepository   repository.ClassRepository
	StudentRepository repository.StudentRepository
	TeacherRepository repository.TeacherRepository
}

// GetAll return all classes
func (r ClassController) GetAll(c echo.Context) error {
	//get data
	classData := r.ClassRepository.GetAll()
	//re-construct this with dto.class
	var all []dto.Class
	for _, v := range classData {
		obj := dto.ReconstructClass(&v)
		all = append(all, obj)
	}
	return c.JSON(http.StatusOK, all)
}

// GetById get class based on class.id
func (r ClassController) GetById(c echo.Context) error {
	//handle param and convert
	id, _ := strconv.Atoi(c.Param("id"))
	//get data
	data := r.ClassRepository.GetById(id)
	//re-construct based on dto.class
	class := dto.ReconstructClass(&data)
	return c.JSON(http.StatusOK, class)
}

func (r ClassController) GetByCode(c echo.Context) error {
	//handle class code
	code := c.Param("code")
	//get data
	data := r.ClassRepository.GetByCode(code)
	//re-construct based on dto.class
	class := dto.ReconstructClass(data)
	return c.JSON(http.StatusOK, class)
}

func (r ClassController) GetStByClCode(c echo.Context) error {
	//empty StudentForClass data
	stClassData := &payload.StudentForClass{}

	//bind data from client request  StudentForClass
	if err := c.Bind(stClassData); err != nil {
		return err
	}
	//get data
	class := r.ClassRepository.GetStByClCode(stClassData.MCode)
	//re-construct code with dto.uniClassData
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
		Class:    dto.ReconstructSimpleClass(class),
		Students: std,
		Teachers: teach,
	}

	return c.JSON(http.StatusOK, uniClassData)
}
