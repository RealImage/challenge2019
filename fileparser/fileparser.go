package fileparser

import (
	"encoding/csv"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/praveenpkg8/challenge2019/customtype"
	"github.com/praveenpkg8/challenge2019/estimator"

	customType "github.com/praveenpkg8/challenge2019/customtype"
)

// DeliveryMap memory hold
var DeliveryMap map[string]customtype.Delivery = make(map[string]customtype.Delivery)

// ParseInput to parse
func ParseInput(filePath string) customType.DeliveryList {
	f, _ := os.Open(filePath)
	defer f.Close()
	csvr := csv.NewReader(f)
	var deliveryList customType.DeliveryList
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
		deliveryList = append(deliveryList, customType.Delivery{
			Delivery: deliveryName,
			Data:     data,
			Theatre:  theatre,
			Price:    0,
			Partner:  ""})
	}
	return deliveryList
}

// GenerateOutputV1 for generating output for second cases
func GenerateOutputV1(filePath string, theaterDetials customtype.Theatre) {

	var deliveryList customtype.DeliveryList = ParseInput(filePath)
	exactDeliveryList := ParseInput(filePath)
	sort.Sort(sort.Reverse(deliveryList))
	file, _ := os.Create("./output2.csv")
	defer file.Close()
	writer := csv.NewWriter(file)

	for _, record := range deliveryList {
		var theatre string = record.Theatre
		var delivery string = record.Delivery
		var data int64 = record.Data
		cost, partner := estimator.GetBillingV1(theaterDetials, theatre, data)
		if cost == 0 {
			partner = ""
		}
		DeliveryMap[delivery] = customtype.Delivery{
			Delivery: delivery,
			Data:     data,
			Theatre:  theatre,
			Price:    cost,
			Partner:  partner,
		}

		estimator.RestCost()
	}
	for _, delivery := range exactDeliveryList {
		_delivery := DeliveryMap[delivery.Delivery]
		billing := strconv.Itoa(int(_delivery.Price))
		partner := _delivery.Partner
		var isPossible string = "true"
		if _delivery.Price == 0 {
			isPossible = "false"
			billing = ""
			partner = ""
		}
		var value []string = []string{_delivery.Delivery, isPossible, partner, billing}
		writer.Write(value)
	}
	writer.Flush()
}

// ParseCapacitiesDetials to parse capacities
func ParseCapacitiesDetials(filePath string) {
	f, _ := os.Open(filePath)
	defer f.Close()
	csvr := csv.NewReader(f)
	csvr.Read()
	var capacities customType.Capacities = make(customType.Capacities)
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
	estimator.Capacities = capacities
}

// GenerateOutput for first case
func GenerateOutput(filePath string, theaterDetials customType.Theatre) {
	f, _ := os.Open(filePath)
	defer f.Close()
	csvr := csv.NewReader(f)
	file, _ := os.Create("./output1.csv")
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
		cost, partner := estimator.GetBilling(theaterDetials, theatre, data)
		billing := strconv.Itoa(int(cost))

		var value []string = []string{operation, "true", partner, billing}
		if cost == 0 {
			value = []string{operation, "false", "", ""}
		}
		writer.Write(value)
		estimator.RestCost()
	}
	writer.Flush()

}

// LoadTheatreDetials get detials
func LoadTheatreDetials(file string) customType.Theatre {
	f, _ := os.Open(file)
	defer f.Close()

	csvr := csv.NewReader(f)
	csvr.Read()
	theaterRecord := make(customType.Theatre)

	for {
		records, err := csvr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		partner, found := theaterRecord[strings.TrimSpace(records[0])]
		minimumCost, _ := strconv.ParseInt(strings.TrimSpace(records[2]), 10, 64)
		fare, _ := strconv.ParseInt(strings.TrimSpace(records[3]), 10, 64)

		if found {
			slot, slotFound := partner[strings.TrimSpace(records[4])]
			if slotFound {
				slot[strings.TrimSpace(records[1])] = customType.Slot{
					Slot:        strings.TrimSpace(records[1]),
					MinimumCost: minimumCost,
					Fare:        fare,
				}
			} else {
				partner[strings.TrimSpace(records[4])] = customType.Slots{
					strings.TrimSpace(records[1]): customType.Slot{
						Slot:        strings.TrimSpace(records[1]),
						MinimumCost: minimumCost,
						Fare:        fare,
					},
				}
			}
		} else {
			theaterRecord[strings.TrimSpace(records[0])] = customType.Partner{
				strings.TrimSpace(records[4]): customType.Slots{
					strings.TrimSpace(records[1]): customType.Slot{
						Slot:        strings.TrimSpace(records[1]),
						MinimumCost: minimumCost,
						Fare:        fare,
					},
				},
			}
		}
	}
	return theaterRecord
}
