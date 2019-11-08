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
	_ = MapPartnerToDelivery(allApplicablePartners, capacityInfo)
	//	Calculate(dpMap)
	//f := permutations(v)
	//fmt.Println(f)
}

func MapPartnerToDelivery(allApplicablePartners []DeliveryAndPartners, partners []CapacityInfo) map[CapacityDetails][]DelAndPartners {

	dpMap := make(map[CapacityDetails][]DelAndPartners)
	for pp, _ := range partners {
		tmp := make([]DelAndPartners, 0)
		for _, r := range allApplicablePartners {
			for _, l := range r.Partners {
				if strings.TrimSpace(partners[pp].PartnerID) == l.PartnerID {
					t := DelAndPartners{r.Delivery.DeliveryID, r.Delivery.DeliverySize, l.Theatre}
					tmp = append(tmp, t)
				}
			}
		}
		tp1 := strings.TrimSpace(partners[pp].PartnerID)
		tp2 := ConvertToInt(partners[pp].Capacity)
		dpMap[CapacityDetails{tp1, tp2}] = tmp
	}
	return dpMap
}

func Calculate(dpMap map[CapacityDetails][]DelAndPartners) {
	finMap := make(map[CapacityDetails][]DelAndPartners)

	for k, v := range dpMap {
		arr := make([]DelAndPartners, 0)
		sum := 0

		fmt.Println(k)
		fmt.Println(v)
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
func perm(m map[CapacityDetails][]DelAndPartners) [][]DelAndPartners {
	var a [][]DelAndPartners
	for _, v := range m {
		a = append(a, v)
	}
	fmt.Println("Array of arrays :", a)
	return a
}

/*
func permutations(arr [][]DelAndPartners) [][]DelAndPartners {
	var helper func([][]DelAndPartners, int)
	res := [][]DelAndPartners{}

	helper = func(arr [][]DelAndPartners, n int) {
		if n == 1 {
			res = append(res, arr...)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
func PermuteDel(parts [][]DelAndPartners) (ret [][]DelAndPartners) {
	{
		var n = 1
		for _, ar := range parts {
			n *= len(ar)
		}
		ret = make([]DelAndPartners, 0, n)
	}
	var at = make([]int, len(parts))
	//var buf bytes.Buffer
loop:
	for {
		// increment position counters
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
		//buf.Reset()
		var tmp DelAndPartners
		for i, ar := range parts {
			var p = at[i]
			if p >= 0 && p < len(ar) {
				tmp = ar[p]
				//			buf.WriteString(ar[p])
			}
		}
		ret = append(ret, tmp)
		fmt.Println(ret)
		at[len(parts)-1]++
	}
	fmt.Println(ret)
	return ret
}
*/
