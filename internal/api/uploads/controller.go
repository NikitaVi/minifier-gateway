package uploads

import "github.com/gin-gonic/gin"

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) RegisterRoutes(rg *gin.RouterGroup) {
	uploadsGr := rg.Group("/uploads")

	uploadsGr.GET("/", c.GetUploads)
	uploadsGr.POST("/", c.GetUploads)
}
