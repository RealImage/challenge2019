package Prob

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func FetchPartnerCapacityFromCSV(filename string) []CapacityInfo {
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var data []CapacityInfo
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		data = append(data, CapacityInfo{
			PartnerID: line[0],
			Capacity:  line[1],
		},
		)
	}

	//Remove Column Titles
	formattedCapacityInfo := append(data[:0], data[0+1:]...)

	return formattedCapacityInfo
}
