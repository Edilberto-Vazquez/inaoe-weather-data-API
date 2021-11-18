package routes

import (
	"net/http"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/schemas"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/services"
	"github.com/gin-gonic/gin"
)

func ElectricFieldWeatherCloudRouter(rg *gin.RouterGroup) {
	efr := rg.Group("/electric-field-and-weather-cloud")
	var service services.ElectricFieldWeatherCloudService
	efr.GET("/records", func(c *gin.Context) {
		var schema schemas.ElectricFieldAndWeatherCloudSchema
		if err := c.ShouldBindQuery(&schema); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		records, rows, err := service.Find(c.Query("firstdate"), c.Query("lastdate"), c.QueryArray("fields"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"rows":    rows,
			"records": records,
		})
	})
}
