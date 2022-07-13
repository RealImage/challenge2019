package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"qube_cinema_test/dtos"
	"strconv"
	"strings"
)

const (
	CsvInputPath      string = "./input.csv"
	CsvOutputPath     string = "./output.csv"
	CsvParterDataPath string = "./partners.csv"
)

func main() {
	listTrans, err := readInputData()
	if err != nil {
		panic(err)
	}
	pData, mapP, err := readPartnersData()
	if err != nil {
		panic(err)
	}

	err = writeOutputCsv(listTrans, pData, mapP)
	if err != nil {
		panic(err)
	}

	fmt.Println("Output csv finished")
}

func writeOutputCsv(listTrans []*dtos.In, pData []*dtos.Partner, mapP map[string][]string) error {
	f, err := os.Create(CsvOutputPath)
	if err != nil {
		return err
	}
	w := csv.NewWriter(f)
	for _, tran := range listTrans {
		rs, err := processTransaction(tran, pData, mapP)
		if err != nil {
			return err
		}
		row := []string{rs.DeliveryID, strconv.FormatBool(rs.IsDeliveryPossible), rs.PartnerID, fmt.Sprintf("%d", rs.Cost)}
		w.Write(row)
		if err := w.Error(); err != nil {
			return err
		}
	}
	w.Flush()
	f.Close()

	return nil
}

func processTransaction(trans *dtos.In, pData []*dtos.Partner, mapP map[string][]string) (*dtos.TransactionResult, error) {
	rs := &dtos.TransactionResult{
		DeliveryID: trans.DeliveryID,
	}
	possiblePartnerIDs := make([]string, 0)
	for k, v := range mapP {
		if isElementExist(v, trans.TheaterID) {
			possiblePartnerIDs = append(possiblePartnerIDs, k)
		}
	}
	for _, partner := range pData {
		if isElementExist(possiblePartnerIDs, partner.PartnerID) {
			rs.PartnerID = partner.PartnerID
			if trans.Size >= partner.MinSize && trans.Size <= partner.MaxSize {
				rs.Cost = trans.Size * partner.CostPerGB
				if rs.Cost < partner.MinimumCost {
					rs.Cost = partner.MinimumCost
				}
				rs.IsDeliveryPossible = true
			}
		} else {
			rs.Cost = 0
		}
	}
	return rs, nil
}

func readInputData() ([]*dtos.In, error) {
	input, err := os.Open(CsvInputPath)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(input)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	arrIn := make([]*dtos.In, 0)
	for _, record := range records {
		entry := &dtos.In{
			DeliveryID: record[0],
			TheaterID:  record[2],
		}
		size, err := strconv.ParseInt(record[1], 10, 64)
		if err != nil {
			return nil, err
		}
		entry.Size = size
		arrIn = append(arrIn, entry)
	}
	return arrIn, nil
}

func readPartnersData() ([]*dtos.Partner, map[string][]string, error) {
	input, err := os.Open(CsvParterDataPath)
	if err != nil {
		return nil, nil, err
	}
	reader := csv.NewReader(input)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	arrIn := make([]*dtos.Partner, 0)
	mapPossibleDelivery := make(map[string][]string, 0)
	for i, record := range records {
		if i == 0 {
			continue
		}
		entry := &dtos.Partner{
			TheatreID: record[0],
			Size:      record[1],
			PartnerID: record[4],
		}
		minCost, err := strconv.ParseInt(strings.TrimSpace(record[2]), 10, 64)
		if err != nil {
			return nil, nil, err
		}
		entry.MinimumCost = minCost
		costPerGB, err := strconv.ParseInt(strings.TrimSpace(record[3]), 10, 64)
		if err != nil {
			return nil, nil, err
		}
		entry.CostPerGB = costPerGB

		sRange := strings.Split(strings.TrimSpace(entry.Size), "-")
		fSize, err := strconv.ParseInt(sRange[0], 10, 64)
		if err != nil {
			return nil, nil, err
		}
		entry.MinSize = fSize
		tSize, err := strconv.ParseInt(sRange[1], 10, 64)
		if err != nil {
			return nil, nil, err
		}
		entry.MaxSize = tSize

		if val, ok := mapPossibleDelivery[entry.PartnerID]; ok {
			theaterID := strings.TrimSpace(entry.TheatreID)
			if !isElementExist(val, theaterID) {
				mapPossibleDelivery[entry.PartnerID] = append(mapPossibleDelivery[entry.PartnerID], theaterID)
			}
		} else {
			arrTheaters := make([]string, 0)
			arrTheaters = append(arrTheaters, strings.TrimSpace(entry.TheatreID))
			mapPossibleDelivery[entry.PartnerID] = arrTheaters
		}

		arrIn = append(arrIn, entry)
	}
	return arrIn, mapPossibleDelivery, nil
}

func isElementExist(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
