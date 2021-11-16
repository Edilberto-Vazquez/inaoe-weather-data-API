package services

import (
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
