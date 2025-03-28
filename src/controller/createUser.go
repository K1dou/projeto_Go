package controller

import (
	resterr "github.com/K1dou/projeto_Go.git/src/configuration/rest_err"
	"github.com/K1dou/projeto_Go.git/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := resterr.NewBadRequestErr("Invalid JSON body" + err.Error())
		c.JSON(restErr.Code, restErr)
		return 
	}

}