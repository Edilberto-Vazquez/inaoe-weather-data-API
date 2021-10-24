package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WeatherRouter(rg *gin.RouterGroup) {
	weather := rg.Group("/weather")

	weather.GET("/:location", func(c *gin.Context) {
		location := c.Param("location")
		c.String(http.StatusOK, "weather data %s", location)
	})
}
