package models

import (
	"time"

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
