package etl

import (
	"strings"
	"time"

	"github.com/Edilberto-Vazquez/inaoe-weather-data-API/db/models"
)

func ProcessElectricFieldMonitorFile(path string) (records map[time.Time]models.ElectricField) {
	records = make(map[time.Time]models.ElectricField)
	file, scanner := OpenFile(path)
	defer file.Close()
	var electricFields []string = make([]string, 0)
	var time string
	var rotorStatus string
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		if time == "" || time == fields[0] {
			electricFields = append(electricFields, fields[1])
			time = fields[0]
			rotorStatus = fields[2]
		} else {
			record := models.ElectricField{
				TimeStamp:     NewTimeStamp("efm", path+" "+time),
				ElectricField: ElectricFieldAvg(electricFields),
				RotorStatus:   NewRotorStatus(rotorStatus),
			}
			records[record.TimeStamp] = record
			electricFields = make([]string, 0)
			electricFields = append(electricFields, fields[1])
			time = fields[0]
		}
	}
	return
}
