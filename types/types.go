package types

type ProblemOps struct {
	InputFile      string
	OutputFile     string
	CapacitiesFile string
	PartnersFile   string
}

type Slab struct {
	MinRange int
	MaxRange int
	MinCost  int
	CostGB   int
}

type Theartre string

type Partner string

type PartnersData map[Partner][]Slab

type WholeData map[Theartre]PartnersData

type CapacityData map[Partner]int

type Combination struct {
	Cost        int
	PartnerComb string
	Possible    bool
	Undel       int
}
