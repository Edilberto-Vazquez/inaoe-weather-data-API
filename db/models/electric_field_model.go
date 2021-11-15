package models

import (
	"time"
)

type ElectricField struct {
	TimeStamp     time.Time
	ElectricField float64
	RotorStatus   bool
}

// func NewElectricField(dateTime string, electricFields []string, rotorStatus string, path string) ElectricField {
// 	return ElectricField{
// 		DateTime:      etl.NewDateTime("efm", dateTime),
// 		ElectricField: etl.ElectricFieldAvg(electricFields),
// 		RotorStatus:   etl.NewRotorStatus(rotorStatus),
// 		PlaceID:       etl.NewPlace(path),
// 	}
// }
