package etl

import "time"

type ElectricField struct {
	DateTime      time.Time
	ElectricField float64
	RotorStatus   bool
	PlaceID       int
}

func NewElectricField(path string, time string, electricFields []string, rotorStatus string) *ElectricField {
	return &ElectricField{
		PlaceID:       NewPlace(path),
		DateTime:      NewTimeStamp("efm", path+" "+time),
		ElectricField: ElectricFieldAvg(electricFields),
		RotorStatus:   NewRotorStatus(rotorStatus),
	}
}
