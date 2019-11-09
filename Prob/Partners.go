package Prob

import (
	"challenge2019/Prob/types"
	"fmt"
)

func FindBestApplicablePartner(data []types.PartnerData, input []types.DeliveryInfo) ([]types.FinalChoice, []types.DeliveryAndPartners) {
	bestPartner, allApplicablePartners := findBestApplicablePartner(data, input)
	return bestPartner, allApplicablePartners
}

func findBestApplicablePartner(data []types.PartnerData, input []types.DeliveryInfo) ([]types.FinalChoice, []types.DeliveryAndPartners) {
	dAndp := make([]types.DeliveryAndPartners, 0)
	fin := make([]types.FinalChoice, 0)
	for v, i := range input {
		applicablePartners := make([]types.PartnerData, 0)
		for _, r := range data {
			if r.Theatre == i.Theatre && (r.Size.Min <= i.DeliverySize) && (i.DeliverySize <= r.Size.Max) {
				r.TotalCost = r.CostPerGB * float64(i.DeliverySize)
				r.DeliveryCost = CalculateDeliveryCost(r.TotalCost, r.MinCost)
				r.Delivery = i
				applicablePartners = append(applicablePartners, r)
			}
		}
		dAndp = append(dAndp, types.DeliveryAndPartners{i, applicablePartners})
		best := FindBestPartner(dAndp[v])
		if len(best.Partners) == 0 {
			fin = append(fin, types.FinalChoice{dAndp[v].Delivery.DeliveryID, false, "", ""})
		} else {
			fin = append(fin, types.FinalChoice{dAndp[v].Delivery.DeliveryID, true, best.Partners[0].PartnerID, fmt.Sprintf("%.0f", best.Partners[0].DeliveryCost)})

		}
	}
	return fin, dAndp
}

func CalculateDeliveryCost(total, min float64) float64 {
	if total > min {
		return total
	}
	return min
}

func FindBestPartner(d types.DeliveryAndPartners) types.DeliveryAndPartners {
	var fin types.DeliveryAndPartners
	if len(d.Partners) == 0 {
		return fin
	}
	min := d.Partners[0].TotalCost
	best := d.Partners[0]
	for i, _ := range d.Partners {
		if min > d.Partners[i].TotalCost {
			min = d.Partners[i].TotalCost
			best = d.Partners[i]
		}
	}
	fin.Partners = append(fin.Partners, best)
	fin.Delivery = d.Delivery

	return fin
}

/*
func findBestApplicablePartnerV2(data []types.PartnerData, input []types.DeliveryInfo) ([]types.FinalChoice, []types.DeliveryAndPartners) {
	dAndp := make([]types.DeliveryAndPartners, 0)
	fin := make([]types.FinalChoice, 0)
	for v, i := range input {
		applicablePartners := make([]types.PartnerData, 0)
		for _, r := range data {
			sum1 := 0.0
			sum2 := 0.0
			if r.Theatre == i.Theatre && (r.Size.Min <= i.DeliverySize) && (i.DeliverySize <= r.Size.Max) {
				v1 := i.DeliverySize - r.Size.Min
				fmt.Println("Total :", i.DeliverySize, "Diff :", v1, "Remaining :", r.Size.Min)
				sum1 = CalculateDeliveryCost(r.CostPerGB*float64(v1), r.MinCost)
				fmt.Println("Range :", r)
				fmt.Println("Sum :", sum1)
				for v := 0; v < len(data); v++ {
					if r.Theatre == data[v].Theatre && r.Size.Min == data[v].Size.Max {
						sum2 = CalculateDeliveryCost(data[v].CostPerGB*float64(r.Size.Min), data[v].MinCost)
						fmt.Println("Range :", data[v])
						fmt.Println("Sum :", sum2)

					}
				}
				sum := sum1 + sum2
				fmt.Println("Total :", sum)
				r.DeliveryCost = sum
				r.Delivery = i
				applicablePartners = append(applicablePartners, r)
			}
		}
		dAndp = append(dAndp, types.DeliveryAndPartners{i, applicablePartners})
		best := FindBestPartner(dAndp[v])
		if len(best.Partners) == 0 {
			fin = append(fin, types.FinalChoice{dAndp[v].Delivery.DeliveryID, false, "", ""})
		} else {
			fin = append(fin, types.FinalChoice{dAndp[v].Delivery.DeliveryID, true, best.Partners[0].PartnerID, fmt.Sprintf("%.0f", best.Partners[0].DeliveryCost)})

		}
	}
	return fin, dAndp
}
*/
