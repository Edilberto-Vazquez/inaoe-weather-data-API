package etl

import (
	"bufio"
	"os"
	"strings"
)

type ElectricField struct {
	lines []*ElectricFieldRow
}

func NewEfmFile() *ElectricField {
	return &ElectricField{
		lines: make([]*ElectricFieldRow, 0),
	}
}

func (ef *ElectricField) ProcessFile(path string, file *os.File) {
	defer file.Close()
	electricFields := make([]string, 0)
	time := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		if time == "" || time == fields[0] {
			electricFields = append(electricFields, fields[1])
			time = fields[0]
		} else {
			ef.lines = append(ef.lines, NewElectricFieldRow(path, fields[0], electricFields, fields[2]))
			electricFields = make([]string, 0)
			electricFields = append(electricFields, fields[1])
			time = fields[0]
		}
	}
}
