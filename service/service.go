package service

import (
	"qube/model"
	"sort"
)

func ProblemStatement1(partners map[string][]model.Partner, deliveries []model.Delivery) ([]model.Result, error) {
	var result []model.Result = make([]model.Result, 0, len(deliveries))

	for _, delivery := range deliveries {
		partnersForTheatre, ok := partners[delivery.Theatre]

		if !ok {
			result = append(result, model.Result{
				Name: delivery.Name,
				Cost: -1,
			})
			continue
		}

		var minCost int = -1
		var tempPartner string = ""

		for _, partner := range partnersForTheatre {
			if partner.MinGB > delivery.Amount || delivery.Amount > partner.MaxGB {
				continue
			}

			cost := delivery.Amount * partner.PerGB

			if cost < partner.MinCost {
				cost = partner.MinCost
			}

			if cost < minCost || minCost == -1 {
				minCost = cost
				tempPartner = partner.Partner
			}

		}

		result = append(result, model.Result{
			Name:    delivery.Name,
			Partner: tempPartner,
			Cost:    minCost,
		})

	}

	return result, nil
}

func ProblemStatement2(partners map[string][]model.Partner, deliveries []model.Delivery, capacities map[string]int) ([]model.Result2, error) {
	var result []model.Result2 = make([]model.Result2, 0, len(deliveries))
	var allCombinations [][]model.Result2 = make([][]model.Result2, 0)

	for _, delivery := range deliveries {
		partnersForTheatre, ok := partners[delivery.Theatre]

		if !ok {
			result = append(result, model.Result2{
				Name: delivery.Name,
				Cost: -1,
			})
			continue
		}

		var partnerCombinations []model.Result2 = make([]model.Result2, 0, len(deliveries))

		for _, partner := range partnersForTheatre {
			if partner.MinGB > delivery.Amount || delivery.Amount > partner.MaxGB {
				continue
			}

			cost := delivery.Amount * partner.PerGB

			if cost < partner.MinCost {
				cost = partner.MinCost
			}

			partnerCombinations = append(partnerCombinations, model.Result2{
				Name:    delivery.Name,
				Amount:  delivery.Amount,
				Partner: partner.Partner,
				Cost:    cost,
			})

		}

		if len(partnerCombinations) == 0 {
			result = append(result, model.Result2{
				Name: delivery.Name,
				Cost: -1,
			})
		} else {
			allCombinations = addToEach(allCombinations, partnerCombinations)
		}

	}

	result = append(checkEach(allCombinations, capacities), result...)

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result, nil
}

func checkEach(allCombinations [][]model.Result2, capacities map[string]int) []model.Result2 {
	var cost int = 0
	var amount = 0
	var indexRes int = 0

	for index, slice := range allCombinations {
		newCost, newAmount := check(slice, capacities)

		if newAmount < amount {
			continue
		}

		if newAmount > amount {
			amount = newAmount
			cost = newCost
			indexRes = index
		} else if newCost < cost {
			cost = newCost
			indexRes = index
		}
	}

	return allCombinations[indexRes]
}

func check(check []model.Result2, capacities map[string]int) (int, int) {
	var amount int = 0
	var cost int = 0
	var takenCapacities map[string]int = make(map[string]int)

	for _, elem := range check {
		if takenCapacities[elem.Partner]+elem.Amount <= capacities[elem.Partner] {
			amount++
			cost += elem.Cost
			takenCapacities[elem.Partner] += elem.Amount
		}
	}

	return cost, amount
}

func addToEach(to [][]model.Result2, from []model.Result2) [][]model.Result2 {
	if len(from) == 0 {
		return to
	}

	var result [][]model.Result2 = make([][]model.Result2, 0)

	if len(to) == 0 {
		for _, elem := range from {
			result = append(result, []model.Result2{elem})
		}
		return result
	}

	for _, array := range to {
		for _, elem := range from {
			result = append(result, append(array, elem))
		}
	}

	return result
}
