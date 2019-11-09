package Prob

import (
	"challenge2019/Prob/types"
	"strings"
)

func FindAllFeasiblePermutations(allPermutations [][]types.PartnerData, cInfo []types.CapacityDetailsStr) [][]types.PartnerData {
	var possible [][]types.PartnerData
	var isPossible bool
	for _, v := range allPermutations {
		isPossible = FindIfPossible(v, cInfo)
		if isPossible == true {
			possible = append(possible, v)
		}
	}
	return possible
}

func FindIfPossible(partners []types.PartnerData, all []types.CapacityDetailsStr) bool {

	for _, i := range all {
		sum := 0
		for _, j := range partners {
			if strings.TrimSpace(i.PartnerID) == j.PartnerID {
				sum = sum + j.Delivery.DeliverySize
			}
		}
		if sum > types.ConvertToInt(i.Capacity) {
			return false
		} else if sum <= types.ConvertToInt(i.Capacity) {
			return true
		}
	}
	return false
}
