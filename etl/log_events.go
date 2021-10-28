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

func (le *LogEvents) ProcessFile(file *os.File) {
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ThereIsLightning(scanner.Text()) {
			le.lines = append(le.lines, NewLogEventsRow(scanner.Text()))
		}
	}
}
