package main

import (
	"encoding/csv"
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
