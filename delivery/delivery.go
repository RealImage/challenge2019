package delivery

type Delivery struct {
	OrderID   int
	Size      int
	TheatreID string
	Possible  bool
	PartnerID string
	Cost      int
}
