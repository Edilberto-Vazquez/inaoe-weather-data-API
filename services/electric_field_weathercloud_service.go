package services

import (
	"strings"
	"time"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/db/models"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/libs"
)

type ElectricFieldWeatherCloudService struct{}

func NewElectricFieldWeatherCloudService() *ElectricFieldWeatherCloudService {
	return &ElectricFieldWeatherCloudService{}
}

func (ElectricFieldWeatherCloudService) CreateRecords(records []*models.ElectricFieldWeatherCloud) (int64, error) {
	result := libs.DBCon.CreateInBatches(records, 1000)
	return result.RowsAffected, result.Error
}

func (ElectricFieldWeatherCloudService) Find(firstDate string, lastDate string, fields []string) ([]*models.ElectricFieldWeatherCloud, int64, error) {
	fd, _ := time.Parse(time.RFC3339, firstDate+"Z")
	ld, _ := time.Parse(time.RFC3339, lastDate+"Z")
	var records []*models.ElectricFieldWeatherCloud
	fdR := strings.Replace(fd.String(), "+0000 UTC", "+00", 1)
	ldR := strings.Replace(ld.String(), "+0000 UTC", "+00", 1)
	result := libs.DBCon.Table("electric_field_weather_clouds").Select(fields).Where("time_stamp BETWEEN ? AND ?", fdR, ldR).Find(&records)
	return records, result.RowsAffected, result.Error
}
