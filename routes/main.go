package routes

import (
	"os"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/middleware"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func getRoutes() {
	v1 := router.Group("/api/v1")
	GroupByRouter(v1)
}

func Run() {
	port := os.Getenv("PORT")
	router.Use(middleware.Cors())
	getRoutes()
	router.Run(":" + port)
}
