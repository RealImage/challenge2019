package challenge2019

import (
	"log"

	"github.com/zeebo/errs"

	"challenge2019/csv"
	"challenge2019/delivery"
	"challenge2019/partner"
)

var (
	// Error is an error class that indicates first task solve error.
	Error = errs.Class("root folder: first task solve error")
)

// TaskOne performs solving first task and parsing answer to file
func TaskOne() error {
	log.Println("Running first problem solve...")

	inputDeliveries, err := csv.ReadDeliveriesInput("./input.csv")
	if err != nil {
		return Error.Wrap(err)
	}

	partnerLists, err := csv.ReadPartners("./partners.csv")
	if err != nil {
		return Error.Wrap(err)
	}

	log.Println("input files parsed successful")

	err = csv.WriteOutput(taskOneAnswer(inputDeliveries, partnerLists))
	return Error.Wrap(err)
}

// taskOneAnswer returns answer to first task
func taskOneAnswer(deliveries map[string]*delivery.Delivery, partnerLists map[string][]partner.Partner) map[string]*delivery.Delivery {
	for id, delivery := range deliveries {
		// Not possible while solution not found.
		deliveries[id].Possible = false

		partners, exist := partnerLists[delivery.TheatreID]
		if !exist {
			// If there is no partners that deliver to this theatre, it is not possible, continue to next one.
			continue
		}

		for _, partner := range partners {
			if deliveries[id].Size < partner.MinAmount || deliveries[id].Size > partner.MaxAmount {
				continue
			}

			// At least one option found.
			deliveries[id].Possible = true
			cost := maxVal(partner.CostPerGB*deliveries[id].Size, partner.MinCost)

			// Setting min cost to first possible solution and then changing it.
			if deliveries[id].Cost == 0 {
				deliveries[id].Cost, deliveries[id].PartnerID = cost, partner.ID
			} else if cost < deliveries[id].Cost {
				deliveries[id].Cost, deliveries[id].PartnerID = cost, partner.ID
			}
		}
	}

	return deliveries
}

// maxVal returning maximum int value.
func maxVal(x, y int) int {
	if x > y {
		return x
	}

	return y
}
