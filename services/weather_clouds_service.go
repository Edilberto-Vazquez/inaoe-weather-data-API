package services

import (
	"bufio"
	"encoding/binary"
	"log"
	"strings"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/db/models"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/etl"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/libs"
)

type WeathercloudService struct{}

func NewWeathercloudService() *WeathercloudService {
	return &WeathercloudService{}
}

func (WeathercloudService) processFile(path string) (records []models.Weathercloud) {
	file := etl.OpenFile(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		if len(scanner.Text()) == 1 {
			continue
		}
		if i == 0 {
			i++
			continue
		}
		utf16, err := etl.DecodeUtf16(scanner.Bytes(), binary.BigEndian)
		if err != nil {
			log.Println(err)
			continue
		}
		fields := strings.Split(utf16, ";")
		records = append(records, models.Weathercloud{
			DateTime: etl.NewDateTime("wc", fields[0]),
			TempIn:   etl.CommaToPoint(fields[1]),
			Temp:     etl.CommaToPoint(fields[2]),
			Chill:    etl.CommaToPoint(fields[3]),
			DewIn:    etl.CommaToPoint(fields[4]),
			Dew:      etl.CommaToPoint(fields[5]),
			HeatIn:   etl.CommaToPoint(fields[6]),
			Heat:     etl.CommaToPoint(fields[7]),
			Humin:    etl.CommaToPoint(fields[8]),
			Hum:      etl.CommaToPoint(fields[9]),
			Wspdhi:   etl.CommaToPoint(fields[10]),
			Wspdavg:  etl.CommaToPoint(fields[11]),
			Wdiravg:  etl.CommaToPoint(fields[12]),
			Bar:      etl.CommaToPoint(fields[13]),
			Rain:     etl.CommaToPoint(fields[14]),
			RainRate: etl.CommaToPoint(fields[15]),
			PlaceID:  etl.NewPlace(path),
		})
	}
	return
}

func (wcs WeathercloudService) CreateMultipleRecords(path string) (error, int64) {
	records := wcs.processFile(path)
	result := libs.DBCon.Create(&records)
	return result.Error, result.RowsAffected
}
