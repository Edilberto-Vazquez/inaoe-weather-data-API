package etl

import (
	"time"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/db/models"
	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/utils"
)

var (
	LogRecords           = make(map[time.Time]models.Log)
	ElectricFieldRecords = make(map[time.Time]models.ElectricField)
	WeathercloudRecords  = make(map[time.Time]models.Weathercloud)
	RecordHashTable      = make(map[time.Time]*utils.Singlylinkedlist)
	RecordSliceTable     = make([]*models.ElectricFieldWeatherCloud, 0)
)

func CreateRecordHashTable() {
	for wcTimeStamp, wcRecord := range WeathercloudRecords {
		if RecordHashTable[wcTimeStamp] == nil {
			RecordHashTable[wcTimeStamp] = &utils.Singlylinkedlist{}
		}
		RecordHashTable[wcTimeStamp].Append(models.ElectricFieldWeatherCloud{
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
	for logTimeStamp, logRecord := range LogRecords {
		roundDate := logTimeStamp.Round(10 * time.Minute)
		efRecord, efok := ElectricFieldRecords[logTimeStamp]
		htRecord, htok := RecordHashTable[roundDate]
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
	LogRecords = make(map[time.Time]models.Log, 0)
	ElectricFieldRecords = make(map[time.Time]models.ElectricField, 0)
	WeathercloudRecords = make(map[time.Time]models.Weathercloud, 0)
}

func CreateRecordSliceTable() {
	for _, key := range RecordHashTable {
		currentNode := key.Head
		for i := 0; i < key.Length(); i++ {
			RecordSliceTable = append(RecordSliceTable, &currentNode.Value)
			currentNode = currentNode.Next
		}
	}
	RecordHashTable = make(map[time.Time]*utils.Singlylinkedlist, 0)
}

func ProcessFiles(logRooot, efmRoot string, wcRoot string) {
	logFiles, _ := ReadDirectory(logRooot, "log")
	efmFiles, _ := ReadDirectory(efmRoot, "efm")
	wcFiles, _ := ReadDirectory(wcRoot, "csv")

	// log File
	clog := make(chan map[time.Time]models.Log)
	go func(path string, c chan<- map[time.Time]models.Log) {
		record := ProcessLogFile(path)
		c <- record
	}(logFiles[0], clog)
	for _, v := range <-clog {
		LogRecords[v.TimeStamp] = v
	}
	close(clog)

	// Electric Field Monitor Files
	cefm := make(chan map[time.Time]models.ElectricField, len(efmFiles))
	for _, efmFile := range efmFiles {
		go func(path string, c chan<- map[time.Time]models.ElectricField) {
			record := ProcessElectricFieldMonitorFile(path)
			c <- record
		}(efmFile, cefm)
	}
	for i := 0; i < len(efmFiles); i++ {
		for _, v := range <-cefm {
			ElectricFieldRecords[v.TimeStamp] = v
		}
	}
	close(cefm)

	// Weathercloud files
	cwc := make(chan map[time.Time]models.Weathercloud, len(wcFiles))
	for _, wcFile := range wcFiles {
		go func(path string, c chan<- map[time.Time]models.Weathercloud) {
			record := ProcessWeathercloudFile(path)
			c <- record
		}(wcFile, cwc)
	}
	for i := 0; i < len(wcFiles); i++ {
		for _, v := range <-cwc {
			WeathercloudRecords[v.TimeStamp] = v
		}
	}
	close(cwc)
	// create record hash table
	CreateRecordHashTable()

	// create slice table
	CreateRecordSliceTable()
}
