package schemas

import "time"

type ElectricFieldSchema struct {
	PlaceID     int       `form:"placeid" binding:"required"`
	InitialDate time.Time `form:"initialdate" binding:"required" time_format:"2006-01-02T15:04:05"`
	FinalDate   time.Time `form:"finaldate" binding:"required" time_format:"2006-01-02T15:04:05"`
}
