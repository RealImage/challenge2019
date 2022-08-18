package entities

type Input struct {
	ID             string
	SizeOfDelivery string
	TheatreID      string
}

type Output struct {
	ID             string
	IsDeliverable  string
	Partner        string
	CostOfDelivery string
}

type Partner struct {
	TheatreID   string
	SizeSlab    string
	MinimumCost string
	CostPerGB   string
	PartnerID   string
}

type PartnerCapacity struct {
	PartnerID string
	Capacity  string
}

type TheatresInInput struct {
	TheatreID string
}

type Capacity struct {
	PartnerID string
	Capacity  string
}
