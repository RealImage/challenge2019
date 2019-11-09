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
	BestPartner, AllApplicablePartners := Prob.FindBestApplicablePartner(partnerData, inputData)
	fmt.Println("Best Applicable Partner :")
	fmt.Println(BestPartner)
	result := Prob.CalculateBestFeasiblePartnerList(AllApplicablePartners, capacityInfo, inputData)
	fmt.Println("Best Permutation of partners :")
	fmt.Println(result)

}
