package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"test-pr/anywr-test-studentProject/dto"
	"test-pr/anywr-test-studentProject/entity"
	"test-pr/anywr-test-studentProject/payload"
	"test-pr/anywr-test-studentProject/repository"
)

type StudentController struct {
	StudentRepository repository.StudentRepository
	ClassRepository   repository.ClassRepository
	TeacherRepository repository.TeacherRepository
}

var (
	ErrNotFound = errors.New("data not found")
)

// GetAll  all student data
func (s StudentController) GetAll(c echo.Context) error {
	//create a handle for student data
	data := s.StudentRepository.GetAll()
	if data == nil {
		return c.JSON(http.StatusBadRequest, ErrNotFound)
	}
	//create response data form
	var students []dto.Student
	for _, v := range data {
		std := dto.ReconstructStudent(v)
		students = append(students, std)
	}
	return c.JSON(http.StatusOK, students)
}

// GetByEmail data of student based on email@
func (s StudentController) GetByEmail(c echo.Context) error {
	//handel param email
	email := c.Param("email")
	//get data
	data := s.StudentRepository.GetByEmail(email)
	if data == nil {
		return c.JSON(http.StatusBadRequest, ErrNotFound)
	}
	//create response data based on dto.student
	std := dto.ReconstructStudent(*data)
	return c.JSON(http.StatusOK, std)
}

func (s StudentController) GetByCode(c echo.Context) error {
	//handle code
	code := c.Param("code")
	//verifie if data of classes.code exist
	data := s.ClassRepository.GetByCode(code)
	if data == nil {
		return c.JSON(http.StatusBadRequest, ErrNotFound)
	}
	//get data of student based on class.id
	stud := s.StudentRepository.GetByCode(data.Id)
	if stud == nil {
		return c.JSON(http.StatusBadRequest, ErrNotFound)
	}
	//re-construct data based on dto.studentByClass
	sdata := dto.StudentByClass{
		Code: data.Code,
		Name: data.Name,
	}
	var stdArray []dto.SimpleStudentData
	for i, v := range stud {
		obj := dto.SimpleStudentData{
			Id:    i,
			Name:  v.Name,
			Email: v.Email,
		}
		stdArray = append(stdArray, obj)
	}
	sdata.Students = stdArray

	return c.JSON(http.StatusOK, sdata)
}

func (s StudentController) GetByThAndClass(c echo.Context) error {
	//empty Filter data
	FilterData := &payload.FilterByTeacherAndClass{}

	//bind data from client request ro StudentForClass
	if err := c.Bind(FilterData); err != nil {
		return err
	}
	//get class data
	class := s.ClassRepository.GetByCode(FilterData.ClassCode)
	if class == nil {
		return c.JSON(http.StatusBadRequest, ErrNotFound)
	}

	teach := s.TeacherRepository.GetByEmail(FilterData.TeachEmail)
	if teach == nil {
		return c.JSON(http.StatusBadRequest, ErrNotFound)
	}

	exist := false
	for _, v := range teach.Classes {
		if class.Id == v.Id {
			exist = true
		}
	}
	FinalData := dto.FilterData{
		Teacher: dto.SimpleTeacherData{
			Id:    teach.Id,
			Name:  teach.Name,
			Email: teach.Email,
		},
	}

	//get student by class code
	var stud []entity.Student
	if exist {
		q := s.StudentRepository.GetByFilter(class.Id)

		if q == nil {
			return c.JSON(http.StatusBadRequest, ErrNotFound)
		}
		stud = q

		FinalData.Class = dto.ReconstructSimpleClass(*class)
	} else {
		stud = nil
	}

	if stud != nil {
		var students []dto.SimpleStudentData
		for _, v := range stud {
			n := dto.SimpleStudentData{
				Id:    v.Id,
				Name:  v.Name,
				Email: v.Email,
			}
			students = append(students, n)
		}

		FinalData.Student = students
	} else {
		FinalData.Student = nil
	}

	return c.JSON(http.StatusOK, FinalData)
}
