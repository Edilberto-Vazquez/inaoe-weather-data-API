package services

import (
	"fmt"
	"time"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/libs"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/schemas"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/utils"
)

type ByDateService struct{}

type Data struct {
	Date          *time.Time `json:"date,omitempty"`
	Distance      *float32   `json:"distance,omitempty"`
	ElectricField *float32   `json:"electric_field,omitempty"`
	TempIn        *float32   `json:"temp_in,omitempty"`
	Temp          *float32   `json:"temp,omitempty"`
	Chill         *float32   `json:"chill,omitempty"`
	DewIn         *float32   `json:"dew_in,omitempty"`
	Dew           *float32   `json:"dew,omitempty"`
	HeatIn        *float32   `json:"heat_in,omitempty"`
	Heat          *float32   `json:"heat,omitempty"`
	Humin         *float32   `json:"hum_in,omitempty"`
	Hum           *float32   `json:"hum,omitempty"`
	Wspdhi        *float32   `json:"wspdhi,omitempty"`
	Wspdavg       *float32   `json:"wspdavg,omitempty"`
	Wdiravg       *float32   `json:"wdiravg,omitempty"`
	Bar           *float32   `json:"bar,omitempty"`
	Rain          *float32   `json:"rain,omitempty"`
	RainRate      *float32   `json:"rain_rate,omitempty"`
}

func (ByDateService) FindByDate(byDateSchema schemas.ByDateSchema) ([]Data, int64, error) {
	var data []Data
	fromDate, fromErr := utils.FormatTimeStamp(byDateSchema.FromDate)
	toDate, toErr := utils.FormatTimeStamp(byDateSchema.ToDate)
	if fromErr != nil || toErr != nil {
		return nil, 0, fmt.Errorf("%v, %v", fromErr, toErr)
	}
	avgQuery := utils.AverageQuery(byDateSchema.Fields)
	res := libs.DBCon.Table(byDateSchema.Table).
		Select("DATE_TRUNC('"+byDateSchema.DatePart+"', time_stamp) date,"+avgQuery).
		Where("time_stamp BETWEEN ? AND ?", fromDate, toDate).
		Group("date").
		Order("date").
		Find(&data)
	return data, res.RowsAffected, res.Error
}
