package p2

import (
	"challenge2019/src/models"
	"challenge2019/src/service"
)

func buildInputQueue(inputPath string) models.MaxQueue {
	inputQueue := models.NewMaxQueue()
	inputList := service.ReadInput(inputPath)
	for _, input := range inputList {
		inputQueue.Insert(input)
	}
	return inputQueue
}

func buildPartnersQueue(partnerList []models.PartnerRecord, input models.Input) models.MinQueue {
	partnerQueue := models.NewMinQueue()
	for _, partner := range partnerList {
		if input.TheatreID == partner.TheatreID {
			if partner.Min < input.Volume && partner.Max >= input.Volume {
				partnerQueue.Insert(partner)
			}
		}
	}
	return partnerQueue
}

func Soultion(partnerPath, inputPath string, capcitiesPath string, outputPath string) {
	partnerList := service.ReadPartnerCsv(partnerPath)
	inputQueue := buildInputQueue(inputPath)
	capacities := service.ReadCapacities(capcitiesPath)
	outputList := make([]models.Output, 0, inputQueue.Size())
	for inputQueue.GetMax() != nil {
		input := inputQueue.ExtractMax()
		partnerQueue := buildPartnersQueue(partnerList, *input)
		if partnerQueue.Size() == -1 {
			output := models.NewOutput(input.DeliveryID, false, "", 0)
			outputList = append(outputList, output)
			continue
		}
		for partnerQueue.GetMin() != nil {
			partner := partnerQueue.ExtractMin()
			if input.Volume <= capacities[partner.PartnerID] {
				totalCost := partner.CostPerGB * input.Volume
				if totalCost < partner.MinCost {
					totalCost = partner.MinCost
				}
				output := models.NewOutput(input.DeliveryID, true, partner.PartnerID, totalCost)
				outputList = append(outputList, output)
				capacities[partner.PartnerID] = capacities[partner.PartnerID] - input.Volume
				break
			}
		}
	}
	service.WriteOutput(outputPath, outputList)
}
