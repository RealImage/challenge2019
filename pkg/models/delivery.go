package models

import (
	"challange2019/tools"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type DeliveryParserConfig struct {
	CsvData        chan *tools.CsvRow
	ParsedDataChan chan *DeliveryInput
	ErrChan        chan error
}

type DeliveryInput struct {
	ID          string
	ContentSize int
	TheaterID   string
}

type DeliveryOutput struct {
	ID         string
	PartnerID  string
	isPossible bool
	Cost       int
}

type DeliveryOutputStorage struct {
	Container []*DeliveryOutput
	sync.RWMutex
}

const (
	DeliveryIdColumnIndex = iota
	ContentSizeColumnIndex
	TheaterIdDeliveryColumnIndex
	DeliveryDataColumnCount
)

func NewDeliveryOutputStorage() *DeliveryOutputStorage {
	return &DeliveryOutputStorage{Container: []*DeliveryOutput{}}
}

func FindCheapestDeliveryOutput(d *DeliveryInput, partnerCh <-chan *Partner, output *DeliveryOutputStorage) {
	do := &DeliveryOutput{ID: d.ID, Cost: -1}

	for p := range partnerCh {
		cost, isPossible := p.CalculateCost(d.ContentSize)

		if isPossible && (cost < do.Cost || do.Cost < 0) {
			do.Cost = cost
			do.PartnerID = p.ID
			do.isPossible = isPossible
		}
	}

	output.Lock()
	output.Container = append(output.Container, do)
	output.Unlock()
}

func (do *DeliveryOutput) String() string {
	cost := strconv.Itoa(do.Cost)
	if do.Cost < 0 {
		cost = ""
	}
	return fmt.Sprintf("%s, %t, %s, %s", do.ID, do.isPossible, do.PartnerID, cost)
}

func NewDeliveryParserConfig(csvData chan *tools.CsvRow, parsedDataChan chan *DeliveryInput, errChan chan error) *DeliveryParserConfig {
	return &DeliveryParserConfig{csvData, parsedDataChan, errChan}
}

func (cfg *DeliveryParserConfig) ParseDeliveriesInputCsv() {
	defer close(cfg.ParsedDataChan)
	defer close(cfg.ErrChan)

	for row := range cfg.CsvData {
		d, err := parseDeliveryFromRow(row.Value)
		if err != nil {
			cfg.ErrChan <- fmt.Errorf("line: %d; can't parse deliveries data: %s", row.LineNumber, err)
			continue
		}

		cfg.ParsedDataChan <- d
	}
}

func parseDeliveryFromRow(row []string) (*DeliveryInput, error) {
	// expected row format:
	// Delivery ID (e.g. D1), Content Size (e.g. 100500), Theater ID (e.g. T1)
	d := &DeliveryInput{}

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
