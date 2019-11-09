package FetchInput

import (
	"bufio"
	"challenge2019/Prob/types"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func FetchPartnerCapacityFromCSV(filename string) []types.CapacityDetailsStr {
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var data []types.CapacityDetailsStr
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		data = append(data, types.CapacityDetailsStr{
			PartnerID: line[0],
			Capacity:  line[1],
		},
		)
	}

	//Remove Column Titles
	formattedCapacityInfo := append(data[:0], data[0+1:]...)

	return formattedCapacityInfo
}
