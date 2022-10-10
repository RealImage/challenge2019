package solve

import (
	"challenge2019/models"
	"fmt"
)

func Solution(f *models.FileDetails) error {
	//Getting input partners and capacities from csv files
	input, err := getInput(f.Input)
	if err != nil {
		return fmt.Errorf("solve/Solution(): error reading input: \n %w", err)
	}
	partners, err := getPartners(f.Partners)
	if err != nil {
		return fmt.Errorf("solve/Solution(): error reading partners: \n %w", err)
	}
	capacityMap, err := getCapacities(f.Capacities)
	if err != nil {
		return fmt.Errorf("solve/Solution(): error reading capacities: \n %w", err)
	}

	output1Map := make([]models.OutputDetails, len(*input))
	output2Map := make([]models.OutputDetails, len(*input))

	//Solving Problem Statements
	for i, in := range *input {
		output1 := models.OutputDetails{
			DeliveryID:  in.DeliveryID,
			Feasibility: false,
			PartnerID:   "",
			Cost:        0,
		}
		output2 := models.OutputDetails{
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
					//solving Problem Statement2
					if in.Size <= capacityMap[FesibilityData.PartnerID] {
						if output2.Cost != 0 {
							if output2.Cost > cost {
								output2.Cost = cost
								output2.PartnerID = FesibilityData.PartnerID
							}
						} else {
							output2.Cost = cost
							output2.PartnerID = FesibilityData.PartnerID
							output2.Feasibility = true
						}
					}
				}

			}
		}
		output1Map[i] = output1
		output2Map[i] = output2
	}
	if err := generateOut(f.Solution1, output1Map); err != nil {
		return fmt.Errorf("solve/Solution(): error generating output files:%s \n %w", f.Solution1, err)
	}
	if err := generateOut(f.Solution2, output2Map); err != nil {
		return fmt.Errorf("solve/Solution(): error generating output files:%s \n %w", f.Solution2, err)
	}
	return nil
}
