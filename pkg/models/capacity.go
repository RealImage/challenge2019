package models

import (
	"challange2019/tools"
	"fmt"
	"strconv"
	"strings"
)

const (
	PartnerIdCapacityColumnIndex = iota
	CapacityColumnIndex
	CapacityDataColumnCount
)

type Capacity struct {
	PartnerId string
	Value     int
}

// ReadCapacityFromCsv reads data from csv, parses it to Capacity instance and sends it to parsedCapacityChan,
// if any error acquired, send it to errChan.
// if error acquired when reading from csv, stops method executing.
func ReadCapacityFromCsv(inputRowChan <-chan *tools.CsvRow, parsedCapacityChan chan<- *Capacity, errChan chan<- error) {
	defer close(parsedCapacityChan)
	defer close(errChan)

	for row := range inputRowChan {
		capacity, err := parseCapacityFromRow(row.Value)
		if err != nil {
			errChan <- fmt.Errorf("line: %d; can't parse capacity data: %s", row.LineNumber, err)
			continue
		}

		parsedCapacityChan <- capacity
	}
}

func parseCapacityFromRow(row []string) (*Capacity, error) {
	// expected row format:
	// Partner ID (e.g. P1), Capacity (e.g. 100500)
	c := &Capacity{}

	if len(row) != CapacityDataColumnCount {
		return nil, fmt.Errorf("capacity data is corrupted: %v", row)
	}

	var err error
	c.PartnerId = strings.TrimSpace(row[PartnerIdCapacityColumnIndex])

	c.Value, err = strconv.Atoi(strings.TrimSpace(row[CapacityColumnIndex]))
	if err != nil {
		return nil, fmt.Errorf("can't parse capacity value: %s", err)
	}

	return c, err

}
