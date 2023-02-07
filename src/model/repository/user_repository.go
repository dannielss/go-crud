package repository

import (
	"github.com/dannielss/go-crud/src/configuration/rest_err"
	"github.com/dannielss/go-crud/src/model"
)

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}