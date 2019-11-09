package Prob

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func FetchInputFromCSV(filename string) []DeliveryInfo {
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var data []DeliveryInfoStr
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		data = append(data, DeliveryInfoStr{
			DeliveryID:   line[0],
			DeliverySize: line[1],
			Theatre:      line[2],
		},
		)
	}

	fmt.Println("INPUT :", data)
	FormattedData := ConvertToDeliveryInfo(data)
	return FormattedData
}

func ConvertToDeliveryInfo(data []DeliveryInfoStr) []DeliveryInfo {
	newArr := []DeliveryInfo{}
	for _, v := range data {
		newArr = append(newArr, DeliveryInfo{
			DeliveryID:   v.DeliveryID,
			DeliverySize: ConvertToInt(v.DeliverySize),
			Theatre:      v.Theatre})
	}
	return newArr
}
