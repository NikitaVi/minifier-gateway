package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (c *Controller) Login(ctx *gin.Context) {
	fmt.Println("Login")
}
