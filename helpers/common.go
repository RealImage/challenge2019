package helpers

import (
	"fmt"
	"github.com/funcoding/qube-challenge-2019/structs"
	"bufio"
	"container/heap"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"path/filepath"
)

func ParseContentsFromCsv(filename string) *csv.Reader{
	csvFile, fileReadingErr := os.Open(filename)
	if fileReadingErr != nil {
		log.Fatalf("Unable to read file %s", filename)
	}
	r := csv.NewReader(bufio.NewReader(csvFile))
	return r
}

func WriteDataToCsv(filename string, data [] structs.ResultFormatStruct) string{
	staticFolderPath, staticFolderPathError := filepath.Abs(filepath.Dir("static/"))
	if staticFolderPathError != nil {
		log.Fatal(staticFolderPathError)
	}

	file, err := os.Create(fmt.Sprintf("%s/%s", staticFolderPath, filename))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to create file %s", fmt.Sprintf("%s/%s", staticFolderPathError, filename)))
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range data {
		temp := [] string{value.DeliveryId, strconv.FormatBool(value.DeliveryPossible), value.PartnerId, value.Cost}
		err := writer.Write(temp)
		if err != nil {
			log.Fatalf("Cannot write to file %s/%s", staticFolderPath,filename)
		}
    }
	return fmt.Sprintf("%s/%s", staticFolderPath, filename)
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


func (pq *PriorityQueue) ToArray() [] Item{
	var temp [] Item
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		temp = append(temp, *item)
	}
	return temp
}