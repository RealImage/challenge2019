package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readCsvFile(path string) [][]string {
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

func writeCsvFile(outputRecords [][]string) {
	file, err := os.Create("myOutput.csv")
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
}

func delivery(inputRecords [][]string, partnerRecords [][]string) [][]string {
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
			size := strings.Split(partnerRecords[index][1], "-")
			slabSize, _ := strconv.Atoi(inputRecords[value][1])
			min, _ := strconv.Atoi(size[0])
			max, _ := strconv.Atoi(size[1])
			if inputRecords[value][2] == partnerRecords[index][0] {
				if slabSize > min && slabSize <= max {
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
					prevCost, _ = strconv.Atoi(array[3])
				}
			}
		}
		output[value] = append(output[value], array[0])
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

/* func deliveryCapcity(inputRecords [][]string, partnerRecords [][]string, capacityRecords [][]string) [][]string {

} */

func main() {
	partnerRecords := readCsvFile("partners.csv")
	inputRecords := readCsvFile("input.csv")
	outputRecords := delivery(inputRecords, partnerRecords)
	writeCsvFile(outputRecords)
	//capacityRecords := readCsvFile("capacities.csv")
	//outputRecords := delivery(inputRecords, partnerRecords,capacityRecords)
}
