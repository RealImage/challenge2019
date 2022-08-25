package main

import (
	"challange2019/pkg/models"
	"challange2019/pkg/services"
	"flag"
	"fmt"
	"log"
)

func main() {
	var inputPath, partnersPath, capacity, outputPath string
	setArgs(&inputPath, &partnersPath, &capacity, &outputPath)

	dSvc := services.NewDeliverySvc(inputPath, partnersPath, capacity)

	outChan := make(chan *models.Output)
	errChan := make(chan error)
	go func() {
		go dSvc.DistributeDeliveriesAmongPartnersByMinCostAndCapacity(outChan, errChan)
		for err := range errChan {
			log.Println(err)
		}
	}()

	oSvc := services.NewOutputService(outputPath)
	oSvc.WriteToCsv(outChan)
}

func setArgs(inputPath, partnersPath, capacity, outputPath *string) {
	usageMsg := `Given a list of content size and Theatre ID, 
Assign deliveries to partners in such a way that all deliveries are possible (Higher Priority)
and overall cost of delivery is minimum (i.e. First make sure no delivery is impossible and then minimise the sum of cost of all the delivery).
If delivery is not possible to a theatre, mark that delivery impossible. Take partner capacity into consideration as well.

The program expected 4 arguments:
1 - input filepath
2 - partners filepath
3 - partners capacity filepath
4 - destination filepath
or no arguments.

If no args, then default values:
input filepath - assets/input.csv
partners filepath - assets/partners.csv
partners capacity filepath - assets/capacity.csv
destination filepath - assets/output.csv
`
	flag.Parse()
	args := flag.Args()
	flag.Usage = func() {
		fmt.Println(usageMsg)
	}

	*inputPath = "assets/input.csv"
	*partnersPath = "assets/partners.csv"
	*capacity = "assets/capacity.csv"
	*outputPath = "assets/output.csv"

	if len(args) > 0 && len(args) != 3 {
		flag.Usage()
		return
	}

	if len(args) > 0 {
		*inputPath = args[0]
		*partnersPath = args[1]
		*capacity = args[2]
		*outputPath = args[3]
	}

}
