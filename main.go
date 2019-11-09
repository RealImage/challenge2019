package main

import (
	"challenge2019/Prob"
	"fmt"
)

func main() {
	/*	input := []Prob.DeliveryInfo{
			{"D1", 150, "T1"},
			{"D2", 325, "T2"},
			{"D3", 510, "T1"},
			{"D4", 700, "T2"},
		}
	*/
	inputData := Prob.FetchInputFromCSV("input.csv")
	partnerData := Prob.FetchPartnerDataFromCSV("partners.csv")
	capacityInfo := Prob.FetchPartnerCapacityFromCSV("capacities.csv")
	_, AllApplicablePartners := Prob.FindBestApplicablePartner(partnerData, inputData)
	fmt.Println(AllApplicablePartners)
	Prob.Prob2(AllApplicablePartners, capacityInfo)

}
