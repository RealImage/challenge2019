package models

import (
	"challange2019/tools"
	"fmt"
	"strconv"
	"strings"
)

const (
	TheaterIdPartnersColumnIndex = iota
	SizeSlabColumnIndex
	MinimumCostColumnIndex
	CostPerGbColumnIndex
	PartnerIdColumnIndex
	PartnersDataColumnCount
)

type PartnerParserConfig struct {
	CsvCfg         *tools.CsvReaderConfig // config to read csv file
	ParsedDataChan chan *Partner          // chan where the parsed Partner instance is send
	ErrChan        chan error             // chan were errors, if acquired, are sent
}

type Partner struct {
	ID        string // is not unique field, it is possible few instances may be with same id, but different theater, cost, etc ...
	TheaterID string // theater the partner is work with
	CostPerGb int    // cost per Gb of traffic
	MinCost   int    // minimum costs for the transportation of the relevant content
	MinSlabGb int    // minimum content size to transport
	MaxSlabGb int    // maximum content size to transport
}

func (p *Partner) String() string {
	return fmt.Sprintf("id: %s, tId: %s, min cost: %d ,slab: %d-%d; ",
		p.ID, p.TheaterID, p.MinCost, p.MinSlabGb, p.MaxSlabGb)
}

func NewPartnerParserConfig(csvCfg *tools.CsvReaderConfig, chanBufferSize int) *PartnerParserConfig {
	return &PartnerParserConfig{
		csvCfg,
		make(chan *Partner, chanBufferSize),
		make(chan error, chanBufferSize),
	}
}

// CalculateCost calculates the cost of the content transporting for given partner
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

// ReadPartnerFromCsv reads data from csv, parses it to Partner instance and sends it to PartnerParserConfig.ParsedDataChan,
// if any error acquired, send it to PartnerParserConfig.ErrChan.
// if error acquired when reading from csv, stops method executing.
func (pp *PartnerParserConfig) ReadPartnerFromCsv() {
	defer close(pp.ParsedDataChan)
	defer close(pp.ErrChan)

	go func() {
		go pp.CsvCfg.ReadLineFromCsv()
		for err := range pp.CsvCfg.ErrChan {
			pp.ErrChan <- err
		}
	}()

	for row := range pp.CsvCfg.RowChan {
		p, err := parsePartnerFromRow(row.Value)
		if err != nil {
			pp.ErrChan <- fmt.Errorf("line: %d; can't parse partner data: %s", row.LineNumber, err)
			continue
		}

		pp.ParsedDataChan <- p
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
