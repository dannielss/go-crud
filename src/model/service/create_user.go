package service

import (
	"github.com/dannielss/go-crud/src/configuration/logger"
	"github.com/dannielss/go-crud/src/configuration/rest_err"
	"github.com/dannielss/go-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)

	if err != nil {
		logger.Error("Error trying to call CreateUser repository", err, zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("CreateUser service executed successfully", zap.String("journey", "createUser"))

	return userDomainRepository, nil
}