package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zeebo/errs"

	"challenge2019/delivery"
	"challenge2019/partner"
)

var (
	// Error is an error class that indicates csv parsing internal error.
	Error = errs.Class("csv internal error")
)

// ReadDeliveriesInput performs parsing input info from csv file to map.
func ReadDeliveriesInput(filePath string) (map[string]*delivery.Delivery, error) {
	csvLines, err := openAndRead(filePath)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	var deliveries = make(map[string]*delivery.Delivery)

	for id, line := range csvLines {
		if len(line) < 3 {
			return nil, Error.New("%s,%d line: input line misses values", filePath, id)
		} else if line[0] == "" || line[1] == "" || line[2] == "" {
			return nil, Error.New("%s,%d line: input line misses values", filePath, id)
		}

		deliveryID := strings.Replace(line[0], " ", "", -1)

		size, err := strconv.Atoi(strings.Replace(line[1], " ", "", -1))
		if err != nil {
			return nil, Error.Wrap(err)
		}

		theatreID := strings.Replace(line[2], " ", "", -1)

		deliveries[deliveryID] = &delivery.Delivery{
			OrderID:   id,
			Size:      size,
			TheatreID: theatreID,
		}
	}

	return deliveries, nil
}

// ReadPartners performs parsing input info from csv file to map.
func ReadPartners(filePath string) (map[string][]partner.Partner, error) {
	csvLines, err := openAndRead(filePath)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	var costLists = make(map[string][]partner.Partner)

	for id, line := range csvLines {
		if id == 0 {
			continue
		}

		if len(line) < 5 {
			return nil, Error.New("%s,%d line: input line misses values", filePath, id)
		} else if line[0] == "" || line[1] == "" || line[2] == "" || line[3] == "" || line[4] == "" {
			return nil, Error.New("%s,%d line: input line misses values", filePath, id)
		}

		theatreID := strings.Replace(line[0], " ", "", -1)

		amounts := strings.Split(strings.Replace(line[1], " ", "", -1), "-")
		minAmount, err := strconv.Atoi(amounts[0])
		if err != nil {
			return nil, Error.Wrap(err)
		}

		maxAmount, err := strconv.Atoi(amounts[1])
		if err != nil {
			return nil, Error.Wrap(err)
		}

		minCost, err := strconv.Atoi(strings.Replace(line[2], " ", "", -1))
		if err != nil {
			return nil, Error.Wrap(err)
		}

		costPerGB, err := strconv.Atoi(strings.Replace(line[3], " ", "", -1))
		if err != nil {
			return nil, Error.Wrap(err)
		}

		partnerID := strings.Replace(line[4], " ", "", -1)

		if _, exist := costLists[theatreID]; !exist {
			costLists[theatreID] = make([]partner.Partner, 0)
		}

		costLists[theatreID] = append(costLists[theatreID], partner.Partner{
			MinAmount: minAmount,
			MaxAmount: maxAmount,
			MinCost:   minCost,
			CostPerGB: costPerGB,
			ID:        partnerID,
		})
	}

	return costLists, nil
}

func WriteOutput(deliveries map[string]*delivery.Delivery) error {
	outputFile, err := os.Create("output.csv")
	if err != nil {
		return Error.Wrap(err)
	}
	defer outputFile.Close()

	csvWriter := csv.NewWriter(outputFile)
	defer csvWriter.Flush()

	lines := make([][]string, len(deliveries))

	for id, delivery := range deliveries {
		lines[delivery.OrderID] = []string{
			id,
			fmt.Sprint(delivery.Possible),
			delivery.PartnerID,
			fmt.Sprint(delivery.Cost),
		}
	}

	err = csvWriter.WriteAll(lines)

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
