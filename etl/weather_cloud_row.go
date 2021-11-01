package etl

import (
	"time"
)

type WeatherCloudRow struct {
	Place    int
	DateTime time.Time
	TempIn   float64
	Temp     float64
	Chill    float64
	DewIn    float64
	Dew      float64
	HeatIn   float64
	Heat     float64
	Humin    float64
	Hum      float64
	Wspdhi   float64
	Wspdavg  float64
	Wdiravg  float64
	Bar      float64
	Rain     float64
	RainRate float64
}

func NewWeatherCloudRow(path string, record string) *WeatherCloudRow {
	fields := splitString(record)
	return &WeatherCloudRow{
		Place:    newPlace(path),
		DateTime: newDateTime("wc", fields[0]),
		TempIn:   commaToPoint(fields[1]),
		Temp:     commaToPoint(fields[2]),
		Chill:    commaToPoint(fields[3]),
		DewIn:    commaToPoint(fields[4]),
		Dew:      commaToPoint(fields[5]),
		HeatIn:   commaToPoint(fields[6]),
		Heat:     commaToPoint(fields[7]),
		Humin:    commaToPoint(fields[8]),
		Hum:      commaToPoint(fields[9]),
		Wspdhi:   commaToPoint(fields[10]),
		Wspdavg:  commaToPoint(fields[11]),
		Wdiravg:  commaToPoint(fields[12]),
		Bar:      commaToPoint(fields[13]),
		Rain:     commaToPoint(fields[14]),
		RainRate: commaToPoint(fields[15]),
	}
}
