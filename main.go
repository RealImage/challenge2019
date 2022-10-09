package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type partners struct {
	ColNameIndexMap map[string]int
}

type PartnersData struct {
	Theatre   string
	SizeSlab  SizeSlabRange
	MiniCost  int
	CostPerGB int
	partnerID string
}

type SizeSlabRange struct {
	min int
	max int
}

func main() {
	p := partners{}
	patnersData := p.getPartnersData()
	outputDetails := findPartnerWithMinimumCost(patnersData)
	createOutputCSV(outputDetails)
}

func (p *partners) getPartnersData() (partnersData []PartnersData) {
	csvReader, f := readFile("./partners.csv")
	defer f.Close()
	rows, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
	}

	for rowNumber, row := range rows {
		if rowNumber != 0 {
			var partnerData PartnersData
			partnerData.Theatre = strings.TrimSpace(GetStringValueCSV(p.ColNameIndexMap, "Theatre", row))
			slabRange := GetStringValueCSV(p.ColNameIndexMap, "Size Slab (in GB)", row)
			partnerData.SizeSlab.min, partnerData.SizeSlab.max = getMinMaxRange(slabRange)
			partnerData.MiniCost, _ = strconv.Atoi(strings.TrimSpace(GetStringValueCSV(p.ColNameIndexMap, "Minimum cost", row)))
			partnerData.CostPerGB, _ = strconv.Atoi(strings.TrimSpace(GetStringValueCSV(p.ColNameIndexMap, "Cost Per GB", row)))
			partnerData.partnerID = strings.TrimSpace(GetStringValueCSV(p.ColNameIndexMap, "Partner ID", row))
			partnersData = append(partnersData, partnerData)

		} else {
			if len(p.ColNameIndexMap) == 0 {
				p.ColNameIndexMap = make(map[string]int)
			}
			for index, col := range row {
				val := strings.TrimSpace(col)
				p.ColNameIndexMap[val] = index
			}
		}

	}
	return partnersData

}
func readFile(path string) (*csv.Reader, *os.File) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("unable to open the file")
	}
	csvReader := csv.NewReader(f)
	return csvReader, f
}

func findPartnerWithMinimumCost(partnerDetails []PartnersData) [][]string {

	csvReader, f := readFile("./input.csv")
	defer f.Close()

	var outputDetails [][]string
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		deliverySize, _ := strconv.Atoi(row[1])

		minimumCost, partnerID, deliveryPossible := getParnerWithMinimumCost(deliverySize, row[2], partnerDetails)
		var outputRow []string
		minimumCostStr := strconv.Itoa(minimumCost)
		if minimumCostStr == "0" {
			minimumCostStr = ""
		}
		outputRow = append(outputRow, row[0], strconv.FormatBool(deliveryPossible), partnerID, minimumCostStr)
		outputDetails = append(outputDetails, outputRow)

	}
	return outputDetails

}

func getParnerWithMinimumCost(deliverySize int, theatreID string, partnerDetails []PartnersData) (int, string, bool) {

	var minimumCost int
	var partnerID string
	deliveryPossible := false
	for _, partnerData := range partnerDetails {
		inRange := checkDeliverySizeIntheRange(partnerData.SizeSlab, deliverySize)
		if partnerData.Theatre == theatreID && inRange {
			deliveryPossible = true
			cost := deliverySize * partnerData.CostPerGB
			if cost <= partnerData.MiniCost {
				cost = partnerData.MiniCost
			}
			if cost < minimumCost || minimumCost == 0 {
				minimumCost = cost
				partnerID = partnerData.partnerID
			}
		}
	}

	return minimumCost, partnerID, deliveryPossible
}

func checkDeliverySizeIntheRange(sizeSlab SizeSlabRange, deliverySize int) bool {
	return deliverySize >= sizeSlab.min && deliverySize <= sizeSlab.max
}

func getMinMaxRange(slabRange string) (min int, max int) {
	rangeArr := strings.Split(slabRange, "-")
	min, _ = strconv.Atoi(strings.TrimSpace(rangeArr[0]))
	max, _ = strconv.Atoi(strings.TrimSpace(rangeArr[1]))
	return min, max
}

func GetStringValueCSV(colNameIndexMap map[string]int, colName string, row []string) string {
	var returnVal string
	if val, ok := colNameIndexMap[colName]; ok {
		if len(row) > val {
			returnVal = row[val]
		}
	}
	return returnVal
}

func createOutputCSV(outputDetails [][]string) {
	csvFile, err := os.Create("output.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, row := range outputDetails {
		_ = csvwriter.Write(row)
	}

	csvwriter.Flush()
	csvFile.Close()
}
