package models

//Partners denotes the attributes on Partners input csv
type Partners struct {
	TheatreID    string `csv:"Theatre"`
	SizeSlabInGB string `csv:"Size_Slab"`
	MinimumCost  int    `csv:"Minimum_cost"`
	CostPerGB    int    `csv:"Cost_Per_GB"`
	PartnerID    string `csv:"Partner_Id"`
	NotUsed      string `csv:"-"`
}

//PartnerConfig denotes attribute with PartnerID as Primary Key
type PartnerConfig struct {
	TID         string
	MinSlabSize int
	MaxSlabSize int
	MinCost     int
	CperGB      int
}
