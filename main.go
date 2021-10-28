package main

import (
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/etl"
)

func main() {
	// routes.Run()
	// etl.ProcesNewFile("/mnt/f/DataSets/Conjuntos-originales/medidor-campo-electrico/INAOE parque-01102019.efm")
	// etl.ProcesNewFile("/mnt/f/DataSets/Conjuntos-originales/medidor-campo-electrico/EFMEvents.log")
	etl.ProcesNewFile("/mnt/f/DataSets/Conjuntos-originales/estacion-meteorologica/Weathercloud Pque Tecnolgico INAOE 2019-02.csv")
}
