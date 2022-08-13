package solve

import (
	"challenge2019/csv"
	"challenge2019/delivery"
	"challenge2019/partner"
	"fmt"
	"github.com/zeebo/errs"
)

var (
	// Error is an error class that indicates first task solve error.
	Error = errs.Class("root folder: first solve error")
)

// costLists represents lists of partners grouped by theatre id as map key.
var costLists map[int][]partner.Partner

// deliveries represents list of deliveries to be performed.
var deliveries []delivery.Delivery

func SolveFirstTask() error {
	deliveries, err := csv.ReadDeliveries("./input.csv")
	if err != nil {
		return Error.Wrap(err)
	}

	err = csv.WriteOutput(*deliveries)
	if err != nil {
		return Error.Wrap(err)
	}

	fmt.Println(deliveries)

	return nil
}
