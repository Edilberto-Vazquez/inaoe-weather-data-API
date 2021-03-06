package routes

import (
	"os"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/middleware"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/validators"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func getRoutes() {
	v1 := router.Group("/api/v1")
	SelectRecordsRouter(v1)
}

func Run() {
	port := os.Getenv("PORT")
	router.Use(middleware.Cors())
	validators.RegisterValidations()
	getRoutes()
	router.Run(":" + port)
}
