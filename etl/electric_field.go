package etl

import (
	"bufio"
	"strings"
)

type ElectricField struct {
	rocords []interface{}
}

func NewElectricFields() *ElectricField {
	return &ElectricField{
		rocords: make([]interface{}, 0),
	}
}

func (ef *ElectricField) ProcessFile(path string) {
	file := OpenFile(path)
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
			ef.rocords = append(ef.rocords, NewElectricFieldRow(path, fields[0], electricFields, fields[2]))
			electricFields = make([]string, 0)
			electricFields = append(electricFields, fields[1])
			time = fields[0]
		}
	}
}

func (ef *ElectricField) GetRecords() []interface{} {
	return ef.rocords
}
