package models

import (
	"time"

	"gorm.io/gorm"
)

type Weathercloud struct {
	gorm.Model
	DateTime time.Time `gorm:"<-:create;not null"`
	TempIn   float32   `gorm:"<-:create;not null"`
	Temp     float32   `gorm:"<-:create;not null"`
	Chill    float32   `gorm:"<-:create;not null"`
	DewIn    float32   `gorm:"<-:create;not null"`
	Dew      float32   `gorm:"<-:create;not null"`
	HeatIn   float32   `gorm:"<-:create;not null"`
	Heat     float32   `gorm:"<-:create;not null"`
	Humin    float32   `gorm:"<-:create;not null"`
	Hum      float32   `gorm:"<-:create;not null"`
	Wspdhi   float32   `gorm:"<-:create;not null"`
	Wspdavg  float32   `gorm:"<-:create;not null"`
	Wdiravg  float32   `gorm:"<-:create;not null"`
	Bar      float32   `gorm:"<-:create;not null"`
	Rain     float32   `gorm:"<-:create;not null"`
	RainRate float32   `gorm:"<-:create;not null"`
	PlaceID  int
	Place    Place `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
