package models

import (
	"time"

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
