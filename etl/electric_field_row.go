package etl

type ElectricFieldRow struct {
	Place         string
	DateTime      string
	ElectricField string
	RotorStatus   bool
}

func NewElectricFieldRow(path string, time string, electricFields []string, rotorStatus string) *ElectricFieldRow {
	return &ElectricFieldRow{
		Place:         newPlace(path),
		DateTime:      newDateTime(path) + " " + time,
		ElectricField: electricFieldAvg(electricFields),
		RotorStatus:   newRotorStatus(rotorStatus),
	}
}
