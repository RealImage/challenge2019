package main

type Parterns struct {
	Theatre     string
	SlabSize    string
	MinimumCost int
	CostPerGB   int
	PartnerID   string
}

type Capacities struct {
	PartnerId string
	Capacity  int
}

type Input struct {
	DeliveryId string
	Size       int
	Theatre    string
}

type Output struct {
	DeliveryId    string
	IsDeliverable bool
	Cost          int
	PartnerId     string
}
