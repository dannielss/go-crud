package converter

import (
	"github.com/dannielss/go-crud/src/model"
	"github.com/dannielss/go-crud/src/model/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Name,
		entity.Password,
		entity.Age,
	)

	domain.SetID(entity.ID.Hex())
	return domain
}

