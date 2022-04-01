package assignment

import (
	"fmt"
	"github.com/sureshn04/challenge2019/internal/io"
	"github.com/sureshn04/challenge2019/internal/utils"
	"github.com/sureshn04/challenge2019/models"
	"strconv"
	"strings"
)

func First() {
	partners := io.ReadPartners("input/partners.csv")
	inputs := io.ReadInput("input/input.csv")

	output := minDeliveryCost(inputs, partners)
	fmt.Println("assigment1 Result:\n", output)
	io.Write("output/output1.csv", output)
}

func minDeliveryCost(inputs []*models.Input, partners []*models.Partners) []models.Output {
	var output []models.Output
	for _, input := range inputs {
		mincost, partnerID := "", ""
		for _, partner := range partners {
			if strings.TrimSpace(input.TheatreID) == strings.TrimSpace(partner.TheatreID) && utils.CheckSlab(partner.SizeSlabInGB, input.Cost) {
				computedCost := input.Cost * partner.CostPerGB
				if computedCost < partner.MinimumCost {
					computedCost = partner.MinimumCost
				}

				if mincost == "" || computedCost < utils.ToInt(mincost) {
					mincost = strconv.Itoa(computedCost)
					partnerID = strings.TrimSpace(partner.PartnerID)
				}
			}
		}

		result := models.Output{
			DistributorID: strings.TrimSpace(input.DistributorID),
		}

		if mincost == "" {
			result.Accepted = false
			result.PartnerID = " "
			result.Cost = " "
		} else {
			result.Accepted = true
			result.PartnerID = partnerID
			result.Cost = mincost
		}

		output = append(output, result)

	}

	return output
}
