package models

import (
	"challange2019/tools"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type DeliveryParserConfig struct {
	CsvCfg         *tools.CsvReaderConfig
	ParsedDataChan chan *DeliveryInput
	ErrChan        chan error
}

type DeliveryInput struct {
	ID          string
	ContentSize int
	TheaterID   string
}

type DeliveryInputList struct {
	sync.RWMutex
	Container []*DeliveryInput
}

type DeliveryOutput struct {
	ID         string
	PartnerID  string
	isPossible bool
	Cost       int
}

type DeliveryOutputList struct {
	sync.RWMutex
	Container []*DeliveryOutput
}

const (
	DeliveryIdColumnIndex = iota
	ContentSizeColumnIndex
	TheaterIdDeliveryColumnIndex
	DeliveryDataColumnCount
)

func (dil *DeliveryInputList) Add(di *DeliveryInput) {
	dil.Lock()
	dil.Container = append(dil.Container, di)
	dil.Unlock()
}

func (dol *DeliveryOutputList) Add(di *DeliveryOutput) {
	dol.Lock()
	dol.Container = append(dol.Container, di)
	dol.Unlock()
}

func FindCheapestDeliveryOutput(d *DeliveryInput, partners []*Partner, output *DeliveryOutputList) {
	do := &DeliveryOutput{ID: d.ID, Cost: -1}

	for _, p := range partners {
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

func NewDeliveryParserConfig(csvCfg *tools.CsvReaderConfig, chanBufferSize int) *DeliveryParserConfig {
	return &DeliveryParserConfig{
		csvCfg,
		make(chan *DeliveryInput, chanBufferSize),
		make(chan error, chanBufferSize),
	}
}

func (dp *DeliveryParserConfig) CloseChannels() {
	close(dp.ErrChan)
	close(dp.ParsedDataChan)
}
func (dp *DeliveryParserConfig) ReadDeliveriesInputFromCsv() {
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
