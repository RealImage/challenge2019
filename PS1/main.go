package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	partnersInfo := readCSV("../partners.csv")
	partnersInfo = partnersInfo[1:]
	theatres := make(map[string]map[string][]*Slap, 0)
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
	deliveries := make([]*Delivery, 0)
	for _, inputInfo := range inputsInfo {
		delivery := &Delivery{
			ID:        strings.TrimSpace(inputInfo[0]),
			TheatreID: strings.TrimSpace(inputInfo[2]),
		}
		delivery.Size = toInt(strings.TrimSpace(inputInfo[1]))
		deliveries = append(deliveries, delivery)
	}

	for _, delivery := range deliveries {
		partnerID, partnerCost := "", 0
		theatre, ok := theatres[delivery.TheatreID]
		if ok {
			for k, slaps := range theatre {
				for _, slap := range slaps {
					if slap.MinSize < delivery.Size && delivery.Size <= slap.MaxSize {
						cost := max(slap.MinCost, slap.Cost*delivery.Size)
						if partnerCost == 0 || (cost < partnerCost) {
							partnerID = k
							partnerCost = cost
						}
					}
				}
			}
		}
		if partnerID != "" {
			fmt.Printf("%v,true,%v,%v\n", delivery.ID, partnerID, partnerCost)
		} else {
			fmt.Printf("%v,false,\"\",\"\"\n", delivery.ID)
		}
	}

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
