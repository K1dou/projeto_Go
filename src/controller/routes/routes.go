package routes

import (
	"github.com/K1dou/projeto_Go.git/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:id",controller.FindUserById)
	r.GET("/getUserByEmail/:email",controller.FindUserByEmail)
	r.POST("/createUser",controller.CreateUser)
	r.PUT("/updateUser/:id",controller.FindUserById)
	r.DELETE("/deleteUser/:id",controller.DeleteUser)
}