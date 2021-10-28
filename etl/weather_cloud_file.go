package etl

import (
	"bufio"
	"os"
)

type WeatherCloudFile struct {
	records []*WeatherCloud
}

func NewWeatherCloudFile() *WeatherCloudFile {
	return &WeatherCloudFile{
		records: make([]*WeatherCloud, 0),
	}
}

func (wcf *WeatherCloudFile) ProcessWeatherCloudFile(file *os.File, path string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 1 {
			continue
		}
		wcf.records = append(wcf.records, NewWeatherCloud(scanner.Text(), path))
	}
	wcf.records = wcf.records[1:]
}
