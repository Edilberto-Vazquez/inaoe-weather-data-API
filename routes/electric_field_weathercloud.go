package routes

import (
	"net/http"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/middleware"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/schemas"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ElectricFieldWeathercloudRouter(rg *gin.RouterGroup) {
	efr := rg.Group("/electric-field-and-weathercloud")
	var efs schemas.ElectricFieldSchema
	var service services.ElectricFieldWeatherCloudService
	efr.GET("/", middleware.ValidatorHandler(&efs, binding.Query), func(c *gin.Context) {
		records, rows, _ := service.Find(c.Query("firstdate"), c.Query("lastdate"), c.QueryArray("fields"))
		c.JSON(http.StatusOK, gin.H{
			"rows":    rows,
			"records": records,
		})
	})
}
