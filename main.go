package main

import (
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/services"
)

func main() {
	// routes.Run()
	// efm := etl.ProcesNewFile("/mnt/f/DataSets/Conjuntos-originales/medidor-campo-electrico/INAOE parque-01102019.efm", "efm")
	// log := etl.ProcesNewFile("/mnt/f/DataSets/Conjuntos-originales/medidor-campo-electrico/EFMEvents.log", "log")
	// wc := etl.ProcesNewFile("/mnt/f/DataSets/Conjuntos-originales/estacion-meteorologica/Weathercloud Pque Tecnolgico INAOE 2019-02.csv", "wc")

	// for _, v := range efm.GetRecords() {
	// 	fmt.Println(v.(*etl.ElectricFieldRow))
	// }
	// db := libs.ConnectionPool()
	// libs.InitCon()
	// wc := services.NewWeathercloudService()
	// wc.CreateMultipleRecords("/mnt/f/DataSets/Conjuntos-originales/estacion-meteorologica/Weathercloud Pque Tecnolgico INAOE 2019-02.csv")
	// efs := services.NewElectricFieldService()
	// efs.CreateMultipleRecords("/mnt/f/DataSets/Conjuntos-originales/medidor-campo-electrico/INAOE parque-01102019.efm")
	log := services.NewLogService()
	log.CreateMultipleRecords("/mnt/f/DataSets/Conjuntos-originales/medidor-campo-electrico/EFMEvents.log")
	// migrations.Migrate()
	// defer libs.CloseCon()
}
