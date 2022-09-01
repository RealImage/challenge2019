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
	openSource     func(string) (io.ReadCloser, error)
}

type CsvRow struct {
	LineNumber int
	Value      []string
}

func NewCsvReaderConfig(sourceFilepath string, skipHeader bool) *CsvReaderConfig {
	return &CsvReaderConfig{
		sourceFilepath,
		skipHeader,
		openFile,
	}
}

// ReadLineFromCsv reads line from source csv file and send it to rowChan,
// if any error acquired, sends it to errChan and stops reading.
func (cfg *CsvReaderConfig) ReadLineFromCsv(rowChan chan<- *CsvRow, errChan chan<- error) {
	defer close(rowChan)
	defer close(errChan)

	f, err := cfg.openSource(cfg.SourceFilepath)
	if err != nil {
		errChan <- fmt.Errorf("can't open source file: {%s}; err: %w", cfg.SourceFilepath, err)
		return
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	lineCounter := 1

	if cfg.SkipHeader {
		//read 1st line to skip header
		_, err = csvReader.Read()
		if err != nil {
			errChan <- fmt.Errorf("source: {%s}; can't read header: %w", cfg.SourceFilepath, err)
			return
		}
		lineCounter++
	}

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			errChan <- fmt.Errorf("source: {%s}; line: %d; can't read data from partners: %w", cfg.SourceFilepath, lineCounter, err)
			return
		}

		rowChan <- &CsvRow{lineCounter, row}
		lineCounter++
	}

}

func openFile(name string) (io.ReadCloser, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	// f implements io.ReadCloser interface as *os.File
	// has Read and Close methods.
	return f, nil
}
