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

func (wcf *WeatherCloud) ProcessFile(file *os.File, path string) {
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 1 {
			continue
		}
		wcf.records = append(wcf.records, NewWeatherCloudRow(path, scanner.Text()))
	}
	wcf.records = wcf.records[1:]
}
