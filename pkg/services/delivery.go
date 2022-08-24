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
	return &DeliverySvc{
		deliverySourceFilepath,
		partnerSourceFilepath,
		partnersCapacitySourceFilepath,
	}
}

func (svc *DeliverySvc) DistributeDeliveriesAmongPartnersByMinCost() (*m.OutputList, []error) {
	output := &m.OutputList{Container: []*m.Output{}}
	errors := &m.Errors{Container: []error{}}
	wg := sync.WaitGroup{}

	dp := m.NewDeliveryParserConfig(tools.NewCsvReaderConfig(svc.DeliveriesSourceFilepath, false, ChanBufferSize), ChanBufferSize)
	go dp.ReadDeliveriesInputFromCsv()

	dList := []*m.Delivery{}
	wg.Add(1)
	go func() {
		fillDeliveryListFromChan(&dList, dp.ParsedDataChan)
		wg.Done()
	}()

	pp := m.NewPartnerParserConfig(tools.NewCsvReaderConfig(svc.PartnersSourceFilepath, true, ChanBufferSize), ChanBufferSize)
	go pp.ReadPartnerFromCsv()

	pMap := make(map[string][]*m.Partner)
	wg.Add(1)
	go func() {
		fillPartnerMapByTheaterFromChan(pMap, pp.ParsedDataChan)
		wg.Done()
	}()

	wg.Wait()

	for _, d := range dList {
		outChan := make(chan *m.Output, ChanBufferSize)
		
		wg.Add(1)
		go func(d *m.Delivery) {
			defer wg.Done()
			m.FindMostProfitableOutput(d, pMap[d.TheaterID], outChan)
			for o := range outChan {
				output.Add(o)
			}
		}(d)

	}

	wg.Wait()

	return output, errors.Container
}

func (svc *DeliverySvc) DistributeDeliveriesAmongPartnersByMinCostAndCapacity() (*m.OutputList, []error) {
	output := &m.OutputList{Container: []*m.Output{}}
	errors := &m.Errors{Container: []error{}}
	wg := sync.WaitGroup{}

	dp := m.NewDeliveryParserConfig(tools.NewCsvReaderConfig(svc.DeliveriesSourceFilepath, false, ChanBufferSize), ChanBufferSize)
	go dp.ReadDeliveriesInputFromCsv()

	dMap := make(map[string][]*m.Delivery)
	wg.Add(1)
	go func() {
		fillDeliveryMapByTheaterFromChan(dMap, dp.ParsedDataChan)
		wg.Done()
	}()

	pp := m.NewPartnerParserConfig(tools.NewCsvReaderConfig(svc.PartnersSourceFilepath, true, ChanBufferSize), ChanBufferSize)
	go pp.ReadPartnerFromCsv()
	pMap := make(map[string][]*m.Partner)
	wg.Add(1)
	go func() {
		fillPartnerMapByTheaterFromChan(pMap, pp.ParsedDataChan)
		wg.Done()
	}()

	cp := m.NewCapacityParserConfig(tools.NewCsvReaderConfig(svc.PartnersCapacitySourceFilepath, true, ChanBufferSize), ChanBufferSize)
	go cp.ReadCapacityFromCsv()
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
		go func(dlist []*m.Delivery, plist []*m.Partner) {
			defer wg.Done()

			outChan := make(chan *m.OutputList)
			go calculateDeliveriesByCostAndCapacity(dlist, plist, cMap, outChan)
			output.AppendFromChan(outChan)
		}(dlist, plist)

	}

	wg.Wait()

	return output, errors.Container
}

func calculateDeliveriesByCostAndCapacity(dList []*m.Delivery, pList []*m.Partner, cMap map[string]int, outListChan chan *m.OutputList) {
	defer close(outListChan)
	wg := sync.WaitGroup{}
	res := &m.OutputList{}

	for _, d := range dList {
		wg.Add(1)

		go func(d *m.Delivery) {
			defer wg.Done()
			oChan := make(chan *m.OutputList)
			go m.FindOutputListSortedByCost(d, pList, oChan)

			// key - partner id, value - how much free capacity has
			availableSpace := make(map[string]int)
			for oList := range oChan {
				distributeDeliveriesByPartnersByCapacityAndCost(res, oList, availableSpace, cMap)
			}
		}(d)
	}
	wg.Wait()

	outListChan <- res

}

func distributeDeliveriesByPartnersByCapacityAndCost(
	result *m.OutputList,
	outVarieties *m.OutputList,
	availableSpace map[string]int,
	capacityMap map[string]int) {

	for i, possibleOut := range outVarieties.Container {
		if possibleOut.Partner == nil {
			result.Add(possibleOut)
			break
		}

		if _, ok := availableSpace[possibleOut.Partner.ID]; ok {
			if availableSpace[possibleOut.Partner.ID]-possibleOut.Delivery.ContentSize >= 0 {
				result.Add(possibleOut)
				availableSpace[possibleOut.Partner.ID] -= possibleOut.Delivery.ContentSize
				break

			}
			continue
		}

		if capacityMap[possibleOut.Partner.ID]-possibleOut.Delivery.ContentSize >= 0 {
			result.Add(possibleOut)
			availableSpace[possibleOut.Partner.ID] = capacityMap[possibleOut.Partner.ID] - possibleOut.Delivery.ContentSize
			break
		}

		if i == len(outVarieties.Container)-1 {
			possibleOut.Partner = nil
			possibleOut.Cost = -1
			possibleOut.IsPossible = false
			result.Add(possibleOut)
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
