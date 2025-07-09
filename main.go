package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-crud/src/controller"
	"go-crud/src/controller/routes"
	"go-crud/src/model/services"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	service := services.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
