package services

// import (
// 	"bufio"

// 	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/db/models"
// 	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/etl"
// 	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/libs"
// )

// type WeathercloudService struct{}

// func NewWeathercloudService() *WeathercloudService {
// 	return &WeathercloudService{}
// }

// func (WeathercloudService) processFile(path string) (records []models.Weathercloud) {
// 	file := etl.OpenFile(path)
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	i := 0
// 	for scanner.Scan() {
// 		if len(scanner.Text()) == 1 {
// 			continue
// 		}
// 		if i == 0 {
// 			i++
// 			continue
// 		}
// 		records = append(records, models.NewWeathercloud(path, scanner.Bytes()))
// 	}
// 	return
// }

// func (wcs WeathercloudService) CreateMultipleRecords(path string) (int64, error) {
// 	records := wcs.processFile(path)
// 	result := libs.DBCon.CreateInBatches(records, 1000)
// 	return result.RowsAffected, result.Error
// }
