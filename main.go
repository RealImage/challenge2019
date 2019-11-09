package main

import (
	"challenge2019/Prob"
)

func main() {
	input := []Prob.DeliveryInfo{
		{"D1", 150, "T1"},
		{"D2", 325, "T2"},
		{"D3", 510, "T1"},
		{"D4", 700, "T2"},
	}

	partnerData := Prob.FetchPartnerDataFromCSV("partners.csv")
	_, AllApplicablePartners := Prob.FindAllPartnerInfo(partnerData, input)
	capacityInfo := Prob.FetchPartnerCapacityFromCSV("capacities.csv")
	Prob.Prob2(AllApplicablePartners, capacityInfo)

}
