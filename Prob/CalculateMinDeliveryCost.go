package Prob

import "fmt"

func FindAllPartnerInfo(data []PartnerData) ([]FinalChoice, []DeliveryAndPartners) {
	input := []DeliveryInfo{
		{"D1", 150, "T2"},
		{"D2", 325, "T2"},
		{"D3", 510, "T2"},
		{"D4", 700, "T2"},
	}
	bestPartner, allApplicablePartners := findAllPartnerInfo(data, input)
	return bestPartner, allApplicablePartners
}

func findAllPartnerInfo(data []PartnerData, input []DeliveryInfo) ([]FinalChoice, []DeliveryAndPartners) {
	dAndp := make([]DeliveryAndPartners, 0)
	fin := make([]FinalChoice, 0)
	for v, i := range input {
		applicablePartners := make([]PartnerData, 0)
		for _, r := range data {
			if r.Theatre == i.Theatre && (r.Size.Min <= i.DeliverySize) && (i.DeliverySize <= r.Size.Max) {
				r.TotalCost = r.CostPerGB * float64(i.DeliverySize)
				r.DeliveryCost = CalculateDeliveryCost(r.TotalCost, r.MinCost)
				applicablePartners = append(applicablePartners, r)
			}
		}
		dAndp = append(dAndp, DeliveryAndPartners{i, applicablePartners})
		best := FindBestPartner(dAndp[v])
		if len(best.Partners) == 0 {
			fin = append(fin, FinalChoice{dAndp[v].Delivery.DeliveryID, false, "", ""})
		} else {
			fin = append(fin, FinalChoice{dAndp[v].Delivery.DeliveryID, true, best.Partners[0].PartnerID, fmt.Sprintf("%f", best.Partners[0].DeliveryCost)})
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

func FindBestPartner(d DeliveryAndPartners) DeliveryAndPartners {
	var fin DeliveryAndPartners
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
