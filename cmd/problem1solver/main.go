package main

import (
	"challange2019/pkg/models"
	"challange2019/pkg/services"
	"log"
)

func main() {
	dSvc := services.NewDeliverySvc(
		"assets/input.csv",
		"assets/partners.csv",
		"assets/capacities.csv",
	)

	outChan := make(chan *models.Output)
	errChan := make(chan error)
	go func() {
		go dSvc.DistributeDeliveriesAmongPartnersByMinCost(outChan, errChan)
		for err := range errChan {
			log.Println(err)
		}
	}()

	oSvc := services.NewOutputService("assets/output.csv")
	oSvc.WriteToCsv(outChan)

}
