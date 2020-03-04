package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var partners, capacities, input [][]string
var cap = make(map[string]int)

func init() {
	partners = InputReader("partners.csv")
	capacities = InputReader("capacities.csv")
	input = InputReader("input.csv")
	for _, capacitiesRow := range capacities[1:] {
		var err error
		cap[capacitiesRow[0]], err = strconv.Atoi(capacitiesRow[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	//fmt.Println(partners, capacities, input)
}

//InputReader ...
func InputReader(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

//PS2 ...
func PS2() [][]string {
	var output [][]string

	for _, inputRow := range input {
		for partner := range cap {
			output = append(output, append(inputRow, partner))
		}
	}

	return output
}

// D1, 100, T1
// D2, 240, T1
// D2, 260, T1

var pResult [][]string

// func allotment(DeliveryLength int) {
// 	if DeliveryLength == 0 {
// 		return
// 	}
// 	for _, v := range pResult {
// 		for partner := range cap {
// 			output = append(output, append(v, partner))
// 		}
// 	}
// 	allotment(DeliveryLength - 1)
// }

//PS1 ...
func PS1() [][]string {
	var output [][]string
	for _, inputRow := range input {
		var mincost, partnerID string

		for _, partnersRow := range partners[1:] {
			if strings.TrimSpace(inputRow[2]) == strings.TrimSpace(partnersRow[0]) && checkSlab(partnersRow[1], inputRow[1]) {
				cost := toInt(inputRow[1]) * toInt(partnersRow[3])
				if partnerMinCost := toInt(partnersRow[2]); cost < partnerMinCost {
					cost = partnerMinCost
				}
				if mincost == "" || cost < toInt(mincost) {
					mincost = strconv.Itoa(cost)
					partnerID = partnersRow[4]
				}
			}
		}

		result := make([]string, 4)
		result[0] = inputRow[0]
		if mincost == "" {
			result[1] = "false"
			result[2], result[3] = ` `, ` `
		} else {
			result[1] = "true"
			result[2], result[3] = partnerID, mincost
		}

		output = append(output, result)

	}

	return output
}

func toInt(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func checkSlab(slab string, value string) bool {
	slabArr := strings.Split(slab, "-")
	if value >= slabArr[0] && value <= slabArr[1] {
		return true
	}
	return false
}

func main() {
	output1 := PS1()
	fmt.Println(output1)
	OutputWriter("output1.csv", output1)
}

//OutputWriter ...
func OutputWriter(filename string, output [][]string) {
	file, _ := os.Create(filename)
	defer file.Close()

	r := csv.NewWriter(file)
	err := r.WriteAll(output)
	if err != nil {
		log.Fatal(err)
	}
}
