package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMessage(err error, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": err.Error(),
	})
	c.Abort()
}
