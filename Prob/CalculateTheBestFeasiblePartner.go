package Prob

import (
	"challenge2019/Prob/types"
)

func CalculateTheBestFeasiblePartner(allApplicablePartners []types.DeliveryAndPartners, capacityInfo []types.CapacityDetailsStr) []types.PartnerData {
	//Update Capacity of each appicable partner
	allApplicablePartners = UpdatePartnersCapacity(allApplicablePartners, capacityInfo)
	//Find All Possible Permutations of all Applicable Partners
	p := FindAllPermutations(allApplicablePartners)
	//Considering, capacity limit of each partner, find all feasible permutation
	feasibleArray := FindAllFeasiblePermutations(p, capacityInfo)
	//Calculate total delivery charge of each feasible permutation
	fMap := make(map[int]int)
	for i, l := range feasibleArray {
		sum := FindTotalDeliveryCharge(l)
		fMap[sum] = i
	}
	//Sort the obtained map in ascending order
	index := SortDeliveryPrices(fMap)
	return feasibleArray[index]

}
