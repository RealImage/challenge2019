package FetchInput

import (
	"bufio"
	"challenge2019/Prob/types"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func FetchInputFromCSV(filename string) []types.DeliveryInfo {
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var data []types.DeliveryInfoStr
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		data = append(data, types.DeliveryInfoStr{
			DeliveryID:   line[0],
			DeliverySize: line[1],
			Theatre:      line[2],
		},
		)
	}

	FormattedData := ConvertToDeliveryInfo(data)
	return FormattedData
}

func ConvertToDeliveryInfo(data []types.DeliveryInfoStr) []types.DeliveryInfo {
	newArr := []types.DeliveryInfo{}
	for _, v := range data {
		newArr = append(newArr, types.DeliveryInfo{
			DeliveryID:   v.DeliveryID,
			DeliverySize: types.ConvertToInt(v.DeliverySize),
			Theatre:      v.Theatre})
	}
	return newArr
}
