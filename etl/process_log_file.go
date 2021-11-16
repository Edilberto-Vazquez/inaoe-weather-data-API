package etl

import (
	"time"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/db/models"
)

func ProcessLogFile(path string) (records map[time.Time]models.Log) {
	records = make(map[time.Time]models.Log, 0)
	file, scanner := OpenFile(path)
	defer file.Close()
	for scanner.Scan() {
		if ThereIsLightning(scanner.Text()) {
			record := models.Log{
				TimeStamp: NewTimeStamp("log", scanner.Text()),
				Lightning: NewLightning(),
				Distance:  NewDistance(scanner.Text()),
			}
			records[record.TimeStamp] = record
		}
	}
	return
}
