package main

import(
	"container/heap"
	"fmt"
	"./src"
	"./src/helpers"
	. "./src/structs"
	"log"
	"io"
	"strconv"
	"strings"
	"math"
	"sort"
)


func binarySearch(item int, array [] qube.PartnerStruct, left int, right int) int {
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

func problem1() {
	r := helpers.ParseContentsFromCsv("./static/input.csv")
	var result [] ResultFormatStruct
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		deliverySize, _ := strconv.Atoi(strings.TrimSpace(record[1]))
		if qube.Partners[strings.TrimSpace(record[2])] == nil{
			
		}
	
		tempResult := make(map [string] string)
		minCost := math.Inf(1)
		for key, item := range qube.Partners[strings.TrimSpace(record[2])]{
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
			ResultFormatStruct{
				DeliveryId: strings.TrimSpace(record[0]), 
				DeliveryPossible:deliverable, 
				PartnerId: tempResult["partner"], 
				Cost: tempResult["cost"]})
	}

	fmt.Println(result)
	helpers.WriteDataToCsv("output1.csv", result)
}


type outputStruct struct {
	deliveryId string
	theatreId string
	deliverySize int
	isDeliverable bool
	partnerId string
	cost string
}

type deliveryOrderStruct struct {
	deliverySize int
	computedCost helpers.PriorityQueue
}

func problem2(){
	capacity := qube.Capacities
	theatreMap := make(map[int] helpers.PriorityQueue)
	var deliveryPriority [] outputStruct
	r := helpers.ParseContentsFromCsv("./static/input.csv")
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
		for partnerId, slabs := range qube.Partners[record[2]]{
			idx := binarySearch(deliverySize, slabs, 0, len(slabs) - 1)
			if idx != -1{
				calculatedCost := float64(deliverySize * slabs[idx].CostPerGB)

				if calculatedCost < float64(slabs[idx].MinCost) {
					calculatedCost = float64(slabs[idx].MinCost)
				}
				heap.Push(&pq, &helpers.Item{Value: partnerId, Priority: int(calculatedCost)})
			}
		}
		theatreMap[deliverySize] = pq
	}

	// Sort
	sort.Slice(deliveryPriority, func(i, j int) bool {
		// Descending order
		return deliveryPriority[i].deliverySize > deliveryPriority[j].deliverySize
	})


	for  i := 0; i < len(deliveryPriority); i++{
		data := &deliveryPriority[i]
		for _, computedHeap := range theatreMap[data.deliverySize] {
			if capacity[computedHeap.Value] >= data.deliverySize {
				capacity[computedHeap.Value] = capacity[computedHeap.Value] - data.deliverySize
				data.isDeliverable = true
				data.partnerId = computedHeap.Value
				data.cost = strconv.Itoa(computedHeap.Priority)
				break
			}
		}
	}

	// Sort
	sort.Slice(deliveryPriority, func(i, j int) bool {
		return deliveryPriority[i].deliverySize < deliveryPriority[j].deliverySize
	})
	fmt.Println(deliveryPriority)
	var result [] ResultFormatStruct
	for _, data := range deliveryPriority{
		result = append(result, ResultFormatStruct{
			DeliveryId: data.deliveryId,
			DeliveryPossible: data.isDeliverable,
			PartnerId: data.partnerId,
			Cost: data.cost,
		})
	}
	helpers.WriteDataToCsv("problem2.csv", result)
}

func main() {
	problem1()
	problem2()	
}