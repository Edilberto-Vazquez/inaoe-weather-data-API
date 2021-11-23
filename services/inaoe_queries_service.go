package services

import (
	"fmt"
	"reflect"
	"time"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/libs"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/utils"
	"gorm.io/gorm"
)

var (
	tables = map[string][]string{
		"weather_clouds": {
			"temp_in",
			"temp",
			"chill",
			"dew_in",
			"heat_in",
			"dew",
			"heat_in",
			"hum_in",
			"hum",
			"wspdhi",
			"wspdavg",
			"wdiravg",
			"bar",
			"rain",
			"rain_rate",
		},
		"electric_fields": {
			"lightning",
			"distance",
			"electric_field",
			"rotor_status",
		},
	}
)

type INAOEQuerysService struct{}

type Result struct {
	Date               *time.Time `json:"date,omitempty"`
	TotalLightning     *int       `json:"total_lightning,omitempty"`
	TotalDistance      *int       `json:"total_distance,omitempty"`
	TotalElectricField *int       `json:"total_electric_field,omitempty"`
	TotalRotorStatus   *int       `json:"total_rotor_status,omitempty"`
	TotalTempIn        *int       `json:"total_temp_in,omitempty"`
	TotalTemp          *int       `json:"total_temp,omitempty"`
	TotalChill         *int       `json:"total_chill,omitempty"`
	TotalDewIn         *int       `json:"total_dew_in,omitempty"`
	TotalDew           *int       `json:"total_dew,omitempty"`
	TotalHeatIn        *int       `json:"total_heat_in,omitempty"`
	TotalHeat          *int       `json:"total_heat,omitempty"`
	TotalHumin         *int       `json:"total_hum_in,omitempty"`
	TotalHum           *int       `json:"total_hum,omitempty"`
	TotalWspdhi        *int       `json:"total_wspdhi,omitempty"`
	TotalWspdavg       *int       `json:"total_wspdavg,omitempty"`
	TotalWdiravg       *int       `json:"total_wdiravg,omitempty"`
	TotalBar           *int       `json:"total_bar,omitempty"`
	TotalRain          *int       `json:"total_rain,omitempty"`
	TotalRainRate      *int       `json:"total_rain_rate,omitempty"`
}

func NewINAOEQuerysService() *INAOEQuerysService {
	return &INAOEQuerysService{}
}

func ConvertToSlice(slice interface{}) (sc []string) {
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		sc = append(sc, s.Index(i).Interface().(string))
	}
	return
}

func Count(fields []string, table string) (count string) {
	if fields != nil {
		for _, field := range fields {
			count += "COUNT(" + field + ") total_" + field + ","
		}
	} else {
		for _, field := range tables[table] {
			count += "COUNT(" + field + ") total_" + field + ","
		}
	}
	return count
}

func (INAOEQuerysService) JoinTypeFind(queryOptions map[string]interface{}) ([]Result, int64, error) {
	var results []Result
	var count string
	datePart := queryOptions["datepart"].(string)
	fromDate, fromErr := utils.ParseTimeStamp(queryOptions["fromdate"].(string))
	toDate, toErr := utils.ParseTimeStamp(queryOptions["todate"].(string))
	if fromErr != nil || toErr != nil {
		return nil, 0, fmt.Errorf("cannot parse %v, %v", fromErr, toErr)
	}
	weatherCloud := ConvertToSlice(queryOptions["weatherclouds"])
	electricField := ConvertToSlice(queryOptions["electricfields"])
	joinType := queryOptions["jointype"].(string)
	count += Count(weatherCloud, "weather_clouds")
	count += Count(electricField, "electric_fields")
	count = count[0 : len(count)-1]
	response := libs.DBCon.Table("weather_clouds").
		Joins(joinType+" JOIN electric_fields ON weather_clouds.time_stamp = electric_fields.time_stamp_round").
		Select("DATE_TRUNC('"+datePart+"', weather_clouds.time_stamp) date,"+count).
		Where("weather_clouds.time_stamp BETWEEN ? AND ?", fromDate, toDate).
		Group("date").
		Order("date").
		Find(&results)
	if response.Error != nil {
		return nil, response.RowsAffected, response.Error
	}
	for _, result := range results {
		*result.Date = result.Date.UTC()
	}
	return results, response.RowsAffected, response.Error
}

func (INAOEQuerysService) Find(queryOptions map[string]interface{}) ([]Result, int64, error) {
	var results []Result
	var response *gorm.DB
	var count string
	datePart := queryOptions["datepart"].(string)
	fromdate, fromErr := utils.ParseTimeStamp(queryOptions["fromdate"].(string))
	toDate, toErr := utils.ParseTimeStamp(queryOptions["todate"].(string))
	weatherCloud := ConvertToSlice(queryOptions["weatherclouds"])
	electricField := ConvertToSlice(queryOptions["electricfields"])
	var table string
	if fromErr != nil || toErr != nil {
		return nil, 0, fmt.Errorf("cannot parse %v, %v", fromErr, toErr)
	}
	if weatherCloud != nil {
		count = Count(weatherCloud, "weather_clouds")
		table = "weather_clouds"
	}
	if electricField != nil {
		count = Count(electricField, "electric_fields")
		table = "electric_fields"
	}
	count = count[0 : len(count)-1]
	response = libs.DBCon.Table(table).
		Select("DATE_TRUNC('"+datePart+"', time_stamp) date,"+count).
		Where("time_stamp BETWEEN ? AND ?", fromdate, toDate).
		Group("date").
		Order("date").
		Find(&results)
	if response.Error != nil {
		return nil, response.RowsAffected, response.Error
	}
	for _, result := range results {
		*result.Date = result.Date.UTC()
	}
	return results, response.RowsAffected, response.Error
}