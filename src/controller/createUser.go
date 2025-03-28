package controller

import (
	resterr "github.com/K1dou/projeto_Go.git/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	err := resterr.NewBadRequestErr("Voce não pode criar um usuario com esse email")
	c.JSON(err.Code,err)

}