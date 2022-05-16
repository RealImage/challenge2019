package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Partner struct {
	theatre     string
	minSizeSlab string
	maxSizeSlab string
	minimumCost string
	costPerGB   string
	partnerID   string
}

type MinimumDelivery struct {
	delivery  string
	status    bool
	partnerID string
	cost      string
}

func main() {
	findMinimumDelivery("input.csv", "partners.csv")
}

/*
	method to convert data to partners
*/
func convertToPartner(data [][]string) []Partner {
	partners := make([]Partner, len(data))

	for i, row := range data {
		sizeSlab := strings.Split(row[1], "-")
		partners[i] = Partner{
			theatre:     strings.TrimSpace(row[0]),
			minSizeSlab: strings.TrimSpace(sizeSlab[0]),
			maxSizeSlab: strings.TrimSpace(sizeSlab[1]),
			minimumCost: strings.TrimSpace(row[2]),
			costPerGB:   strings.TrimSpace(row[3]),
			partnerID:   strings.TrimSpace(row[4]),
		}
	}

	return partners
}

/*
	method to read data from a csv file
*/
func readDataFromCsv(filePath string, headingOn bool) [][]string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	// skipping the first line if headingOn is on
	if headingOn {
		row1, err := bufio.NewReader(file).ReadSlice('\n')
		if err != nil {
			log.Fatal(err)
		}

		_, err = file.Seek(int64(len(row1)), io.SeekStart)
		if err != nil {
			log.Fatal(err)
		}
	}

	// reading file
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	file.Close()

	if err != nil {
		log.Fatal(err)
	}

	return data
}

/*
	method to find minimum delivery
*/
func findMinimumDelivery(inputFilePath string, partnersFilePath string) {
	data := readDataFromCsv(inputFilePath, false)
	partners := convertToPartner(readDataFromCsv(partnersFilePath, true))
	var finalOutput []MinimumDelivery

	for _, d := range data {
		dSize := d[1]
		dTheatre := d[2]
		dSizeInt, _ := strconv.Atoi(dSize)
		var partnersFinalCost []MinimumDelivery

		for _, partner := range partners {
			minSizeSlabInt, _ := strconv.Atoi(partner.minSizeSlab)
			maxSizeSlabInt, _ := strconv.Atoi(partner.maxSizeSlab)
			costPerGBInt, _ := strconv.Atoi(partner.costPerGB)
			minimumCostInt, _ := strconv.Atoi(partner.minimumCost)
			totalCost := dSizeInt * costPerGBInt

			var finalCost int

			// check if we have data for the theatre and check the size slab conditions
			if partner.theatre == dTheatre && (dSizeInt >= minSizeSlabInt && dSizeInt <= maxSizeSlabInt) {

				// check total cost and minimum cost
				if totalCost < minimumCostInt {
					finalCost = minimumCostInt
				} else {
					finalCost = totalCost
				}

				minimunDeliverCost := []MinimumDelivery{
					{
						delivery:  d[0],
						status:    true,
						partnerID: partner.partnerID,
						cost:      strconv.Itoa(finalCost),
					},
				}

				partnersFinalCost = append(partnersFinalCost, minimunDeliverCost...)
			}
		}

		// get the minimum cost for the delivery
		finalMinimumCostDelivery := lo.MinBy(partnersFinalCost, func(item MinimumDelivery, min MinimumDelivery) bool {
			itemCostInt, _ := strconv.Atoi(item.cost)
			minCostInt, _ := strconv.Atoi(min.cost)

			return itemCostInt < minCostInt
		})

		// check the item can be delivered
		if finalMinimumCostDelivery.status {
			finalOutput = append(finalOutput, finalMinimumCostDelivery)
		} else {
			finalOutput = append(finalOutput, MinimumDelivery{
				delivery:  d[0],
				status:    false,
				partnerID: "''",
				cost:      "''",
			})
		}
	}

	// open the output file
	outputFile, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}

	csvWriter := csv.NewWriter(outputFile)

	// loop through the output
	lo.ForEach(finalOutput, func(item MinimumDelivery, _ int) {
		// write to csv file
		csvWriter.Write([]string{item.delivery, strconv.FormatBool(item.status), item.partnerID, item.cost})

		// print output
		fmt.Printf("%s, %t, %s, %s\n", item.delivery, item.status, item.partnerID, item.cost)
	})

	csvWriter.Flush()
	outputFile.Close()
}
