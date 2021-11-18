package schemas

import "time"

type ElectricFieldAndWeatherCloudSchema struct {
	FirstDate time.Time `form:"firstdate" binding:"required" time_format:"2006-01-02 15:04:05"`
	LastDate  time.Time `form:"lastdate" time_format:"2006-01-02 15:04:05"`
	Fields    []string  `form:"fields"`
}

func (ElectricFieldAndWeatherCloudSchema) NewSchema() ElectricFieldAndWeatherCloudSchema {
	return ElectricFieldAndWeatherCloudSchema{}
}
