package tools

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type CsvReaderConfig struct {
	SourceFilepath string
	SkipHeader     bool
	RowChan        chan *CsvRow
	ErrChan        chan error
}

type CsvRow struct {
	LineNumber int
	Value      []string
}

func NewCsvReaderConfig(sourceFilepath string, skipHeader bool, chanBufferSize int) *CsvReaderConfig {
	return &CsvReaderConfig{
		sourceFilepath,
		skipHeader,
		make(chan *CsvRow, chanBufferSize),
		make(chan error, chanBufferSize),
	}
}

func (cfg *CsvReaderConfig) CloseChannels() {
	close(cfg.RowChan)
	close(cfg.ErrChan)
}

func (cfg *CsvReaderConfig) ReadLineFromCsv() {
	defer cfg.CloseChannels()
	f, err := os.Open(cfg.SourceFilepath)
	if err != nil {
		cfg.ErrChan <- fmt.Errorf("source: {%s}; can't open deliveries data: {%s}", cfg.SourceFilepath, err)
		return
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	if cfg.SkipHeader {
		//read 1st line to skip header
		_, err = csvReader.Read()
		if err != nil {
			cfg.ErrChan <- fmt.Errorf(fmt.Sprintf("source: {%s}; can't read header: %s", cfg.SourceFilepath, err))
			return
		}
	}

	lineCounter := 1
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			cfg.ErrChan <- fmt.Errorf(fmt.Sprintf("source: {%s}; line: %d; can't read data from partners: %s", cfg.SourceFilepath, lineCounter, err))
			return
		}

		cfg.RowChan <- &CsvRow{lineCounter, row}
		lineCounter++
	}

}
