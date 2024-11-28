package handler

import (
	"errors"
	"log"
	"net/http"
	models "purebaCrud/API/Models"
	service "purebaCrud/API/Service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentHandler interface {
	GetStudents(c *gin.Context)
	CreateStudent(c *gin.Context)
	GetStudentByID(c *gin.Context)
	UpdateStudent(c *gin.Context)
	DeleteStudent(c *gin.Context)
}

type studentHandler struct {
	studentService service.StudentService
}

func CreateStudentHandler(studentService service.StudentService) StudentHandler {
	return &studentHandler{
		studentService: studentService,
	}
}

func (h studentHandler) GetStudents(c *gin.Context) {
	err := authToken(c)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	students, err := h.studentService.GetStudents()
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, errors.New("error al cargar"))
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h studentHandler) CreateStudent(c *gin.Context) {
	err := authToken(c)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	student := models.NewStudent{}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.studentService.CreateStudent(student)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("errors during student creation"))
		return
	}

	c.JSON(http.StatusCreated, "success")

}

func (h studentHandler) GetStudentByID(c *gin.Context) {
	err := authToken(c)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		log.Println("Error while parsing id:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := h.studentService.GetById(id)

	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)

}

func (h studentHandler) UpdateStudent(c *gin.Context) {

	err := authToken(c)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	student := models.Student{}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.studentService.UpdateStudent(student)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("error al crear"))
		return
	}

	c.JSON(http.StatusCreated, "student updated")
}

func (h studentHandler) DeleteStudent(c *gin.Context) {
	err := authToken(c)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	student := models.Student{}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.studentService.DeleteStudent(student)

	if err != nil {
		log.Println("Error al realizar la consulta:", err)
		c.JSON(http.StatusInternalServerError, errors.New("error al cargar"))
		return
	}

	c.JSON(http.StatusOK, "student deleted")
}
