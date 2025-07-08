package model

import (
	"crypto/md5"
	"encoding/hex"
	"go-crud/src/configuration/rest_err"
)

func NewUserDomain(name, email, password string, age int8) *UserDomain {
	return &UserDomain{
		Name:     name,
		Email:    email,
		Age:      age,
		Password: password,
	}
}

type UserDomain struct {
	Name     string
	Email    string
	Age      int8
	Password string
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	GetUserByID(id string) (*UserDomain, *rest_err.RestErr)
	GetUserByEmail(email string) (*UserDomain, *rest_err.RestErr)
	UpdateUser(id string) *rest_err.RestErr
	DeleteUser(id string) *rest_err.RestErr
}
