package services

import (
	"fmt"
	"go-crud/src/configuration/rest_err"
	"go-crud/src/model"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetEmail())
	return nil
}
