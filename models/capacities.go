package models

//Capacity input data format for capacity
type Capacity struct {
	PartnerID    string `csv:"Partner_ID"`
	CapacityInGB int    `csv:"Capacity_in_GB"`
	NotUsed      string `csv:"-"`
}
