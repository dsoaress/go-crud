package routes

import (
	"github.com/gin-gonic/gin"
	"go-crud/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users/email/:email", controller.FindUserByEmail)
	r.GET("/users/:id", controller.FindUserByID)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)
}
