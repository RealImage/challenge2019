package main

import (
	"challenge2019/Prob"

	f "challenge2019/Prob/FetchInput"
	"fmt"
)

func main() {
	inputData := f.FetchInputFromCSV("input.csv")
	partnerData := f.FetchPartnerDataFromCSV("partners.csv")
	capacityInfo := f.FetchPartnerCapacityFromCSV("capacities.csv")
	_, AllApplicablePartners := Prob.FindBestApplicablePartner(partnerData, inputData)
	fmt.Println(AllApplicablePartners)
	Prob.Prob2(AllApplicablePartners, capacityInfo)

}
