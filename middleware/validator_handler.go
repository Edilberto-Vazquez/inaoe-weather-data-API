package middleware

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ValidatorHandler(schema interface{}, b binding.Binding) gin.HandlerFunc {
	return func(c *gin.Context) {
		rType := reflect.TypeOf(schema)
		newSchema := reflect.New(rType)
		sc := newSchema.Elem().MethodByName("NewSchema").Call([]reflect.Value{})
		a := sc[0]
		log.Println(a)
		if err := c.ShouldBindWith(&a, b); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
	}
}
