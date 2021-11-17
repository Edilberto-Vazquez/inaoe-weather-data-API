package services

import (
	"log"

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

func (ElectricFieldWeatherCloudService) Find(firstDate string, lastDate string, fields []string) ([]models.ElectricFieldWeatherCloud, int64, error) {
	log.Println(fields)
	var records []models.ElectricFieldWeatherCloud
	result := libs.DBCon.Select(fields).Where("time_stamp BETWEEN ? AND ?", firstDate, lastDate).Find(&records)
	return records, result.RowsAffected, result.Error
}
