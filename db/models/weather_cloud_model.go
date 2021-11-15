package models

import (
	"time"
)

type Weathercloud struct {
	TimeStamp time.Time
	TempIn    float64
	Temp      float64
	Chill     float64
	DewIn     float64
	Dew       float64
	HeatIn    float64
	Heat      float64
	Humin     float64
	Hum       float64
	Wspdhi    float64
	Wspdavg   float64
	Wdiravg   float64
	Bar       float64
	Rain      float64
	RainRate  float64
}

// func NewWeathercloud(path string, records []byte) Weathercloud {
// 	utf16, _ := etl.DecodeUtf16(records, binary.BigEndian)
// 	fields := strings.Split(utf16, ";")
// 	return Weathercloud{
// 		DateTime: etl.NewDateTime("wc", fields[0]),
// 		TempIn:   etl.CommaToPoint(fields[1]),
// 		Temp:     etl.CommaToPoint(fields[2]),
// 		Chill:    etl.CommaToPoint(fields[3]),
// 		DewIn:    etl.CommaToPoint(fields[4]),
// 		Dew:      etl.CommaToPoint(fields[5]),
// 		HeatIn:   etl.CommaToPoint(fields[6]),
// 		Heat:     etl.CommaToPoint(fields[7]),
// 		Humin:    etl.CommaToPoint(fields[8]),
// 		Hum:      etl.CommaToPoint(fields[9]),
// 		Wspdhi:   etl.CommaToPoint(fields[10]),
// 		Wspdavg:  etl.CommaToPoint(fields[11]),
// 		Wdiravg:  etl.CommaToPoint(fields[12]),
// 		Bar:      etl.CommaToPoint(fields[13]),
// 		Rain:     etl.CommaToPoint(fields[14]),
// 		RainRate: etl.CommaToPoint(fields[15]),
// 		PlaceID:  etl.NewPlace(path),
// 	}
// }
