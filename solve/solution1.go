package solve

import "challenge2019/models"

func solution1(input []models.InputDetails, partners models.TheatreMap) ([]models.OutputDetails, map[string]models.TotalDataPerPartner) {
	//Solving Problem Statements
	output1Map := make([]models.OutputDetails, len(input))
	totalDataPerPartner := make(map[string]models.TotalDataPerPartner)
	for i, in := range input {
		output1 := models.OutputDetails{
			DeliveryID:  in.DeliveryID,
			Feasibility: false,
			PartnerID:   "",
			Cost:        0,
		}
		for _, theatreData := range partners[in.TheatreID] {
			for _, FesibilityData := range theatreData {
				//solving Problem Statement 1
				if in.Size >= FesibilityData.SizeSlab.Min && in.Size <= FesibilityData.SizeSlab.Max {
					cost := FesibilityData.CostPerGB * in.Size
					if cost < FesibilityData.MinimumCost {
						cost = FesibilityData.MinimumCost
					}
					if output1.Cost != 0 {
						if output1.Cost > cost {
							output1.Cost = cost
							output1.PartnerID = FesibilityData.PartnerID
						}
					} else {
						output1.Cost = cost
						output1.PartnerID = FesibilityData.PartnerID
						output1.Feasibility = true
					}
				}

			}
		}
		output1Map[i] = output1
		if output1.PartnerID == "" {
			continue
		}
		totaldata, ok := totalDataPerPartner[output1.PartnerID]
		if ok {
			totaldata.Data += in.Size
			totaldata.DataUnitsMap[in.Size] = i
			totalDataPerPartner[output1.PartnerID] = totaldata
		} else {
			newData := models.TotalDataPerPartner{}
			newData.Data = in.Size
			dataMap := make(map[int]int)
			dataMap[in.Size] = i
			newData.DataUnitsMap = dataMap
			totalDataPerPartner[output1.PartnerID] = newData
		}

	}
	return output1Map, totalDataPerPartner
}
