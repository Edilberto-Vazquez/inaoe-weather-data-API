package etl

import "log"

var (
	filesType = map[string]FilesTypes{
		"log": NewLogEvents(),
		"efm": NewElectricFields(),
		"wc":  NewWeatherCloud(),
	}
)

func ProcesNewFile(path string, fileType string) FilesTypes {
	ft, ok := filesType[fileType]
	if !ok {
		log.Printf("Can not open the file %s\n", path)
		return nil
	}
	ft.ProcessFile(path)
	return ft
}
