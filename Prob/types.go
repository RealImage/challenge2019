package Prob

type PartnerDataStr struct {
	Theatre   string
	Size      string
	MinCost   string
	CostPerGB string
	PartnerID string
}

type PartnerData struct {
	Delivery     DeliveryInfo
	Theatre      string
	Size         SizeSlab
	MinCost      float64
	CostPerGB    float64
	PartnerID    string
	TotalCost    float64
	DeliveryCost float64
	Capacity     int
}
type DelAndPartners struct {
	DeliveryID   string
	DeliverySize int
	TheatreID    string
	PartnerID    string
	//Capacity     int
}
type SizeSlab struct {
	Min int
	Max int
}

type DeliveryInfo struct {
	DeliveryID   string
	DeliverySize int
	Theatre      string
}

type DeliveryAndPartners struct {
	Delivery DeliveryInfo
	Partners []PartnerData
}

type FinalChoice struct {
	DeliveryID    string
	IsPossible    bool
	ChosenPartner string
	MinimumCost   string
}

type PartnerInfo struct {
	PartnerID    string
	DeliveryCost float64
	Capacity     int
}

type DeliveryDetails struct {
	DeliveryID   string
	DeliverySize int
}

type CapacityInfo struct {
	PartnerID string
	Capacity  string
}

type CapacityDetails struct {
	PartnerID string
	Capacity  int
}
