package luffy

import "github.com/gin-gonic/gin"

const (
	luffy = "D Monkey"
)

func Luffy(c *gin.Context) {
	c.String(200, luffy)
}