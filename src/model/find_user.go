package model

import (
	"github.com/dannielss/go-crud/src/configuration/rest_err"
)

func (ud *userDomain) FindUser(id string) (*userDomain, *rest_err.RestErr) {
	return &userDomain{}, nil
}