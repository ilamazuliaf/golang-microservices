package utils

import "github.com/gin-gonic/gin"

func Respond(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
	}
	c.JSON(status,body)
}

func RespondError(c *gin.Context, err *ApplicationError) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(err.StatusCode, err)
	}
	c.JSON(err.StatusCode, err)
}