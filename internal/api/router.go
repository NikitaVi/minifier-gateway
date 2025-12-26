package api

import "github.com/gin-gonic/gin"

type Router interface {
	RegisterRoutes(rg *gin.RouterGroup)
}
