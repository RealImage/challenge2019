package populate

import (
	"github.com/purush7/challenge2019/v1/types"
	"github.com/purush7/challenge2019/v1/util"
)

func PopulateData(partnersFile string) (data types.WholeData) {

	records := util.ReadCsv(partnersFile)

	for _, record := range records {
		theatre := types.Theartre(record[0])
		//Todo: check if below code is required or not?
		// _, exists := data[theatre]
		// if !exists {
		// 	data[theatre] = types.PartnersData{}
		// }
		partnerData := data[theatre]

		//slabRange := record[1]
		//Todo: split the slab range to min and max
		minRange := 0
		maxRange := 0

		minCost := util.StringToInt(record[2])
		costPerGB := util.StringToInt(record[3])
		partner := types.Partner(record[4])
		newSlab := types.Slabs{minRange, maxRange, minCost, costPerGB}
		partnerData[partner] = append(partnerData[partner], newSlab)

	}

	return data
}
