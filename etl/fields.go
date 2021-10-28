package etl

import (
	"fmt"
	"strconv"
	"strings"
)

// -----chech if the log contains a lightning
func ThereIsLightning(str string) bool {
	ok := regexList["lightning"].MatchString(str)
	if !ok {
		return ok
	}
	return ok
}

// -----date-----
func newDateTime(str string) string {
	if regexList["date"].MatchString(str) {
		re := regexList["date"]
		date := re.FindString(str)
		re = regexList["time"]
		time := re.FindString(str)
		split := strings.Split(date, "/")
		date = split[2] + "-" + split[0] + "-" + split[1]
		return date + " " + time
	}
	split := strings.Split(str, "/")
	split = strings.Split(split[len(split)-1], "-")
	date := split[1]
	return date[4:8] + "-" + date[0:2] + "-" + date[2:4]
}

// -----place-----
func newPlace(str string) string {
	re := regexList["place"]
	return re.FindString(str)
}

func newLightning() bool {
	return true
}

// -----distance-----
func newDistance(str string) (distance int) {
	re := regexList["distance"]
	match := re.FindString(str)
	split := strings.Split(match, " ")
	distance, _ = strconv.Atoi(split[1])
	return
}

// calc the average of the electric field
func electricFieldAvg(electricFields []string) string {
	sum := 0.0
	divisor := float64(len(electricFields))
	for _, value := range electricFields {
		float, _ := strconv.ParseFloat(value, 32)
		sum += float
	}
	// return (math.Round((sum / divisor * 100)) / 100) // the function return a float64 if the rounding is with the function Round
	return fmt.Sprintf("%0.2f", (sum / divisor)) // the function return a string if the roundig is with the function Sprintf

}

// electric field
func newRotorStatus(rotorStatus string) bool {
	return rotorStatus != "0"
}

// weather cloud
func splitString(record string) []string {
	return strings.Split(record, ";")
}

func commaToPoint(str string) string {
	if str == "" {
		return "null"
	}
	return strings.Replace(str, ",", ".", 1)
}
