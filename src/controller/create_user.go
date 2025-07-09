package controller

import (
	"github.com/gin-gonic/gin"
	"go-crud/src/configuration/logger"
	"go-crud/src/configuration/validation"
	"go-crud/src/controller/model/request"
	"go-crud/src/model"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error binding user request", err, zap.String("journey", "CreateUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Name,
		userRequest.Email,
		userRequest.Password,
		userRequest.Age)

	if err := uc.service.CreateUser(domain); err != nil {
		logger.Error("Error creating user", err, zap.String("journey", "CreateUser"))
		c.JSON(err.Code, err)
		return
	}
}
