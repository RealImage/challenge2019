package models

//Output is output data format
type Output struct {
	DeliveryID       string `csv:"deliveryId"`
	DeliveryPossible bool   `csv:"deliveryPossible"`
	PartnerID        string `csv:"partnerId"`
	CostOfDelivery   string `csv:"costOfDelivery"`
}
