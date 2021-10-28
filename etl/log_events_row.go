package etl

type LogEventsRow struct {
	dateTime  string
	place     string
	lightning bool
	distance  int
}

func NewLogEventsRow(str string) *LogEventsRow {
	return &LogEventsRow{
		dateTime:  newDateTime(str),
		place:     newPlace(str),
		lightning: newLightning(),
		distance:  newDistance(str),
	}
}
