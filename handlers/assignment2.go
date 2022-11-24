package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/niroopreddym/realimage/models"
	"github.com/niroopreddym/realimage/services"
)

type deliveryPartners struct {
	Partners []string
	Prices   []int
	Cost     int
}
type partnerinput struct {
	PartnerID string
	input     int
}

func loadPartnerData(fileName string) map[string][]*models.PartnerConfig {
	partners := services.ReadCSVRecordsPartners(fileName)
	data := map[string][]*models.PartnerConfig{}
	for _, partner := range partners {
		slabArr := strings.Split(partner.SizeSlabInGB, "-")
		data[partner.PartnerID] = append(data[partner.PartnerID], &models.PartnerConfig{
			TID:         partner.TheatreID,
			MinSlabSize: toInt(slabArr[0]),
			MaxSlabSize: toInt(slabArr[1]),
			MinCost:     partner.MinimumCost,
			CperGB:      partner.CostPerGB,
		})
	}

	return data
}

//Assignment2 ...
func Assignment2() {
	inputs := services.ReadCSVRecordsInputs("input.csv")
	partnerData := services.ReadCSVRecordsPartners("partners.csv")
	capacityProviders := services.ReadCSVRecordsCapacity("capacities.csv")

	providerNames := []string{}
	for _, capacityData := range capacityProviders {
		providerNames = append(providerNames, strings.TrimSpace(capacityData.PartnerID))
	}

	//find capacity providers for theatres trim down other thetares providers apart from input to reduce the iterations
	trimmedPartners := []*models.Partners{}
	for _, input := range inputs {
		for _, partner := range partnerData {
			if strings.TrimSpace(input.TheatreID) == strings.TrimSpace(partner.TheatreID) {
				trimmedPartners = append(trimmedPartners, partner)
			}
		}
	}

	//find cost for partnerID and Distrbution combination
	//calculate the capcity of input
	requiredCapacity := calculateCapacityRequired(inputs)
	inputPartnersCost := priceEngine(inputs, trimmedPartners, providerNames)

	deliveryPartnersEstimatedCost := getDeliveryCombinations(inputPartnersCost, providerNames, inputs)
	deliveryPartnersData := bestCostCombo(deliveryPartnersEstimatedCost, capacityProviders, requiredCapacity)

	output2 := getOutputBasedOnPriceFeasibility(deliveryPartnersData, inputs, trimmedPartners)
	fmt.Println("assigment2 Out:", output2)
	services.WriteDataToCSV("output2.csv", output2)
}

func getOutputBasedOnPriceFeasibility(deliveryPartnersData deliveryPartners, inputs []*models.Input, partners []*models.Partners) []models.Output {
	lstOutputs := []models.Output{}
	for index, input := range inputs {
		partnerID := deliveryPartnersData.Partners[index]
		costOfDelivery := strconv.Itoa(deliveryPartnersData.Prices[index])
		fmt.Println("totalcost:", costOfDelivery)

		possibleDelivery := findDeliveryPossibility(input, partners, partnerID)
		if !possibleDelivery {
			partnerID = " "
			costOfDelivery = " "
		}

		output := models.Output{
			DeliveryID:       input.DeliveryID,
			DeliveryPossible: possibleDelivery,
			PartnerID:        partnerID,
			CostOfDelivery:   costOfDelivery,
		}

		lstOutputs = append(lstOutputs, output)
	}

	return lstOutputs
}

func findDeliveryPossibility(input *models.Input, partners []*models.Partners, partnerID string) bool {
	for _, partner := range partners {
		if strings.TrimSpace(partner.TheatreID) == input.TheatreID && strings.TrimSpace(partner.PartnerID) == partnerID {
			minCost := partner.MinimumCost
			if minCost > input.SizeOfDelivery*partner.CostPerGB {
				return false
			}

			return true
		}
	}

	return true
}

//returns the best Combo by verifying their capacity and cost
func bestCostCombo(deliveryPartnersEstimatedCost []deliveryPartners, capacityProviders []*models.Capacity, requiredCapacity int) deliveryPartners {
	expectedCombo := deliveryPartners{}
	minValue := 999999
	for _, deliveryPartner := range deliveryPartnersEstimatedCost {
		isValidCombo := checkCapacity(deliveryPartner.Partners, capacityProviders, requiredCapacity)
		if isValidCombo && minValue > deliveryPartner.Cost {
			minValue = deliveryPartner.Cost
			expectedCombo = deliveryPartner
		}
	}

	return expectedCombo
}

func checkCapacity(combo []string, capacityProviders []*models.Capacity, requiredCapacity int) bool {
	capMap := map[string]int{}
	for _, value := range capacityProviders {
		capMap[strings.TrimSpace(value.PartnerID)] = value.CapacityInGB
	}

	occuredMap := map[string]bool{}
	capacity := 0
	for _, provider := range combo {
		if _, isExists := occuredMap[provider]; !isExists {
			capacity = capacity + capMap[provider]
			occuredMap[provider] = true
		}
	}

	if capacity < requiredCapacity {
		return false
	}

	return true
}

func getDeliveryCombinations(inputPartnersCost map[partnerinput]int, patnerIDs []string, inputs []*models.Input) []deliveryPartners {
	partnerPlacementCount := len(inputs)
	arr := make([]string, partnerPlacementCount)

	lstCombos := [][]string{}
	permutate(partnerPlacementCount, arr, 0, patnerIDs, &lstCombos)

	//using lstcombo and input find the cost in inputPartnersCost and calc total cost
	lstDeliveryPartnersCost := []deliveryPartners{}
	for _, combo := range lstCombos {
		priceCombos := []int{}
		cost := 0
		for index, partnerID := range combo {
			p := partnerinput{
				input:     inputs[index].SizeOfDelivery,
				PartnerID: partnerID,
			}

			priceCombos = append(priceCombos, inputPartnersCost[p])
			cost = cost + inputPartnersCost[p]
		}

		lstDeliveryPartnersCost = append(lstDeliveryPartnersCost, deliveryPartners{
			Partners: combo,
			Prices:   priceCombos,
			Cost:     cost,
		})
	}

	return lstDeliveryPartnersCost
}

func permutate(n int, arr []string, index int, capacityProvidors []string, comboList *[][]string) {
	if index == n {
		combo := make([]string, n)
		copy(combo, arr)
		*comboList = append(*comboList, combo)
		return
	}

	for _, value := range capacityProvidors {
		arr[index] = value
		permutate(n, arr, index+1, capacityProvidors, comboList)
	}

}

func priceEngine(inputs []*models.Input, partners []*models.Partners, patnerIds []string) map[partnerinput]int {
	output := map[partnerinput]int{}
	for _, input := range inputs {
		for _, partnerID := range patnerIds {
			cost := calculatePrice(input, partners, partnerID)
			p := partnerinput{
				input:     input.SizeOfDelivery,
				PartnerID: partnerID,
			}

			output[p] = cost
		}
	}

	return output
}

func calculatePrice(input *models.Input, partners []*models.Partners, partnerID string) int {
	sizeOfDelivery := input.SizeOfDelivery
	for _, partner := range partners {
		min, max := slabRange(partner.SizeSlabInGB)
		if sizeOfDelivery >= min && sizeOfDelivery <= max && partnerID == partner.PartnerID {
			return sizeOfDelivery * partner.CostPerGB
		}
	}

	return -1
}

func slabRange(slabRange string) (int, int) {
	arr := strings.Split(slabRange, "-")
	return toInt(arr[0]), toInt(arr[1])
}

func calculateCapacityRequired(inputs []*models.Input) int {
	total := 0
	for _, input := range inputs {
		total = total + input.SizeOfDelivery
	}
	return total
}
