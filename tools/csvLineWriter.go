package tools

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvWriterConfig struct {
	DestinationFilepath string
	ErrChan             chan error
}

func NewCsvWriterConfig(destinationFilepath string, chanBufferSize int) *CsvWriterConfig {
	return &CsvWriterConfig{
		destinationFilepath,
		make(chan error, chanBufferSize),
	}
}

// WriteLineToCsv writes row to destination csv file.
// Row is received from chan []string, if any error acquired, send it to CsvWriterConfig.ErrChan and stops writing.
func (cfg *CsvWriterConfig) WriteLineToCsv(rowChan chan []string) {
	defer close(cfg.ErrChan)
	f, err := os.Create(cfg.DestinationFilepath)
	if err != nil {
		cfg.ErrChan <- fmt.Errorf("can't create destination file: {%s}; err: %s", cfg.DestinationFilepath, err)
		return
	}
	defer f.Close()

	csvWriter := csv.NewWriter(f)

	for row := range rowChan {
		err := csvWriter.Write(row)
		if err != nil {
			cfg.ErrChan <- fmt.Errorf("can't write row {%s} to csv: %s", row, err)
			return
		}

	}

	csvWriter.Flush()
}
