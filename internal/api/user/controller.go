package user

import "github.com/gin-gonic/gin"

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) RegisterRoutes(rg *gin.RouterGroup) {
	userGr := rg.Group("/user")

	userGr.POST("/premium", c.IsPremium)
}
