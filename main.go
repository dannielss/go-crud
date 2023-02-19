package main

import (
	"context"
	"log"

	"github.com/dannielss/go-crud/src/configuration/database/mongodb"
	"github.com/dannielss/go-crud/src/configuration/logger"
	"github.com/dannielss/go-crud/src/controllers/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())

	if err != nil {
		log.Fatalf("error trying to connect to database, error=%s", err.Error())
	}

	//Init dependencies
	controller := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, controller)
	
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}