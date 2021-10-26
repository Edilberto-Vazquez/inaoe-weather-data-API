package etl

import (
	"log"
	"os"
)

func ProcesNewFile(path string) *EfmFile {
	// path := "/mnt/f/DataSets/Conjuntos-originales/medidor-campo-electrico/INAOE parque-01102019.efm"
	newEFM := NewEfmFile()
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	newEFM.ProcessFile(file, path)

	return newEFM
}
