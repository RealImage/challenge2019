package types

type ProblemOps struct {
	InputFile      string
	OutputFile     string
	CapacitiesFile string
	PartnersFile   string
}

type Slabs struct {
	MinRange int
	MaxRange int
	MinCost  int
	CostGB   int
}

type Theartre string

type Partner string

type PartnersData map[Partner][]Slabs

type WholeData map[Theartre]PartnersData
