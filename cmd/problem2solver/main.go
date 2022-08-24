package main

import (
	"challange2019/pkg/services"
	"fmt"
)

func main() {
	svc := services.NewDeliverySvc(
		"assets/input.csv",
		"assets/partners.csv",
		"assets/capacities.csv",
	)
	output, errors := svc.DistributeDeliveriesAmongPartnersByMinCostAndCapacity()

	for _, v := range output.Container {
		fmt.Println(v)
	}

	for err := range errors {
		fmt.Println(err)
	}
}
