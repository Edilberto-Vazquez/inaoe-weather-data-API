package routes

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func getRoutes() {
	v1 := router.Group("/api/v1")
	ElectricFieldRouter(v1)
}

func Run() {
	getRoutes()
	router.Run(":3000")
}
