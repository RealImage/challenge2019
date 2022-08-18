package service

import (
	"challenge2019/db/entities"
	"challenge2019/helper"
	"encoding/csv"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

type ServiceInterface interface {
	//  Problem 1
	PartnerDetails() ([]entities.Partner, bool)
	InputDetails() ([]entities.Input, bool)
	Deliverable(Partner []entities.Partner, Input []entities.Input, PartnerCapacity []entities.PartnerCapacity) (os.File, []entities.Output, bool)
	DeliveryCheck(Partner []entities.Partner, ID string, TheatreID string, SizeOfDelivery string) (string, string, string, string)
	SizeSlabComparison(Minimum, Maximum int, InputSizeSlab string) bool

	// Problem 2
	PartnerCapacityDetails() ([]entities.PartnerCapacity, bool)
	PartnerCapacity(PartnerCapacity []entities.PartnerCapacity)
	checkIfPartnerAllowed(PartnerCapacity []entities.PartnerCapacity, Input []entities.Input)
}
type ServiceStruct struct {
	Helper helper.HelperInterface
}

func NewService(Helper helper.HelperInterface) ServiceInterface {
	return &ServiceStruct{Helper: Helper}
}
func (s *ServiceStruct) PartnerDetails() ([]entities.Partner, bool) {
	records, err := s.Helper.ReadCsvFile("partners.csv", true)
	if err != nil {
		log.Error("Failed to read PartnerDetails")
		return nil, false
	}
	var data []entities.Partner
	for i, _ := range records {
		tempRec := entities.Partner{
			TheatreID:   strings.TrimSpace(records[i][0]),
			SizeSlab:    strings.TrimSpace(records[i][1]),
			MinimumCost: strings.TrimSpace(records[i][2]),
			CostPerGB:   strings.TrimSpace(records[i][3]),
			PartnerID:   strings.TrimSpace(records[i][4]),
		}
		data = append(data, tempRec)
	}
	log.Info("Read PartnerDetails Successfully")
	return data, true
}

func (s *ServiceStruct) InputDetails() ([]entities.Input, bool) {
	records, err := s.Helper.ReadCsvFile("input.csv", false)
	if err != nil {
		log.Error("Failed to read Input")
		return nil, false
	}
	var data []entities.Input
	for i, _ := range records {
		tempRec := entities.Input{
			ID:             strings.TrimSpace(records[i][0]),
			SizeOfDelivery: strings.TrimSpace(records[i][1]),
			TheatreID:      strings.TrimSpace(records[i][2]),
		}
		data = append(data, tempRec)
	}
	return data, true
}

func (s *ServiceStruct) PartnerCapacityDetails() ([]entities.PartnerCapacity, bool) {
	records, err := s.Helper.ReadCsvFile("capacities.csv", true)
	if err != nil {
		log.Error("Failed to read Input")
		return nil, false
	}
	var data []entities.PartnerCapacity
	for i, _ := range records {
		tempRec := entities.PartnerCapacity{
			PartnerID: strings.TrimSpace(records[i][0]),
			Capacity:  strings.TrimSpace(records[i][1]),
		}
		data = append(data, tempRec)
	}
	return data, true
}

//  the inside deliverycheck function returns the output for each delivery id
// this function takes the output from delivery check and convert that output into csv file

func (s *ServiceStruct) Deliverable(Partner []entities.Partner, Input []entities.Input, PartnerCapacity []entities.PartnerCapacity) (os.File, []entities.Output, bool) {
	var OP []entities.Output
	for _, rec := range Input {
		id, iDP, pID, prc := s.DeliveryCheck(Partner, rec.ID, rec.TheatreID, rec.SizeOfDelivery)
		tempOp := entities.Output{
			ID:             id,
			IsDeliverable:  iDP,
			Partner:        pID,
			CostOfDelivery: prc,
		}
		OP = append(OP, tempOp)
	}
	csvOutput, err := os.Create("output.csv")
	if err != nil {
		log.Error("Not able to create output csv file", err.Error())
		return *csvOutput, OP, false
	}
	defer csvOutput.Close()

	writer := csv.NewWriter(csvOutput)
	for _, rec := range OP {
		var row []string
		row = append(row, rec.ID)
		row = append(row, rec.IsDeliverable)
		row = append(row, rec.CostOfDelivery)
		row = append(row, rec.Partner)
		writer.Write(row)
	}
	writer.Flush()
	return *csvOutput, OP, true
}

// check if delivery is possible
// if possible calculates the cost
// if cost is less than minimum cost it returns the minimum cost

func (s *ServiceStruct) DeliveryCheck(Partner []entities.Partner, ID string, TheatreID string, SizeOfDelivery string) (string, string, string, string) {
	for _, name := range Partner {
		if strings.ToLower(strings.TrimSpace(name.TheatreID)) == strings.ToLower(strings.TrimSpace(TheatreID)) {

			sizeSlab := strings.Split(name.SizeSlab, "-")
			Minimum := s.Helper.Sscan(sizeSlab[0])
			Maximum := s.Helper.Sscan(sizeSlab[1])

			if s.SizeSlabComparison(Minimum, Maximum, SizeOfDelivery) {

				cost := s.Helper.Sscan(name.CostPerGB)
				mincost := s.Helper.Sscan(name.MinimumCost)
				size := s.Helper.Sscan(SizeOfDelivery)

				TotalCost := cost * size

				if TotalCost < mincost {
					TotalCost = mincost
				}

				TotalCostInStr := strconv.Itoa(TotalCost)

				return ID, "true", name.PartnerID, TotalCostInStr
			}
		}
	}
	return ID, "false", "", ""
}

// compares the given Input Size with minimum and maximum range for the theatre if it's lying in the range
// that partner is able to deliver that content to theatre

func (s *ServiceStruct) SizeSlabComparison(Minimum, Maximum int, InputSizeSlab string) bool {
	input := s.Helper.Sscan(InputSizeSlab)
	if Minimum <= input && input <= Maximum {
		return true
	}
	return false
}

func (s *ServiceStruct) PartnerCapacity(PartnerCapacity []entities.PartnerCapacity) {
	//s.checkIfPartnerAllowed(PartnerCapacity)
}

func (s *ServiceStruct) checkIfPartnerAllowed(PartnerCapacity []entities.PartnerCapacity, Input []entities.Input) {
	// need input

	// use set here or map , cause here it can cause double entry too

	// three conditions for granting thing to partner
	// low charge
	// available
	// allowed to deliver to that theatre

	// one more thing to consider total cost should be minimum
	var TII []entities.TheatresInInput
	for _, rec := range Input {
		temp := entities.TheatresInInput{TheatreID: rec.TheatreID}
		TII = append(TII, temp)
	}

}
