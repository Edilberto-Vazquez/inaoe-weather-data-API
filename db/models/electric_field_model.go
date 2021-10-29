package models

import (
	"time"

	"gorm.io/gorm"
)

type ElectricField struct {
	gorm.Model
	DateTime      time.Time `gorm:"<-:create;not null"`
	ElectricField float32   `gorm:"<-:create;not null"`
	RotorStatus   bool      `gorm:"<-:create;not null"`
	PlaceID       int
	Place         Place `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
