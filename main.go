package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open input.csv partnersFile
	inputFile, err := os.Open("input.csv")
	if err != nil {
		log.Fatal(err)
	}
	// close inputFile at the end of main
	defer inputFile.Close()

	// create output.csv file
	outputFile, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	// close output file at the end of main
	defer outputFile.Close()

	// create a csv writer for output file
	outputFileWriter := csv.NewWriter(outputFile)
	// flush writer at the end
	defer outputFileWriter.Flush()

	// read csv input.csv using a csv.Reader
	inputFileCsvReader := csv.NewReader(inputFile)

	// iterate over input and partners to minimum cost of delivery
	for {
		inputRecord, err := inputFileCsvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		inputTheaterId := inputRecord[2]
		deliverySize, err := strconv.Atoi(inputRecord[1])
		if err != nil {
			log.Fatal("Error while converting deliverySize string to int: ", err)
		}

		// open partners.csv partnersFile
		partnersFile, err := os.Open("partners.csv")
		if err != nil {
			log.Fatal(err)
		}
		// read csv partner.csv using a csv.Reader
		partnerFileCsvReader := csv.NewReader(partnersFile)

		//skip first record
		if _, err := partnerFileCsvReader.Read(); err != nil {
			fmt.Println("Error reading first line of partner.csv: ", err)
		}

		minDeliveryCost := 0
		isDeliveryPossible := false
		deliveryPartner := ""

		var outputRecord []string

		// write deliveryId to output
		outputRecord = append(outputRecord, inputRecord[0])

		for {
			partnerRecord, err := partnerFileCsvReader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			sizeSlab := strings.Split(partnerRecord[1], "-")
			sizeSlabMin, err := strconv.Atoi(strings.TrimSpace(sizeSlab[0]))
			if err != nil {
				log.Fatal(err)
			}
			sizeSlabMax, err := strconv.Atoi(strings.TrimSpace(sizeSlab[1]))
			if err != nil {
				log.Fatal(err)
			}

			if inputTheaterId == strings.TrimSpace(partnerRecord[0]) && (deliverySize >= sizeSlabMin && deliverySize <= sizeSlabMax) {
				isDeliveryPossible = true
				rate, err := strconv.Atoi(strings.TrimSpace(partnerRecord[3]))
				if err != nil {
					log.Fatal(err)
				}
				deliveryCost := deliverySize * rate
				minCost, err := strconv.Atoi(strings.TrimSpace(partnerRecord[2]))
				if err != nil {
					log.Fatal(err)
				}
				if deliveryCost < minCost {
					deliveryCost = minCost
				}
				if minDeliveryCost == 0 || deliveryCost < minDeliveryCost {
					minDeliveryCost = deliveryCost
					deliveryPartner = partnerRecord[4]
				}
			}
		}
		if !isDeliveryPossible {
			outputRecord = append(outputRecord, fmt.Sprintf("%t", isDeliveryPossible))
			outputRecord = append(outputRecord, "")
			outputRecord = append(outputRecord, "")
		} else {
			outputRecord = append(outputRecord, fmt.Sprintf("%t", isDeliveryPossible), deliveryPartner, fmt.Sprintf("%d", minDeliveryCost))
		}
		partnersFile.Close()
		outputFileWriter.Write(outputRecord)
	}
}
