package main

import (
	"github.com/dannielss/go-crud/src/controllers/userController"
	"github.com/dannielss/go-crud/src/model/repository"
	"github.com/dannielss/go-crud/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) userController.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	controller := userController.NewUserControllerInterface(service)

	return controller
}