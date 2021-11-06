package routes

import (
	"net/http"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/middleware"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/schemas"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ElectricFieldRouter(rg *gin.RouterGroup) {
	efr := rg.Group("/electric-field")
	var efs schemas.ElectricFieldSchema
	efr.GET("/date-range", middleware.ValidatorHandler(&efs, binding.Query), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"initialdate": c.Query("initialdate"),
			"finaldate":   c.Query("finaldate"),
			// todo try QueryArray for multiple fields
		})
	})
}
