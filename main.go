package main

import (
	"challenge2019/Prob"
	"fmt"
)

func main() {
	fmt.Println("PARTNER DATA")
	partnerData := Prob.FetchPartnerDataFromCSV("partners.csv")
	fmt.Println("CalculateMinDeliveryCost")
	Prob.CalculateMinDeliveryCost(partnerData)
	/*	fmt.Println("INPUT DATA")
		Prob.FetchDataFromCSV("input.csv")
	*/
}
