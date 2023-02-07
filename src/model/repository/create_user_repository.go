package repository

import (
	"os"

	"github.com/dannielss/go-crud/src/configuration/logger"
	"github.com/dannielss/go-crud/src/configuration/rest_err"
	"github.com/dannielss/go-crud/src/model"
)

const (
	USER_COLLECTION = "USER_COLLECTION"
)
func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository")

	collection_name := os.Getenv(USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)
}