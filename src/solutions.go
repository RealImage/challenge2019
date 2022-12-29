package src

import (
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/shreeyashnaik/challenge2019/common/db"
	"github.com/shreeyashnaik/challenge2019/common/schemas"
	"github.com/shreeyashnaik/challenge2019/common/utils"
)

func ProblemStatementOne(inputPath, outputPath string) error {
	inputRows, err := utils.ReadCsv(inputPath)
	if err != nil {
		return err
	}

	outputLines := []string{}
	for _, row := range inputRows {
		theatreID := utils.Trim(row[2])
		size := utils.ToInt(row[1])

		costOfDelivery := math.MaxInt
		partnerID := "-1"
		for _, partnerDetail := range db.TheatrePartner[theatreID] {
			cost := ComputeCost(partnerDetail.SizeSlab, size, partnerDetail.CostPerGB, partnerDetail.MinimumCost)
			if costOfDelivery > cost {
				costOfDelivery = cost
				partnerID = partnerDetail.PartnerID
			}
			log.Println(cost)
		}

		outputLine := ""
		if partnerID != "-1" {
			outputLine = fmt.Sprintf("%s,true ,%s,%d", utils.Trim(row[0]), partnerID, costOfDelivery)
		} else {
			outputLine = fmt.Sprintf("%s,false,,", utils.Trim(row[0]))
		}

		outputLines = append(outputLines, outputLine)
	}

	// Save solution to my_output1.go
	if err := utils.WriteToCsv(outputPath, outputLines); err != nil {
		return err
	}

	return nil
}

func ProblemStatementTwo(inputPath, outputPath string) error {
	inputRows, err := utils.ReadCsv(inputPath)
	if err != nil {
		return err
	}

	// For each delivery_id in input.csv stores size_requirement, number of valid available Partner Options, details of each partner options.
	deliveryChoices := []schemas.DeliveryChoices{}
	for _, row := range inputRows {
		deliveryID := utils.Trim(row[0])
		size := utils.ToInt(row[1])
		theatreID := utils.Trim(row[2])

		choice := schemas.DeliveryChoices{
			DeliveryID:      deliveryID,
			SizeRequirement: int(size),
		}
		numberOfOptions := 0
		for _, partnerDetail := range db.TheatrePartner[theatreID] {
			cost := ComputeCost(partnerDetail.SizeSlab, size, partnerDetail.CostPerGB, partnerDetail.MinimumCost)

			if cost != math.MaxInt {
				numberOfOptions += 1
				choice.PartnerOptions = append(choice.PartnerOptions,
					struct {
						PartnerID      string
						CostOfDelivery int
					}{
						PartnerID:      partnerDetail.PartnerID,
						CostOfDelivery: cost,
					},
				)
			}
			log.Println(cost)
		}

		choice.NumberOfOptions = numberOfOptions
		deliveryChoices = append(deliveryChoices, choice)
	}

	// Sort the above array on the basis number of available options.
	sort.SliceStable(deliveryChoices, func(i, j int) bool {
		return deliveryChoices[i].NumberOfOptions < deliveryChoices[j].NumberOfOptions
	})

	log.Println(deliveryChoices)

	// Delivery with least number of options will get processed first.
	output := make([]string, len(inputRows))
	for _, delivery := range deliveryChoices {
		finalCost := math.MaxInt
		finalPartner := "-1"

		for _, partnerChoice := range delivery.PartnerOptions {
			if partnerChoice.CostOfDelivery < finalCost && delivery.SizeRequirement <= db.Capacities[partnerChoice.PartnerID] {
				finalCost = partnerChoice.CostOfDelivery
				finalPartner = partnerChoice.PartnerID

				db.Capacities[partnerChoice.PartnerID] -= delivery.SizeRequirement
			}
		}

		deliveryIdx := utils.ToInt(delivery.DeliveryID[1:]) - 1

		if delivery.NumberOfOptions > 0 && finalPartner != "-1" {
			output[deliveryIdx] = fmt.Sprintf("%s,true,%s,%d", delivery.DeliveryID, finalPartner, finalCost)
		} else {
			output[deliveryIdx] = fmt.Sprintf("%s,false,,", delivery.DeliveryID)
		}
	}

	// Save solution to my_output2.go
	if err := utils.WriteToCsv(outputPath, output); err != nil {
		return err
	}

	return nil
}
