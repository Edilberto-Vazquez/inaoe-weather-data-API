package models

type Place struct {
	ID   int    `gorm:"<-:create;primaryKey;unique;autoIncrement;not null"`
	Name string `gorm:"<-:create;unique;not null"`
}
