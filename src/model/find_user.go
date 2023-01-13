package model

import (
	"github.com/dannielss/go-crud/src/configuration/rest_err"
)

func (ud *UserDomain) FindUser(id string) (*UserDomain, *rest_err.RestErr) {
	return &UserDomain{}, nil
}