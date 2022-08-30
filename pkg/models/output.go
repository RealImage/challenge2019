package models

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

const (
	OutPutFieldsCount = 4
)

type Output struct {
	Delivery   *Delivery
	Partner    *Partner
	IsPossible bool
	Cost       int
}

type OutputList struct {
	sync.RWMutex
	Container []*Output
}

func (ol *OutputList) Add(do *Output) {
	ol.Lock()
	ol.Container = append(ol.Container, do)
	ol.Unlock()
}

func (ol *OutputList) AppendFromChan(doChan chan *OutputList) {
	for o := range doChan {
		ol.Append(o)
	}
}

func (ol *OutputList) Append(v *OutputList) {
	ol.Lock()
	ol.Container = append(ol.Container, v.Container...)
	ol.Unlock()
}

// EncodeOutputToCsvRow encodes Output's data to []string to write it into csv file
func EncodeOutputToCsvRow(outputChan <-chan *Output, rowChan chan<- []string) {
	defer close(rowChan)

	for o := range outputChan {
		rowChan <- encodeOutputDataToStringArray(o)
	}
}

func encodeOutputDataToStringArray(o *Output) []string {
	res := make([]string, OutPutFieldsCount)
	res[0] = o.Delivery.ID
	res[1] = fmt.Sprintf("%t", o.IsPossible)

	res[2] = ""
	if o.Partner != nil {
		res[2] = o.Partner.ID
	}

	res[3] = ""
	if o.Cost >= 0 {
		res[3] = strconv.Itoa(o.Cost)
	}

	return res
}

// FindMostProfitableOutput finds the cheapest offer among the []*Partner for given Delivery,
// sends result to chan *Output
func FindMostProfitableOutput(d *Delivery, partners []*Partner, outChan chan<- *Output) {
	defer close(outChan)
	o := &Output{Delivery: d, Cost: -1}

	for _, p := range partners {
		cost, isPossible := p.CalculateCost(d.ContentSize)

		if isPossible && (cost < o.Cost || o.Cost < 0) {
			o.Cost = cost
			o.Partner = p
			o.IsPossible = isPossible
		}
	}

	outChan <- o
}

// GetOutputListSortedByCost creates *OutputList, where container of *Output is sorted by Output.Cost ascending,
// the impossible *Output's instances are put in the end of *OutputList,
// sends the resulting *OutputList to chan *OutputList
func GetOutputListSortedByCost(d *Delivery, partners []*Partner, sortedOutList chan<- *OutputList) {
	defer close(sortedOutList)
	res := &OutputList{}
	for _, p := range partners {
		cost, isPossible := p.CalculateCost(d.ContentSize)
		if isPossible {
			o := &Output{Delivery: d, Cost: cost, Partner: p, IsPossible: isPossible}
			res.Add(o)
		}
	}

	if len(res.Container) == 0 {
		o := &Output{Delivery: d, Cost: -1, Partner: nil, IsPossible: false}
		res.Add(o)
	}

	sort.Slice(res.Container, func(i, j int) bool {
		if !res.Container[i].IsPossible {
			return false
		}
		if !res.Container[j].IsPossible {
			return true
		}
		return res.Container[i].Cost < res.Container[j].Cost
	})

	sortedOutList <- res
}

func (o *Output) String() string {
	cost := ""
	pId := ""
	if o.IsPossible {
		cost = strconv.Itoa(o.Cost)
		pId = o.Partner.ID
	}

	return fmt.Sprintf("%s, %t, %s, %s", o.Delivery.ID, o.IsPossible, pId, cost)
}
