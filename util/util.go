package util

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func GetInput(input string) (output string) {
	output = ""
	fmt.Println("Enter the input for", input)
	fmt.Scan(&output)
	return output
}

func GetAbsPaths(path string) (output string) {
	if path == "" {
		return path
	}
	path, err := filepath.Abs(path)
	if err != nil {
		log.Fatal("error while getting absoulte filepath: ", err)
	}
	return path
}

func StringToInt(input string) (result int) {
	result, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err, " for input ", input)
	}
	return result
}

func ReadCSV(fileName string) (records [][]string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	records, err = reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

func WriteCSV(fileName string, records [][]string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(file)
	err = writer.WriteAll(records)
	if err != nil {
		log.Fatal(err)
	}
}
