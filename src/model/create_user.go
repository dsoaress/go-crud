package model

import "go-crud/src/configuration/rest_err"

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	ud.EncryptPassword()
	return nil
}
