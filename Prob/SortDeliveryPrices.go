package Prob

import "sort"

func SortDeliveryPrices(fmap map[int]int) int {
	sMap := make(map[int]int)
	keys := make([]int, 0, len(fmap))
	for k := range fmap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		sMap[k] = fmap[k]
	}
	return sMap[keys[0]]
}
