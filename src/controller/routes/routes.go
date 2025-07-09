package routes

import (
	"github.com/gin-gonic/gin"
	"go-crud/src/controller"
)

func InitRoutes(r *gin.RouterGroup, uc controller.UserControllerInterface) {

	r.GET("/users/email/:email", uc.FindUserByEmail)
	r.GET("/users/:id", uc.FindUserByID)
	r.POST("/users", uc.CreateUser)
	r.PUT("/users/:id", uc.UpdateUser)
	r.DELETE("/users/:id", uc.DeleteUser)
}
