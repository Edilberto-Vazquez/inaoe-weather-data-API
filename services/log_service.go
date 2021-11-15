package services

// import (
// 	"bufio"

// 	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/db/models"
// 	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/etl"
// 	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/libs"
// )

// type LogService struct{}

// func NewLogService() *LogService {
// 	return &LogService{}
// }

// func (LogService) processFile(path string) (records []models.Log) {
// 	file := etl.OpenFile(path)
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		if etl.ThereIsLightning(scanner.Text()) {
// 			records = append(records, models.NewLog(scanner.Text()))
// 		}
// 	}
// 	return
// }

// func (log LogService) CreateMultipleRecords(path string) (int64, error) {
// 	records := log.processFile(path)
// 	result := libs.DBCon.CreateInBatches(records, 1000)
// 	return result.RowsAffected, result.Error
// }
