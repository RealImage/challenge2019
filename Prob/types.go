package Prob

type PartnerDataStr struct {
	Theatre   string
	Size      string
	MinCost   string
	CostPerGB string
	PartnerID string
}

type PartnerData struct {
	Theatre      string
	Size         SizeSlab
	MinCost      float64
	CostPerGB    float64
	PartnerID    string
	TotalCost    float64
	DeliveryCost float64
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
