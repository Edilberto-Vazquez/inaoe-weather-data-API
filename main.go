package main

import (
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/libs"
)

func main() {
	libs.InitCon()
	defer libs.CloseCon()
	// migrations.Migrate()
	// wc := services.NewWeathercloudService()
	// wc.CreateMultipleRecords("/mnt/f/DataSets/Conjuntos-originales/estacion-meteorologica/Weathercloud Pque Tecnolgico INAOE 2019-02.csv")
	// efs := services.NewElectricFieldService()
	// efs.CreateMultipleRecords("/mnt/f/DataSets/Conjuntos-originales/medidor-campo-electrico/INAOE parque-01102019.efm")
	// log := services.NewLogService()
	// log.CreateMultipleRecords("/mnt/f/DataSets/Conjuntos-originales/medidor-campo-electrico/EFMEvents.log")
}
