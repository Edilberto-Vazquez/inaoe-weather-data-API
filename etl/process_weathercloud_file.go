package etl

import (
	"encoding/binary"
	"strings"
	"time"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/db/models"
)

func ProcessWeathercloudFile(path string) (records map[time.Time]models.Weathercloud) {
	records = make(map[time.Time]models.Weathercloud, 0)
	file, scanner := OpenFile(path)
	defer file.Close()
	i := 0
	for scanner.Scan() {
		utf16, _ := DecodeUtf16(scanner.Bytes(), binary.BigEndian)
		fields := strings.Split(utf16, ";")
		if len(scanner.Text()) == 1 {
			continue
		}
		if i == 0 {
			i++
			continue
		}
		record := models.Weathercloud{
			TimeStamp: NewTimeStamp("wc", fields[0]),
			TempIn:    CommaToPoint(fields[1]),
			Temp:      CommaToPoint(fields[2]),
			Chill:     CommaToPoint(fields[3]),
			DewIn:     CommaToPoint(fields[4]),
			Dew:       CommaToPoint(fields[5]),
			HeatIn:    CommaToPoint(fields[6]),
			Heat:      CommaToPoint(fields[7]),
			Humin:     CommaToPoint(fields[8]),
			Hum:       CommaToPoint(fields[9]),
			Wspdhi:    CommaToPoint(fields[10]),
			Wspdavg:   CommaToPoint(fields[11]),
			Wdiravg:   CommaToPoint(fields[12]),
			Bar:       CommaToPoint(fields[13]),
			Rain:      CommaToPoint(fields[14]),
			RainRate:  CommaToPoint(fields[15]),
		}
		records[record.TimeStamp] = record
	}
	return
}
