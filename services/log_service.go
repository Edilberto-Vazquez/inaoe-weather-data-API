package services

import (
	"bufio"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/db/models"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/etl"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/libs"
)

type LogService struct{}

func NewLogService() *LogService {
	return &LogService{}
}

func newLog(str string) models.Log {
	return models.Log{
		DateTime:  etl.NewDateTime("log", str),
		Lightning: etl.NewLightning(),
		Distance:  etl.NewDistance(str),
		PlaceID:   etl.NewPlace(str),
	}
}

func (LogService) processFile(path string) (records []models.Log) {
	file := etl.OpenFile(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if etl.ThereIsLightning(scanner.Text()) {
			records = append(records, newLog(scanner.Text()))
		}
	}
	return
}

func (log LogService) CreateMultipleRecords(path string) (int64, error) {
	records := log.processFile(path)
	result := libs.DBCon.Create(&records)
	return result.RowsAffected, result.Error
}
