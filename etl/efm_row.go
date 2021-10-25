package etl

import (
	"fmt"
	"strconv"
	"strings"
)

type EfmRows struct {
	place         string
	dateTime      string
	electricField string
	rotorStatus   bool
}

func NewEfmRows(fileName string, time string, electricFields []string, rotorStatus string) *EfmRows {
	return &EfmRows{
		place:         newPlace(fileName),
		dateTime:      newDateTime(fileName, time),
		electricField: electricFieldAvg(electricFields),
		rotorStatus:   newRotorStatus(rotorStatus),
	}
}

func newPlace(fileName string) (place string) {
	split := strings.Split(fileName, "/")
	split = strings.Split(split[len(split)-1], "-")
	// split = strings.Split(split[len(split)-1], " ")
	place = split[0]
	return
}

func newDateTime(fileName string, time string) (dateTime string) {
	split := strings.Split(fileName, "/")
	split = strings.Split(split[len(split)-1], "-")
	date := split[1]
	dateTime = date[4:8] + "-" + date[0:2] + "-" + date[2:4] + " " + time
	return
}

func newRotorStatus(rotorStatus string) (status bool) {
	if rotorStatus == "0" {
		status = false
	} else {
		status = true
	}
	return
}

func electricFieldAvg(electricFields []string) (electricField string) {
	sum := 0.0
	divisor := float64(len(electricFields))
	for _, value := range electricFields {
		float, _ := strconv.ParseFloat(value, 32)
		sum += float
	}
	// electricField = (math.Round((sum / divisor * 100)) / 100) // the function return a float64 if the rounding is with the function Round
	electricField = fmt.Sprintf("%0.2f", (sum / divisor)) // the function return a string if the roundig is with the function Sprintf
	return
}
