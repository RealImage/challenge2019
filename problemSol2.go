package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var count int = 1

func readCsvFile(path string) [][]string { //to read the csv file and return as array
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Unable to read input file "+path, err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Unable to parse file as CSV for "+path, err)
	}

	return records
}

func writeCsvFile(outputRecords [][]string) { //to write the output record into csv file
	file, err := os.Create("myOutput" + strconv.Itoa(count) + ".csv")
	if err != nil {
		fmt.Println("Unable to create CSV file", err)
	}
	defer file.Close()
	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()
	for _, data := range outputRecords {
		err := csvWriter.Write(data)
		if err != nil {
			fmt.Println("Unable to write CSV file ", err)
		}
	}
	count++
}

func delivery(inputRecords [][]string, partnerRecords [][]string) [][]string { //method for problem statement-1
	output := make([][]string, len(inputRecords))
	for value := range inputRecords {
		array := [4]string{"", "", "", ""}
		array[0] = inputRecords[value][0]
		prevCost := 0
		for index := 1; index < len(partnerRecords); index++ {
			partnerRecords[index][0] = strings.Trim(partnerRecords[index][0], " ")
			partnerRecords[index][1] = strings.Trim(partnerRecords[index][1], " ")
			partnerRecords[index][2] = strings.Trim(partnerRecords[index][2], " ")
			partnerRecords[index][3] = strings.Trim(partnerRecords[index][3], " ")
			inputRecords[value][2] = strings.Trim(inputRecords[value][2], " ")
			slabSize, _ := strconv.Atoi(inputRecords[value][1])
			if inputRecords[value][2] == partnerRecords[index][0] { //check the Theatre
				if checkSize(slabSize, partnerRecords[index][1]) { //comparing the total cost with the minimum cost
					array[1] = "true"
					cost, _ := strconv.Atoi(partnerRecords[index][3])
					minCost, _ := strconv.Atoi(partnerRecords[index][2])
					if (cost * slabSize) > minCost {
						if prevCost == 0 {
							array[3] = strconv.Itoa(cost * slabSize)
							array[2] = partnerRecords[index][4]
						} else if prevCost > cost*slabSize {
							array[3] = strconv.Itoa(cost * slabSize)
							array[2] = partnerRecords[index][4]
						}
					} else if prevCost == 0 || prevCost > minCost {
						array[3] = partnerRecords[index][2]
						array[2] = partnerRecords[index][4]
					}
					prevCost, _ = strconv.Atoi(array[3]) //storing the calculated cost to previous cost to campare with other partners
				}
			}
		}
		output[value] = append(output[value], array[0]) //appending the result array to the output records
		if array[1] != "true" {
			output[value] = append(output[value], "false")
			output[value] = append(output[value], "")
			output[value] = append(output[value], "")
		} else {
			output[value] = append(output[value], array[1])
			output[value] = append(output[value], array[2])
			output[value] = append(output[value], array[3])
		}

	}
	return output
}

func deliveryCapacity(inputRecords [][]string, partnerRecords [][]string, capacityRecords [][]string) [][]string { ///method for problem statement-2
	var output [][]string
	var capMap = make(map[string]int)
	for i, cap := range capacityRecords {
		if i != 0 {
			actualCapacity, _ := strconv.Atoi(cap[1])
			capMap[strings.Trim(cap[0], " ")] = actualCapacity
		}

	}
	output = readCsvFile("myOutput1.csv")
	partners := getPartners(inputRecords, partnerRecords)

	if checkCapacity(output, inputRecords, partnerRecords, capacityRecords, capMap) {
		return output
	} else {
		for index := 0; index < len(output); index++ {
			for _, pId := range partners {
				if output[0][2] != pId {
					output[0][2] = pId
					if checkCapacity(output, inputRecords, partnerRecords, capacityRecords, capMap) {
						return output
					}
				}
			}
		}
	}
	return output
}

func checkCapacity(output [][]string, inputRecords [][]string, partnerRecords [][]string, capacityRecords [][]string, capMap map[string]int) bool {
	for i, cap := range capacityRecords {
		if i != 0 {
			actualCapacity, _ := strconv.Atoi(cap[1])
			capMap[strings.Trim(cap[0], " ")] = actualCapacity
		}

	}
	for i, oRecord := range output {
		c, _ := strconv.Atoi(inputRecords[i][1])
		capMap[oRecord[2]] = capMap[oRecord[2]] - c
		if capMap[oRecord[2]] < 0 {
			return false
		}
	}
	return true
}

func getPartners(inputRecords [][]string, partnerRecords [][]string) []string {
	var partners []string
	for value := range inputRecords {
		for index := range partnerRecords {
			if index != 0 {
				inputRecords[value][2] = strings.Trim(inputRecords[value][2], " ")
				partnerRecords[index][0] = strings.Trim(partnerRecords[index][0], " ")
				if inputRecords[value][2] == partnerRecords[index][0] && !contains(partners, partnerRecords[index][4]) {
					partners = append(partners, partnerRecords[index][4])
				}
			}
		}
	}
	return partners
}

func getRecords(partnerRecords [][]string, array []string, capacity int, minimumCost int, pInput []string, actualCapacity int, partnerId string, totalMinCost int) ([]string, int) {
	for i, pPartner := range partnerRecords {
		if i != 0 {
			pPartner[1] = strings.Trim(pPartner[1], " ")
			slabSize, _ := strconv.Atoi(pInput[1])
			if (pInput[2] == pPartner[0]) && checkSize(slabSize, pPartner[1]) && actualCapacity >= capacity {
				costPerGB, _ := strconv.Atoi(pPartner[3])
				cost := slabSize * costPerGB
				minCost, _ := strconv.Atoi(pPartner[2])
				if cost < minCost {
					cost = minCost
				}
				if minimumCost == 0 || cost < minimumCost {
					minimumCost = cost
					partnerId = pPartner[4]
				}
				capacity = capacity - minimumCost
				totalMinCost = totalMinCost + minimumCost
				array[1] = "true"
				array[2] = partnerId
				array[3] = strconv.Itoa(minimumCost)
			}
		}
	}
	return array, capacity
}

func getCapacity(inputRecords [][]string) int { // to get the partners eligible to deliver
	totalCapacity := 0
	output1Records := readCsvFile("myOutput1.csv")
	for index, pOutput := range output1Records {
		if pOutput[1] == "true" {
			data, _ := strconv.Atoi(inputRecords[index][1])
			totalCapacity += data
		}
	}
	return totalCapacity
}

func contains(partners []string, str string) bool {
	for _, value := range partners {
		if value == str {
			return true
		}
	}
	return false
}

func checkSize(size int, slabSize string) bool {
	sizeRange := strings.Split(slabSize, "-")
	min, _ := strconv.Atoi(sizeRange[0])
	max, _ := strconv.Atoi(sizeRange[1])
	return size >= min && size <= max
}

func main() {
	partnerRecords := readCsvFile("partners.csv")
	inputRecords := readCsvFile("input.csv")
	output1Records := delivery(inputRecords, partnerRecords)
	writeCsvFile(output1Records)
	capacityRecords := readCsvFile("capacities.csv")
	output2Records := deliveryCapacity(inputRecords, partnerRecords, capacityRecords)
	fmt.Println(output2Records)
}
