package controller

import (
	"github.com/gin-gonic/gin"
	"go-crud/src/model/services"
)

func NewUserControllerInterface(
	serviceInterface services.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service services.UserDomainService
}
