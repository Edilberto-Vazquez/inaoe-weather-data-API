package etl

type LogEventsRow struct {
	dateTime  string
	place     string
	lightning bool
	distance  int
}

func NewLogEventsRow(str string) *LogEventsRow {
	return &LogEventsRow{
		dateTime:  newDateTime(str),
		place:     newPlace(str),
		lightning: newLightning(),
		distance:  newDistance(str),
	}
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
