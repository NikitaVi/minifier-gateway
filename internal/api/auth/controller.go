package auth

import "github.com/gin-gonic/gin"

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) RegisterRoutes(rg *gin.RouterGroup) {
	userGr := rg.Group("/auth")

	userGr.POST("/login", c.Login)
	userGr.POST("/register", c.Register)
}
