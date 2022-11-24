package delivery

// Delivery holds information about Delivery.
type Delivery struct {
	OrderID   int
	Size      int
	TheatreID string
	Possible  bool
	PartnerID string
	Cost      int
}
