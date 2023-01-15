package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/dannielss/go-crud/src/configuration/rest_err"
)

func NewUserDomain(
	email, 
	password, 
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email,
		password,
		name,
		age,
	}
}

type userDomain struct {
	email string
	password string
	name string
	age int8
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*userDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}