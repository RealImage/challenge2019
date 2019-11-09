package FetchInput

import (
	"bufio"
	"challenge2019/Prob/types"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

func FetchPartnerDataFromCSV(filename string) []types.PartnerData {
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var data []types.PartnerDataStr
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		data = append(data, types.PartnerDataStr{
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

func RemoveArrayElement(arr []types.PartnerDataStr, index int) []types.PartnerDataStr {
	return append(arr[:index], arr[index+1:]...)
}

func ConvertToPartnerData(arr []types.PartnerDataStr) (newArr []types.PartnerData) {
	for _, r := range arr {
		res := types.SplitString(r.Size)
		theatre := strings.TrimSpace(r.Theatre)
		newArr = append(newArr, types.PartnerData{
			Theatre:   theatre,
			Size:      types.SizeSlab{Min: types.ConvertToInt(res[0]), Max: types.ConvertToInt(res[1])},
			MinCost:   types.ConvertToFloat(r.MinCost),
			CostPerGB: types.ConvertToFloat(r.CostPerGB),
			PartnerID: r.PartnerID,
		})
	}
	return
}
