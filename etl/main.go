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
	newEfmEvents := NewEfmEvents()
	// wcf := NewWeatherCloud()
	file := OpenFile(path)
	// newEFM.ProcessFile(file, path)
	newEfmEvents.ProcessFile(file)
	// wcf.ProcessFile(file, path)
	for _, v := range newEfmEvents.lines {
		fmt.Println(v)
	}
}
