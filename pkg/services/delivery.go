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
	DeliveriesSourceFilepath       string
	PartnersSourceFilepath         string
	PartnersCapacitySourceFilepath string
}

func NewDeliverySvc(deliverySourceFilepath, partnerSourceFilepath, partnersCapacitySourceFilepath string) *DeliverySvc {
	return &DeliverySvc{deliverySourceFilepath, partnerSourceFilepath, partnersCapacitySourceFilepath}
}

// DistributeDeliveriesAmongPartnersByMinCost distributes given deliveries (read from DeliverySvc.DeliveriesSourceFilepath)
// among the partners (read from DeliverySvc.PartnersSourceFilepath) by minimum cost, packing result to *models.Output
// and sends it to  chan *models.Output, if any error acquired, sends it to  chan error
func (svc *DeliverySvc) DistributeDeliveriesAmongPartnersByMinCost(outChan chan<- *m.Output, errChan chan<- error) {
	defer close(errChan)
	defer close(outChan)
	wg := sync.WaitGroup{}

	dp := m.NewDeliveryParserConfig(tools.NewCsvReaderConfig(svc.DeliveriesSourceFilepath, false, ChanBufferSize), ChanBufferSize)
	go func() {
		go dp.ReadDeliveriesFromCsv()
		for e := range dp.ErrChan {
			errChan <- e
		}
	}()

	dList := []*m.Delivery{}
	wg.Add(1)
	go func() {
		fillDeliveryListFromChan(&dList, dp.ParsedDataChan)
		wg.Done()
	}()

	pp := m.NewPartnerParserConfig(tools.NewCsvReaderConfig(svc.PartnersSourceFilepath, true, ChanBufferSize), ChanBufferSize)
	go func() {
		go pp.ReadPartnerFromCsv()
		for e := range pp.ErrChan {
			errChan <- e
		}
	}()

	pMap := make(map[string][]*m.Partner)
	wg.Add(1)
	go func() {
		fillPartnerMapByTheaterFromChan(pMap, pp.ParsedDataChan)
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

	dp := m.NewDeliveryParserConfig(tools.NewCsvReaderConfig(svc.DeliveriesSourceFilepath, false, ChanBufferSize), ChanBufferSize)
	go func() {
		go dp.ReadDeliveriesFromCsv()
		for err := range dp.ErrChan {
			errChan <- err
		}
	}()

	dMap := make(map[string][]*m.Delivery)
	wg.Add(1)
	go func() {
		fillDeliveryMapByTheaterFromChan(dMap, dp.ParsedDataChan)
		wg.Done()
	}()

	pp := m.NewPartnerParserConfig(tools.NewCsvReaderConfig(svc.PartnersSourceFilepath, true, ChanBufferSize), ChanBufferSize)
	go func() {
		go pp.ReadPartnerFromCsv()
		for err := range pp.ErrChan {
			errChan <- err
		}
	}()

	pMap := make(map[string][]*m.Partner)
	wg.Add(1)
	go func() {
		fillPartnerMapByTheaterFromChan(pMap, pp.ParsedDataChan)
		wg.Done()
	}()

	cp := m.NewCapacityParserConfig(tools.NewCsvReaderConfig(svc.PartnersCapacitySourceFilepath, true, ChanBufferSize), ChanBufferSize)
	go func() {
		go cp.ReadCapacityFromCsv()
		for err := range cp.ErrChan {
			errChan <- err
		}
	}()

	cMap := make(map[string]int)
	wg.Add(1)
	go func() {
		fillCapacityMapByPartnerFromChan(cMap, cp.ParsedDataChan)
		wg.Done()
	}()

	wg.Wait()

	for theaterId, _ := range dMap {
		dlist := dMap[theaterId]
		plist := pMap[theaterId]

		wg.Add(1)
		go func(dList []*m.Delivery, plist []*m.Partner) {
			defer wg.Done()

			outTransmitterChan := make(chan *m.Output)
			go calculateDeliveriesByCostAndCapacity(dList, plist, cMap, outTransmitterChan)

			for o := range outTransmitterChan {
				resChan <- o
			}
		}(dlist, plist)
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
