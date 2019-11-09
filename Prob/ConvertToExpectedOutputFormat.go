package Prob

import (
	"challenge2019/Prob/types"
	"fmt"
)

func Output(feasibleArray [][]types.PartnerData, inputData []types.DeliveryInfo) []types.FinalChoice {
	result := []types.FinalChoice{}
	for _, r := range feasibleArray[0] {
		result = append(result, types.FinalChoice{
			DeliveryID:    r.Delivery.DeliveryID,
			IsPossible:    true,
			ChosenPartner: r.PartnerID,
			MinimumCost:   fmt.Sprintf("%.0f", r.DeliveryCost),
		},
		)
	}

	if len(inputData) > len(feasibleArray[0]) {
		for _, v := range inputData {
			isExists := CheckExistence(v.DeliveryID, result)
			if !isExists {
				result = append(result, types.FinalChoice{
					DeliveryID:    v.DeliveryID,
					IsPossible:    false,
					ChosenPartner: "",
					MinimumCost:   "",
				},
				)
			}
		}

	}
	return result
}
func DeliverySet(a1 []types.DeliveryInfo) (dArray []string) {
	for _, r := range a1 {
		dArray = append(dArray, r.DeliveryID)
	}
	return
}

func CheckExistence(a string, res []types.FinalChoice) bool {
	for _, r := range res {
		if r.DeliveryID == a {
			return true
		}
	}
	return false
}
