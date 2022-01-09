package services

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/niroopreddym/realimage/models"
)

//ReadCSVRecords reads the data out of a csv file
func ReadCSVRecordsPartners(filename string) []*models.Partners {
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

	partners := []*models.Partners{}

	if err := gocsv.UnmarshalFile(partnersFile, &partners); err != nil { // Load clients from file
		panic(err)
	}

	return partners
}

//ReadCSVRecordsInputs reads the data out of a csv file
func ReadCSVRecordsInputs(filename string) []*models.Input {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	inptFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer inptFile.Close()

	inputs := []*models.Input{}

	if err := gocsv.UnmarshalFile(inptFile, &inputs); err != nil { // Load clients from file
		panic(err)
	}

	return inputs
}

//ReadCSVRecordsCapacity Reads a csv and returns capacity struct
func ReadCSVRecordsCapacity(filename string) []*models.Capacity {
	inputFile := readCSVFile(filename)
	defer inputFile.Close()

	capacities := []*models.Capacity{}

	if err := gocsv.UnmarshalFile(inputFile, &capacities); err != nil { // Load clients from file
		panic(err)
	}

	return capacities
}

func readCSVFile(filename string) *os.File {
	inptFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	return inptFile
}

//WriteDataToCSV writes the data to CSV file
func WriteDataToCSV(filename string, output []models.Output) {
	file, _ := os.Create(filename)
	defer file.Close()

	err := gocsv.MarshalFile(&output, file) // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}
}
