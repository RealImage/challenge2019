package main

import (
	"challenge2019/Prob"
)

func main() {
	partnerData := Prob.FetchPartnerDataFromCSV("partners.csv")
	_, AllApplicablePartners := Prob.FindAllPartnerInfo(partnerData)
	capacityInfo := Prob.FetchPartnerCapacityFromCSV("capacities.csv")
	Prob.UpdatePartnersCapacity(AllApplicablePartners, capacityInfo)

}
