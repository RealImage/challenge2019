package Models

type DeliveryPartnerCostSlab struct {
	TheatreName string
	SizeSlab    string
	MinimumCost string
	CostPerGb   string
	PartnerID   string
}

type Input struct {
	ID          string
	SizeSlab    string
	TheatreName string
}

type Output struct {
	ID                 string
	DeliveryIsPossible string
	PartnerID          string
	Price              string
}
