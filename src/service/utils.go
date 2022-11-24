package service

import (
	"challenge2019/src/models"
	"strconv"
	"strings"
)

func ReadPartnerCsv(path string) []models.PartnerRecord {
	partnerRecords := readCSVFile(path)
	partnersList := make([]models.PartnerRecord, 0, len(partnerRecords)-1)
	for index, record := range partnerRecords {
		if index == 0 {
			continue
		}
		partner := models.NewPartnerRecord(record)
		partnersList = append(partnersList, partner)
	}
	return partnersList
}

func ReadInput(path string) []models.Input {
	inputRecords := readCSVFile(path)
	inputList := make([]models.Input, 0, len(inputRecords))
	for _, record := range inputRecords {
		input := models.NewInput(record)
		inputList = append(inputList, input)
	}
	return inputList
}

func WriteOutput(path string, outputList []models.Output) bool {
	stringfyOutputList := make([][]string, 0, len(outputList))
	for _, val := range outputList {
		stringfyOutputList = append(stringfyOutputList, val.String())
	}
	return writeCSVFile(path, stringfyOutputList)
}

func ReadCapacities(path string) map[string]int {
	capcitiesRecord := readCSVFile(path)
	capMap := make(map[string]int, 0)
	var partnerID string
	var capacity int
	for index, record := range capcitiesRecord {
		if index == 0 {
			continue
		}
		partnerID = strings.TrimSpace(record[0])
		capacity, _ = strconv.Atoi(strings.TrimSpace(record[1]))
		capMap[partnerID] = capacity
	}
	return capMap
}
