package etl

import (
	"encoding/binary"
	"time"
)

type Weathercloud struct {
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
	PlaceID  int
}

func NewWeathercloud(path string, record []byte) *Weathercloud {
	utf16, _ := DecodeUtf16(record, binary.BigEndian)
	fields := SplitString(utf16)
	return &Weathercloud{
		DateTime: NewDateTime("wc", fields[0]),
		TempIn:   CommaToPoint(fields[1]),
		Temp:     CommaToPoint(fields[2]),
		Chill:    CommaToPoint(fields[3]),
		DewIn:    CommaToPoint(fields[4]),
		Dew:      CommaToPoint(fields[5]),
		HeatIn:   CommaToPoint(fields[6]),
		Heat:     CommaToPoint(fields[7]),
		Humin:    CommaToPoint(fields[8]),
		Hum:      CommaToPoint(fields[9]),
		Wspdhi:   CommaToPoint(fields[10]),
		Wspdavg:  CommaToPoint(fields[11]),
		Wdiravg:  CommaToPoint(fields[12]),
		Bar:      CommaToPoint(fields[13]),
		Rain:     CommaToPoint(fields[14]),
		RainRate: CommaToPoint(fields[15]),
		PlaceID:  NewPlace(path),
	}
}
