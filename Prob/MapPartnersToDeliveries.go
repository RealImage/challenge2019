package Prob

import (
	"fmt"
	"strings"
)

func Prob2(allApplicablePartners []DeliveryAndPartners, capacityInfo []CapacityInfo) {
	allApplicablePartners = UpdatePartnersCapacity(allApplicablePartners, capacityInfo)
	partners := FindAllPartners(capacityInfo)
	fmt.Println(allApplicablePartners)
	fmt.Println(partners)

	p := FindAllPermutations(allApplicablePartners)
	feasibleArray := FindAllFeasiblePermutations(p, capacityInfo)
	fmt.Println(feasibleArray)
	for _, l := range feasibleArray {
		dMap := make(map[[]PartnerData]int)
		sum := FindTotalDeliveryCharge(l)
		dMap[l] = sum
		fmt.Println(dMap)
	}
}
func FindAllFeasiblePermutations(allPermutations [][]PartnerData, cInfo []CapacityInfo) [][]PartnerData {
	var possible [][]PartnerData
	var isPossible bool
	for _, v := range allPermutations {
		isPossible = FindIfPossible(v, cInfo)
		if isPossible == true {
			possible = append(possible, v)
		}
	}
	return possible
}
func FindAllPartners(capacityInfo []CapacityInfo) []string {
	pArray := []string{}
	for _, r := range capacityInfo {
		p := strings.TrimSpace(r.PartnerID)
		pArray = append(pArray, p)
	}
	return pArray
}

func UpdatePartnersCapacity(allApplicablePartners []DeliveryAndPartners, capacityInfo []CapacityInfo) []DeliveryAndPartners {
	for i, a := range allApplicablePartners {
		for j, vv := range a.Partners {
			for _, pp := range capacityInfo {
				if strings.TrimSpace(pp.PartnerID) == vv.PartnerID {
					allApplicablePartners[i].Partners[j].Capacity = ConvertToInt(pp.Capacity)
				}
			}
		}
	}
	return allApplicablePartners
}

func FindAllPermutations(allApplicablePartners []DeliveryAndPartners) [][]PartnerData {
	partners := [][]PartnerData{}
	for _, r := range allApplicablePartners {
		partners = append(partners, r.Partners)
	}

	res := findAllPermutations(partners...)
	for _, s := range res {
		fmt.Println("-->", s)
	}
	return res

}

func findAllPermutations(parts ...[]PartnerData) (ret [][]PartnerData) {
	{
		var n = 1
		for _, ar := range parts {
			n *= len(ar)
		}
		ret = [][]PartnerData{}
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
		tmp := []PartnerData{}
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

func FindIfPossible(partners []PartnerData, all []CapacityInfo) bool {

	for _, i := range all {
		sum := 0
		for _, j := range partners {
			if strings.TrimSpace(i.PartnerID) == j.PartnerID {
				sum = sum + j.Delivery.DeliverySize
			}
		}
		if sum > ConvertToInt(i.Capacity) {
			return false
		} else if sum <= ConvertToInt(i.Capacity) {
			return true
		}
	}
	return false
}

func FindTotalDeliveryCharge(fPartners []PartnerData) int {
	sum := 0
	for _, i := range fPartners {
		sum = sum + int(i.DeliveryCost)
	}
	return sum
}

/*
func MapPartnerToDelivery(allApplicablePartners []DeliveryAndPartners, partners []CapacityInfo) map[CapacityDetails][]DelAndPartners {
	fmt.Println("APPLICABLE PARTNERS", allApplicablePartners)
	dpMap := make(map[CapacityDetails][]DelAndPartners)
	for pp, _ := range partners {
		tmp := make([]DelAndPartners, 0)
		for _, r := range allApplicablePartners {
			for _, l := range r.Partners {
				if strings.TrimSpace(partners[pp].PartnerID) == l.PartnerID {
	P				t := DelAndPartners{r.Delivery.DeliveryID, r.Delivery.DeliverySize, l.Theatre, l.PartnerID}
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
func ArrayOfPartners(m map[CapacityDetails][]DelAndPartners) [][]DelAndPartners {
	var a [][]DelAndPartners
	for _, v := range m {
		a = append(a, v)
	}
	return a
}
*/
