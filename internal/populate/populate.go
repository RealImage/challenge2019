package populate

import (
	"log"
	"strings"

	"github.com/purush7/challenge2019/v1/types"
	"github.com/purush7/challenge2019/v1/util"
)

func PopulatePartnersData(partnersFile string) (data types.WholeData) {

	data = make(types.WholeData)
	records := util.ReadCSV(partnersFile)

	for _, record := range records[1:] {
		theatre := types.Theartre(strings.TrimSpace(record[0]))
		_, exists := data[theatre]
		if !exists {
			data[theatre] = make(types.PartnersData)
		}
		partnerData := data[theatre]

		slabRange := strings.Split(strings.TrimSpace(record[1]), "-")
		if len(slabRange) != 2 {
			log.Fatal("slabRange len isn't 2 of having range value", record[1], "and theartre", theatre)
		}
		minRange := util.StringToInt(slabRange[0])
		maxRange := util.StringToInt(slabRange[1])

		minCost := util.StringToInt(strings.TrimSpace(record[2]))
		costPerGB := util.StringToInt(strings.TrimSpace(record[3]))
		partner := types.Partner(strings.TrimSpace(record[4]))
		newSlab := types.Slab{MinRange: minRange, MaxRange: maxRange, MinCost: minCost, CostGB: costPerGB}
		partnerData[partner] = append(partnerData[partner], newSlab)
	}
	return data
}

func PopulateCapacitiesData(capacitiesFile string) (data types.CapacityData) {
	data = make(types.CapacityData)
	records := util.ReadCSV(capacitiesFile)

	for _, record := range records[1:] {
		partner := types.Partner(strings.TrimSpace(record[0]))
		data[partner] = util.StringToInt(strings.TrimSpace(record[1]))
	}
	return data
}
