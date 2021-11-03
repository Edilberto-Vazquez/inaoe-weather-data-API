package etl

import (
	"bufio"
)

type LogEventsRecords struct {
	records []interface{}
}

func NewLogEventsRecords() *LogEventsRecords {
	return &LogEventsRecords{
		records: make([]interface{}, 0),
	}
}

func (le *LogEventsRecords) ProcessFile(path string) {
	file := OpenFile(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ThereIsLightning(scanner.Text()) {
			le.records = append(le.records, NewLogEvents(scanner.Text()))
		}
	}
}

func (le *LogEventsRecords) GetRecords() []interface{} {
	return le.records
}
