package models

type OutputDetails struct {
	DeliveryID  string
	Feasibility bool
	PartnerID   string
	Cost        int
}

type TotalDataPerPartner struct {
	Data int
	//indivisible data unit map, maps dataunit to output 1
	DataUnitsMap map[int]int
}
