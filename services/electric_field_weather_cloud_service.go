package services

import (
	"bufio"
	"encoding/binary"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/db/models"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/etl"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/utils"
)

type ElectricFieldWeatherCloudService struct {
	RecordHashTable      map[time.Time]*utils.Singlylinkedlist
	LogRecords           map[time.Time]*models.Log
	ElectricFieldRecords map[time.Time]models.ElectricField
	WeatherCloudRecords  map[time.Time]*models.Weathercloud
	Mux                  sync.RWMutex
}

func NewElectricFieldWeatherCloudService() *ElectricFieldWeatherCloudService {
	return &ElectricFieldWeatherCloudService{
		RecordHashTable:      make(map[time.Time]*utils.Singlylinkedlist),
		LogRecords:           make(map[time.Time]*models.Log),
		ElectricFieldRecords: make(map[time.Time]models.ElectricField),
		WeatherCloudRecords:  make(map[time.Time]*models.Weathercloud),
	}
}

func ReadLines(path string) (records map[time.Time]models.ElectricField) {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	records = make(map[time.Time]models.ElectricField, 0)
	electricFields := make([]string, 0)
	var time string
	var rotorStatus string
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		if time == "" || time == fields[0] {
			electricFields = append(electricFields, fields[1]) // pasar a ReadLines
			time = fields[0]
			rotorStatus = fields[2]
		} else {
			record := models.ElectricField{
				TimeStamp:     etl.NewTimeStamp("efm", path+" "+time),
				ElectricField: etl.ElectricFieldAvg(electricFields),
				RotorStatus:   etl.NewRotorStatus(rotorStatus),
			}
			records[record.TimeStamp] = record
			electricFields = make([]string, 0)
			electricFields = append(electricFields, fields[1])
			time = fields[0]
		}
	}
	return
}

func ProcessLogFile(path string) (records map[time.Time]*models.Log) {
	records = make(map[time.Time]*models.Log, 0)
	file, scanner := etl.OpenFile(path)
	defer file.Close()
	for scanner.Scan() {
		if etl.ThereIsLightning(scanner.Text()) {
			record := models.Log{
				TimeStamp: etl.NewTimeStamp("log", scanner.Text()),
				Lightning: etl.NewLightning(),
				Distance:  etl.NewDistance(scanner.Text()),
			}
			records[record.TimeStamp] = &record
		}
	}
	return
}

func (s *ElectricFieldWeatherCloudService) ProcessElectricFieldFile(path string) {
	for _, line := range ReadLines(path) {
		s.Mux.Lock()
		s.ElectricFieldRecords[line.TimeStamp] = line
		s.Mux.Unlock()
	}
}

func ProcessWeatherCloudFile(path string) (records map[time.Time]*models.Weathercloud) {
	records = make(map[time.Time]*models.Weathercloud, 0)
	file, scanner := etl.OpenFile(path)
	defer file.Close()
	i := 0
	for scanner.Scan() {
		utf16, _ := etl.DecodeUtf16(scanner.Bytes(), binary.BigEndian)
		fields := strings.Split(utf16, ";")
		if len(scanner.Text()) == 1 {
			continue
		}
		if i == 0 {
			i++
			continue
		}
		record := models.Weathercloud{
			TimeStamp: etl.NewTimeStamp("wc", fields[0]),
			TempIn:    etl.CommaToPoint(fields[1]),
			Temp:      etl.CommaToPoint(fields[2]),
			Chill:     etl.CommaToPoint(fields[3]),
			DewIn:     etl.CommaToPoint(fields[4]),
			Dew:       etl.CommaToPoint(fields[5]),
			HeatIn:    etl.CommaToPoint(fields[6]),
			Heat:      etl.CommaToPoint(fields[7]),
			Humin:     etl.CommaToPoint(fields[8]),
			Hum:       etl.CommaToPoint(fields[9]),
			Wspdhi:    etl.CommaToPoint(fields[10]),
			Wspdavg:   etl.CommaToPoint(fields[11]),
			Wdiravg:   etl.CommaToPoint(fields[12]),
			Bar:       etl.CommaToPoint(fields[13]),
			Rain:      etl.CommaToPoint(fields[14]),
			RainRate:  etl.CommaToPoint(fields[15]),
		}
		records[record.TimeStamp] = &record
	}
	return
}

func (s *ElectricFieldWeatherCloudService) ProcessNewFiles(logPath string, efPath string, wcPath string) {
	for wcTimeStamp, wcRecord := range s.WeatherCloudRecords {
		if s.RecordHashTable[wcTimeStamp] == nil {
			s.RecordHashTable[wcTimeStamp] = &utils.Singlylinkedlist{}
		}
		s.RecordHashTable[wcTimeStamp].Append(models.ElectricFieldWeatherCloud{
			TimeStamp: wcTimeStamp,
			TempIn:    wcRecord.TempIn,
			Temp:      wcRecord.Temp,
			Chill:     wcRecord.Chill,
			DewIn:     wcRecord.DewIn,
			Dew:       wcRecord.Dew,
			HeatIn:    wcRecord.HeatIn,
			Heat:      wcRecord.Heat,
			Humin:     wcRecord.Humin,
			Hum:       wcRecord.Hum,
			Wspdhi:    wcRecord.Wspdhi,
			Wspdavg:   wcRecord.Wspdavg,
			Wdiravg:   wcRecord.Wdiravg,
			Bar:       wcRecord.Bar,
			Rain:      wcRecord.Rain,
			RainRate:  wcRecord.RainRate,
		})
	}
	for logTimeStamp, logRecord := range s.LogRecords {
		roundDate := logTimeStamp.Round(10 * time.Minute)
		efRecord, efok := s.ElectricFieldRecords[logTimeStamp]
		htRecord, htok := s.RecordHashTable[roundDate]
		if efok && htok {
			if !htRecord.Head.Value.Lightning {
				htRecord.Head.Value.Lightning = logRecord.Lightning
				htRecord.Head.Value.Distance = logRecord.Distance
				htRecord.Head.Value.ElectricField = efRecord.ElectricField
				htRecord.Head.Value.RotorStatus = efRecord.RotorStatus
			} else {
				htRecord.Append(models.ElectricFieldWeatherCloud{
					TimeStamp:     roundDate,
					Lightning:     logRecord.Lightning,
					Distance:      logRecord.Distance,
					ElectricField: efRecord.ElectricField,
					RotorStatus:   efRecord.RotorStatus,
					TempIn:        htRecord.Head.Value.TempIn,
					Temp:          htRecord.Head.Value.Temp,
					Chill:         htRecord.Head.Value.Chill,
					DewIn:         htRecord.Head.Value.DewIn,
					Dew:           htRecord.Head.Value.Dew,
					HeatIn:        htRecord.Head.Value.HeatIn,
					Heat:          htRecord.Head.Value.Heat,
					Humin:         htRecord.Head.Value.Humin,
					Hum:           htRecord.Head.Value.Hum,
					Wspdhi:        htRecord.Head.Value.Wspdhi,
					Wspdavg:       htRecord.Head.Value.Wspdavg,
					Wdiravg:       htRecord.Head.Value.Wdiravg,
					Bar:           htRecord.Head.Value.Bar,
					Rain:          htRecord.Head.Value.Rain,
					RainRate:      htRecord.Head.Value.RainRate,
				})
			}
		}
	}
}
