package models

import (
	"time"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/etl"
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	DateTime  time.Time `gorm:"<-:create;not null"`
	Lightning bool      `gorm:"<-:create;not null"`
	Distance  int64     `gorm:"<-:create;not null"`
	PlaceID   int
	Place     Place `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLog(str string) Log {
	return Log{
		DateTime:  etl.NewDateTime("log", str),
		Lightning: etl.NewLightning(),
		Distance:  etl.NewDistance(str),
		PlaceID:   etl.NewPlace(str),
	}
}
