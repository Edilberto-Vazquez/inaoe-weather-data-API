package utils

import (
	"reflect"
	"strings"
	"time"
)

func FormatTimeStamp(timeStamp time.Time) (string, error) {
	return strings.Replace(timeStamp.String(), " +0000 UTC", "", 1), nil
}

func ConvertToSlice(slice interface{}) (sc []string) {
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		sc = append(sc, s.Index(i).Interface().(string))
	}
	return
}
