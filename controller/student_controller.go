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

func StudCont(studRep repository.StudentRepository, classRep repository.ClassRepository, teachRep repository.TeacherRepository) *StudentController {
	return &StudentController{
		StudentRepository: studRep,
		ClassRepository:   classRep,
		TeacherRepository: teachRep,
	}
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
	sData := dto.StudentByClass{
		Code: data.Code,
		Name: data.Name,
	}
	var stdArray []dto.SimpleData
	for i, v := range stud {
		obj := dto.ReconstructSimpleData(i, v.Name, v.Email)
		stdArray = append(stdArray, obj)
	}
	sData.Students = stdArray

	return c.JSON(http.StatusOK, sData)
}

func (s StudentController) GetByThAndClass(c echo.Context) error {
	//empty Filter data
	FilterData := &payload.FilterByTeacherAndClass{}
	//bind data from client request ro StudentForClass
	if err := c.Bind(FilterData); err != nil {
		return err
	}
	//get class data & get teacher data
	class := s.ClassRepository.GetByCode(FilterData.ClassCode)
	teach := s.TeacherRepository.GetByEmail(FilterData.TeachEmail)
	if class == nil || teach == nil {
		return c.JSON(http.StatusBadRequest, ErrNotFound)
	}
	//get student by class code
	var stud []entity.Student
	stud = s.StudentRepository.GetByFilter(class.Id, FilterData.Page, FilterData.PageSize, FilterData.TeachEmail)
	if stud == nil {
		return c.JSON(http.StatusBadRequest, ErrNotFound)
	}
	//re-construct final data
	FinalData := dto.ReconstructFilterData(*teach, *class, stud)

	return c.JSON(http.StatusOK, FinalData)
}
