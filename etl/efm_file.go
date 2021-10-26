package etl

import (
	"bufio"
	"os"
	"strings"
)

type EfmFile struct {
	lines []*EfmRows
}

func NewEfmFile() *EfmFile {
	return &EfmFile{
		lines: make([]*EfmRows, 0),
	}
}

func (efm *EfmFile) ProcessFile(file *os.File, path string) {
	electricFields := make([]string, 0)
	time := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		if time == "" || time == fields[0] {
			electricFields = append(electricFields, fields[1])
			time = fields[0]
		} else {
			efm.lines = append(efm.lines, NewEfmRows(path, fields[0], electricFields, fields[2]))
			electricFields = make([]string, 0)
			electricFields = append(electricFields, fields[1])
			time = fields[0]
		}
	}
}
