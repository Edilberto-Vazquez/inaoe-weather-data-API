package utils

import (
	"strings"
	"time"
)

func ParseTimeStamp(timeStamp string) (string, error) {
	ts, err := time.Parse("2006-01-02T15:04:05", timeStamp)
	if err != nil {
		return ts.String(), err
	}
	return strings.Replace(ts.String(), " +0000 UTC", "", 1), nil
}
