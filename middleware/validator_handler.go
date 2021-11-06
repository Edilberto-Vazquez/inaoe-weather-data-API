package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ValidatorHandler(schema interface{}, b binding.Binding) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindWith(schema, b); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
	}
}
