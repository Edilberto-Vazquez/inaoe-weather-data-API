package models

import (
	"time"
)

type ElectricFieldWeatherCloud struct {
	TimeStamp     *time.Time `json:"timestamp,omitempty"`
	Lightning     *bool      `json:"lightning,omitempty"`
	Distance      *uint8     `json:"distance,omitempty"`
	ElectricField *float32   `json:"electricfield,omitempty"`
	RotorStatus   *bool      `json:"rotorstatus,omitempty"`
	TempIn        *float32   `json:"tampin,omitempty"`
	Temp          *float32   `json:"temp,omitempty"`
	Chill         *float32   `json:"chill,omitempty"`
	DewIn         *float32   `json:"dewin,omitempty"`
	Dew           *float32   `json:"dew,omitempty"`
	HeatIn        *float32   `json:"hetin,omitempty"`
	Heat          *float32   `json:"heat,omitempty"`
	Humin         *float32   `json:"humin,omitempty"`
	Hum           *float32   `json:"hum,omitempty"`
	Wspdhi        *float32   `json:"wsphi,omitempty"`
	Wspdavg       *float32   `json:"wspdavg,omitempty"`
	Wdiravg       *float32   `json:"wdiravg,omitempty"`
	Bar           *float32   `json:"bar,omitempty"`
	Rain          *float32   `json:"rain,omitempty"`
	RainRate      *float32   `json:"rainrate,omitempty"`
}
