package csv

import (
	"challenge2019/partner"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zeebo/errs"

	"challenge2019/delivery"
)

var (
	// Error is an error class that indicates csv parsing internal error.
	Error = errs.Class("csv internal error")
)

func ReadDeliveries(filePath string) (*[]delivery.Delivery, error) {
	csvLines, err := openAndRead(filePath)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	var deliveries = make([]delivery.Delivery, len(csvLines))

	for id, line := range csvLines {
		size, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, Error.Wrap(err)
		}

		deliveries[id] = delivery.Delivery{
			ID:        line[0],
			Size:      size,
			TheatreID: line[2],
		}
	}

	return &deliveries, nil
}

func ReadPartners(filePath string) (*[]delivery.Delivery, error) {
	csvLines, err := openAndRead(filePath)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	var deliveries = make([]partner.Partner, len(csvLines)-1)

	for id, line := range csvLines {
		minCost, err := strconv.Atoi(line[2])
		if err != nil {
			return nil, Error.Wrap(err)
		}

		costPerGB, err := strconv.Atoi(line[3])
		if err != nil {
			return nil, Error.Wrap(err)
		}

		amounts := strings.Split(line[1], "-")
		minAmount, err := strconv.Atoi(amounts[0])
		if err != nil {
			return nil, Error.Wrap(err)
		}

		maxAmount, err := strconv.Atoi(amounts[1])
		if err != nil {
			return nil, Error.Wrap(err)
		}

		deliveries[id] = partner.Partner{
			TheatreID: line[0],
			MinAmount: minAmount,
			MaxAmount: maxAmount,
			MinCost:   minCost,
			CostPerGB: costPerGB,
			ID:        line[4],
		}
	}

	return &deliveries, nil
}

func WriteOutput(deliveries []delivery.Delivery) error {
	outputFile, err := os.Create("output.csv")
	if err != nil {
		return Error.Wrap(err)
	}
	defer outputFile.Close()

	csvWriter := csv.NewWriter(outputFile)
	defer csvWriter.Flush()

	for _, delivery := range deliveries {
		err = csvWriter.Write([]string{
			delivery.ID,
			fmt.Sprint(delivery.Possible),
			delivery.PartnerID,
			fmt.Sprint(delivery.Cost),
		})
	}

	return Error.Wrap(err)
}

func openAndRead(filePath string) ([][]string, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return [][]string{}, err
	}

	defer csvFile.Close()
	return csv.NewReader(csvFile).ReadAll()
}
