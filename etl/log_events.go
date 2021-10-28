package etl

import (
	"bufio"
)

type LogEvents struct {
	records []interface{}
}

func NewLogEvents() *LogEvents {
	return &LogEvents{
		records: make([]interface{}, 0),
	}
}

func (le *LogEvents) ProcessFile(path string) {
	file := OpenFile(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ThereIsLightning(scanner.Text()) {
			le.records = append(le.records, NewLogEventsRow(scanner.Text()))
		}
	}
}

func (le *LogEvents) GetRecords() []interface{} {
	return le.records
}
