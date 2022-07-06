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
	fmt.Scanln("%s", output)
	return output
}

func GetAbsPaths(path string) (output string) {
	if path == "" {
		return path
	}
	path, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	return path
}

func StringToInt(input string) (result int) {
	result, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func ReadCsv(fileName string) (records [][]string) {
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
