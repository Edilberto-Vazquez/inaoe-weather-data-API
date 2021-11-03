package models

import (
	"encoding/binary"
	"strings"
	"time"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/etl"
	"gorm.io/gorm"
)

type Weathercloud struct {
	gorm.Model
	DateTime time.Time `gorm:"<-:create;not null"`
	TempIn   float64   `gorm:"<-:create"`
	Temp     float64   `gorm:"<-:create"`
	Chill    float64   `gorm:"<-:create"`
	DewIn    float64   `gorm:"<-:create"`
	Dew      float64   `gorm:"<-:create"`
	HeatIn   float64   `gorm:"<-:create"`
	Heat     float64   `gorm:"<-:create"`
	Humin    float64   `gorm:"<-:create"`
	Hum      float64   `gorm:"<-:create"`
	Wspdhi   float64   `gorm:"<-:create"`
	Wspdavg  float64   `gorm:"<-:create"`
	Wdiravg  float64   `gorm:"<-:create"`
	Bar      float64   `gorm:"<-:create"`
	Rain     float64   `gorm:"<-:create"`
	RainRate float64   `gorm:"<-:create"`
	PlaceID  int
	Place    Place `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewWeathercloud(path string, records []byte) Weathercloud {
	utf16, _ := etl.DecodeUtf16(records, binary.BigEndian)
	fields := strings.Split(utf16, ";")
	return Weathercloud{
		DateTime: etl.NewDateTime("wc", fields[0]),
		TempIn:   etl.CommaToPoint(fields[1]),
		Temp:     etl.CommaToPoint(fields[2]),
		Chill:    etl.CommaToPoint(fields[3]),
		DewIn:    etl.CommaToPoint(fields[4]),
		Dew:      etl.CommaToPoint(fields[5]),
		HeatIn:   etl.CommaToPoint(fields[6]),
		Heat:     etl.CommaToPoint(fields[7]),
		Humin:    etl.CommaToPoint(fields[8]),
		Hum:      etl.CommaToPoint(fields[9]),
		Wspdhi:   etl.CommaToPoint(fields[10]),
		Wspdavg:  etl.CommaToPoint(fields[11]),
		Wdiravg:  etl.CommaToPoint(fields[12]),
		Bar:      etl.CommaToPoint(fields[13]),
		Rain:     etl.CommaToPoint(fields[14]),
		RainRate: etl.CommaToPoint(fields[15]),
		PlaceID:  etl.NewPlace(path),
	}
}
