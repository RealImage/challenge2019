package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateMinimumCost(record []string, contentSize int, minCost float64, deliveryPartner string) (float64, string) {
	sizeSlab := strings.Split(strings.TrimSpace(record[1]), "-")
	minRange, err := strconv.Atoi(sizeSlab[0])
	if err != nil {
		log.Fatal(err)
	}
	maxRange, err := strconv.Atoi(sizeSlab[1])
	if err != nil {
		log.Fatal(err)
	}
	// Checking if delivery is possible or not
	if contentSize >= minRange && contentSize <= maxRange {
		a, err := strconv.Atoi(strings.TrimSpace(record[2]))
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.Atoi(strings.TrimSpace(record[3]))
		if err != nil {
			log.Fatal(err)
		}
		c := math.Max(float64(a), float64(b)*float64(contentSize))
		//  Calculating minimum cost
		if minCost > c {
			// Assigning delivery partner
			deliveryPartner = record[4]
			minCost = c
		}
	}
	return minCost, deliveryPartner
}

func main() {
	f2, err := os.Open("input.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()
	p := csv.NewReader(f2)
	allRecords := [][]string{}

	for {
		f, err := os.Open("partners.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		r := csv.NewReader(f)
		// Skip the first row
		if _, err := r.Read(); err != nil {
			fmt.Println("Error", err)
		}
		var minCost float64 = math.Pow10(10)
		var deliveryPartner string
		outputRecord := []string{}
		// Reading the input from input.csv
		input, err := p.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		contentSize, err := strconv.Atoi(strings.TrimSpace(input[1]))
		if err != nil {
			log.Fatal(err)
		}
		for {
			// Reading each record from partners.csv
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			minCost, deliveryPartner = calculateMinimumCost(record, contentSize, minCost, deliveryPartner)
		}
		if minCost == math.Pow10(10) {
			// Delivery not possible
			outputRecord = append(outputRecord, input[0])
			outputRecord = append(outputRecord, "false")
			outputRecord = append(outputRecord, deliveryPartner)
			outputRecord = append(outputRecord, "")
			allRecords = append(allRecords, outputRecord)
		} else {
			// Delivery possible
			outputRecord = append(outputRecord, input[0])
			outputRecord = append(outputRecord, "true")
			outputRecord = append(outputRecord, deliveryPartner)
			outputRecord = append(outputRecord, strconv.Itoa(int(minCost)))
			allRecords = append(allRecords, outputRecord)
		}
	}
	// Create or if exists open output file
	f3, err := os.OpenFile("output.csv", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer f3.Close()
	// Writing to the output.csv file
	w := csv.NewWriter(f3)
	defer w.Flush()
	for _, record := range allRecords {
		if err := w.Write(record); err != nil {
			log.Fatalln("Error writing record to file!")
		}
	}
}
