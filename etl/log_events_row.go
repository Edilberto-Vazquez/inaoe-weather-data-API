package etl

type LogEventsRow struct {
	DateTime  string
	Place     string
	Lightning bool
	Distance  int
}

func NewLogEventsRow(str string) *LogEventsRow {
	return &LogEventsRow{
		DateTime:  newDateTime(str),
		Place:     newPlace(str),
		Lightning: newLightning(),
		Distance:  newDistance(str),
	}
}
