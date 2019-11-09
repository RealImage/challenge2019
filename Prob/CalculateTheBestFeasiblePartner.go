package Prob

import (
	"challenge2019/Prob/types"
)

func CalculateBestFeasiblePartnerList(allApplicablePartners []types.DeliveryAndPartners,
	capacityInfo []types.CapacityDetailsStr, inputData []types.DeliveryInfo) []types.FinalChoice {

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

	//Sort the obtained permutation array in ascending order
	SortFeasibleArray(feasibleArray)

	//Converting result into expected output format
	result := Output(feasibleArray, inputData)

	return result
}
