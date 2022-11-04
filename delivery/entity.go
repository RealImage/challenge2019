package delivery

import (
	"strconv"
	"strings"
)

type Input struct {
	DeliveryID TrimString `csv:"delivery_id"`
	Amount     int        `csv:"amount"`
	TheatreID  TrimString `csv:"theatre_id"`
}

type Output struct {
	DeliveryID TrimString `csv:"delivery_id"`
	IsPossible bool       `csv:"is_possible"`
	PartnerID  TrimString `csv:"partner_id"`
	Cost       int        `csv:"cost"`
}

type Partner struct {
	TheatreID TrimString `csv:"theatre_id"`
	Slab      Slab       `csv:"slab"`
	MinCost   int        `csv:"min_cost"`
	CostPerGB int        `csv:"cost_gb"`
	PartnerID TrimString `csv:"partner_id"`
}

type Slab struct {
	MinSlab int
	MaxSlab int
}

type TrimString string

func (s *TrimString) UnmarshalCSV(csv string) (err error) {
	*s = TrimString(strings.TrimSpace(csv))
	return
}

func (s *Slab) UnmarshalCSV(csv string) (err error) {
	ss := strings.Split(strings.TrimSpace(csv), "-")
	s.MinSlab, err = strconv.Atoi(ss[0])
	if err != nil {
		return err
	}
	s.MaxSlab, err = strconv.Atoi(ss[1])
	if err != nil {
		return err
	}
	return nil
}
