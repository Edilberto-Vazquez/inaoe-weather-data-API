package etl

import (
	"bufio"
	"os"
)

type EfmEvents struct {
	lines []*EfmEventsRow
}

func NewEfmEvents() *EfmEvents {
	return &EfmEvents{
		lines: make([]*EfmEventsRow, 0),
	}
}

func (events *EfmEvents) ProcessEfmEvents(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ThereIsLightning(scanner.Text()) {
			events.lines = append(events.lines, NewEfmEventsRow(scanner.Text()))
		}
	}
}
