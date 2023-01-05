package routes

import (
	"github.com/dannielss/go-crud/src/controllers/userController"
	"github.com/gin-gonic/gin"
)
func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:userId", userController.ShowUser)
	r.GET("/userByEmail/:userEmail", userController.ShowUserByEmail)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user/:userId", userController.UpdateUser)
	r.DELETE("/user/:userId", userController.DestroyUser)
}