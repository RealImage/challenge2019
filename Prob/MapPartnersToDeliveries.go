package Prob

import (
	"fmt"
	"strings"
)

func UpdatePartnersCapacity(allApplicablePartners []DeliveryAndPartners, capacityInfo []CapacityInfo) {
	/*
		for i, a := range allApplicablePartners {
			for j, vv := range a.Partners {
				for _, pp := range capacityInfo {
					if strings.TrimSpace(pp.PartnerID) == vv.PartnerID {
						allApplicablePartners[i].Partners[j].Capacity = ConvertToInt(pp.Capacity)
					}
				}
			}
		}
	*/
	//Remove
	dpMap := MapPartnerToDelivery(allApplicablePartners, capacityInfo)
	//	Calculate(dpMap)
	arr := ArrayOfPartners(dpMap)
	_ = PermuteDel(arr...)
}

func MapPartnerToDelivery(allApplicablePartners []DeliveryAndPartners, partners []CapacityInfo) map[CapacityDetails][]DelAndPartners {
	fmt.Println("APPLICABLE PARTNERS", allApplicablePartners)
	dpMap := make(map[CapacityDetails][]DelAndPartners)
	for pp, _ := range partners {
		tmp := make([]DelAndPartners, 0)
		for _, r := range allApplicablePartners {
			for _, l := range r.Partners {
				if strings.TrimSpace(partners[pp].PartnerID) == l.PartnerID {
					t := DelAndPartners{r.Delivery.DeliveryID, r.Delivery.DeliverySize, l.Theatre, l.PartnerID}
					tmp = append(tmp, t)
				}
			}
		}
		tp1 := strings.TrimSpace(partners[pp].PartnerID)
		tp2 := ConvertToInt(partners[pp].Capacity)
		dpMap[CapacityDetails{tp1, tp2}] = tmp
	}
	fmt.Println(dpMap)
	return dpMap
}

func Calculate(dpMap map[CapacityDetails][]DelAndPartners) {
	finMap := make(map[CapacityDetails][]DelAndPartners)

	for k, v := range dpMap {
		arr := make([]DelAndPartners, 0)
		sum := 0

		for _, r := range v {

			if sum < k.Capacity {
				sum = sum + r.DeliverySize
				if sum < k.Capacity {
					arr = append(arr, r)
				} else {
					sum = sum - r.DeliverySize

				}
			}
		}

		finMap[CapacityDetails{k.PartnerID, sum}] = arr
	}

	//fmt.Println(finMap)
	//fmt.Println(dpMap)
}

func ArrayOfPartners(m map[CapacityDetails][]DelAndPartners) [][]DelAndPartners {
	var a [][]DelAndPartners
	for _, v := range m {
		a = append(a, v)
	}
	return a
}

func PermuteDel(parts ...[]DelAndPartners) (ret [][]DelAndPartners) {
	{
		var n = 1
		for _, ar := range parts {
			n *= len(ar)
		}
		ret = [][]DelAndPartners{}
	}
	var at = make([]int, len(parts))
loop:
	for {
		for i := len(parts) - 1; i >= 0; i-- {
			if at[i] > 0 && at[i] >= len(parts[i]) {
				if i == 0 || (i == 1 && at[i-1] == len(parts[0])-1) {
					break loop
				}
				at[i] = 0
				at[i-1]++
			}
		}
		// construct permutated string
		tmp := []DelAndPartners{}
		for i, ar := range parts {
			var p = at[i]
			if p >= 0 && p < len(ar) {
				tmp = append(tmp, ar[p])
			}
		}
		ret = append(ret, tmp)
		at[len(parts)-1]++
	}
	return ret
}
