package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	// "reflect"
)

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
	slot        string
	minimumCost int64
	fare        int64
}

var _partner string = ""
var _cost int64
var _capacities Capacities
var _deliveryMap map[string]Delivery = make(map[string]Delivery)

func parseInputFile(filePath string, theaterDetials Theatre) {
	f, _ := os.Open(filePath)
	defer f.Close()

	csvr := csv.NewReader(f)
	// csvr.Read()
	_, err := os.Stat("./result.csv")
	if err != nil {
		os.Create("./result.csv")
	}
	// file, _ := os.OpenFile("./result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	file, _ := os.Create("./result.csv")
	defer file.Close()
	writer := csv.NewWriter(file)

	for {
		records, err := csvr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		var theatre string = strings.TrimSpace(records[2])
		var operation string = strings.TrimSpace(records[0])
		data, _ := strconv.ParseInt(strings.TrimSpace(records[1]), 10, 64)
		cost, partner := getBilling(theaterDetials, theatre, data)
		billing := strconv.Itoa(int(cost))

		var value []string = []string{operation, "true", partner, billing}
		if cost == 0 {
			value = []string{operation, "false", "", ""}
		}
		writer.Write(value)
		_cost = 0
	}
	writer.Flush()

}

func extractOutputV1(filePath string, theaterDetials Theatre) {

	var deliveryList DeliveryList = parseInput(filePath)
	var exactDeliveryList DeliveryList = parseInput(filePath)
	fmt.Println(exactDeliveryList)
	sort.Sort(sort.Reverse(deliveryList))
	fmt.Println(exactDeliveryList)

	_, err := os.Stat("./result1.csv")
	if err != nil {
		os.Create("./result1.csv")
	}
	// file, _ := os.OpenFile("./result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	file, _ := os.Create("./result1.csv")
	defer file.Close()
	writer := csv.NewWriter(file)

	for _, record := range deliveryList {
		var theatre string = record.theatre
		var delivery string = record.delivery
		var data int64 = record.data
		cost, partner := getBillingV1(theaterDetials, theatre, data)
		if cost == 0 {
			partner = ""
		}
		_deliveryMap[delivery] = Delivery{delivery, data, theatre, cost, partner}
		_cost = 0
	}
	for _, delivery := range exactDeliveryList {
		_delivery := _deliveryMap[delivery.delivery]
		billing := strconv.Itoa(int(_delivery.price))
		partner := _delivery.partner
		var isPossible string = "true"
		if _delivery.price == 0 {
			isPossible = "false"
			billing = ""
			partner = ""
		}
		var value []string = []string{_delivery.delivery, isPossible, partner, billing}
		writer.Write(value)
	}
	writer.Flush()
}

func parseCapacitiesDetials(filePath string) Capacities {
	f, _ := os.Open(filePath)
	defer f.Close()

	csvr := csv.NewReader(f)
	csvr.Read()
	var capacities Capacities = make(Capacities)
	for {
		records, err := csvr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		var partner string = strings.TrimSpace(records[0])
		var limit string = strings.TrimSpace(records[1])
		capacities[partner], _ = strconv.ParseInt(limit, 10, 64)
	}
	return capacities
}

func updateTheatreDetials(file string) Theatre {
	f, _ := os.Open(file)
	defer f.Close()

	csvr := csv.NewReader(f)
	csvr.Read()
	sample := make(Theatre)

	for {
		records, err := csvr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		partner, found := sample[strings.TrimSpace(records[0])]
		minimumCost, _ := strconv.ParseInt(strings.TrimSpace(records[2]), 10, 64)
		fare, _ := strconv.ParseInt(strings.TrimSpace(records[3]), 10, 64)
		if found {
			slot, slotFound := partner[strings.TrimSpace(records[4])]
			if slotFound {
				slot[strings.TrimSpace(records[1])] = Slot{
					slot:        strings.TrimSpace(records[1]),
					minimumCost: minimumCost,
					fare:        fare,
				}
			} else {
				partner[strings.TrimSpace(records[4])] = Slots{
					strings.TrimSpace(records[1]): Slot{
						slot:        strings.TrimSpace(records[1]),
						minimumCost: minimumCost,
						fare:        fare,
					},
				}
			}
		} else {
			sample[strings.TrimSpace(records[0])] = Partner{
				strings.TrimSpace(records[4]): Slots{
					strings.TrimSpace(records[1]): Slot{
						slot:        strings.TrimSpace(records[1]),
						minimumCost: minimumCost,
						fare:        fare,
					},
				},
			}
		}
	}
	return sample
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

func calculateCost(slot Slot, data int64) int64 {
	cost := data * slot.fare
	if cost < slot.minimumCost {
		return slot.minimumCost
	}
	return cost
}

func estimateFare(partner string, slots Slots, data int64) (int64, string) {
	slotList := reflect.ValueOf(slots).MapKeys()
	for _, slot := range slotList {
		if inRange(slot.String(), data) {
			cost := calculateCost(slots[slot.String()], data)
			if _cost == 0 {
				_cost = cost
				_partner = partner
			}
			if cost < _cost {
				_cost = cost
				_partner = partner
			}
		}
	}
	return _cost, _partner
}
func estimateFareByCapacities(partner string, slots Slots, data int64) (int64, string) {
	slotList := reflect.ValueOf(slots).MapKeys()
	for _, slot := range slotList {
		if inRange(slot.String(), data) {
			cost := calculateCost(slots[slot.String()], data)
			if _cost == 0 {
				_cost = cost
				_partner = partner
			}
			if cost < _cost {
				_cost = cost
				_partner = partner
			}
		}
	}
	return _cost, _partner
}

func estimateFareByCapacitiesV1(partner string, slots Slots, data int64) (int64, string) {
	slotList := reflect.ValueOf(slots).MapKeys()
	if _capacities[partner] < data {
		return 0, partner
	}
	for _, slot := range slotList {
		if inRange(slot.String(), data) {
			cost := calculateCost(slots[slot.String()], data)
			if _cost == 0 {
				_cost = cost
				_partner = partner
			}
			if cost < _cost {
				_cost = cost
				_partner = partner
			}
		}
	}
	_capacities[_partner] = _capacities[_partner] - data
	return _cost, _partner
}

func getBilling(theatreDetials Theatre, theatre string, data int64) (int64, string) {
	partners := theatreDetials[theatre]
	partnerKeys := reflect.ValueOf(partners).MapKeys()
	var _partner string
	for _, partner := range partnerKeys {
		slots := partners[partner.String()]
		_cost, _partner = estimateFare(partner.String(), slots, data)
	}
	return _cost, _partner
}

func getBillingV1(theatreDetials Theatre, theatre string, data int64) (int64, string) {
	partners := theatreDetials[theatre]
	partnerKeys := reflect.ValueOf(partners).MapKeys()
	var _partner string
	for _, partner := range partnerKeys {
		slots := partners[partner.String()]
		if _capacities[partner.String()] < data {
			continue
		}
		_cost, _partner = estimateFare(partner.String(), slots, data)
	}
	_capacities[_partner] = _capacities[_partner] - data
	return _cost, _partner
}

// Delivery test
type Delivery struct {
	delivery string
	data     int64
	theatre  string
	price    int64
	partner  string
}

// DeliveryList det
type DeliveryList []Delivery

func (p DeliveryList) Len() int           { return len(p) }
func (p DeliveryList) Less(i, j int) bool { return p[i].data < p[j].data }
func (p DeliveryList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func parseInput(filePath string) DeliveryList {
	f, _ := os.Open(filePath)
	defer f.Close()
	csvr := csv.NewReader(f)
	var deliveryList DeliveryList
	for {
		record, err := csvr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		var deliveryName string = strings.TrimSpace(record[0])
		data, _ := strconv.ParseInt(strings.TrimSpace(record[1]), 10, 64)
		var theatre string = strings.TrimSpace(record[2])
		deliveryList = append(deliveryList, Delivery{deliveryName, data, theatre, 0, ""})
	}
	return deliveryList
}

func main() {
	fmt.Println("hello world")
	const inputFilePath string = "./input.csv"
	const capcitiesFilePath string = "./capacities.csv"
	const parternsFilePath string = "./partners.csv"
	var theatreDetials Theatre = updateTheatreDetials(parternsFilePath)
	_capacities = parseCapacitiesDetials(capcitiesFilePath)
	parseInputFile(inputFilePath, theatreDetials)
	extractOutputV1(inputFilePath, theatreDetials)
	var deliveryList DeliveryList = parseInput(inputFilePath)
	sort.Sort(sort.Reverse(deliveryList))
	fmt.Println(_capacities)
}
