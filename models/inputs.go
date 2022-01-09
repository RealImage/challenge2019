package models

//Input input data format
type Input struct {
	DistributorID string `csv:"distributorId"`
	MinCost       int    `csv:"minCost"`
	TheatreID     string `csv:"theatreId"`
	NotUsed       string `csv:"-"`
}
