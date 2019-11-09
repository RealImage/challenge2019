package Prob

import (
	"challenge2019/Prob/types"
	"sort"
)

func SortFeasibleArray(a [][]types.PartnerData) {

	for _, m := range a {
		for i, j := 0, len(m)-1; i < j; i, j = i+1, j-1 {
			m[i], m[j] = m[j], m[i]
		}
	}

	sort.Slice(a[:], func(i, j int) bool {
		for x := range a[i] {
			if a[i][x] == a[j][x] {
				continue
			}
			return a[i][x].DeliveryCost < a[j][x].DeliveryCost
		}
		return false
	})

}
