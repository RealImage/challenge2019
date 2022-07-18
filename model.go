package main

type Rate struct {
	Theatre     string
	Partner     string
	CostPerGB   int
	MinimumCost int
	Lower       int
	Upper       int
}

type DeliveryRequest struct {
	ID      string
	Size    int
	Theatre string
}

type DeliveryResponse struct {
	ID       string
	Accepted bool
	Partner  string
	Cost     int
	Request  *DeliveryRequest
}

type DeliveryResponses []*DeliveryResponse

func (dr DeliveryResponses) Len() int           { return len(dr) }
func (dr DeliveryResponses) Less(i, j int) bool { return dr[i].Request.Size < dr[j].Request.Size }
func (dr DeliveryResponses) Swap(i, j int)      { dr[i], dr[j] = dr[j], dr[i] }
