package services

import (
	m "challange2019/pkg/models"
	"challange2019/tools"
	"sync"
)

const (
	ChanBufferSize = 42
)

type DeliverySvc struct {
	DeliveryReader tools.CsvReader
	PartnerReader  tools.CsvReader
	CapacityReader tools.CsvReader
}

func NewDeliverySvc(deliveryReader, partnerReader, capacityReader tools.CsvReader) *DeliverySvc {
	return &DeliverySvc{deliveryReader, partnerReader, capacityReader}
}

// DistributeDeliveriesAmongPartnersByMinCost distributes given deliveries (read from DeliverySvc.DeliveriesSourceFilepath)
// among the partners (read from DeliverySvc.PartnersSourceFilepath) by minimum cost, packing result to *models.Output
// and sends it to  chan *models.Output, if any error acquired, sends it to  chan error
func (svc *DeliverySvc) DistributeDeliveriesAmongPartnersByMinCost(outChan chan<- *m.Output, errChan chan<- error) {
	defer close(errChan)
	defer close(outChan)
	wg := sync.WaitGroup{}

	dList := []*m.Delivery{}
	dpChan := make(chan *m.Delivery, ChanBufferSize)
	wg.Add(1)

	go func() {
		go parseDeliveries(svc, errChan, dpChan)
		fillDeliveryListFromChan(&dList, dpChan)
		wg.Done()
	}()

	pMap := make(map[string][]*m.Partner)
	ppChan := make(chan *m.Partner, ChanBufferSize)
	wg.Add(1)

	go func() {
		go parsePartners(svc, errChan, ppChan)
		fillPartnerMapByTheaterFromChan(pMap, ppChan)
		wg.Done()
	}()

	wg.Wait()

	for _, d := range dList {
		outBufferChan := make(chan *m.Output, ChanBufferSize)

		wg.Add(1)
		go func(d *m.Delivery) {
			defer wg.Done()
			m.FindMostProfitableOutput(d, pMap[d.TheaterID], outBufferChan)
			for o := range outBufferChan {
				outChan <- o
			}
		}(d)

	}

	wg.Wait()
}

// DistributeDeliveriesAmongPartnersByMinCostAndCapacity distributes given deliveries (read from DeliverySvc.DeliveriesSourceFilepath)
// among the partners (read from DeliverySvc.PartnersSourceFilepath) by minimum cost, takes partner
// capacity (read from DeliverySvc.PartnersCapacitySourceFilepath) into consideration as well. Packing result to *models.Output
// and sends it to  chan *models.Output, if any error acquired, sends it to  chan error
func (svc *DeliverySvc) DistributeDeliveriesAmongPartnersByMinCostAndCapacity(resChan chan<- *m.Output, errChan chan<- error) {
	defer close(resChan)
	defer close(errChan)
	wg := sync.WaitGroup{}

	dMap := make(map[string][]*m.Delivery)
	dpChan := make(chan *m.Delivery, ChanBufferSize)

	wg.Add(1)
	go func() {
		go parseDeliveries(svc, errChan, dpChan)
		fillDeliveryMapByTheaterFromChan(dMap, dpChan)
		wg.Done()
	}()

	pMap := make(map[string][]*m.Partner)
	ppChan := make(chan *m.Partner, ChanBufferSize)
	wg.Add(1)

	go func() {
		go parsePartners(svc, errChan, ppChan)
		fillPartnerMapByTheaterFromChan(pMap, ppChan)
		wg.Done()
	}()

	cChan := make(chan *m.Capacity, ChanBufferSize)
	cMap := make(map[string]int)
	wg.Add(1)

	go func() {
		go parseCapacity(svc, errChan, cChan)
		fillCapacityMapByPartnerFromChan(cMap, cChan)
		wg.Done()
	}()

	wg.Wait()

	for theaterId, _ := range dMap {
		dList := dMap[theaterId]
		pList := pMap[theaterId]

		wg.Add(1)
		go func(dList []*m.Delivery, plist []*m.Partner) {
			defer wg.Done()

			outTransmitterChan := make(chan *m.Output)
			go calculateDeliveriesByCostAndCapacity(dList, plist, cMap, outTransmitterChan)

			for o := range outTransmitterChan {
				resChan <- o
			}
		}(dList, pList)
	}

	wg.Wait()

}

// calculateDeliveriesByCostAndCapacity finds the best partner delivery option according to the cost and partner's capacity,
// if delivery is impossible, marks it appropriately. Sends the result to chan *models.Output
func calculateDeliveriesByCostAndCapacity(dList []*m.Delivery, pList []*m.Partner, cMap map[string]int, outChan chan *m.Output) {
	defer close(outChan)

	wg := sync.WaitGroup{}
	for _, d := range dList {

		wg.Add(1)
		go func(d *m.Delivery) {
			defer wg.Done()
			sortedOutListChan := make(chan *m.OutputList)
			go m.GetOutputListSortedByCost(d, pList, sortedOutListChan)

			for sortedOutList := range sortedOutListChan {
				distributeDeliveriesByPartnersByCapacityAndCost(outChan, sortedOutList, cMap)
			}
		}(d)
	}

	wg.Wait()

}

func distributeDeliveriesByPartnersByCapacityAndCost(
	result chan *m.Output,
	outVarieties *m.OutputList,
	capacityMap map[string]int) {

	// key - partner id, value - how much free capacity has
	availableSpace := make(map[string]int)

	for i, possibleOut := range outVarieties.Container {
		if possibleOut.Partner == nil {
			result <- possibleOut
			break
		}

		if _, ok := availableSpace[possibleOut.Partner.ID]; ok {
			if availableSpace[possibleOut.Partner.ID]-possibleOut.Delivery.ContentSize >= 0 {
				result <- possibleOut
				availableSpace[possibleOut.Partner.ID] -= possibleOut.Delivery.ContentSize
				break

			}
			continue
		}

		if capacityMap[possibleOut.Partner.ID]-possibleOut.Delivery.ContentSize >= 0 {
			result <- possibleOut
			availableSpace[possibleOut.Partner.ID] = capacityMap[possibleOut.Partner.ID] - possibleOut.Delivery.ContentSize
			break
		}

		if i == len(outVarieties.Container)-1 {
			possibleOut.Partner = nil
			possibleOut.Cost = -1
			possibleOut.IsPossible = false
			result <- possibleOut
			break
		}
	}
}

func parseDeliveries(svc *DeliverySvc, errChan chan<- error, result chan<- *m.Delivery) {
	// read deliveries from csv
	drRowChan := make(chan *tools.CsvRow, ChanBufferSize)
	drErrChan := make(chan error, ChanBufferSize)
	go func() {
		go svc.DeliveryReader.ReadLineFromCsv(drRowChan, drErrChan)
		for e := range drErrChan {
			errChan <- e
		}
	}()

	// parse deliveries from csv
	dpErrChan := make(chan error, ChanBufferSize)
	go m.ParseDeliveryFromCsvRow(drRowChan, result, dpErrChan)
	for e := range drErrChan {
		errChan <- e
	}
}

func parseCapacity(svc *DeliverySvc, errChan chan<- error, result chan<- *m.Capacity) {
	// read capacity from csv
	rowChan := make(chan *tools.CsvRow, ChanBufferSize)
	crErrChan := make(chan error, ChanBufferSize)
	go func() {
		go svc.CapacityReader.ReadLineFromCsv(rowChan, crErrChan)
		for e := range crErrChan {
			errChan <- e
		}
	}()

	// parse capacity from csv
	cpErrChan := make(chan error, ChanBufferSize)
	go m.ReadCapacityFromCsv(rowChan, result, cpErrChan)
	for e := range crErrChan {
		errChan <- e
	}
}

func parsePartners(svc *DeliverySvc, errChan chan<- error, result chan<- *m.Partner) {
	// read partners from csv
	rowChan := make(chan *tools.CsvRow, ChanBufferSize)
	rErrChan := make(chan error, ChanBufferSize)
	go func() {
		go svc.PartnerReader.ReadLineFromCsv(rowChan, rErrChan)
		for e := range rErrChan {
			errChan <- e
		}
	}()

	// parse partners from csv
	ppErrChan := make(chan error, ChanBufferSize)
	go m.ParsePartnerFromCsvRow(rowChan, result, ppErrChan)
	for e := range ppErrChan {
		errChan <- e
	}
}

func fillDeliveryListFromChan(dList *[]*m.Delivery, c chan *m.Delivery) {
	for d := range c {
		*dList = append(*dList, d)
	}
}

func fillDeliveryMapByTheaterFromChan(dMap map[string][]*m.Delivery, c chan *m.Delivery) {
	for d := range c {
		if _, ok := dMap[d.TheaterID]; ok {
			dMap[d.TheaterID] = append(dMap[d.TheaterID], d)
		} else {
			dMap[d.TheaterID] = []*m.Delivery{d}
		}
	}
}

func fillCapacityMapByPartnerFromChan(cMap map[string]int, c chan *m.Capacity) {
	for capacity := range c {
		cMap[capacity.PartnerId] = capacity.Value
	}
}

func fillPartnerMapByTheaterFromChan(pMap map[string][]*m.Partner, c chan *m.Partner) {
	for p := range c {
		if _, ok := pMap[p.TheaterID]; ok {
			pMap[p.TheaterID] = append(pMap[p.TheaterID], p)
		} else {
			pMap[p.TheaterID] = []*m.Partner{p}
		}
	}
}
