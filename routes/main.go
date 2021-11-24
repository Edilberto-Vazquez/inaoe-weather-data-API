package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func getRoutes() {
	v1 := router.Group("/inaoe/v1")
	GroupByRouter(v1)
}

func Run() {
	port := os.Getenv("PORT")
	getRoutes()
	router.Run(":" + port)
}
