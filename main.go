package main

import (
	"fmt"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/db/migrations"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/etl"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/libs"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/services"
)

func main() {
	libs.InitCon()
	defer libs.CloseCon()
	migrations.Migrate()
	logRoot := "/mnt/d/DataSets/Conjuntos-originales/medidor-campo-electrico/"
	efmRoot := "/mnt/d/DataSets/Conjuntos-originales/medidor-campo-electrico/"
	wcRoot := "/mnt/d/DataSets/Conjuntos-originales/estacion-meteorologica/"
	etl.ProcessFiles(logRoot, efmRoot, wcRoot)
	service := services.NewElectricFieldWeatherCloudService()
	rows, _ := service.CreateRecords(etl.RecordSliceTable)
	fmt.Println(rows)
	// routes.Run()
}
