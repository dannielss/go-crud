package userController

import (
	"fmt"

	"github.com/dannielss/go-crud/src/configuration/validation"
	"github.com/dannielss/go-crud/src/controllers/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}