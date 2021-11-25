package schemas

import "time"

type INAOEQuerySchema struct {
	DatePart       string    `form:"datepart" binding:"required"`
	FromDate       time.Time `form:"fromdate" binding:"required" time_format:"2006-01-02T15:04:05"`
	ToDate         time.Time `form:"todate" binding:"required" time_format:"2006-01-02T15:04:05"`
	WeatherClouds  []string  `form:"weatherclouds"`
	ElectricFields []string  `form:"electricfields"`
	JoinType       string    `form:"jointype"`
}
