package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/niroopreddym/realimage/models"
	"github.com/niroopreddym/realimage/services"
)

var cap = make(map[string]int)

//Assignment2 ..
func Assignment2() {
	inputs := services.ReadCSVRecordsInputs("input.csv")
	partnerData := loadPartnerData("partners.csv")
	partnerids := []string{}
	for key := range partnerData {
		partnerids = append(partnerids, key)
	}

	distribution := make([][]int, len(inputs))
	for inputIndex, inputData := range inputs {
		distribution[inputIndex] = make([]int, len(partnerids))
		for partnerIndex, partnerID := range partnerids {
			distribution[inputIndex][partnerIndex] = findCost(partnerData, inputData, partnerID)
		}
	}

	for _, value := range distribution {
		fmt.Println(value)
	}

	cummilative(distribution, partnerids, inputs, partnerData)
}

func cummilative(distributions [][]int, partnerIds []string, inputs []*models.Input, partnerData map[string][]*models.PartnerConfig) {
	var res []models.Group
	var length = 0
	for _, dist := range distributions {
		res = add(res, dist, length, partnerIds)
		length = len(res)
	}

	for index, value := range res {
		if !checkOrder(value.Order, inputs) {
			res[index].Check = false
		}
	}

	var mincost int = 99999999
	var final string
	final += ""
	for _, value := range res {
		if value.Check && mincost > value.Sum {
			mincost = value.Sum
			final = value.Order
		}
	}

	fmt.Println(final)

	var partnerRes []string
	for i := 2; i <= len(final); i = i + 2 {
		partnerRes = append(partnerRes, final[i-2:i])
	}

	var output1 [][]string

	for index, inputData := range inputs {
		//var resultrow []string
		resultrow := []string{inputData.DistributorID}
		if partnerRes[index] == "  " {
			resultrow = append(resultrow, "false", ` `, ` `)
		} else {
			resultrow = append(resultrow, "true", partnerRes[index], strconv.Itoa(findCost(partnerData, inputData, partnerRes[index])))
			//uncomment above linbe
		}
		output1 = append(output1, resultrow)
	}

	fmt.Println(output1)
	fmt.Println("assigment2 Out:", output1)
	// OutputWriter("output2.csv", output1)

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

func findCost(data map[string][]*models.PartnerConfig, input *models.Input, pid string) int {
	configarr := data[pid]
	for _, value := range configarr {
		if input.MinCost >= value.MinSlabSize && input.MinCost <= value.MaxSlabSize {
			c := value.CperGB * input.MinCost
			if c <= value.MinCost {
				return value.MinCost
			}
			return c
		}
	}
	return -1
}

func loadcap() {

	capacities := services.ReadCSVRecordsCapacity("capacities.csv")
	for _, capacitiesRow := range capacities {
		cap[strings.TrimSpace(capacitiesRow.PartnerID)] = capacitiesRow.CapacityInGB
	}
}

func checkOrder(seq string, inputs []*models.Input) bool {
	defer loadcap()
	for i := 2; i <= len(seq); i = i + 2 {
		if seq[i-2:i] == "  " {
			return true
		}
		value := cap[seq[i-2:i]] - inputs[(i-2)/2].MinCost
		if value < 0 {
			return false
		}
		cap[seq[i-2:i]] = value
	}
	return true
}

func add(result []models.Group, input []int, length int, partnerids []string) []models.Group {
	if len(result) != 0 {
		presult := result
		result = result[length:]
		for _, p1 := range presult {
			for index, p2 := range input {
				if p2 == -1 {
					if check(input) {
						result = append(result, models.Group{Sum: p1.Sum + 0, Order: p1.Order + "  ", Check: true})
					} else {
						result = append(result, models.Group{Sum: p1.Sum + p2, Order: p1.Order + partnerids[index], Check: false})

					}
					continue
				}
				result = append(result, models.Group{Sum: p1.Sum + p2, Order: p1.Order + partnerids[index], Check: true})
			}
		}
	} else {
		for index, p2 := range input {
			result = append(result, models.Group{Sum: p2, Order: partnerids[index], Check: true})
		}
	}

	return result
}

func check(input []int) bool {
	for _, v := range input {
		if v != -1 {
			return false
		}
	}
	return true
}
