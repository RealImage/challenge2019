package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readCsvFile(filePath string) [][]string {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Cannot open the file "+filePath, err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("Cannot parse CSV file "+filePath, err)
	}

	return data
}

func writeToCsvFile(filePath, fileName string, data [][]string, headers []string) error {
	file, err := os.Create(filePath + fileName + ".csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to create file", err)
		return err
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(headers)

	for row := range data {
		writer.Write(data[row])
	}
	return nil
}

func Solution1(inputRecords []*Input, partners map[string][]*TheatrePartnerRelation) error {

	for i := range inputRecords {
		partner, ok := partners[inputRecords[i].TheatreId]
		if !ok || len(partner) == 0 {
			log.Printf("No Partner for theatre Id :%s. Skipping this theatre\n", inputRecords[i].TheatreId)
			continue
		}

		var minimumCost int64 = -1
		var partnerId string

		for j := range partner {
			if inputRecords[i].DeliverySize >= partner[j].SizeSlabLowerLimit && inputRecords[i].DeliverySize <= partner[j].SizeSlabUpperLimit {
				var tempMinCost int64
				if partner[j].MinimumCost > inputRecords[i].DeliverySize*partner[j].CostPerGB {
					tempMinCost = partner[j].MinimumCost
				} else {
					tempMinCost = inputRecords[i].DeliverySize * partner[j].CostPerGB
				}
				if minimumCost == -1 || tempMinCost < minimumCost {
					minimumCost = tempMinCost
					partnerId = partner[j].PartnerId
				}
			}
		}
		fmt.Println(minimumCost, partnerId)
	}
	return nil
}

type Input struct {
	DeliveryPartnerId string
	DeliverySize      int64
	TheatreId         string
}

type TheatrePartnerRelation struct {
	PartnerId          string
	TheatreId          string
	SizeSlabLowerLimit int64
	SizeSlabUpperLimit int64
	MinimumCost        int64
	CostPerGB          int64
}

type DeliveryPartner struct {
	Id       string
	Capacity int64
}

func main() {
	inputFile := readCsvFile("./input.csv")
	var inputBody []*Input

	var err error = nil

	for i := range inputFile {
		inputRow := Input{}

		inputRow.DeliveryPartnerId = inputFile[i][0]

		inputRow.DeliverySize, err = strconv.ParseInt(inputFile[i][1], 10, 64)
		if err != nil {
			log.Println("not able to parse input size to int64 ", inputFile[i][1], err)
			return
		}

		inputRow.TheatreId = inputFile[i][2]

		inputBody = append(inputBody, &inputRow)
	}

	partnerFile := readCsvFile("./partners.csv")

	partnerMap := make(map[string][]*TheatrePartnerRelation)

	for i := range partnerFile {
		if i == 0 {
			continue
		}
		partner := TheatrePartnerRelation{}

		partner.TheatreId = partnerFile[i][0]

		slabSize := strings.Split(partnerFile[i][1], "-")

		if len(slabSize) != 2 {
			log.Println("Invalid slab size exiting program", partnerFile[i][1])
			return
		}

		partner.SizeSlabLowerLimit, err = strconv.ParseInt(slabSize[0], 10, 64)
		if err != nil {
			log.Println("not able to parse SizeSlabLowerLimit to int64 returning", strings.Trim(slabSize[0], " "), err)
			return
		}

		partner.SizeSlabUpperLimit, err = strconv.ParseInt(strings.Trim(slabSize[1], " "), 10, 64)
		if err != nil {
			log.Println("not able to parse SizeSlabUpperLimit to int64 returning", slabSize[1], err)
			return
		}

		partner.MinimumCost, err = strconv.ParseInt(strings.Trim(partnerFile[i][2], " "), 10, 64)
		if err != nil {
			log.Println("not able to parse MinimumCost to int64 returning", partnerFile[i][2], err)
			return
		}

		partner.CostPerGB, err = strconv.ParseInt(strings.Trim(partnerFile[i][3], " "), 10, 64)
		if err != nil {
			log.Println("not able to parse CostPerGB to int64 returning", partnerFile[i][3], err)
			return
		}

		partner.PartnerId = partnerFile[i][4]

		if _, ok := partnerMap[strings.Trim(partner.TheatreId, " ")]; !ok {
			partnerMap[strings.Trim(partner.TheatreId, " ")] = []*TheatrePartnerRelation{&partner}
		} else {
			partnerMap[strings.Trim(partner.TheatreId, " ")] = append(partnerMap[strings.Trim(partner.TheatreId, " ")], &partner)
		}

	}

	capacitieFile := readCsvFile("./capacities.csv")

	deliveryMap := make(map[string]*DeliveryPartner)

	for i := range capacitieFile {
		if i == 0 {
			continue
		}
		var delivery DeliveryPartner

		delivery.Id = capacitieFile[i][0]

		delivery.Capacity, err = strconv.ParseInt(strings.Trim(capacitieFile[i][1], " "), 10, 64)
		if err != nil {
			log.Println("not able to parse  Capacity to int64 returning", capacitieFile[i][1], err)
			return
		}

		deliveryMap[delivery.Id] = &delivery
	}

	Solution1(inputBody, partnerMap)

}
