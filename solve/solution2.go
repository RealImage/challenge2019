package solve

import (
	"challenge2019/models"
	"sort"
)

func solution2(input []models.InputDetails, capacityMap models.CapacityMap, totalDataPerPartner map[string]models.TotalDataPerPartner, output1Map []models.OutputDetails, partners models.TheatreMap) []models.OutputDetails {
	for partner, dm := range totalDataPerPartner {
		if capacityMap[partner] < dm.Data {
			//geting first i.e smallest suitable element from dataunit map
			//sorting the keys
			keys := make([]int, 0, len(dm.DataUnitsMap))
			for k := range dm.DataUnitsMap {
				keys = append(keys, k)
			}
			sort.Ints(keys)
			for k, v := range dm.DataUnitsMap {
				newKey, ok := checkIfAddPartner(partners, input[v], output1Map[v], capacityMap, totalDataPerPartner)

				if ok {
					delete(dm.DataUnitsMap, k)
					dm.Data -= k
					totalDataPerPartner[partner] = dm
					update, updateok := totalDataPerPartner[newKey]
					if updateok {
						update.Data += k
						update.DataUnitsMap[k] = v
						totalDataPerPartner[newKey] = update
					} else {
						newupdate := models.TotalDataPerPartner{}
						newupdate.Data = k
						newDataUnitmap := make(map[int]int)
						newDataUnitmap[k] = v
						newupdate.DataUnitsMap = newDataUnitmap
						totalDataPerPartner[newKey] = newupdate
					}
					if capacityMap[partner] < dm.Data {
						continue
					} else {
						break
					}
				} else {
					output1Map[v].Feasibility = false
					output1Map[v].Cost = 0
					output1Map[v].PartnerID = ""
				}
			}

		}
	}
	//mutating output
	outputMap := updateTrueOutput(input, output1Map, totalDataPerPartner, partners)
	return outputMap
}

func checkIfAddPartner(partners models.TheatreMap, in models.InputDetails, output models.OutputDetails, capacityMap models.CapacityMap, totalDataPerPartner map[string]models.TotalDataPerPartner) (string, bool) {
	Tid := in.TheatreID
	Pid := output.PartnerID
	replacement := ""
	status := false
	cost := 0
	for i, patnerDetailsPerTheatre := range partners[Tid] {
		if i != Pid {
			for _, Tdata := range patnerDetailsPerTheatre {
				if Tdata.SizeSlab.Max >= in.Size && Tdata.SizeSlab.Min <= in.Size {

					if cost != 0 {

						Tempcost := Tdata.CostPerGB * in.Size
						if Tempcost < Tdata.MinimumCost {
							Tempcost = Tdata.MinimumCost
						}

						if Tempcost < cost {
							cost = Tdata.CostPerGB * in.Size
							if cost < Tdata.MinimumCost {
								cost = Tdata.MinimumCost
								if capacityMap[i] >= (totalDataPerPartner[i].Data + in.Size) {
									replacement = i
									status = true
								}

							}
						}

					} else {

						cost = Tdata.CostPerGB * in.Size
						if cost < Tdata.MinimumCost {
							cost = Tdata.MinimumCost
						}

						if capacityMap[i] >= (totalDataPerPartner[i].Data + in.Size) {
							replacement = i
							status = true
						}
					}

				}
			}
		}

	}
	return replacement, status
}

func updateTrueOutput(input []models.InputDetails,
	output1Map []models.OutputDetails,
	totalDataPerPartner map[string]models.TotalDataPerPartner,
	partners models.TheatreMap) []models.OutputDetails {
	for i, out := range output1Map {
		if out.Feasibility {
			for pid, v := range totalDataPerPartner {

				var cost int
				//finding the cost for updated value
				for _, value := range partners[input[i].TheatreID][pid] {
					for key, value2 := range v.DataUnitsMap {
						if value2 == i {
							if key >= value.SizeSlab.Min && key <= value.SizeSlab.Max {
								cost = key * value.CostPerGB
								if cost < value.MinimumCost {
									cost = value.MinimumCost
								}
								out.PartnerID = pid
								out.Cost = cost
								out.Feasibility = true
								output1Map[i] = out
								//can be replaced with multiple flag and break
								goto label
							}
						}

					}

				}

			}
		label:
		}
	}
	return output1Map
}
