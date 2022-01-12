package handlers

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/niroopreddym/realimage/models"
	"github.com/niroopreddym/realimage/services"
)

//Assignment1 ...
func Assignment1() {
	partners := services.ReadCSVRecordsPartners("partners.csv")
	inputs := services.ReadCSVRecordsInputs("input.csv")

	output1 := deliveryEngine(inputs, partners)
	fmt.Println("assigment1 Out:", output1)
	services.WriteDataToCSV("output1.csv", output1)
}

//DeliveryEngine decides whether the delivery is possible or not
func deliveryEngine(inputs []*models.Input, partners []*models.Partners) []models.Output {
	var output []models.Output
	for _, input := range inputs {
		mincost := 0
		partnerID := ""
		for _, partner := range partners {
			if strings.TrimSpace(input.TheatreID) == strings.TrimSpace(partner.TheatreID) && checkSlab(partner.SizeSlabInGB, input.SizeOfDelivery) {
				computedCost := input.SizeOfDelivery * partner.CostPerGB
				if computedCost < partner.MinimumCost {
					computedCost = partner.MinimumCost
				}

				if mincost == 0 || computedCost < mincost {
					mincost = computedCost
					partnerID = strings.TrimSpace(partner.PartnerID)
				}

				if computedCost < 2000 {
					mincost = 2000
					partnerID = strings.TrimSpace(partner.PartnerID)
				}
			}
		}

		result := models.Output{
			DeliveryID: strings.TrimSpace(input.DeliveryID),
		}

		if mincost == 0 {
			result.DeliveryPossible = false
			result.PartnerID = " "
			result.CostOfDelivery = " "
		} else {
			result.DeliveryPossible = true
			result.PartnerID = partnerID
			result.CostOfDelivery = strconv.Itoa(mincost)
		}

		output = append(output, result)
	}

	return output
}

func toInt(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Fatal(err)
	}

	return n
}

func checkSlab(slab string, value int) bool {
	slabArr := strings.Split(slab, "-")
	if value >= toInt(slabArr[0]) && value <= toInt(slabArr[1]) {
		return true
	}
	return false
}
