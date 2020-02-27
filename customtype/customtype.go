package customtype

// Theatre detials
type Theatre map[string]Partner

// Partner detials
type Partner map[string]Slots

// Slots detials
type Slots map[string]Slot

// Capacities detials
type Capacities map[string]int64

// Slot detials
type Slot struct {
	Slot        string
	MinimumCost int64
	Fare        int64
}

// Delivery detials
type Delivery struct {
	Delivery string
	Data     int64
	Theatre  string
	Price    int64
	Partner  string
}

// DeliveryList det
type DeliveryList []Delivery

func (p DeliveryList) Len() int           { return len(p) }
func (p DeliveryList) Less(i, j int) bool { return p[i].Data < p[j].Data }
func (p DeliveryList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
