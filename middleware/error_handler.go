package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		log.Println("log errors:")
	}
}

// func OrmErrorHandler(err error) gin.HandlerFunc  {
// 	return  func(c *gin.Context) {
// 		if err != nil {

// 		}
// 	}
// }

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.MustGet("error")
		if err == nil {
			c.Abort()
			return
		}
		log.Println("error handler")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.(error).Error(),
		})
	}
}
