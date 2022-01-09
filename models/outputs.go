package models

//Output is output data format
type Output struct {
	DistributorID string `csv:"distributorId"`
	SLAAccepted   bool   `csv:"slaAccepted"`
	PartnerID     string `csv:"partnerId"`
	CostAgreed    string `csv:"costAgreed"`
}
