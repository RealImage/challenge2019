package Prob

import "challenge2019/Prob/types"

func FindTotalDeliveryCharge(fPartners []types.PartnerData) int {
	sum := 0
	for _, i := range fPartners {
		sum = sum + int(i.DeliveryCost)
	}
	return sum
}
