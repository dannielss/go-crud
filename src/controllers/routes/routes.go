package routes

import (
	"github.com/dannielss/go-crud/src/controllers/userController"
	"github.com/gin-gonic/gin"
)
func InitRoutes(r *gin.RouterGroup, controller userController.UserControllerInterface) {
	r.GET("/user/:userId", controller.ShowUser)
	r.GET("/userByEmail/:userEmail", controller.ShowUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DestroyUser)
}