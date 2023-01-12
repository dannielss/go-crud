package userController

import (
	"net/http"

	"github.com/dannielss/go-crud/src/configuration/logger"
	"github.com/dannielss/go-crud/src/configuration/validation"
	"github.com/dannielss/go-crud/src/controllers/model/request"
	"github.com/dannielss/go-crud/src/controllers/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser func", zapcore.Field{
		Key: "journey",
		String: "createUser",
	})

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("error trying to validate user info", err, zapcore.Field{
			Key: "journey",
			String: "createUser",
		})

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	response := response.UserResponse{
		ID: "1",
		Email: userRequest.Email,
		Name: userRequest.Name,
		Age: userRequest.Age,
	}

	logger.Info("User created successfully")

	c.JSON(http.StatusCreated, response)
}