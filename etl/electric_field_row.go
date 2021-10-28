package etl

type ElectricFieldRow struct {
	place         string
	dateTime      string
	electricField string
	rotorStatus   bool
}

func NewElectricFieldRow(path string, time string, electricFields []string, rotorStatus string) *ElectricFieldRow {
	return &ElectricFieldRow{
		place:         newPlace(path),
		dateTime:      newDateTime(path) + " " + time,
		electricField: electricFieldAvg(electricFields),
		rotorStatus:   newRotorStatus(rotorStatus),
	}
}
