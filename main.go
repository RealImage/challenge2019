package main

import (
	"fmt"
	"qube/service"
)

func main() {
	partners := service.ReadPartners()

	delivery := service.ReadDelivery()

	result, err := service.ProblemStatement1(partners, delivery)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	service.PrintProblemStatement1(result)

	capacities := service.ReadCapacity()

	result2, err := service.ProblemStatement2(partners, delivery, capacities)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	service.PrintProblemStatement2(result2)
}
