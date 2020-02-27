package estimator

import (
	"reflect"
	"strconv"
	"strings"

	customType "github.com/praveenpkg8/challenge2019/customtype"
)

// Capacities memmory hold
var Capacities customType.Capacities = make(customType.Capacities)

// Partner memmory hold
var Partner string = ""

// Cost memmory hold
var Cost int64

// GetBilling to caculate output
func GetBilling(theatreDetials customType.Theatre, theatre string, data int64) (int64, string) {
	partners := theatreDetials[theatre]
	partnersList := reflect.ValueOf(partners).MapKeys()
	for _, partner := range partnersList {
		slots := partners[partner.String()]
		Cost, Partner = estimateFare(partner.String(), slots, data)
	}
	return Cost, Partner
}

// GetBillingV1 to caculate output
func GetBillingV1(theatreDetials customType.Theatre, theatre string, data int64) (int64, string) {
	partners := theatreDetials[theatre]
	partnersList := reflect.ValueOf(partners).MapKeys()
	for _, partner := range partnersList {
		slots := partners[partner.String()]
		if Capacities[partner.String()] < data {
			continue
		}
		Cost, Partner = estimateFare(partner.String(), slots, data)
	}
	Capacities[Partner] = Capacities[Partner] - data
	return Cost, Partner
}

func estimateFare(partner string, slots customType.Slots, data int64) (int64, string) {
	slotList := reflect.ValueOf(slots).MapKeys()
	for _, slot := range slotList {
		if inRange(slot.String(), data) {
			cost := calculateCost(slots[slot.String()], data)
			if Cost == 0 {
				Cost = cost
				Partner = partner
			}
			if cost < Cost {
				Cost = cost
				Partner = partner
			}
		}
	}
	return Cost, Partner
}

func inRange(_range string, data int64) bool {
	splitedRange := strings.Split(_range, "-")
	startRange, _ := strconv.ParseInt(splitedRange[0], 10, 64)
	stopRange, _ := strconv.ParseInt(splitedRange[1], 10, 64)
	if startRange <= data && data < stopRange {
		return true
	}
	return false
}

func calculateCost(slot customType.Slot, data int64) int64 {
	cost := data * slot.Fare
	if cost < slot.MinimumCost {
		return slot.MinimumCost
	}
	return cost
}

// RestCost reset Cost once billing done for entity
func RestCost() {
	Cost = 0
}
