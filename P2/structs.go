package main

// Slap :
type Slap struct {
	MinSize int
	MaxSize int
	MinCost int
	Cost    int
}

// Delivery :
type Delivery struct {
	ID        string
	Size      int
	TheatreID string
}

// Solution :
type Solution struct {
	Deliveries map[string]*DeliverySolution
}

// DeliverySolution :
type DeliverySolution struct {
	PartnerID string
	Cost      int
}
