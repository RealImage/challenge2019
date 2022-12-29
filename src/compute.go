package src

import (
	"math"
	"strings"

	"github.com/shreeyashnaik/challenge2019/common/utils"
)

func ComputeCost(sizeSlab string, size, costPerGB, minimumCost int) int {
	// Parse Upper & Lower bound from size slab
	bound := strings.Split(sizeSlab, "-")

	// If size is not between Size Slab
	if size < utils.ToInt(bound[0]) || size > utils.ToInt(bound[1]) {
		return math.MaxInt
	}

	// Compute cost
	cost := size * costPerGB
	if cost <= minimumCost {
		return int(minimumCost)
	}

	return int(size * costPerGB)
}
