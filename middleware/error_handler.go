package middleware

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.MustGet("error")
		httpStatus := c.MustGet("httpStatus")
		c.JSON(httpStatus.(int), gin.H{
			"message": err.(error).Error(),
		})
	}
}
