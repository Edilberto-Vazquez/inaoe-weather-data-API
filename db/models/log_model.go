package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	DateTime  string `gorm:"<-:create;not null"`
	Lightning bool   `gorm:"<-:create;not null"`
	Distance  uint8  `gorm:"<-:create;not null"`
	PlaceID   int
	Place     Place `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
