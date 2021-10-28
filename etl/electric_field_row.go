package etl

type ElectricFieldRow struct {
	place         string
	dateTime      string
	electricField string
	rotorStatus   bool
}

func NewElectricFieldRow(fileName string, time string, electricFields []string, rotorStatus string) *ElectricFieldRow {
	return &ElectricFieldRow{
		place:         newPlace(fileName),
		dateTime:      newDateTime(fileName) + " " + time,
		electricField: electricFieldAvg(electricFields),
		rotorStatus:   newRotorStatus(rotorStatus),
	}
}
