package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Partner struct {
	theatre         string
	minimumSizeSlab int
	maximumSizeSlab int
	minimumCost     int
	costPerGB       int
	ID              string
}

type Input struct {
	ID      string
	size    int
	theatre string
}

type Output struct {
	ID         string
	indication bool
	partner    Partner
	totalCost  int
}

func getMinimumCost(partner Partner, input Input) int {

	calculatedCost := input.size * partner.costPerGB // calucation of total cost for the input size
	if calculatedCost <= partner.minimumCost {
		return partner.minimumCost
	}
	return calculatedCost
}

func main() {
	content, err := ioutil.ReadFile("../../partners.csv") //realtive path for partners csv file
	if err != nil {
		log.Fatal(err)
	}

	masterData := string(content)
	csvReader := csv.NewReader(strings.NewReader(masterData))
	records, err := csvReader.ReadAll() //reading partners csv file
	if err != nil {
		log.Fatal(err)
	}
	partnersData := []Partner{}      //declaring array of Partner
	for i, record := range records { //looping through the records of partners csv file
		if i == 0 {
			continue
		}
		theatre := strings.TrimSpace(record[0])
		costPerGB, err := strconv.Atoi(strings.TrimSpace(record[3]))

		if err != nil {
			log.Fatal(err)
		}
		ID := strings.TrimSpace(record[4])
		minimumCost, err := strconv.Atoi(strings.TrimSpace(record[2]))
		if err != nil {
			log.Fatal(err)
		}
		cost := (strings.TrimSpace(record[1]))
		costs := strings.Split(cost, "-")
		minimumSizeSlab, err := strconv.Atoi(costs[0])
		if err != nil {
			log.Fatal(err)
		}
		maximumSizeSlab, err := strconv.Atoi(costs[1])
		if err != nil {
			log.Fatal(err)
		}
		partnersData = append(partnersData, Partner{ //appending partners data with partner struct
			theatre:         theatre,
			minimumSizeSlab: minimumSizeSlab,
			maximumSizeSlab: maximumSizeSlab,
			costPerGB:       costPerGB,
			minimumCost:     minimumCost,
			ID:              ID,
		})

	}

	inputDataContent, err := ioutil.ReadFile("../../input.csv")
	if err != nil {
		log.Fatal(err)
	}

	inputData := string(inputDataContent)
	inputCsvReader := csv.NewReader(strings.NewReader(inputData)) //reading input csv data
	inputRecords, err := inputCsvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	csvWriter := csv.NewWriter(file)

	for _, inputData := range inputRecords {

		ID := strings.TrimSpace(inputData[0])

		size, err := strconv.Atoi(strings.TrimSpace(inputData[1]))
		if err != nil {
			log.Fatal(err)
		}
		theatre := strings.TrimSpace(inputData[2])

		inputDataStruct := Input{ //assigning input csv datas to input struct
			ID:      ID,
			size:    size,
			theatre: theatre,
		}
		outputStruct := Output{}
		bestTotalCost := 0
		bestPartner := ""

		for _, partnerData := range partnersData {
			if inputDataStruct.theatre == partnerData.theatre {
				if inputDataStruct.size < partnerData.minimumSizeSlab || inputDataStruct.size > partnerData.maximumSizeSlab { //checking condition for false indication
					outputStruct.indication = false
				} else {
					outputStruct.indication = true
					newMinimalCost := getMinimumCost(partnerData, inputDataStruct) //creating function to calculate minimum cost
					if bestTotalCost == 0 {
						bestTotalCost = newMinimalCost // assigning the minimal cost as besttotalcost when the bestest is 0
						bestPartner = partnerData.ID   // bestpartner is assigned for the partner who holds besttotalcost which is calculated above
					} else if newMinimalCost < bestTotalCost {
						bestTotalCost = newMinimalCost //checking whether the newminimalcost is lesser than the previous cost and assigning as besttotalcost
						bestPartner = partnerData.ID   // bestpartner is assigned for the partner who holds besttotalcost which is calculated above
					}
				}
			}
		}
		var bestCost string
		if bestTotalCost > 0 {
			bestCost = strconv.Itoa(bestTotalCost)
		}
		outputRecord := []string{
			inputDataStruct.ID,
			strconv.FormatBool(outputStruct.indication),
			bestCost,
			bestPartner,
		}
		if err := csvWriter.Write(outputRecord); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
		csvWriter.Flush()
		if err := csvWriter.Error(); err != nil {
			log.Fatal(err)
		}

	}
}
