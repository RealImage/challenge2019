package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mohae/deepcopy"
)

var (
	theatres   = make(map[string]map[string][]*Slap, 0)
	deliveries = make([]*Delivery, 0)
	capacities = make(map[string]int, 0)
	solution   *Solution
)

func main() {
	partnersInfo := readCSV("../partners.csv")
	partnersInfo = partnersInfo[1:]
	for _, partnerInfo := range partnersInfo {
		theatreID := strings.TrimSpace(partnerInfo[0])
		theatre, ok := theatres[theatreID]
		if !ok {
			theatre = make(map[string][]*Slap, 0)
		}

		partnerID := strings.TrimSpace(partnerInfo[4])
		slaps, ok := theatre[partnerID]
		if !ok {
			slaps = make([]*Slap, 0)
		}

		slap := &Slap{}
		sizeSplit := strings.Split(strings.TrimSpace(partnerInfo[1]), "-")
		slap.MinSize = toInt(sizeSplit[0])
		slap.MaxSize = toInt(sizeSplit[1])
		slap.MinCost = toInt(strings.TrimSpace(partnerInfo[2]))
		slap.Cost = toInt(strings.TrimSpace(partnerInfo[3]))

		slaps = append(slaps, slap)
		theatre[partnerID] = slaps
		theatres[theatreID] = theatre
	}

	inputsInfo := readCSV("../input.csv")
	for _, inputInfo := range inputsInfo {
		delivery := &Delivery{
			ID:        strings.TrimSpace(inputInfo[0]),
			TheatreID: strings.TrimSpace(inputInfo[2]),
		}
		delivery.Size = toInt(strings.TrimSpace(inputInfo[1]))
		deliveries = append(deliveries, delivery)
	}

	capacitiesInfo := readCSV("../capacities.csv")
	capacitiesInfo = capacitiesInfo[1:]
	for _, capacityInfo := range capacitiesInfo {
		capacities[strings.TrimSpace(capacityInfo[0])] = toInt(strings.TrimSpace(capacityInfo[1]))
	}

	find(&Solution{Deliveries: make(map[string]*DeliverySolution, 0)},
		make(map[string]int, 0), 0)

	for _, delivery := range deliveries {
		deliverySolution := solution.Deliveries[delivery.ID]
		if deliverySolution.PartnerID != "" {
			fmt.Printf("%v,true,%v,%v\n", delivery.ID, deliverySolution.PartnerID, deliverySolution.Cost)
		} else {
			fmt.Printf("%v,false,\"\",\"\"\n", delivery.ID)
		}
	}

}

func find(tempSolution *Solution, partnerSize map[string]int, deliveryIndex int) {
	if len(tempSolution.Deliveries) < deliveryIndex {
		return
	}
	for i := deliveryIndex; i < len(deliveries); i++ {
		delivery := deliveries[i]
		theatre, ok := theatres[delivery.TheatreID]
		if ok {
			for k, slaps := range theatre {
				for _, slap := range slaps {
					if slap.MinSize < delivery.Size && delivery.Size <= slap.MaxSize {
						cost := max(slap.MinCost, slap.Cost*delivery.Size)
						size := partnerSize[k]
						capacity := capacities[k]
						if size+delivery.Size <= capacity {
							tempSolution.Deliveries[delivery.ID] = &DeliverySolution{PartnerID: k, Cost: cost}
							partnerSize[k] = size + delivery.Size
							find(tempSolution, partnerSize, i+1)
							delete(tempSolution.Deliveries, delivery.ID)
							partnerSize[k] = size
						}
					}
				}
			}
		}
		tempSolution.Deliveries[delivery.ID] = &DeliverySolution{PartnerID: ""}
		find(tempSolution, partnerSize, i+1)
		delete(tempSolution.Deliveries, delivery.ID)
	}
	if len(tempSolution.Deliveries) == len(deliveries) {
		if solution == nil {
			solution = deepcopy.Copy(tempSolution).(*Solution)
		} else {
			successfulDelivery, cost := getDeliveryCountAndCost(solution)
			tempSuccessfulDelivery, tempCost := getDeliveryCountAndCost(tempSolution)
			if tempSuccessfulDelivery > successfulDelivery ||
				(tempSuccessfulDelivery == successfulDelivery && tempCost < cost) {
				solution = deepcopy.Copy(tempSolution).(*Solution)
			}
		}
	}
}

func getDeliveryCountAndCost(solution *Solution) (int, int) {
	successfulDelivery, cost := 0, 0
	for _, delivery := range deliveries {
		deliverySolution := solution.Deliveries[delivery.ID]
		if deliverySolution.PartnerID != "" {
			successfulDelivery++
			cost += deliverySolution.Cost
		}
	}
	return successfulDelivery, cost
}

func readCSV(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	return data
}

func toInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return intValue
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
