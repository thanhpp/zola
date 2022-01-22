package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
		return nil, err
	}
	defer f.Close()
	var record [][]string
	scanner := bufio.NewScanner(f)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		if string(line[0]) == "#" {
			continue
		} else {
			record = append(record, strings.Split(line, ", "))
		}
	}

	return record, nil
}
