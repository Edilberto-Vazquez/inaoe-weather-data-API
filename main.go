package main

import (
	"fmt"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/etl"
)

func main() {
	// routes.Run()
	efm := etl.ProcesNewFile("/mnt/f/DataSets/Conjuntos-originales/medidor-campo-electrico/INAOE parque-01102019.efm", "efm")
	// log := etl.ProcesNewFile("/mnt/f/DataSets/Conjuntos-originales/medidor-campo-electrico/EFMEvents.log", "log")
	// wc := etl.ProcesNewFile("/mnt/f/DataSets/Conjuntos-originales/estacion-meteorologica/Weathercloud Pque Tecnolgico INAOE 2019-02.csv", "wc")

	for _, v := range efm.GetRecords() {
		fmt.Println(v.(*etl.ElectricFieldRow))
	}
}
