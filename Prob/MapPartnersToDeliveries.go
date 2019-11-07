package Prob

import (
	"fmt"
	"strings"
)

func UpdatePartnersCapacity(allApplicablePartners []DeliveryAndPartners, capacityInfo []CapacityInfo) {
	var partners []DelAndPartners
	dpMap := make(map[string][]DelAndPartners)
	for i, a := range allApplicablePartners {
		for j, vv := range a.Partners {
			for _, pp := range capacityInfo {
				if strings.TrimSpace(pp.PartnerID) == vv.PartnerID {
					allApplicablePartners[i].Partners[j].Capacity = ConvertToInt(pp.Capacity)
					partners = append(partners, DelAndPartners{a.Delivery.DeliveryID, a.Delivery.DeliverySize, vv.Theatre, vv.Capacity})
					dpMap[pp.PartnerID] = partners
				}

			}
		}
	}
	fmt.Println(dpMap)
	//MapPartnerToDelivery(allApplicablePartners, partners)
}

/*
func MapPartnerToDelivery(allApplicablePartners []DeliveryAndPartners, partners []string) {
	dpMap := make(map[string][]DelAndPartners)
	var tmp []DelAndPartners
	for _, r := range allApplicablePartners {
		for _, l := range r.Partners {
			for _,pp:= range partners{
if l

			tmp = append(tmp, DelAndPartners{r.Delivery.DeliveryID, r.Delivery.DeliverySize, l.Theatre, l.Capacity})
			dpMap[l.PartnerID] = tmp
		}
	}

	fmt.Println(dpMap)
}

/*
func FindAllCombinations(allApplicablePartners []DeliveryAndPartners, capacityInfo []CapacityInfo) {

	DToPMap := make(map[DeliveryDetails][]PartnerInfo)

	for _, a := range allApplicablePartners {
		pArray := make([]PartnerInfo, 0)
		for _, vv := range a.Partners {
			var p PartnerInfo
			p = PartnerInfo{vv.PartnerID, vv.DeliveryCost, vv.Capacity}
			for _, pp := range capacityInfo {
				if strings.TrimSpace(pp.PartnerID) == p.PartnerID {
					p.Capacity = ConvertToInt(pp.Capacity)
				}
			}
			pArray = append(pArray, p)
		}

		DToPMap[DeliveryDetails{a.Delivery.DeliveryID, a.Delivery.DeliverySize}] = pArray
	}
	var temp [][]PartnerInfo
	for _, v := range DToPMap {
		for i := 1; i < len(v); i++ {
			temp = HeapPerm(v, len(v))
			fmt.Println(temp)
		}
	}
}
*/

func HeapPerm(arr []PartnerInfo, size int) [][]PartnerInfo {
	var heap func([]PartnerInfo, int)
	res := [][]PartnerInfo{}

	heap = func(arr []PartnerInfo, n int) {
		if n == 1 {
			res = append(res, arr)
		} else {
			for i := 0; i < n; i++ {
				heap(arr, n-1)
				if n%2 == 1 {
					arr[i], arr[n-1] = arr[n-1], arr[i]

				} else {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				}
			}
		}
	}
	heap(arr, len(arr))
	return res
}
