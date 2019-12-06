package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

var partners []Parterns
var capacities []Capacities
var input []Input
var output []Output

func main() {
	readPartnersData()
	readCapacityData()
	readInput()
	processProblemStmt1()
	processProblemStmt2()
}

/*Reads the actual partner data */
func readPartnersData() {
	csvfile, _ := os.Open("partners.csv")
	reader := csv.NewReader(csvfile)
	for {
		if property, err := reader.Read(); err == nil {
			data := Parterns{strings.TrimSpace(property[0]),
				strings.TrimSpace(property[1]),
				IntValue(property[2]),
				IntValue(property[3]),
				strings.TrimSpace(property[4]),
			}
			partners = append(partners, data)
		} else {
			break
		}
	}
}

/*Reads the actual capacity data */
func readCapacityData() {
	csvfile, _ := os.Open("capacities.csv")
	reader := csv.NewReader(csvfile)
	for {
		if property, err := reader.Read(); err == nil {
			data := Capacities{
				strings.TrimSpace(property[0]),
				IntValue(property[1]),
			}
			capacities = append(capacities, data)
		} else {
			break
		}
	}
}

/*Reads the input data */
func readInput() {
	csvfile, _ := os.Open("input.csv")
	reader := csv.NewReader(csvfile)

	for {
		if property, err := reader.Read(); err == nil {
			data := Input{strings.TrimSpace(property[0]),
				IntValue(property[1]),
				strings.TrimSpace(property[2]),
			}
			input = append(input, data)
		} else {
			break
		}
	}
}

/*Reads the input data */
func readOutput(fileName string) {
	csvfile, _ := os.Open(fileName)
	reader := csv.NewReader(csvfile)
	for {
		if property, err := reader.Read(); err == nil {
			data := Output{strings.TrimSpace(property[0]),
				BooleanValue(property[1]),
				IntValue(property[2]),
				strings.TrimSpace(property[3]),
			}
			output = append(output, data)
		} else {
			break
		}
	}
}

/*Problem statement 1*/
func processProblemStmt1() {
	csvfile, _ := os.Create("myoutput1.csv")
	for _, pInput := range input {
		var minimumCost = 0
		var partnerId = ""
		for _, pPartner := range partners {
			if pInput.Theatre == pPartner.Theatre && IsSizeUnderSlab(pInput.Size, pPartner.SlabSize) {
				cost := pInput.Size * pPartner.CostPerGB
				if cost < pPartner.MinimumCost {
					cost = pPartner.MinimumCost
				}
				if minimumCost == 0 || cost < minimumCost {
					minimumCost = cost
					partnerId = pPartner.PartnerID
				}
			}
		}
		writeOutputFile(csvfile, pInput.DeliveryId, strconv.FormatBool(minimumCost > 0), partnerId, minimumCost)
	}
	csvfile.Close()
}

/*Problem statement 2*/
func processProblemStmt2() {
	readOutput("myoutput1.csv")
	csvfile, _ := os.Create("myoutput2.csv")
	var actualCapacity int
	var partnerId string
	var totalSizeToBeDelivered = getTotalSizeToBeDelivered()
	lowestCost := totalSizeToBeDelivered
	for _, pCapacity := range capacities {
		cost := totalSizeToBeDelivered - pCapacity.Capacity
		if cost < lowestCost {
			lowestCost = cost
			partnerId = pCapacity.PartnerId
			actualCapacity = pCapacity.Capacity
		}
	}
	totalMinCost := 0
	for _, pInput := range input {
		minimumCost := 0
		capacity := actualCapacity
		for _, pPartner := range partners {
			if (partnerId == pPartner.PartnerID && pInput.Theatre == pPartner.Theatre) && IsSizeUnderSlab(pInput.Size, pPartner.SlabSize) && actualCapacity >= capacity {
				cost := pInput.Size * pPartner.CostPerGB
				if cost < pPartner.MinimumCost {
					cost = pPartner.MinimumCost
				}
				if minimumCost == 0 || cost < minimumCost {
					minimumCost = cost
					partnerId = pPartner.PartnerID
				}
				capacity = capacity - minimumCost
				totalMinCost = totalMinCost + minimumCost
				writeOutputFile(csvfile, pInput.DeliveryId, strconv.FormatBool(minimumCost > 0), partnerId, minimumCost)
			}
		}
	}
	csvfile.Close()
}

/*Writes the output to corresponding files*/
func writeOutputFile(file *os.File, deliveryID string, isDeliverable string, partnerID string, cost int) {
	csvwriter := csv.NewWriter(file)
	var costString = ""
	if cost > 0 {
		costString = strconv.Itoa(cost)
	}
	record := []string{
		deliveryID,
		isDeliverable,
		partnerID,
		costString,
	}
	csvwriter.Write(record)
	csvwriter.Flush()
}

func getTotalSizeToBeDelivered() int {
	var totalSizeToBeDelivered int
	var count int
	for _, pOutput := range output {
		if pOutput.IsDeliverable {
			totalSizeToBeDelivered = totalSizeToBeDelivered + input[count].Size
		}
		count++
	}
	return totalSizeToBeDelivered
}

func IsSizeUnderSlab(size int, slabSize string) bool {
	sizeRange := strings.Split(slabSize, "-")
	return size >= IntValue(sizeRange[0]) && size <= IntValue(sizeRange[1])
}

/* Converts string to int */
func IntValue(stringValue string) int {
	stringValue = strings.TrimSpace(stringValue)
	i, err := strconv.Atoi(stringValue)
	if err == nil {
		return i
	}
	return 0
}

/* Converts string to int */
func BooleanValue(stringValue string) bool {
	stringValue = strings.TrimSpace(stringValue)
	i, err := strconv.ParseBool(stringValue)
	if err == nil {
		return i
	}
	return false
}
