package model

import (
	"github.com/dannielss/go-crud/src/configuration/logger"
	"github.com/dannielss/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	return nil
}