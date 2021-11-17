package routes

import (
	"net/http"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/middleware"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/schemas"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ElectricFieldWeatherCloudRouter(rg *gin.RouterGroup) {
	efr := rg.Group("/electric-field-and-weather-cloud")
	var efs schemas.ElectricFieldSchema
	var service services.ElectricFieldWeatherCloudService
	efr.GET("/", middleware.ValidatorHandler(&efs, binding.Query), func(c *gin.Context) {
		records, rows, err := service.Find(c.Query("firstdate"), c.Query("lastdate"), c.QueryArray("fields"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"rows":    rows,
				"records": records,
			})
		}
	})
}
