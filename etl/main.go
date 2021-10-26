package etl

import (
	"log"
	"os"
)

func ProcesNewFile(path string) *EfmEvents {
	// path := "/mnt/f/DataSets/Conjuntos-originales/medidor-campo-electrico/INAOE parque-01102019.efm"
	// newEFM := NewEfmFile()
	newEfmEvents := NewEfmEvents()
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	// newEFM.ProcessFile(file, path)
	newEfmEvents.ProcessEfmEvents(file)

	return newEfmEvents
}
