package etl

import "time"

type LogEventsRow struct {
	DateTime  time.Time
	Lightning bool
	Distance  int64
	Place     int
}

func NewLogEventsRow(str string) *LogEventsRow {
	return &LogEventsRow{
		DateTime:  newDateTime("log", str),
		Place:     newPlace(str),
		Lightning: newLightning(),
		Distance:  newDistance(str),
	}
}
