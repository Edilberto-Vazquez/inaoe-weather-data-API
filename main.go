package main

import (
	"fmt"
	"time"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/etl"
)

func main() {
	// libs.InitCon()
	// defer libs.CloseCon()
	// migrations.Migrate()
	// wc := services.NewWeathercloudService()
	// wc.CreateMultipleRecords("/mnt/f/DataSets/Conjuntos-originales/estacion-meteorologica/Weathercloud Pque Tecnolgico INAOE 2019-02.csv")
	// efs := services.NewElectricFieldService()
	// efs.CreateMultipleRecords("/mnt/f/DataSets/Conjuntos-originales/medidor-campo-electrico/INAOE parque-01102019.efm")
	// log := services.NewLogService()
	// log.CreateMultipleRecords("/mnt/f/DataSets/Conjuntos-originales/medidor-campo-electrico/EFMEvents.log")
	// efwcs := services.NewElectricFieldWeatherCloudService()
	// efwcs.ProcessNewFiles("/mnt/d/DataSets/Conjuntos-originales/medidor-campo-electrico/EFMEvents.log", "/mnt/d/DataSets/Conjuntos-originales/medidor-campo-electrico/INAOE parque-03032019.efm", "/mnt/d/DataSets/Conjuntos-originales/estacion-meteorologica/Weathercloud Pque Tecnolgico INAOE 2019-03.csv")
	// routes.Run()
	start := time.Now()
	logRoot := "/mnt/d/DataSets/Conjuntos-originales/medidor-campo-electrico/"
	efmRoot := "/mnt/d/DataSets/Conjuntos-originales/medidor-campo-electrico/"
	wcRoot := "/mnt/d/DataSets/Conjuntos-originales/estacion-meteorologica/"
	etl.ProcessFiles(logRoot, efmRoot, wcRoot)
	fmt.Println(len(etl.LogRecords))
	fmt.Println(len(etl.ElectricFieldRecords))
	fmt.Println(len(etl.WeatherCloudRecords))
	fmt.Println(len(etl.RecordHashTable))
	duration := time.Since(start)
	fmt.Println(duration)
}
