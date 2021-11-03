package etl

import "log"

var (
	filesType = map[string]FilesTypes{
		"log": NewLogEventsRecords(),
		"efm": NewElectricFieldRecords(),
		"wc":  NewWeathercloudRecords(),
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
