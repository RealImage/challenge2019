package algos

import (
	"strconv"
	"strings"

	"github.com/purush7/challenge2019/v1/internal/populate"
	"github.com/purush7/challenge2019/v1/types"
	"github.com/purush7/challenge2019/v1/util"
)

func MinimizeCost(ops types.ProblemOps) {
	data := populate.PopulatePartnersData(ops.PartnersFile)
	cdata := populate.PopulateCapacitiesData(ops.CapacitiesFile)

	var partners []types.Partner
	partnerPosition := make(map[types.Partner]int)
	var i int = 0
	for part := range cdata {
		partnerPosition[part] = i
		partners = append(partners, part)
		i++
	}

	inputRecords := util.ReadCSV(ops.InputFile)
	partnerSize := len(partners)
	deliverySize := len(inputRecords)
	table := make([][]int, deliverySize)

	//populate table
	for ind, del := range inputRecords {
		table[ind] = make([]int, partnerSize)
		theartre := types.Theartre(strings.TrimSpace(del[2]))
		partnersData := data[theartre]
		for partnerInd, partner := range partners {
			table[ind][partnerInd] = findMinCost(util.StringToInt(strings.TrimSpace(del[1])), partnersData[partner])
		}
	}

	//populateAllPossibilites
	var allComb []types.Combination
	var leng = 0
	for _, row := range table {
		allComb = add(allComb, row, leng, partners)
		leng = len(allComb)
	}

	//checkCombinations
	for ind, comb := range allComb {
		if !checkComb(comb.PartnerComb, inputRecords, cdata) {
			comb.Possible = false
		}
		allComb[ind] = comb
	}
	var found bool = false
	var minCost int = 0
	var finalComb string = ""

	// get finalCombination
	mnUndel := 9999999
	for _, comb := range allComb {
		if comb.Possible {
			if mnUndel > comb.Undel {
				mnUndel = comb.Undel
			}
		}
	}

	for _, comb := range allComb {
		if comb.Possible && mnUndel == comb.Undel && (!found || minCost > comb.Cost) {
			minCost = comb.Cost
			finalComb = comb.PartnerComb
			found = true
			continue
		}
	}

	//result
	res := strings.Split(finalComb, "_")
	var outputRecords [][]string
	for i := 0; i < deliverySize; i++ {
		var outputRecord []string
		outputRecord = append(outputRecord, strings.TrimSpace(inputRecords[i][0]))
		partString := res[i]
		if partString == "" {
			outputRecord = append(outputRecord, strconv.FormatBool(false))
			outputRecord = append(outputRecord, partString)
			outputRecord = append(outputRecord, string(""))
		} else {
			outputRecord = append(outputRecord, strconv.FormatBool(true))
			outputRecord = append(outputRecord, partString)
			outputRecord = append(outputRecord, strconv.Itoa(table[i][partnerPosition[types.Partner(partString)]]))
		}
		outputRecords = append(outputRecords, outputRecord)
	}

	//write

	util.WriteCSV(ops.OutputFile, outputRecords)
}

func checkComb(comb string, inputData [][]string, cdata types.CapacityData) bool {

	content := make(map[types.Partner]int)
	res := strings.Split(comb, "_")

	for ind := range res {
		if res[ind] == "" {
			continue
		}
		part := types.Partner(res[ind])
		content[part] += util.StringToInt(strings.TrimSpace(inputData[ind][1]))
		if content[part] > cdata[part] {
			return false
		}
	}

	return true
}

func add(allComb []types.Combination, row []int, leng int, partners []types.Partner) []types.Combination {
	if leng != 0 {
		pres := allComb
		allComb = allComb[leng:]

		for _, preComb := range pres {
			pos := checkRow(row)
			for ind, p2 := range row {
				if p2 == -1 {
					if pos {
						allComb = append(allComb, types.Combination{Cost: preComb.Cost, PartnerComb: preComb.PartnerComb + "_", Possible: true, Undel: preComb.Undel})
					} else {
						allComb = append(allComb, types.Combination{Cost: preComb.Cost, PartnerComb: preComb.PartnerComb + "_" + string(" "), Possible: true, Undel: preComb.Undel + 1})
					}
					continue
				}
				allComb = append(allComb, types.Combination{Cost: preComb.Cost + p2, PartnerComb: preComb.PartnerComb + "_" + string(partners[ind]), Possible: true, Undel: preComb.Undel})
			}
		}
	} else {

		pos := checkRow(row)
		for ind, p2 := range row {

			if p2 == -1 {
				if pos {
					allComb = append(allComb, types.Combination{Cost: 0, PartnerComb: string(""), Possible: true, Undel: 0})
				} else {
					allComb = append(allComb, types.Combination{Cost: 0, PartnerComb: string(""), Possible: true, Undel: 1})
				}
				continue
			}
			allComb = append(allComb, types.Combination{Cost: p2, PartnerComb: string(partners[ind]), Possible: true, Undel: 0})
		}
	}

	return allComb
}

func checkRow(row []int) bool {
	for _, val := range row {
		if val != -1 {
			return false
		}
	}
	return true
}
