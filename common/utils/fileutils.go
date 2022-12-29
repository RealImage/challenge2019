package utils

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
)

// Reads CSV file from given filepath
func ReadCsv(filepath string) ([][]string, error) {
	csvFile, err := os.Open(filepath)
	if err != nil {
		log.Print("Unable to read csv file ", err)
		return nil, err
	}

	defer csvFile.Close()

	rows, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Print("Unable to parse csv file ", err)
		return nil, err
	}

	return rows, nil
}

// Writes content to given filepath
func WriteToCsv(filepath string, content []string) error {
	outputCsv, err := os.Create(filepath)
	if err != nil {
		log.Print("Unable to create output file ", err)
		return err
	}

	defer outputCsv.Close()

	outputWriter := bufio.NewWriter(outputCsv)
	defer outputWriter.Flush()
	for _, line := range content {
		outputWriter.WriteString(line + "\n")
	}

	return nil
}
