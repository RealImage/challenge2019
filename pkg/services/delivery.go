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

func (svc *DeliverySvc) DistributeDeliveriesAmongPartnersByMinCost() (*models.DeliveryOutputList, []error) {
	output := &models.DeliveryOutputList{Container: []*models.DeliveryOutput{}}
	errors := &models.Errors{Container: []error{}}
	wg := sync.WaitGroup{}

	dp := models.NewDeliveryParserConfig(tools.NewCsvReaderConfig(svc.DeliveriesSourceFilepath, false, ChanBufferSize), ChanBufferSize)
	go dp.ReadDeliveriesInputFromCsv()

	diList := []*models.DeliveryInput{}
	wg.Add(1)
	go func() {
		fillDeliveryListFromChan(&diList, dp.ParsedDataChan)
		wg.Done()
	}()

	pp := models.NewPartnerParserConfig(tools.NewCsvReaderConfig(svc.PartnersSourceFilepath, true, ChanBufferSize), ChanBufferSize)
	go pp.ReadPartnerFromCsv()

	pMap := make(map[string][]*models.Partner)
	wg.Add(1)
	go func() {
		fillPartnerMapByTheaterFromChan(pMap, pp.ParsedDataChan)
		wg.Done()
	}()

	wg.Wait()

	for _, di := range diList {
		wg.Add(1)
		go func(di *models.DeliveryInput) {
			models.FindCheapestDeliveryOutput(di, pMap[di.TheaterID], output)
			wg.Done()
		}(di)
	}

	wg.Wait()

	return output, errors.Container
}

func (svc *DeliverySvc) DistributeDeliveriesAmongPartnersByMinCostAndCapacity() (*models.DeliveryOutputList, []error) {
	output := &models.DeliveryOutputList{Container: []*models.DeliveryOutput{}}
	errors := &models.Errors{Container: []error{}}

	return output, errors.Container
}

func fillDeliveryListFromChan(diList *[]*models.DeliveryInput, c chan *models.DeliveryInput) {
	for di := range c {
		*diList = append(*diList, di)
	}
}

func fillPartnerMapByTheaterFromChan(pMap map[string][]*models.Partner, c chan *models.Partner) {
	for p := range c {
		if _, ok := pMap[p.TheaterID]; ok {
			pMap[p.TheaterID] = append(pMap[p.TheaterID], p)
		} else {
			pMap[p.TheaterID] = []*models.Partner{p}
		}
	}
}
