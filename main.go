package main

import (
	"container/heap"
	"fmt"
	"github.com/funcoding/challenge2019/helpers"
	"io"
	"log"
	"math"
	"github.com/funcoding/challenge2019/preprocess"
	"github.com/funcoding/challenge2019/structs"
	"sort"
	"strconv"
	"strings"
	"go/build"
)


func binarySearch(item int, array [] structs.PartnerStruct, left int, right int) int {
	if left > right{
		return -1
	}
	middle := (left + right) / 2
	if item >= array[middle].MinSlab && item <= array[middle].MaxSlab {
		return middle
	}
	if item <= array[middle].MinSlab {
		right = middle
	}else if item >= array[middle].MaxSlab {
		left = middle+1
	}
	return binarySearch(item, array, left, right)
}

/**
Logic:
1. Csv input file is read line by line.
2. After reading each input line
	a) corresponding partners are fetched for a theatre from Partner hash map.
	b) For every partner using modified binary search the corresponding struct holding slab, minimum cost, costPerGB
		details is obtained and minimum cost is calculated for the corresponding theatreId.
 */
func problemOne() {
	outputFileName := "output1.csv"
	r := helpers.ParseContentsFromCsv(build.Default.GOPATH+"/src/github.com/funcoding/challenge2019/static/input.csv")
	var result [] structs.ResultFormatStruct
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		deliverySize, _ := strconv.Atoi(strings.TrimSpace(record[1]))
	
		tempResult := make(map [string] string)
		minCost := math.Inf(1)
		for key, item := range preprocess.Partners[strings.TrimSpace(record[2])]{
			index := binarySearch(deliverySize, item, 0, len(item) - 1)
			if index == -1 {
				continue
			} else {
				calculatedCost := float64(deliverySize * item[index].CostPerGB)

				if calculatedCost < float64(item[index].MinCost) {
					calculatedCost = float64(item[index].MinCost)
				}
	
				if float64(calculatedCost) < minCost {
					tempResult["partner"] = key
					tempResult["cost"] = strconv.FormatFloat(calculatedCost, 'f', 0, 64)
					minCost = calculatedCost
				}
			}
		}
		
		_, deliverable := tempResult["partner"]
		result = append(
			result, 
			structs.ResultFormatStruct{
				DeliveryId: strings.TrimSpace(record[0]), 
				DeliveryPossible:deliverable, 
				PartnerId: tempResult["partner"], 
				Cost: tempResult["cost"]})
	}


	response := helpers.WriteDataToCsv(outputFileName, result)
	fmt.Println(fmt.Sprintf("Output file %s created for problem 1.", response))
}


type outputStruct struct {
	deliveryId string
	theatreId string
	deliverySize int
	isDeliverable bool
	partnerId string
	cost string
}


/**
For this problem statement I would suggest min-cost max-flow algorithm. Since my exposure to the algorithm is less, I have
come up with another logic.
@Logic:
1. The csv file is read line by line
	a) Required input data is stored in slice of outPutStruct and sorted in descending order according to deliverySize.
		By sorting in descending order we will be able to ensure that
	b) A hash map of computedPartnerCost is used which holds slice containing the minimum cost by the partner for the delivery quantity.
		Min heap is made use because it holds the minimum computed costs in ascending order. Another use of keeping the costs sorted
		in ascending order is to ensure that when the partner capacity is reached, the next minimum cost could be chosen from the slice ie the next index.
		This logic holds good for most use cases only if the traversing of the delivery orders are in descending.
2. After storing the required data in computedPartnerCost and deliveryPriority; required output data are calculated by iterating deliveryPriority
	and applying calculations on corresponding computedPartnerCost.

 */
func problemTwo(){
	outputFileName := "output2.csv"
	capacity := preprocess.Capacities
	computedPartnerCost := make(map[int] []helpers.Item)
	var deliveryPriority [] outputStruct
	r := helpers.ParseContentsFromCsv(build.Default.GOPATH+"/src/github.com/funcoding/challenge2019/static/input.csv")
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		deliverySize, _ := strconv.Atoi(record[1])
		deliveryPriority = append(deliveryPriority, outputStruct{
			deliveryId: record[0],
			theatreId: record[2],
			deliverySize: deliverySize,
		})

		// Compute cost for all partners
		pq := helpers.PriorityQueue{}
		for partnerId, slabs := range preprocess.Partners[record[2]]{
			idx := binarySearch(deliverySize, slabs, 0, len(slabs) - 1)
			if idx != -1{
				calculatedCost := float64(deliverySize * slabs[idx].CostPerGB)

				if calculatedCost < float64(slabs[idx].MinCost) {
					calculatedCost = float64(slabs[idx].MinCost)
				}
				heap.Push(&pq, &helpers.Item{Value: partnerId, Priority: int(calculatedCost)})
			}
		}
		computedPartnerCost[deliverySize] = pq.ToArray()
	}

	// Sort
	sort.Slice(deliveryPriority, func(i, j int) bool {
		return deliveryPriority[i].deliverySize > deliveryPriority[j].deliverySize
	})

	// Confusing part. Refer @Logic 1) b)
	for  i := 0; i < len(deliveryPriority); i++{
		data := &deliveryPriority[i]
		for _, computedHeap := range computedPartnerCost[data.deliverySize] {
			if capacity[computedHeap.Value] >= data.deliverySize {
				capacity[computedHeap.Value] = capacity[computedHeap.Value] - data.deliverySize
				data.isDeliverable = true
				data.partnerId = computedHeap.Value
				data.cost = strconv.Itoa(computedHeap.Priority)
				break
			}
		}
	}

	//Optional
	sort.Slice(deliveryPriority, func(i, j int) bool {
		return deliveryPriority[i].deliverySize < deliveryPriority[j].deliverySize
	})

	var result [] structs.ResultFormatStruct
	for _, data := range deliveryPriority{
		result = append(result, structs.ResultFormatStruct{
			DeliveryId: data.deliveryId,
			DeliveryPossible: data.isDeliverable,
			PartnerId: data.partnerId,
			Cost: data.cost,
		})
	}
	response := helpers.WriteDataToCsv(outputFileName, result)
	fmt.Println(fmt.Sprintf("Output file %s created for problem 2.", response))
}

func main() {
	problemOne()
	problemTwo()
}