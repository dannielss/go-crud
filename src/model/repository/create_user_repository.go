package repository

import (
	"context"
	"os"

	"github.com/dannielss/go-crud/src/configuration/logger"
	"github.com/dannielss/go-crud/src/configuration/rest_err"
	"github.com/dannielss/go-crud/src/model"
	"github.com/dannielss/go-crud/src/model/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	 value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	return converter.ConvertEntityToDomain(*value), nil
}