package services

import (
	"challange2019/pkg/models"
	"challange2019/tools"
)

type OutputServiceInterface interface {
	WriteToCsv(outChan <-chan *models.Output, errChan chan<- error)
}

type OutputService struct {
	tools.CsvWriter
}

func NewOutputService(csvWriter tools.CsvWriter) *OutputService {
	return &OutputService{csvWriter}
}

// WriteToCsv writes models.Output to csv
func (svc *OutputService) WriteToCsv(outChan <-chan *models.Output, errChan chan<- error) {
	defer close(errChan)

	rowChan := make(chan []string)
	go models.EncodeOutputToCsvRow(outChan, rowChan)

	wErrChan := make(chan error, ChanBufferSize)
	go svc.WriteLineToCsv(rowChan, wErrChan)
	for e := range wErrChan {
		errChan <- e
	}
}
