package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	partnerRecords := readCsvFile("partners.csv")
	inputRecords := readCsvFile("input.csv")
	output1Records := delivery(inputRecords, partnerRecords)
	writeCsvFile(output1Records)
	capacityRecords := readCsvFile("capacities.csv")
	output2Records := deliveryCapacity(inputRecords, partnerRecords, capacityRecords)
	writeCsvFile(output2Records)
}

var count int = 1
var totalCost int
var finalCombination []string

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
					prevCost, _ = strconv.Atoi(array[3]) //storing the calculated cost to previous cost to compare with other partners
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

var capMap = make(map[string]int)

func deliveryCapacity(inputRecords [][]string, partnerRecords [][]string, capacityRecords [][]string) [][]string { ///method for problem statement-2
	var output [][]string
	var mapp = make(map[string][]int)
	totalCapacity, combinationCount := getCapacity(inputRecords)
	var actualCapacity int
	var partnerId string
	lowestCost := totalCapacity
	for i, pCapacity := range capacityRecords {
		if i != 0 {
			capacityinGB, _ := strconv.Atoi(pCapacity[1])
			cost := totalCapacity - capacityinGB
			if cost < lowestCost {
				lowestCost = cost
				partnerId = strings.Trim(pCapacity[0], " ")
				actualCapacity = capacityinGB
			}
		}
	}
	totalMinCost := 0
	partners := getPartners(inputRecords, partnerRecords)
	pArray := getRecordsByPartner(partnerRecords, inputRecords, partnerId, actualCapacity, totalMinCost, &mapp)
	//totalCost := 0
	for _, array := range pArray {
		tCost, _ := strconv.Atoi(array[3])
		totalCost += tCost
	}
	//fmt.Println(totalCost)
	for _, pId := range partners {
		for i, pCapacity := range capacityRecords {
			if i != 0 {
				if strings.Trim(pCapacity[0], " ") == pId {
					actualCapacity, _ = strconv.Atoi(pCapacity[1])
					capMap[pId] = actualCapacity
				}
			}
		}
		if pId != partnerId {
			_ = getRecordsByPartner(partnerRecords, inputRecords, pId, actualCapacity, totalMinCost, &mapp)
		}

	}
	//fmt.Println(mapp)
	//fmt.Println(pArray)

	checkPermutation(partners, inputRecords, partnerRecords, capacityRecords, mapp, totalCapacity, combinationCount) //compare total cost with all combinations
	//fmt.Println(finalCombination)
	for index, partner := range finalCombination {
		pArray[index][2] = partner
		pArray[index][3] = strconv.Itoa(mapp[partner][index])
	}
	output = pArray // assign the partners with minimum totalcost to output
	return output
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

func getRecordsByPartner(partnerRecords [][]string, inputRecords [][]string, pId string, actualCapacity int, totalMinCost int, mapp *map[string][]int) [][]string {
	var pArray [][]string
	for _, pInput := range inputRecords {
		minimumCost := 0
		capacity := actualCapacity
		array := []string{"", "", "", ""}
		array[0] = pInput[0]
		for i, pPartner := range partnerRecords {
			if i != 0 {
				pPartner[1] = strings.Trim(pPartner[1], " ")
				slabSize, _ := strconv.Atoi(pInput[1])
				if (pId == pPartner[4] && pInput[2] == pPartner[0]) && checkSize(slabSize, pPartner[1]) && actualCapacity >= capacity {
					costPerGB, _ := strconv.Atoi(pPartner[3])
					cost := slabSize * costPerGB
					minCost, _ := strconv.Atoi(pPartner[2])
					if cost < minCost {
						cost = minCost
					}
					if minimumCost == 0 || cost < minimumCost {
						minimumCost = cost
						pId = pPartner[4]
					}
					capacity = capacity - minimumCost
					totalMinCost = totalMinCost + minimumCost
					array[1] = "true"
					array[2] = pId
					array[3] = strconv.Itoa(minimumCost)
				}
			}
		}
		if array[1] != "true" {
			array[1] = "false"
		}
		(*mapp)[pId] = append((*mapp)[pId], minimumCost)
		pArray = append(pArray, array)
	}
	return pArray
}

func getCapacity(inputRecords [][]string) (int, int) { // to get the partners eligible to deliver
	totalCapacity := 0
	combinationCount := 0
	output1Records := readCsvFile("myOutput1.csv")
	for index, pOutput := range output1Records {
		if pOutput[1] == "true" {
			data, _ := strconv.Atoi(inputRecords[index][1])
			totalCapacity += data
			combinationCount++
		}
	}
	return totalCapacity, combinationCount
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

func checkPermutation(set []string, inputRecords [][]string, partnerRecords [][]string, capacityRecords [][]string, mapp map[string][]int, totalCapacity int, combinationCount int) {
	checkPermutationRec(set, "", len(set), combinationCount, inputRecords, partnerRecords, capacityRecords, mapp, totalCapacity)
}

func checkPermutationRec(set []string, prefix string, n int, k int, inputRecords [][]string, partnerRecords [][]string, capacityRecords [][]string, mapp map[string][]int, totalCapacity int) {

	if k == 0 {
		for i, pCapacity := range capacityRecords {
			if i != 0 {
				actualCapacity, _ := strconv.Atoi(pCapacity[1])
				capMap[strings.Trim(pCapacity[0], " ")] = actualCapacity
			}
		}
		var arr []string
		cost := 0
		isCapable := true
		for i := 0; i < len(prefix); i = i + 2 {
			arr = append(arr, prefix[i:i+2])
		}
		for index, arrValue := range arr {
			cost += mapp[arrValue][index]
			tempCap, _ := strconv.Atoi(inputRecords[index][1])
			capMap[arrValue] = capMap[arrValue] - tempCap
			if mapp[arrValue][index] == 0 || capMap[arrValue] < 0 {
				isCapable = false
			}
		}
		if cost < totalCost && isCapable != false {
			finalCombination = arr
			totalCost = cost
		}
		return
	}

	for i := 0; i < n; i = i + 1 {

		newPrefix := prefix + set[i]

		checkPermutationRec(set, newPrefix, n, k-1, inputRecords, partnerRecords, capacityRecords, mapp, totalCapacity)
	}
	return
}
