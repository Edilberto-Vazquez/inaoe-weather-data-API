package models

import (
	"time"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/etl"
	"gorm.io/gorm"
)

type ElectricField struct {
	gorm.Model
	DateTime      time.Time `gorm:"<-:create;not null"`
	ElectricField float64   `gorm:"<-:create;not null"`
	RotorStatus   bool      `gorm:"<-:create;not null"`
	PlaceID       int
	Place         Place `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewElectricField(dateTime string, electricFields []string, rotorStatus string, path string) ElectricField {
	return ElectricField{
		DateTime:      etl.NewDateTime("efm", dateTime),
		ElectricField: etl.ElectricFieldAvg(electricFields),
		RotorStatus:   etl.NewRotorStatus(rotorStatus),
		PlaceID:       etl.NewPlace(path),
	}
}
