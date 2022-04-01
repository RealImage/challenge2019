package models

//Input data format
type Input struct {
	DistributorID string `csv:"distributorId"`
	Cost          int    `csv:"cost"`
	TheatreID     string `csv:"theatreId"`
}
