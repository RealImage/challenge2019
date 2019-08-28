package helpers

import (
	"container/heap"
	"os"
	"encoding/csv"
	"bufio"
	"log"
	. "../structs"
	//"fmt"
	"strconv"
)

func ParseContentsFromCsv(filename string) *csv.Reader{
	csvFile, _ := os.Open(filename)
	r := csv.NewReader(bufio.NewReader(csvFile))
	return r
}

func WriteDataToCsv(filename string, data [] ResultFormatStruct) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Unable to create file")
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range data {
		temp := [] string{value.DeliveryId, strconv.FormatBool(value.DeliveryPossible), value.PartnerId, value.Cost}
		err := writer.Write(temp)
		if err != nil {
			log.Fatal("Cannot write to file")
		}
    }
}


// Min priority heap
type Item struct {
	Value    string
	Priority int
	Index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}