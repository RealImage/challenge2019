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

type Delivery struct {
	ID          string // unique delivery's id
	ContentSize int    // content size
	TheaterID   string // client theater
}

func (d *Delivery) String() string {
	return fmt.Sprintf("id: %s, tId: %s, ctx size: %d;", d.ID, d.TheaterID, d.ContentSize)
}

// ParseDeliveryFromCsvRow reads data from csv, parses it to Delivery instance and sends it to parsedDeliveryChan,
// if any error acquired, send it to errChan.
// if error acquired when reading from csv, stops method executing.
func ParseDeliveryFromCsvRow(inputRowChan <-chan *tools.CsvRow, parsedDeliveryChan chan<- *Delivery, errChan chan<- error) {
	defer close(parsedDeliveryChan)
	defer close(errChan)

	for row := range inputRowChan {
		d, err := parseDeliveryFromRow(row.Value)
		if err != nil {
			errChan <- fmt.Errorf("line: %d; can't parse deliveries data: %s", row.LineNumber, err)
			continue
		}

		parsedDeliveryChan <- d
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
