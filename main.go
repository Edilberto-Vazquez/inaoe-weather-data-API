package main

import (
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/libs"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/routes"
)

func main() {
	libs.InitCon()
	defer libs.CloseCon()
	routes.Run()
}
