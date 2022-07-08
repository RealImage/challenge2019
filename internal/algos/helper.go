package algos

import "github.com/purush7/challenge2019/v1/types"

func findMinCost(content int, slabSlice []types.Slab) (minCost int) {
	var found bool = false
	var cost int = 0
	var tmpCost = 0

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
			}
			found = true
		}
	}

	if !found {
		minCost = -1
	}
	return minCost
}
