package schemas

type PartnerCsvColumns struct {
	TheatreID   int
	SizeSlab    string
	MinimumCost int
	CostPerGB   int
	PartnerID   int
}

// Details of each partner_id that is to be stored in DB for specific theatre_id
type PartnerDetail struct {
	PartnerID   string
	SizeSlab    string
	MinimumCost int
	CostPerGB   int
}

/**
	Specific to Problem Statement 2
	For each delivery_id in input.csv stores size_requirement, number of valid available Partner Options, details of each partner options.
**/
type DeliveryChoices struct {
	DeliveryID      string
	SizeRequirement int
	NumberOfOptions int
	PartnerOptions  []struct {
		PartnerID      string
		CostOfDelivery int
	}
}
