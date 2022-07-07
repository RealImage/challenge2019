package algos

import (
	"log"
	"strconv"

	"github.com/purush7/challenge2019/v1/internal/populate"
	"github.com/purush7/challenge2019/v1/types"
	"github.com/purush7/challenge2019/v1/util"
)

type findPartnerOps struct {
	data     types.WholeData
	content  int
	theartre types.Theartre
	possible bool
	partner  types.Partner
	cost     int
}

func BestPartner(ops types.ProblemOps) {
	//populate data

	data := populate.PopulateData(ops.PartnersFile)
	inputRecords := util.ReadCSV(ops.InputFile)

	var outputRecords [][]string

	partnerOps := findPartnerOps{data: data}

	for _, record := range inputRecords {
		outputRecord := make([]string, 4)
		outputRecord[0] = record[0]
		content := util.StringToInt(record[1])
		theartre := types.Theartre(record[2])

		//fill partnerOps
		partnerOps.content = content
		partnerOps.theartre = theartre
		partnerOps.possible = false
		partnerOps.partner = types.Partner("")
		partnerOps.cost = -1

		findPartner(&partnerOps)

		outputRecord[1] = strconv.FormatBool(partnerOps.possible)
		if partnerOps.cost == -1 {
			outputRecord[3] = ""
			outputRecord[2] = ""
		} else {
			outputRecord[2] = string(partnerOps.partner)
			outputRecord[3] = strconv.Itoa(partnerOps.cost)
		}
		outputRecords = append(outputRecords, outputRecord)
	}
	util.WriteCSV(ops.OutputFile, outputRecords)
}

func findPartner(ops *findPartnerOps) {

	if ops == nil {
		return
	}
	_, exits := ops.data[ops.theartre]

	if !exits {
		log.Fatal("theartre", ops.theartre, "isn't present in the partners file provided")
	}

	partnersData := ops.data[ops.theartre]
	content := ops.content

	var found bool = false
	var minCost int = 0
	var cost int = 0
	var tmpCost = 0
	var finalPartner types.Partner

	for partner, slabSlice := range partnersData {
		for _, slab := range slabSlice {
			if content >= slab.MinRange && content <= slab.MaxRange {
				tmpCost = content * slab.CostGB
				if tmpCost >= slab.MinCost {
					cost = tmpCost
				} else {
					cost = slab.MinCost
				}
				if !found || minCost >= cost {
					minCost = cost
					finalPartner = partner
				}
				found = true
			}
		}
	}

	ops.possible = found
	if found {
		ops.cost = minCost
		ops.partner = finalPartner
	} else {
		ops.partner = types.Partner(``)
		ops.cost = -1
	}

}
