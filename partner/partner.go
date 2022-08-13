package partner

// Partner holds detailed information about a Partner's delivery capability.
type Partner struct {
	ID        string
	TheatreID string
	MinCost   int
	CostPerGB int
	MinAmount int
	MaxAmount int
}
