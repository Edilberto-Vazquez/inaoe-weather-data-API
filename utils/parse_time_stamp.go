package utils

import (
	"reflect"
	"regexp"
	"time"
)

func FormatTimeStamp(timeStamp time.Time) (string, error) {
	r := regexp.MustCompile(`\s(\+|\-)\d\d\d\d\s\w\w\w`)
	return r.ReplaceAllString(timeStamp.String(), ""), nil
}

func ConvertToSlice(slice interface{}) (sc []string) {
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		sc = append(sc, s.Index(i).Interface().(string))
	}
	return
}
