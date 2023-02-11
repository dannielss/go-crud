package converter

import (
	"github.com/dannielss/go-crud/src/model"
	"github.com/dannielss/go-crud/src/model/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email: domain.GetEmail(),
		Password: domain.GetPassword(),
		Name: domain.GetName(),
		Age: domain.GetAge(),
	}
}

