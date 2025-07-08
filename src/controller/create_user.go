package controller

import (
	"github.com/gin-gonic/gin"
	"go-crud/src/configuration/logger"
	"go-crud/src/configuration/validation"
	"go-crud/src/model"
	"go-crud/src/model/request"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
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

	if err := domain.CreateUser(); err != nil {
		logger.Error("Error creating user", err, zap.String("journey", "CreateUser"))
		c.JSON(err.Code, err)
		return
	}
}
