package models

import (
	"strconv"
)

type Output struct {
	DeliveryID     string
	isPossible     bool
	PartnerID      string
	CostOfDelivery int
}

func NewOutput(
	DeliveryID string,
	isPossible bool,
	PartnerID string,
	CostOfDelivery int) Output {
	return Output{
		DeliveryID:     DeliveryID,
		isPossible:     isPossible,
		PartnerID:      PartnerID,
		CostOfDelivery: CostOfDelivery,
	}
}

func (o *Output) String() []string {
	costOfDelivery := ""
	if o.CostOfDelivery != 0 {
		costOfDelivery = strconv.Itoa(o.CostOfDelivery)
	}
	return []string{o.DeliveryID, strconv.FormatBool(o.isPossible), o.PartnerID, costOfDelivery}
}
