package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type EssentialData struct {
	TheatreID string
	SizeSlab  string
	MinCost   int
	CostPerGB int
	PartnerID string
}
type InputDetails struct {
	DeliveryID string
	Size       int
	TheatreID  string
}

func main() {
	mydir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	database, err := ParseDatabase(fmt.Sprintf("%s/%s", mydir, "pretty.csv"))
	if err != nil {
		log.Printf("error reading csv file %w", err)
		return
	}
	inputData, err := ReadInputData(fmt.Sprintf("%s/%s", mydir, "input.csv"))
	if err != nil {
		log.Printf("error reading input csv file %w", err)
		return
	}
	result := FindValidPartner(database, inputData)
	WriteOutput(result)
}

func WriteOutput(output [][]string) {
	csvFile, err := os.Create("output.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range output {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()
}

// FindValidPartner :- function to generate effective partners of corrsponding theaters
func FindValidPartner(database []EssentialData, inputData []InputDetails) (outputList [][]string) {
	outputList = append(outputList, []string{"delivery_id", "delivery_possible", "cost_of_delivery", "partner_id"})
	for _, input := range inputData {
		var (
			validCost int
		)

		activeIndex := -1
		for index, dataRow := range database {
			sizes := strings.Split(dataRow.SizeSlab, "-")
			low, _ := strconv.Atoi(sizes[0])
			high, _ := strconv.Atoi(sizes[1])
			if low <= input.Size && high >= input.Size && dataRow.TheatreID == input.TheatreID {
				deliveryCost := input.Size * dataRow.CostPerGB
				if validCost > deliveryCost || validCost == 0 {
					if deliveryCost < dataRow.MinCost {
						validCost = dataRow.MinCost
					} else {
						validCost = deliveryCost
					}

					activeIndex = index
				}
			}
		}
		if activeIndex != -1 {
			outputList = append(outputList, []string{input.DeliveryID, "true", fmt.Sprintf("%d", validCost), database[activeIndex].PartnerID})
		} else {
			outputList = append(outputList, []string{input.DeliveryID, "false"})
		}

	}
	return
}

// ReadDataFromInput :- read headers and conents from csv
func ReadDataFromInput(fileName string) ([]map[string]string, error) {
	var inputData []map[string]string
	csvfile, err := os.Open(fileName)
	if err != nil {
		log.Printf("error reading csv file due to error %w", err)
		return inputData, err
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Printf("error reading row data due to error %w", err)
		return inputData, err
	}
	header := []string{}
	for lineNum, line := range rawCSVdata {
		if lineNum == 0 {
			for _, title := range line {
				header = append(header, strings.TrimSpace(title))
			}
		} else {
			rowValues := map[string]string{}
			for i, value := range line {
				rowValues[header[i]] = value
			}
			inputData = append(inputData, rowValues)
		}
	}
	return inputData, nil
}

// ParseDatabase :- extract the contents from csv and parse it to the variable
func ParseDatabase(fileName string) ([]EssentialData, error) {
	var inputData []EssentialData
	csvfile, err := os.Open(fileName)
	if err != nil {
		log.Printf("error reading csv file due to error %w", err)
		return inputData, err
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Printf("error reading row data due to error %w", err)
		return inputData, err
	}

	for lineNum, line := range rawCSVdata {
		if lineNum > 0 {
			cost, _ := strconv.Atoi(line[2])
			costPerGB, _ := strconv.Atoi(line[3])
			rowValues := EssentialData{
				TheatreID: line[0],
				SizeSlab:  line[1],
				MinCost:   cost,
				CostPerGB: costPerGB,
				PartnerID: line[4],
			}
			inputData = append(inputData, rowValues)
		}
	}
	return inputData, nil

}

//ReadInputData :- function to read input csv file
func ReadInputData(fileName string) ([]InputDetails, error) {
	var inputData []InputDetails
	csvfile, err := os.Open(fileName)
	if err != nil {
		log.Printf("error reading csv file due to error %w", err)
		return inputData, err
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Printf("error reading row data due to error %w", err)
		return inputData, err
	}

	for lineNum, line := range rawCSVdata {
		if lineNum > 0 {
			size, _ := strconv.Atoi(strings.TrimSpace((line[1])))
			rowValues := InputDetails{
				DeliveryID: strings.TrimSpace(line[0]),
				Size:       size,
				TheatreID:  strings.TrimSpace(line[2]),
			}
			inputData = append(inputData, rowValues)
		}
	}
	return inputData, nil
}
