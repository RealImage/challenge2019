package Prob

import (
	"fmt"
	"strings"
)

func UpdatePartnersCapacity(allApplicablePartners []DeliveryAndPartners, capacityInfo []CapacityInfo) {
	for i, a := range allApplicablePartners {
		for j, vv := range a.Partners {
			for _, pp := range capacityInfo {
				if strings.TrimSpace(pp.PartnerID) == vv.PartnerID {
					allApplicablePartners[i].Partners[j].Capacity = ConvertToInt(pp.Capacity)
				}
			}
		}
	}
	//Remove
	dpMap := MapPartnerToDelivery(allApplicablePartners, capacityInfo)
	Calculate(dpMap)
}

func MapPartnerToDelivery(allApplicablePartners []DeliveryAndPartners, partners []CapacityInfo) map[CapacityDetails][]DelAndPartners {

	dpMap := make(map[CapacityDetails][]DelAndPartners)
	for pp, _ := range partners {
		tmp := make([]DelAndPartners, 0)
		for _, r := range allApplicablePartners {
			for _, l := range r.Partners {
				if strings.TrimSpace(partners[pp].PartnerID) == l.PartnerID {
					t := DelAndPartners{r.Delivery.DeliveryID, r.Delivery.DeliverySize, l.Theatre, l.Capacity}
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

		for _, r := range v {
			sum = sum + r.DeliverySize
			if sum >= k.Capacity {
				break
			}
			arr = append(arr, r)
		}

		finMap[CapacityDetails{k.PartnerID, sum}] = arr
	}
	fmt.Println(finMap)
	fmt.Println(dpMap)
}
