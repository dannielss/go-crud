package userController

import (
	"github.com/dannielss/go-crud/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	ShowUser(c *gin.Context)
	ShowUserByEmail(c *gin.Context)
	
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DestroyUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}