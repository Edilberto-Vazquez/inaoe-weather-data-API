package schemas

import "time"

type ElectricFieldSchema struct {
	FirstDate time.Time `form:"firstdate" binding:"required" time_format:"2006-01-02 15:04:05"` // todo checar este formato con la base de datos
	LastDate  time.Time `form:"lastdate" binding:"required" time_format:"2006-01-02 15:04:05"`
	Fields    []string  `form:"fields"`
}
