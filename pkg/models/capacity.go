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

type CapacityParserConfig struct {
	CsvCfg         *tools.CsvReaderConfig
	ParsedDataChan chan *Capacity
	ErrChan        chan error
}

type Capacity struct {
	PartnerId string
	Value     int
}

func NewCapacityParserConfig(csvCfg *tools.CsvReaderConfig, chanBufferSize int) *CapacityParserConfig {
	return &CapacityParserConfig{
		csvCfg,
		make(chan *Capacity, chanBufferSize),
		make(chan error, chanBufferSize),
	}
}

func (c *CapacityParserConfig) ReadCapacityFromCsv() {
	defer close(c.ParsedDataChan)
	defer close(c.ErrChan)

	go func() {
		go c.CsvCfg.ReadLineFromCsv()
		for err := range c.CsvCfg.ErrChan {
			c.ErrChan <- err
		}
	}()

	for row := range c.CsvCfg.RowChan {
		capacity, err := parseCapacityFromRow(row.Value)
		if err != nil {
			c.ErrChan <- fmt.Errorf("line: %d; can't parse capacity data: %s", row.LineNumber, err)
			continue
		}

		c.ParsedDataChan <- capacity
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
