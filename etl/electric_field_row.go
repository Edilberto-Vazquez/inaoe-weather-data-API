package etl

import "time"

type ElectricFieldRow struct {
	DateTime      time.Time
	ElectricField float64
	RotorStatus   bool
	Place         int
}

func NewElectricFieldRow(path string, time string, electricFields []string, rotorStatus string) *ElectricFieldRow {
	return &ElectricFieldRow{
		Place:         newPlace(path),
		DateTime:      newDateTime("efm", path+" "+time),
		ElectricField: electricFieldAvg(electricFields),
		RotorStatus:   newRotorStatus(rotorStatus),
	}
}
