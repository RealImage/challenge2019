package models

//Output is output data
type Output struct {
	DistributorID string `csv:"distributorId"`
	Accepted      bool   `csv:"accepted"`
	PartnerID     string `csv:"partnerId"`
	Cost          string `csv:"cost"`
}
