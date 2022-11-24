package service

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readCSVFile(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("error opening %s file: "+err.Error(), path))
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(fmt.Sprintf("error reading the %s file: ", path) + err.Error())
	}
	return records
}

func writeCSVFile(path string, records [][]string) bool {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("error opening/creating %s file: ", path) + err.Error())
		return false
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	err = writer.WriteAll(records)
	if err != nil {
		log.Fatal(fmt.Sprintf("error writing to %s file: ", path) + err.Error())
		return false
	}
	return true
}
