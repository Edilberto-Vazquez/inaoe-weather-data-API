package routes

import (
	"net/http"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/schemas"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/services"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/utils"
	"github.com/gin-gonic/gin"
)

func SelectRecordsRouter(rg *gin.RouterGroup) {
	efr := rg.Group("/select-records")
	efr.GET("/by-date", func(c *gin.Context) {
		var byDateSchema schemas.ByDateSchema
		var data []services.Data
		var rows int64
		var err error
		var byDateService services.ByDateService
		if err = c.ShouldBindQuery(&byDateSchema); err != nil {
			utils.ErrorMessage(err, c)
			return
		}
		data, rows, err = byDateService.FindByDate(byDateSchema)
		if err != nil {
			utils.ErrorMessage(err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
			"rows": rows,
		})
	})
}
