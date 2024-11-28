package main

import (
	"fmt"
	"log"
	handler "purebaCrud/API/Handler"
	repository "purebaCrud/API/Repository"
	router "purebaCrud/API/Router"
	service "purebaCrud/API/Service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Print("Iniciando")
	buildApp()

}

func buildApp() {
	//cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error al cargar variables de entorno desde .env")
	}

	//create router
	routerGin := gin.Default()

	//create sql conection
	db := service.CreateConection()

	//create repos
	studentRepo := repository.CreateSqlGestor(db)

	//create service
	studentService := service.CreateStudentService(studentRepo)

	//create handlers
	studentHandler := handler.CreateStudentHandler(studentService)
	authHandler := handler.CreateAuthHandler()
	//giving handler to router and creating routes
	router.Routes(routerGin, studentHandler, authHandler)

	//start server
	routerGin.Run()
}
