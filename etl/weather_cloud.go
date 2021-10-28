package etl

import (
	"bufio"
	"os"
)

type WeatherCloud struct {
	records []*WeatherCloudRow
}

func NewWeatherCloud() *WeatherCloud {
	return &WeatherCloud{
		records: make([]*WeatherCloudRow, 0),
	}
}

func (wc *WeatherCloud) ProcessFile(file *os.File, path string) {
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
