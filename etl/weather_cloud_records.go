package etl

import (
	"bufio"
)

type WeathercloudRows struct {
	records []interface{}
}

func NewWeathercloudRecords() *WeathercloudRows {
	return &WeathercloudRows{
		records: make([]interface{}, 0),
	}
}

func (wc *WeathercloudRows) ProcessFile(path string) {
	file := OpenFile(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		if len(scanner.Text()) == 1 {
			continue
		}
		if i == 0 {
			i++
			continue
		}
		wc.records = append(wc.records, NewWeathercloud(path, scanner.Bytes()))
	}
	wc.records = wc.records[1:]
}

func (wc *WeathercloudRows) GetRecords() []interface{} {
	return wc.records
}
