package services

import (
	"challange2019/pkg/models"
	"challange2019/tools"
)

type OutputService struct {
	DestinationFilepath string
	ErrChan             chan error
}

func NewOutputService(destinationFile string) *OutputService {
	return &OutputService{destinationFile, make(chan error, ChanBufferSize)}
}

// WriteToCsv writes models.Output to csv
func (svc *OutputService) WriteToCsv(outChan <-chan *models.Output) {
	defer close(svc.ErrChan)

	rowChan := make(chan []string)
	go models.EncodeOutputToCsvRow(outChan, rowChan)

	csvCfg := tools.NewCsvWriterConfig(svc.DestinationFilepath, ChanBufferSize)
	go csvCfg.WriteLineToCsv(rowChan)

	for err := range csvCfg.ErrChan {
		svc.ErrChan <- err
	}
}
