package etl

import (
	"bufio"
	"strings"
)

type ElectricFieldRecords struct {
	rocords []interface{}
}

func NewElectricFieldRecords() *ElectricFieldRecords {
	return &ElectricFieldRecords{
		rocords: make([]interface{}, 0),
	}
}

func (ef *ElectricFieldRecords) ProcessFile(path string) {
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
			ef.rocords = append(ef.rocords, NewElectricField(path, fields[0], electricFields, fields[2]))
			electricFields = make([]string, 0)
			electricFields = append(electricFields, fields[1])
			time = fields[0]
		}
	}
}

func (ef *ElectricFieldRecords) GetRecords() []interface{} {
	return ef.rocords
}
