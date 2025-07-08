package model

import (
	"crypto/md5"
	"encoding/hex"
	"go-crud/src/configuration/rest_err"
)

func NewUserDomain(name, email, password string, age int8) UserDomainInterface {
	return &userDomain{
		name,
		email,
		age,
		password,
	}
}

type userDomain struct {
	name     string
	email    string
	age      int8
	password string
}

func (ud *userDomain) encryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	GetUserByID(id string) (*userDomain, *rest_err.RestErr)
	GetUserByEmail(email string) (*userDomain, *rest_err.RestErr)
	UpdateUser(id string) *rest_err.RestErr
	DeleteUser(id string) *rest_err.RestErr
}
