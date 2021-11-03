package services

import (
	"bufio"
	"strings"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/db/models"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/etl"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/libs"
)

type ElectricFieldService struct{}

func NewElectricFieldService() *ElectricFieldService {
	return &ElectricFieldService{}
}

func (ElectricFieldService) processFile(path string) (records []models.ElectricField) {
	file := etl.OpenFile(path)
	defer file.Close()
	electricFields := make([]string, 0)
	var time string
	var rotorStatus string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		if time == "" || time == fields[0] {
			electricFields = append(electricFields, fields[1])
			time = fields[0]
			rotorStatus = fields[2]
		} else {
			records = append(records, models.ElectricField{
				DateTime:      etl.NewDateTime("efm", path+" "+time),
				ElectricField: etl.ElectricFieldAvg(electricFields),
				RotorStatus:   etl.NewRotorStatus(rotorStatus),
				PlaceID:       etl.NewPlace(path),
			})
			electricFields = make([]string, 0)
			electricFields = append(electricFields, fields[1])
			time = fields[0]
		}
	}
	return
}

func (efs ElectricFieldService) CreateMultipleRecords(path string) (int64, error) {
	records := efs.processFile(path)
	result := libs.DBCon.CreateInBatches(records, 1000)
	return result.RowsAffected, result.Error
}
