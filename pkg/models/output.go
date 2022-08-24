package models

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
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
	v.Lock()
	ol.Container = append(ol.Container, v.Container...)
	ol.Unlock()
	v.Unlock()
}

func FindMostProfitableOutput(d *Delivery, partners []*Partner, outChan chan *Output) {
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

func FindOutputListSortedByCost(d *Delivery, partners []*Partner, sortedOutList chan *OutputList) {
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
