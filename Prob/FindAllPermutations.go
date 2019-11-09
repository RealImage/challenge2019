package Prob

import (
	"challenge2019/Prob/types"
)

func FindAllPermutations(allApplicablePartners []types.DeliveryAndPartners) [][]types.PartnerData {
	partners := [][]types.PartnerData{}
	for _, r := range allApplicablePartners {
		partners = append(partners, r.Partners)
	}

	res := findAllPermutations(partners...)
	return res
}

func findAllPermutations(parts ...[]types.PartnerData) (ret [][]types.PartnerData) {
	{
		var n = 1
		for _, ar := range parts {
			n *= len(ar)
		}
		ret = [][]types.PartnerData{}
	}
	var at = make([]int, len(parts))
loop:
	for {
		for i := len(parts) - 1; i >= 0; i-- {
			if at[i] > 0 && at[i] >= len(parts[i]) {
				if i == 0 || (i == 1 && at[i-1] == len(parts[0])-1) {
					break loop
				}
				at[i] = 0
				at[i-1]++
			}
		}
		tmp := []types.PartnerData{}
		for i, ar := range parts {
			var p = at[i]
			if p >= 0 && p < len(ar) {
				tmp = append(tmp, ar[p])
			}
		}
		ret = append(ret, tmp)
		at[len(parts)-1]++
	}
	return ret
}
