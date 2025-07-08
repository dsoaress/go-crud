package model

import "go-crud/src/configuration/rest_err"

func (ud *userDomain) CreateUser() *rest_err.RestErr {
	ud.encryptPassword()
	return nil
}
