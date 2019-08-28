/*
@author: Vinay Kumar
*/
package qube

import(
	"strings"
	"strconv"
 	"io"
 	"log"
	"./helpers"
)

type Capacity map[string] int

type PartnerStruct struct{
	MinSlab int
	MaxSlab int
	MinCost int
	CostPerGB int
}

type Partner map[string] map[string] [] PartnerStruct

// To export
var Capacities Capacity = createCapacityHashTable()
var Partners Partner = createPartnerHashTable()

func createPartnerHashTable() Partner {
	r := helpers.ParseContentsFromCsv("static/partners.csv")
	partner := Partner{}
	line := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		
		if line == 0{
			line++
			continue
		}
		theatreId := strings.TrimSpace(record[0])
		partnerId := strings.TrimSpace(record[4])
		minCost, minCostConversionErr := strconv.Atoi(strings.TrimSpace(record[2]))
		costPerGB, costPerGBconversionErr := strconv.Atoi(strings.TrimSpace(record[3]))
		usageSlab := strings.Split(strings.TrimSpace(record[1]), "-")
		minSlab, _ := strconv.Atoi(usageSlab[0])
		maxSlab, _ := strconv.Atoi(usageSlab[1])
		if minCostConversionErr != nil || costPerGBconversionErr != nil {
			log.Fatal("Something went wrong", minCostConversionErr, costPerGBconversionErr)
		}
		partnerData := PartnerStruct{MinSlab: minSlab, MaxSlab: maxSlab, MinCost: minCost, CostPerGB: costPerGB}
		if partner[theatreId] == nil {
			partner[theatreId] = map [string] [] PartnerStruct{}
		}
		partner[theatreId][partnerId] = append(partner[theatreId][partnerId], partnerData)
		line++
	}
	return partner
}

func createCapacityHashTable() Capacity{
	r := helpers.ParseContentsFromCsv("static/capacities.csv")
	capacity := Capacity{}
	line := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		
		if line == 0{
			line++
			continue
		}
		number, err := strconv.Atoi(record[1])
		if err != nil{
			log.Fatal(err)
		}
		capacity[strings.TrimSpace(record[0])] = number
		line++
	}
	return capacity
}