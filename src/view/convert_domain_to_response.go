package view

import (
	"github.com/dannielss/go-crud/src/controllers/model/response"
	"github.com/dannielss/go-crud/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID: userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name: userDomain.GetName(),
		Age: userDomain.GetAge(),
	}
}