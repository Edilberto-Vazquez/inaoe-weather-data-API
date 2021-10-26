package etl

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	regexList = map[string]*regexp.Regexp{
		"date":      regexp.MustCompile(`\d\d/\d\d/\d\d\d\d`),
		"time":      regexp.MustCompile(`\d\d:\d\d:\d\d`),
		"place":     regexp.MustCompile(`INAOE parque|UPTLAX`),
		"lightning": regexp.MustCompile(`Lightning Detected`),
		"distance":  regexp.MustCompile(`at\s\d\d\skm|at\s\d\skm`),
	}
)

func ThereIsLightning(str string) bool {
	ok := regexList["lightning"].MatchString(str)
	if !ok {
		return ok
	}
	return ok
}

type EfmEventsRow struct {
	dateTime  string
	place     string
	lightning bool
	distance  int
}

func NewEfmEventsRow(str string) *EfmEventsRow {
	return &EfmEventsRow{
		dateTime:  formatDateTime(str),
		place:     formatPlace(str),
		lightning: formatLightning(),
		distance:  formatDistance(str),
	}
}

func formatDateTime(dateTime string) string {
	re := regexList["date"]
	date := re.FindString(dateTime)
	re = regexList["time"]
	time := re.FindString(dateTime)
	split := strings.Split(date, "/")
	date = split[2] + "-" + split[0] + "-" + split[1]
	return date + " " + time
}

func formatPlace(str string) (place string) {
	re := regexList["place"]
	place = re.FindString(str)
	return
}

func formatLightning() (lightning bool) {
	lightning = true
	return
}

func formatDistance(str string) (distance int) {
	re := regexList["distance"]
	match := re.FindString(str)
	split := strings.Split(match, " ")
	distance, _ = strconv.Atoi(split[1])
	return
}

// func GetField(str string, fields ...string) map[string]string {
// 	matchingFields := make(map[string]string)
// 	for _, field := range fields {
// 		re := regexList[field]
// 		if field == "dateTime" {
// 			matchingFields[field] = formatDateTime(str)
// 			continue
// 		}
// 		matchingFields[field] = re.FindString(str)
// 	}
// 	matchingFields["lightning"] = "true"
// 	return matchingFields
// }
