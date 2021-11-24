package schemas

import "time"

type INAOEQuerySchema struct {
	DatePart       string    `json:"datepart" binding:"required"`
	FromDate       time.Time `json:"fromdate" binding:"required" time_format:"2006-01-02T15:04:05"`
	ToDate         time.Time `json:"todate" binding:"required" time_format:"2006-01-02T15:04:05"`
	WeatherClouds  []string  `json:"weatherclouds"`
	ElectricFields []string  `json:"electricfields"`
	JoinType       string    `json:"jointype"`
}
