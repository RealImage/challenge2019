package Prob

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

func FetchPartnerDataFromCSV(filename string) []PartnerData {
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var data []PartnerDataStr
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		data = append(data, PartnerDataStr{
			Theatre:   line[0],
			Size:      line[1],
			MinCost:   line[2],
			CostPerGB: line[3],
			PartnerID: line[4],
		},
		)
	}

	//Remove Column Titles
	data = RemoveArrayElement(data, 0)

	FormattedData := ConvertToPartnerData(data)
	return FormattedData
}

func RemoveArrayElement(arr []PartnerDataStr, index int) []PartnerDataStr {
	return append(arr[:index], arr[index+1:]...)
}

func ConvertToPartnerData(arr []PartnerDataStr) (newArr []PartnerData) {
	for _, r := range arr {
		res := SplitString(r.Size)
		theatre := strings.TrimSpace(r.Theatre)
		newArr = append(newArr, PartnerData{
			Theatre:   theatre,
			Size:      SizeSlab{Min: ConvertToInt(res[0]), Max: ConvertToInt(res[1])},
			MinCost:   ConvertToFloat(r.MinCost),
			CostPerGB: ConvertToFloat(r.CostPerGB),
			PartnerID: r.PartnerID,
		})
	}
	return
}

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
