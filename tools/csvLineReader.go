package tools

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type CsvReader interface {
	ReadLineFromCsv(rowChan chan<- *CsvRow, errChan chan<- error)
}

type CsvReaderConfig struct {
	SourceFilepath string
	SkipHeader     bool
}

type CsvRow struct {
	LineNumber int
	Value      []string
}

func NewCsvReaderConfig(sourceFilepath string, skipHeader bool) *CsvReaderConfig {
	return &CsvReaderConfig{
		sourceFilepath,
		skipHeader,
	}
}

// ReadLineFromCsv reads line from source csv file and send it to CsvReaderConfig.RowChan,
// if any error acquired, sends it to CsvReaderConfig.ErrChan and stops reading.
func (cfg *CsvReaderConfig) ReadLineFromCsv(rowChan chan<- *CsvRow, errChan chan<- error) {
	defer close(rowChan)
	defer close(errChan)

	f, err := os.Open(cfg.SourceFilepath)
	if err != nil {
		errChan <- fmt.Errorf("can't open source file: {%s}; err: %s", cfg.SourceFilepath, err)
		return
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	if cfg.SkipHeader {
		//read 1st line to skip header
		_, err = csvReader.Read()
		if err != nil {
			errChan <- fmt.Errorf(fmt.Sprintf("source: {%s}; can't read header: %s", cfg.SourceFilepath, err))
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
			errChan <- fmt.Errorf(fmt.Sprintf("source: {%s}; line: %d; can't read data from partners: %s", cfg.SourceFilepath, lineCounter, err))
			return
		}

		rowChan <- &CsvRow{lineCounter, row}
		lineCounter++
	}

}
