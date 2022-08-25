package models

import (
	"challange2019/tools"
	"fmt"
	"strconv"
	"strings"
)

const (
	DeliveryIdColumnIndex = iota
	ContentSizeColumnIndex
	TheaterIdDeliveryColumnIndex
	DeliveryDataColumnCount
)

type DeliveryParserConfig struct {
	CsvCfg         *tools.CsvReaderConfig // config to read csv file
	ParsedDataChan chan *Delivery         // chan where the parsed Delivery instance is send
	ErrChan        chan error             // chan were errors, if acquired, are sent
}

type Delivery struct {
	ID          string // unique delivery's id
	ContentSize int    // content size
	TheaterID   string // client theater
}

func (d *Delivery) String() string {
	return fmt.Sprintf("id: %s, tId: %s, ctx size: %d;", d.ID, d.TheaterID, d.ContentSize)
}

func NewDeliveryParserConfig(csvCfg *tools.CsvReaderConfig, chanBufferSize int) *DeliveryParserConfig {
	return &DeliveryParserConfig{
		csvCfg,
		make(chan *Delivery, chanBufferSize),
		make(chan error, chanBufferSize),
	}
}

func (dp *DeliveryParserConfig) CloseChannels() {
	close(dp.ErrChan)
	close(dp.ParsedDataChan)
}

// ReadDeliveriesFromCsv reads data from csv, parses it to Delivery instance and sends it to Delivery.ParsedDataChan,
// if any error acquired, send it to Delivery.ErrChan.
// if error acquired when reading from csv, stops method executing.
func (dp *DeliveryParserConfig) ReadDeliveriesFromCsv() {
	defer dp.CloseChannels()

	go func() {
		go dp.CsvCfg.ReadLineFromCsv()
		for err := range dp.CsvCfg.ErrChan {
			dp.ErrChan <- err
		}
	}()

	for row := range dp.CsvCfg.RowChan {
		d, err := parseDeliveryFromRow(row.Value)
		if err != nil {
			dp.ErrChan <- fmt.Errorf("line: %d; can't parse deliveries data: %s", row.LineNumber, err)
			continue
		}

		dp.ParsedDataChan <- d
	}
}

func parseDeliveryFromRow(row []string) (*Delivery, error) {
	// expected row format:
	// Delivery ID (e.g. D1), Content Size (e.g. 100500), Theater ID (e.g. T1)
	d := &Delivery{}

	if len(row) != DeliveryDataColumnCount {
		return nil, fmt.Errorf("deliveries data is corrupted: %v", row)
	}

	var err error
	d.ID = strings.TrimSpace(row[DeliveryIdColumnIndex])
	d.TheaterID = strings.TrimSpace(row[TheaterIdDeliveryColumnIndex])

	d.ContentSize, err = strconv.Atoi(strings.TrimSpace(row[ContentSizeColumnIndex]))
	if err != nil {
		return nil, fmt.Errorf("can't parse content size: %s", err)
	}

	return d, err
}
