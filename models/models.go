package models

type DeliveryRequest struct {
	DeliveryId string `csv:"delivery_id"`
	Amount     int    `csv:"amount"`
	TheatreId  string `csv:"theatre_id"`
}

type Capacity struct {
	PartnerId    string `csv:"partner_id"`
	CapacityInGb int    `csv:"capacity_in_gb"`
}

type DeliveryResponse struct {
	DeliveryId string `csv:"delivery_id"`
	IsPossible bool   `csv:"is_possible"`
	PartnerId  string `csv:"partner_id"`
	Cost       string `csv:"cost"`
}

type Partner struct {
	TheatreId   string `csv:"theatre_id"`
	Slab        string `csv:"size_slab"`
	CostMinimal int    `csv:"minimum_cost"`
	CostPerGB   int    `csv:"cost_per_gb"`
	Id          string `csv:"partner_id"`
}
