package etl

import (
	"bufio"
	"os"
)

type LogEvents struct {
	lines []*LogEventsRow
}

func NewEfmEvents() *LogEvents {
	return &LogEvents{
		lines: make([]*LogEventsRow, 0),
	}
}

func (events *LogEvents) ProcessEfmEvents(file *os.File) {
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ThereIsLightning(scanner.Text()) {
			events.lines = append(events.lines, NewLogEventsRow(scanner.Text()))
		}
	}
}
