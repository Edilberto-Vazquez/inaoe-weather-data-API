package models

import (
	"time"
)

type Log struct {
	TimeStamp time.Time
	Lightning bool
	Distance  int64
}

// func NewLog(str string) Log {
// 	return Log{
// 		DateTime:  etl.NewDateTime("log", str),
// 		Lightning: etl.NewLightning(),
// 		Distance:  etl.NewDistance(str),
// 		PlaceID:   etl.NewPlace(str),
// 	}
// }
