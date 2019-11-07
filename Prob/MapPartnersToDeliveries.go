package Prob

import (
	"fmt"
	"strings"
)

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
	for k, v := range DToPMap {
		fmt.Println(k, v)
	}
}
