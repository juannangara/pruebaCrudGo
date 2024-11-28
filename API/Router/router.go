package router

import (
	"net/http"
	handlers "purebaCrud/API/Handler"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, studentHandler handlers.StudentHandler, authHandler handlers.AuthHandler) {
	//health check
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//Routes users
	student := router.Group("/student")
	student.GET("/getStudents", studentHandler.GetStudents)
	student.GET("/getStudent", studentHandler.GetStudentByID)
	student.PUT("/updateStudent", studentHandler.UpdateStudent)
	student.DELETE("/deleteStudent", studentHandler.DeleteStudent)
	student.POST("/createStudent", studentHandler.CreateStudent)

	//Routes login
	router.POST("/login", authHandler.Login)
}
