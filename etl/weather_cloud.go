package etl

import (
	"bufio"
)

type WeatherCloud struct {
	records []interface{}
}

func NewWeatherCloud() *WeatherCloud {
	return &WeatherCloud{
		records: make([]interface{}, 0),
	}
}

func (wc *WeatherCloud) ProcessFile(path string) {
	file := OpenFile(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 1 {
			continue
		}
		wc.records = append(wc.records, NewWeatherCloudRow(path, scanner.Text()))
	}
	wc.records = wc.records[1:]
}

func (wc *WeatherCloud) GetRecords() []interface{} {
	return wc.records
}
