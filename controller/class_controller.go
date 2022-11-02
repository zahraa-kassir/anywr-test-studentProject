package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"test-pr/anywr-test-studentProject/dto"
	"test-pr/anywr-test-studentProject/payload"
	"test-pr/anywr-test-studentProject/repository"
)

type ClassController struct {
	ClassRepository   repository.ClassRepository
	StudentRepository repository.StudentRepository
	TeacherRepository repository.TeacherRepository
}

func ClassCont(classRep repository.ClassRepository, studRep repository.StudentRepository, teachRep repository.TeacherRepository) *ClassController {
	return &ClassController{
		ClassRepository:   classRep,
		StudentRepository: studRep,
		TeacherRepository: teachRep,
	}
}

// GetAll return all classes
func (r ClassController) GetAll(c echo.Context) error {
	//get data
	classData := r.ClassRepository.GetAll()
	//at this level you are mapping from entity to dto not re-constructing
	//re-construct this with dto.class
	//ToClassDTO()

	//try to give accurate names for all variables
	//var classDTOs []dto.Class
	//for _, class := range classData {
	//	classDTO := dto.ReconstructClass(&class)
	//	classDTOs = append(classDTOs, classDTO)
	//}
	var all []dto.Class
	for _, v := range classData {
		//also for method names lets use something more relatable
		//like something related to mapping
		obj := dto.ReconstructClass(&v)
		all = append(all, obj)
	}
	return c.JSON(http.StatusOK, all)
}

// GetById get class based on class.id
func (r ClassController) GetById(c echo.Context) error {
	//handle param and convert
	//if this returns an error on failure we should always handle it not ignore it
	id, _ := strconv.Atoi(c.Param("id"))
	//get data
	data := r.ClassRepository.GetById(id)
	//re-construct based on dto.class
	class := dto.ReconstructClass(&data)
	return c.JSON(http.StatusOK, class)
}

// GetByCode get class based on class.code
func (r ClassController) GetByCode(c echo.Context) error {
	//handle class code
	code := c.Param("code")
	//get data
	data := r.ClassRepository.GetByCode(code)
	//re-construct based on dto.class
	class := dto.ReconstructClass(data)
	return c.JSON(http.StatusOK, class)
}

// GetStByClCode get class students based on class.code
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
	uniClassData := dto.ReconstructUniClassData(class)

	return c.JSON(http.StatusOK, uniClassData)
}

/*Business model

Student Entity
Student personal details Entity

Student Model ()

-----------------------------

usecase ->
get -> query
saving -> command


controller -> usecase -> repo Interface -> repo Impl -> DB
interface (api, service) -> usecase (query, command) -> domain (repository, model) -> adapter (db(repository, entity), university(repository, ))

adapter -> entity (adapter returns model) -> mapping from entity to model
usecase (takes models w returns models)
interface(api) takes model w returns dto -> mapping from model to


--------------------------------------
Project Structure:
->anywr
 -> cmd (main)
 -> internal (adapter, domain, interface, usecase, route)
 -> config
 -> container (database, wire..)
 -> pkg
  -> utils (validate email, split string)

*/
