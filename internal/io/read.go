package io

import (
	"encoding/csv"
	"github.com/sureshn04/challenge2019/models"
	"log"
	"os"
	"strconv"

	"github.com/gocarina/gocsv"
)

// ReadPartners GetPartners reads the data out of a csv file
func ReadPartners(filename string) []*models.Partners {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	partnersFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer partnersFile.Close()

	var partners []*models.Partners

	if err := gocsv.UnmarshalFile(partnersFile, &partners); err != nil { // Load clients from file
		panic(err)
	}

	return partners
}

func ReadInput(fileName string) []*models.Input {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}

	var inputRecord []*models.Input

	for _, record := range records {
		cost, err := strconv.Atoi(record[1])
		if err != nil {
			panic(err)
		}
		data := &models.Input{
			DistributorID: record[0],
			Cost:          cost,
			TheatreID:     record[2],
		}
		inputRecord = append(inputRecord, data)
	}
	return inputRecord
}
