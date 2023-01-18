package userController

import (
	"net/http"

	"github.com/dannielss/go-crud/src/configuration/logger"
	"github.com/dannielss/go-crud/src/configuration/validation"
	"github.com/dannielss/go-crud/src/controllers/model/request"
	"github.com/dannielss/go-crud/src/model"
	"github.com/dannielss/go-crud/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

var (
	UserDomainInterface model.UserDomainInterface
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

	domain := model.NewUserDomain(
		userRequest.Email, 
		userRequest.Password, 
		userRequest.Name,
		 userRequest.Age,
	)

	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully")

	c.String(http.StatusCreated, "")
}