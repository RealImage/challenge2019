package models

//Input input data format
type Input struct {
	DeliveryID     string `csv:"distributorId"`
	SizeOfDelivery int    `csv:"minCost"`
	TheatreID      string `csv:"theatreId"`
	NotUsed        string `csv:"-"`
}
