package schemas

import "time"

type ElectricFieldSchema struct {
	FirstDate time.Time `form:"firstdate" binding:"required" time_format:"2006-01-02T15:04:05"`
	LastDate  time.Time `form:"lastdate" binding:"required" time_format:"2006-01-02T15:04:05"`
	Fields    []string  `form:"fields"`
}
