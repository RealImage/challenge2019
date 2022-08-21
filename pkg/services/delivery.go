package services

import (
	"challange2019/pkg/models"
	"challange2019/tools"
	"sync"
)

const (
	ChanBufferSize = 42
)

type DeliverySvc struct {
	DeliveriesSourceFilepath       string
	PartnersSourceFilepath         string
	PartnersCapacitySourceFilepath string
}

func NewDeliverySvc(deliverySourceFilepath, partnerSourceFilepath, partnersCapacitySourceFilepath string) *DeliverySvc {
	return &DeliverySvc{
		deliverySourceFilepath,
		partnerSourceFilepath,
		partnersCapacitySourceFilepath,
	}
}

func (svc *DeliverySvc) DistributeDeliveriesAmongPartnersByMinCost() (*models.DeliveryOutputStorage, []error) {
	output := &models.DeliveryOutputStorage{Container: []*models.DeliveryOutput{}}
	errors := &models.Errors{Container: []error{}}

	deliveryCsvCfg := tools.NewCsvReaderConfig(
		false,
		svc.DeliveriesSourceFilepath,
		make(chan *tools.CsvRow, ChanBufferSize),
		make(chan error))
	go csvRead(deliveryCsvCfg, errors)

	deliveryParserCfg := models.NewDeliveryParserConfig(
		deliveryCsvCfg.RowChan,
		make(chan *models.DeliveryInput, ChanBufferSize),
		make(chan error))
	go func(e *models.Errors) {
		go deliveryParserCfg.ParseDeliveriesInputCsv()
		for err := range deliveryParserCfg.ErrChan {
			e.Add(err)
		}
	}(errors)

	wg := sync.WaitGroup{}
	for di := range deliveryParserCfg.ParsedDataChan {
		wg.Add(1)

		go func(di *models.DeliveryInput) {
			defer wg.Done()

			partnersReaderCfg := tools.NewCsvReaderConfig(
				true,
				svc.PartnersSourceFilepath,
				make(chan *tools.CsvRow, ChanBufferSize),
				make(chan error))
			go csvRead(partnersReaderCfg, errors)

			partnerParserCfg := models.NewPartnerParserConfig(
				partnersReaderCfg.RowChan,
				make(chan *models.Partner, ChanBufferSize),
				make(chan error))

			go func(e *models.Errors) {
				go partnerParserCfg.ParsePartnerByTheaterFromCsv(di.TheaterID)
				for err := range partnerParserCfg.ErrChan {
					e.Add(err)
				}
			}(errors)

			models.FindCheapestDeliveryOutput(di, partnerParserCfg.ParsedDataChan, output)
		}(di)

	}

	wg.Wait()

	return output, errors.Container
}

func csvRead(csvCfg *tools.CsvReaderConfig, e *models.Errors) {
	go csvCfg.ReadLineFromCsv()
	for err := range csvCfg.ErrChan {
		e.Add(err)
	}
}
