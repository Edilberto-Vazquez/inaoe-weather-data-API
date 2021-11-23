package utils

import (
	"reflect"
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

func ConvertToSlice(slice interface{}) (sc []string) {
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		sc = append(sc, s.Index(i).Interface().(string))
	}
	return
}
