package userController

import (
	"net/http"

	"github.com/dannielss/go-crud/src/configuration/logger"
	"github.com/dannielss/go-crud/src/configuration/validation"
	"github.com/dannielss/go-crud/src/controllers/model/request"
	"github.com/dannielss/go-crud/src/model"
	"github.com/dannielss/go-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
			
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

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))

	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domain))
}