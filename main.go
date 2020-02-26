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

type deliveryDetails struct {
	theatre         string
	minSizeSlab     string
	maximumSizeSlab string
	minimumCost     string
	costPerGB       string
	partnerID       string
}

var deliveryDetailsArray []deliveryDetails

func main() {
	fmt.Println("Qube Cinemas Code challanges")
	readRawData()
	problemStatement1()
}

func readRawData() {
	csvfile, err := os.Open("partners.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)
	i := 0
	for {
		value, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if i == 0 {
			i++
			continue
		}

		sizeSlab := strings.Split(value[1], "-")
		deliveryObj := deliveryDetails{
			theatre:         strings.TrimSpace(value[0]),
			minSizeSlab:     strings.TrimSpace(sizeSlab[0]),
			maximumSizeSlab: strings.TrimSpace(sizeSlab[1]),
			minimumCost:     strings.TrimSpace(value[2]),
			costPerGB:       strings.TrimSpace(value[3]),
			partnerID:       strings.TrimSpace(value[4]),
		}
		deliveryDetailsArray = append(deliveryDetailsArray, deliveryObj)
	}
}

func problemStatement1() {
	fmt.Println("Problem Statement 1")
	csvfile, err := os.Open("input.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)
	for {
		value, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		deliveryId := value[0]
		sizeOfdelivery, _ := strconv.Atoi(value[1])
		thaetreId := value[2]
		flag := false
		for _, val := range deliveryDetailsArray {
			minSizeSlab, _ := strconv.Atoi(val.minSizeSlab)
			maximumSizeSlab, _ := strconv.Atoi(val.maximumSizeSlab)
			if thaetreId == val.theatre && minSizeSlab < sizeOfdelivery && maximumSizeSlab > sizeOfdelivery {
				flag = true
				costPerGB, _ := strconv.Atoi(val.costPerGB)
				fmt.Println(deliveryId, true, val.partnerID, (sizeOfdelivery * costPerGB))
				break
			}
		}
		if !flag {
			fmt.Println(deliveryId, false, "''", "''")
		}
	}
}

func problemStatement2() {

}
