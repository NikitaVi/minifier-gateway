package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (c *Controller) IsPremium(ctx *gin.Context) {
	fmt.Println("isPremium")
}
