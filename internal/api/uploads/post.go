package uploads

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (c *Controller) PostUploads(ctx *gin.Context) {
	fmt.Println("post uploads")
}
