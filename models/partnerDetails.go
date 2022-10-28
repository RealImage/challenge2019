package models

type PartnerDetails struct {
	TheatreID   string
	SizeSlab    Slab
	MinimumCost int
	CostPerGB   int
	PartnerID   string
}

type Slab struct {
	Min int
	Max int
}
type PartnerMap map[string][]PartnerDetails
type TheatreMap map[string]PartnerMap
