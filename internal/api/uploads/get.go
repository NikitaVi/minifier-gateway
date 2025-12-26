package uploads

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (c *Controller) GetUploads(ctx *gin.Context) {
	fmt.Println("get uploads")
}
