package models

import (
	"strconv"
)

type DeliveryInput struct {
	Delivery  string
	Size      int
	TheatreID string
}

type Cost int
type PartnerID string

type DeliveryOutput struct {
	Delivery   string
	IsPossible bool
	PartnerID  PartnerID
	Cost       Cost
}

type DeliveryOutputs []*DeliveryOutput

func (do *DeliveryOutputs) GetPartnersDeals(pId PartnerID) DeliveryOutputs {
	partnersDeals := []*DeliveryOutput{}
	for _, output := range *do {
		if output.PartnerID == pId {
			partnersDeals = append(partnersDeals, output)
		}
	}
	return partnersDeals
}

func (c *Cost) MarshalCSV() (string, error) {
	if *c == 0 {
		return " ", nil
	}
	return strconv.Itoa(int(*c)), nil
}

func (p *PartnerID) MarshalCSV() (string, error) {
	if *p == "" {
		return " ", nil
	}
	return string(*p), nil
}
