package model

type Partner struct {
	MinGB   int
	MaxGB   int
	MinCost int
	PerGB   int
	Partner string
}

type Delivery struct {
	Name    string
	Theatre string
	Amount  int
}

type Result struct {
	Name    string
	Partner string
	Cost    int
}

type Result2 struct {
	Name    string
	Amount  int
	Partner string
	Cost    int
}
