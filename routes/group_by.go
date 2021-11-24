package routes

import (
	"net/http"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/schemas"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/services"
	"github.com/gin-gonic/gin"
)

func GroupByRouter(rg *gin.RouterGroup) {
	efr := rg.Group("/group-by-date")
	efr.GET("/", func(c *gin.Context) {
		var schema schemas.INAOEQuerySchema
		var results []services.Result
		var rows int64
		var err error
		service := services.NewINAOEQuerysService()
		if err = c.ShouldBindJSON(&schema); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		if schema.WeatherClouds == nil && schema.ElectricFields == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "no field has been selected",
			})
			c.Abort()
			return
		}
		if schema.JoinType != "" && schema.WeatherClouds != nil && schema.ElectricFields != nil {
			results, rows, err = service.JoinTypeFind(schema)
		} else {
			results, rows, err = service.Find(schema)
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"results": results,
			"rows":    rows,
		})
	})
}
