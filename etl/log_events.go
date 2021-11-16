package etl

import "time"

type LogEvents struct {
	DateTime  time.Time
	Lightning bool
	Distance  int64
	PlaceID   int
}

func NewLogEvents(str string) *LogEvents {
	return &LogEvents{
		DateTime:  NewTimeStamp("log", str),
		PlaceID:   NewPlace(str),
		Lightning: NewLightning(),
		Distance:  NewDistance(str),
	}
}
