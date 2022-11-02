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
	// you can directly here return student based on class code without passing by class repo first
	stud, _ := s.StudentRepository.GetByCode(data.Id, "", "")
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
	filterData := &payload.FilterByTeacherAndClass{}
	//bind data from client request ro StudentForClass
	if err := c.Bind(filterData); err != nil {
		return err
	}
	//try to make it one repository method which takes all filters
	switch {
	case filterData.RankBy != "":
		//by rank (min , max) 2 case for this:
		//{} one by class code
		//{} student for every class
		//you can check the existence of the filter on the repo level (applies for all filters)
		if filterData.ClassCode == "" {
			//get student by class code
			var stud []entity.Student
			var numb int64
			stud, numb = s.StudentRepository.GetByRank(0, *filterData)
			if stud == nil {
				return c.JSON(http.StatusBadRequest, ErrNotFound)
			}
			//re-construct final data
			FinalData := dto.ReconstructFilterData(stud, *filterData, numb, "rank")

			return c.JSON(http.StatusOK, FinalData)

		} else {
			//get class data
			class := s.ClassRepository.GetByCode(filterData.ClassCode)
			if class == nil {
				return c.JSON(http.StatusBadRequest, "no data")
			}
			//get student by class code
			var stud []entity.Student
			var numb int64
			stud, numb = s.StudentRepository.GetByRank(class.Id, *filterData)
			if stud == nil {
				return c.JSON(http.StatusBadRequest, ErrNotFound)
			}
			//re-construct final data
			FinalData := dto.ReconstructFilterData(stud, *filterData, numb, "rank")

			return c.JSON(http.StatusOK, FinalData)

		}

	case filterData.TeachEmail != "" && filterData.ClassCode != "":
		//get class data & get teacher data
		class := s.ClassRepository.GetByCode(filterData.ClassCode)
		teach := s.TeacherRepository.GetBy2(class.Id, filterData.TeachEmail)
		if class == nil || teach == nil {
			return c.JSON(http.StatusBadRequest, "no data with this option")
		}
		//get student by class code
		var stud []entity.Student
		var numb int64
		stud, numb = s.StudentRepository.GetByFilterTeacherAndClass(class.Id, *filterData, *teach)
		if stud == nil {
			return c.JSON(http.StatusBadRequest, ErrNotFound)
		}
		//re-construct final data
		FinalData := dto.ReconstructFilterData(stud, *filterData, numb, "teacher:email - class:code")

		return c.JSON(http.StatusOK, FinalData)

	case filterData.ClassCode != "":
		//get class data
		class := s.ClassRepository.GetByCode(filterData.ClassCode)
		if class == nil {
			return c.JSON(http.StatusBadRequest, "no data with this class")
		}
		//get student by class code
		var stud []entity.Student
		var numb int64
		stud, numb = s.StudentRepository.GetByCode(class.Id, filterData.SortBy.StudentName, filterData.SortBy.TeacherName)
		if stud == nil {
			return c.JSON(http.StatusBadRequest, ErrNotFound)
		}
		//re-construct final data
		FinalData := dto.ReconstructFilterData(stud, *filterData, numb, "class:code")

		return c.JSON(http.StatusOK, FinalData)

	}
	//you can send nil if no response
	return c.JSON(http.StatusOK, "no data with this option")
}
