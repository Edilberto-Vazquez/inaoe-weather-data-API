package models

import (
	"time"
)

type ElectricFieldWeatherCloud struct {
	TimeStamp     time.Time
	Lightning     bool
	Distance      int64
	ElectricField float64
	RotorStatus   bool
	TempIn        float64
	Temp          float64
	Chill         float64
	DewIn         float64
	Dew           float64
	HeatIn        float64
	Heat          float64
	Humin         float64
	Hum           float64
	Wspdhi        float64
	Wspdavg       float64
	Wdiravg       float64
	Bar           float64
	Rain          float64
	RainRate      float64
}
