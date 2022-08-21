package models

import (
	"challange2019/tools"
	"fmt"
	"strconv"
	"strings"
)

type PartnerParserConfig struct {
	CsvData        chan *tools.CsvRow
	ParsedDataChan chan *Partner
	ErrChan        chan error
}

type Partner struct {
	ID        string
	TheaterID string
	CostPerGb int
	MinCost   int
	MinSlabGb int
	MaxSlabGb int
	Capacity  int
}

const (
	TheaterIdPartnersColumnIndex = iota
	SizeSlabColumnIndex
	MinimumCostColumnIndex
	CostPerGbColumnIndex
	PartnerIdColumnIndex
	PartnersDataColumnCount
)

func NewPartnerParserConfig(csvData chan *tools.CsvRow, parsedDataChan chan *Partner, errChan chan error) *PartnerParserConfig {
	return &PartnerParserConfig{csvData, parsedDataChan, errChan}
}

func (p *Partner) CalculateCost(contentSize int) (int, bool) {
	if contentSize < p.MinSlabGb || contentSize > p.MaxSlabGb {
		return -1, false
	}

	actualCost := contentSize * p.CostPerGb
	if actualCost < p.MinCost {
		return p.MinCost, true
	}

	return actualCost, true
}

func (p *Partner) ReadCapacity(capacityPath string, logger tools.SomeLogger) {

}

func (cfg *PartnerParserConfig) ParsePartnerByTheaterFromCsv(theaterId string) {
	defer close(cfg.ParsedDataChan)
	defer close(cfg.ErrChan)

	for row := range cfg.CsvData {
		p, err := parsePartnerFromRow(row.Value)
		if err != nil {
			cfg.ErrChan <- fmt.Errorf("line: %d; can't parse partner data: %s", row.LineNumber, err)
			continue
		}

		if p.TheaterID == theaterId {
			cfg.ParsedDataChan <- p
		}
	}
}

func parsePartnerFromRow(row []string) (*Partner, error) {
	// expected row format:
	// Theatre ID (e.g. T1), Size Slab (e.g. 0-200), Minimum cost (e.g. 100500), Cost Per GB (e.g. 20), Partner ID (e.g. P1)
	p := &Partner{}

	if len(row) != PartnersDataColumnCount {
		return nil, fmt.Errorf("parnters data is corrupted: %v", row)
	}

	var err error
	p.ID = strings.TrimSpace(row[PartnerIdColumnIndex])
	p.TheaterID = strings.TrimSpace(row[TheaterIdPartnersColumnIndex])

	p.CostPerGb, err = strconv.Atoi(strings.TrimSpace(row[CostPerGbColumnIndex]))
	if err != nil {
		return nil, fmt.Errorf("can't parse cost per Gb: %s", err)
	}

	p.MinCost, err = strconv.Atoi(strings.TrimSpace(row[MinimumCostColumnIndex]))
	if err != nil {
		return nil, fmt.Errorf("can't parse minimum cost: %s", err)
	}

	sizeSlabValues := strings.Split(row[SizeSlabColumnIndex], "-")
	if len(sizeSlabValues) != 2 {
		return nil, fmt.Errorf("size slab values corrupted")
	}

	p.MinSlabGb, err = strconv.Atoi(strings.TrimSpace(sizeSlabValues[0]))
	if err != nil {
		return nil, fmt.Errorf("can't parse minimum slab value: %s", err)
	}

	p.MaxSlabGb, err = strconv.Atoi(strings.TrimSpace(sizeSlabValues[1]))
	if err != nil {
		return nil, fmt.Errorf("can't parse maximum slab value: %s", err)
	}

	return p, nil
}
