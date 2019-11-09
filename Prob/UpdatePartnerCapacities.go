package Prob

import (
	"challenge2019/Prob/types"
	"strings"
)

func UpdatePartnersCapacity(allApplicablePartners []types.DeliveryAndPartners, capacityInfo []types.CapacityDetailsStr) []types.DeliveryAndPartners {
	for i, a := range allApplicablePartners {
		for j, vv := range a.Partners {
			for _, pp := range capacityInfo {
				if strings.TrimSpace(pp.PartnerID) == vv.PartnerID {
					allApplicablePartners[i].Partners[j].Capacity = types.ConvertToInt(pp.Capacity)
				}
			}
		}
	}
	return allApplicablePartners
}
