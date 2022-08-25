package main

import (
	"challange2019/pkg/models"
	"challange2019/pkg/services"
	"flag"
	"fmt"
	"log"
)

func main() {
	var inputPath, partnersPath, outputPath string

	setArgs(&inputPath, &partnersPath, &outputPath)

	dSvc := services.NewDeliverySvc(inputPath, partnersPath, "")

	outChan := make(chan *models.Output)
	errChan := make(chan error)
	go func() {
		go dSvc.DistributeDeliveriesAmongPartnersByMinCost(outChan, errChan)
		for err := range errChan {
			log.Println(err)
		}
	}()

	oSvc := services.NewOutputService(outputPath)
	oSvc.WriteToCsv(outChan)

}

func setArgs(inputPath, partnersPath, outputPath *string) {
	usageMsg := `Given a list of content size and Theatre ID, Find the partner for each delivery where cost of delivery is minimum. 
If delivery is not possible, mark that delivery impossible.

The program expected 3 arguments:
1 - input filepath
2 - partners filepath
3 - destination filepath
or no arguments.

If no args, then default values:
input filepath - assets/input.csv
partners filepath - assets/partners.csv
destination filepath - assets/output.csv
`
	flag.Parse()
	args := flag.Args()
	flag.Usage = func() {
		fmt.Println(usageMsg)
	}

	*inputPath = "assets/input.csv"
	*partnersPath = "assets/partners.csv"
	*outputPath = "assets/output.csv"

	if len(args) > 0 && len(args) != 3 {
		flag.Usage()
		return
	}

	if len(args) > 0 {
		*inputPath = args[0]
		*partnersPath = args[1]
		*outputPath = args[2]
	}

}
