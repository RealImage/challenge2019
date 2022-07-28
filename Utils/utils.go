package Utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"../Models"
)

var Maximum string
var Minimum string

//Contains Reading the partners.csv logic
func ReadPartnersDetails() (bool, []Models.DeliveryPartnerCostSlab) {
	var users []Models.DeliveryPartnerCostSlab
	rows, err := ReadData(PartnerDetails)
	if err != nil {
		log.Println("Cannot open CSV file:", err)
		return false, users
	}
	for _, row := range rows {
		user := Models.DeliveryPartnerCostSlab{TheatreName: row[0],
			SizeSlab:    row[1],
			MinimumCost: row[2],
			CostPerGb:   row[3],
			PartnerID:   row[4],
		}
		users = append(users, user)

	}
	return true, users
}

//Contains Delivery is possible checking Logic
func DeliveryIsPossibleCheck(PartnerDetails []Models.DeliveryPartnerCostSlab, InputDetails []Models.Input) bool {
	var output []Models.Output
	for _, name := range InputDetails {
		s1, s2, s3, s4 := DeliveryCheck(PartnerDetails, name.SizeSlab, name.TheatreName, name.ID)
		user1 := Models.Output{
			ID:                 s1,
			DeliveryIsPossible: s2,
			PartnerID:          s3,
			Price:              s4,
		}
		output = append(output, user1)
	}
	fmt.Println("output:", output)
	csvFile, err := os.Create(OutputDetails)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
		return false
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	for _, usance := range output {
		var row []string
		row = append(row, usance.ID)
		row = append(row, usance.DeliveryIsPossible)
		row = append(row, usance.Price)
		row = append(row, usance.PartnerID)
		writer.Write(row)
	}
	writer.Flush()
	return true
}

//Contains the PrizeCalculation and SizeSlabComparision is possible logic
func DeliveryCheck(PartnerDetails []Models.DeliveryPartnerCostSlab, SizeSlab string, TheatreName string, ID string) (string, string, string, string) {
	for _, name1 := range PartnerDetails {
		if strings.ToLower(strings.TrimSpace(name1.TheatreName)) == strings.ToLower(strings.TrimSpace(TheatreName)) {
			status := sizeSlabComparision(name1, SizeSlab)
			if status {
				CostPerGb := name1.CostPerGb
				cost := 0
				mincost := 0
				_, _ = fmt.Sscan(CostPerGb, &cost)
				_, _ = fmt.Sscan(name1.MinimumCost, &mincost)
				size, _ := strconv.Atoi(SizeSlab)
				Totalcost := cost * size
				if Totalcost < mincost {
					Totalcost = mincost
				}
				TotalcostinStr := strconv.Itoa(Totalcost)
				return ID, "true", name1.PartnerID, TotalcostinStr
			}
		}
	}
	return ID, "false", "", ""
}

func sizeSlabComparision(PartnerDetails Models.DeliveryPartnerCostSlab, SizeSlab string) bool {
	if Compare(PartnerDetails.SizeSlab, SizeSlab) {
		return true
	}
	return false
}

//Comparing Input size with Partners.csv sizeSlab
func Compare(SizeSlab string, InputSizeslab string) bool {
	s1 := strings.Split(SizeSlab, "-")
	for i, size := range s1 {
		if i == 0 {
			Minimum = size
		} else {
			Maximum = size
		}
	}
	min, max, input := IntConv(Minimum, Maximum, InputSizeslab)
	if input > min && input < max {
		return true
	}
	return false
}

//Converting String to int
func IntConv(MinValue string, MaxValue string, Inputvalue string) (int, int, int) {
	min, _ := strconv.Atoi(MinValue)
	max := 0
	_, _ = fmt.Sscan(MaxValue, &max)
	input, _ := strconv.Atoi(Inputvalue)
	return min, max, input
}

//Contains Reading the input.csv logic
func ReadInput() (bool, []Models.Input) {
	var ip []Models.Input
	input, err := ReadInputData(InputDetails)
	if err != nil {
		fmt.Println("Cannot open CSV file:", err)
		return false, ip
	}
	for _, row := range input {
		user1 := Models.Input{ID: row[0],
			SizeSlab:    row[1],
			TheatreName: row[2],
		}
		ip = append(ip, user1)
	}
	return true, ip
}

//To read the partners.csv file
func ReadData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Unable to open the file:", err)
		return [][]string{}, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	if _, err := r.Read(); err != nil {
		fmt.Println("Unable to skip the first line:", err)
		return [][]string{}, err
	}
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Unable to Read the file:", err)
		return [][]string{}, err
	}
	return records, nil
}

//To read the input.csv file
func ReadInputData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Unable to open the file:", err)
		return [][]string{}, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Unable to read the file:", err)
		return [][]string{}, err
	}
	return records, nil
}
