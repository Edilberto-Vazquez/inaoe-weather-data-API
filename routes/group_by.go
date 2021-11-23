package routes

import (
	"net/http"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/schemas"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/services"
	"github.com/gin-gonic/gin"
)

func GroupByRouter(rg *gin.RouterGroup) {
	efr := rg.Group("/group-by")
	efr.GET("/", func(c *gin.Context) {
		schema := schemas.INAOEQuerySchema{}
		service := services.NewINAOEQuerysService()
		var results []services.Result
		var rows int64
		var err error
		if err = c.ShouldBindQuery(&schema); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		queryOptions := map[string]interface{}{
			"datepart":       c.Query("datepart"),
			"fromdate":       c.Query("fromdate"),
			"todate":         c.Query("todate"),
			"weatherclouds":  c.QueryArray("weatherclouds"),
			"electricfields": c.QueryArray("electricfields"),
			"jointype":       c.Query("jointype"),
		}
		if c.QueryArray("weatherclouds") == nil && c.QueryArray("electricfields") == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "no field has been selected",
			})
			c.Abort()
			return
		}
		if c.Query("jointype") != "" && c.QueryArray("weatherclouds") != nil && c.QueryArray("electricfields") != nil {
			results, rows, err = service.JoinTypeFind(queryOptions)
		} else {
			results, rows, err = service.Find(queryOptions)
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
