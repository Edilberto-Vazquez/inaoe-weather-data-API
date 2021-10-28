package etl

import (
	"fmt"
	"log"
	"os"
)

func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	return file
}

func ProcesNewFile(path string) {
	// path := "/mnt/f/DataSets/Conjuntos-originales/medidor-campo-electrico/INAOE parque-01102019.efm"
	// newEFM := NewEfmFile()
	// newEfmEvents := NewEfmEvents()
	wcf := NewWeatherCloudFile()
	file := OpenFile(path)
	defer file.Close()
	// newEFM.ProcessFile(file, path)
	// newEfmEvents.ProcessEfmEvents(file)
	wcf.ProcessWeatherCloudFile(file, path)
	for _, v := range wcf.records {
		fmt.Println(v)
	}
}
