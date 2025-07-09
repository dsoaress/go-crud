package services

import (
	"go-crud/src/configuration/rest_err"
	"go-crud/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct{}

type UserDomainService interface {
	CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr
	GetUserByID(id string) (*model.UserDomainInterface, *rest_err.RestErr)
	GetUserByEmail(email string) (*model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(id string, userDomain model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(id string) *rest_err.RestErr
}
