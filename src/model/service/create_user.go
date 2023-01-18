package service

import (
	"github.com/dannielss/go-crud/src/configuration/logger"
	"github.com/dannielss/go-crud/src/configuration/rest_err"
	"github.com/dannielss/go-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	return nil
}