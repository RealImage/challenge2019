package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Partners struct {
	TheatreId   string
	Min         int
	Max         int
	MinimumCost int
	CostPerGB   int
	PartnerId   string
}

type Input struct {
	DeliveryId   string
	DeliverySize int
	TheatreId    string
}

type Output struct {
	DeliveryId    string
	Possible      bool
	TheatreId     string
	MinimumAmount int
}

func main() {
	var s string
	var minCost []int
	var partnerId, deliveryId []string
	var possible []bool
	var poss bool
	var miniCost, maxDeliverySize int
	var Allinput, sortedInput []Input
	var Alldata []Partners

	//getting data from partner.csv and saving in local structure i.e. []Partners

	lines, err := AccessCsv("partners.csv")
	if err != nil {
		fmt.Println(err)
	}
	for i, line := range lines {
		if i == 0 {
			continue
		}
		minCostSpace := strings.Replace(line[2], " ", "", -1)
		minCost, _ := strconv.Atoi(minCostSpace)
		perCost := strings.Replace(line[3], " ", "", -1)
		CostPergb, _ := strconv.Atoi(perCost)
		ranges := strings.Split(line[1], "-")
		upp := strings.Replace(ranges[1], " ", "", -1)
		var lower, upper int
		lower, _ = strconv.Atoi(ranges[0])
		upper, _ = strconv.Atoi(upp)

		data := Partners{
			TheatreId:   line[0],
			Min:         lower,
			Max:         upper,
			MinimumCost: minCost,
			CostPerGB:   CostPergb,
			PartnerId:   line[4],
		}
		Alldata = append(Alldata, data)
	}

	//getting data from capacitites.csv and saving in capacityMap

	capacityMap := make(map[string]int)

	lines, err = AccessCsv("capacities.csv")
	if err != nil {
		fmt.Println(err)
	}
	for i, line := range lines {
		if i == 0 {
			continue
		}
		partnerid := strings.Replace(line[0], " ", "", -1)
		caps := strings.Replace(line[1], " ", "", -1)
		Capacity, _ := strconv.Atoi(caps)
		capacityMap[partnerid] = Capacity
	}

	//getting data from input.csv and saving in local structure i.e. []Input

	lines, err = AccessCsv("input.csv")
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range lines {
		Size, _ := strconv.Atoi(line[1])
		data := Input{
			DeliveryId:   line[0],
			DeliverySize: Size,
			TheatreId:    line[2],
		}
		Allinput = append(Allinput, data)
	}

	//sorting the input data based on size of delivery

	index := -1
	length := len(Allinput)

	for i := 0; i < length; i++ {
		maxDeliverySize = 0
		for j, v := range Allinput {
			if maxDeliverySize < v.DeliverySize {
				maxDeliverySize = v.DeliverySize
				index = j
			}
		}
		sortedInput = append(sortedInput, Allinput[index])
		Allinput = RemoveIndex(Allinput, index)
	}

	//checking for the conditions

	for _, vinput := range sortedInput {
		miniCost = 9223372036854775807
		poss = false
		s = " "
		for j, vdata := range Alldata {
			if j == 0 {
				continue
			}
			replace := strings.Replace(vdata.TheatreId, " ", "", -1)
			if vinput.TheatreId == replace {
				if vinput.DeliverySize >= vdata.Min && vinput.DeliverySize <= vdata.Max && capacityMap[vdata.PartnerId] > vinput.DeliverySize {
					poss = true
					cost := vinput.DeliverySize * vdata.CostPerGB
					if cost < vdata.MinimumCost {
						cost = vdata.MinimumCost
					}
					if cost < miniCost {
						miniCost = cost
						s = vdata.PartnerId
					}
				}
			}
		}
		if poss == false {
			miniCost = 0
		}
		if v, ok := capacityMap[s]; ok {
			capacityMap[s] = v - vinput.DeliverySize
		}

		//Appending the answer for the specific input into the arrays
		deliveryId = append(deliveryId, vinput.DeliveryId)
		minCost = append(minCost, miniCost)
		partnerId = append(partnerId, s)
		possible = append(possible, poss)
	}

	//Combining all the arrays in 2D array of string to convert into CSV
	outputData := make([][]string, len(deliveryId))
	for i := 0; i < len(deliveryId); i++ {
		var temp []string
		Cost := strconv.Itoa(minCost[i])
		Poss := strconv.FormatBool(possible[i])
		temp = append(temp, deliveryId[i], Poss, partnerId[i], Cost)
		outputData[i] = temp
	}

	//Writing the ouput data to result.csv and console

	fmt.Println(outputData)
	WriteToCsv(outputData)
}

func AccessCsv(name string) ([][]string, error) {

	f, err := os.Open(name)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func WriteToCsv(outputdata [][]string) {
	outputfile, err := os.Create("result.csv")
	if err != nil {
		fmt.Println("cannot create file", err)
	}
	defer outputfile.Close()

	writer := csv.NewWriter(outputfile)
	defer writer.Flush()

	for _, value := range outputdata {
		err := writer.Write(value)
		if err != nil {
			fmt.Println("cannot write to file", err)
		}
	}
}

func RemoveIndex(s []Input, index int) []Input {
	return append(s[:index], s[index+1:]...)
}
