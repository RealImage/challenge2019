package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	result := [][]string{}
	for _, j := range Data {
		Did := j.Did
		Found := false
		Pid := ""
		Amt := 0
		data, ok := Arranged[j.Tid]
		if ok {

			for _, k := range data {
				// fmt.Println(k)

				if j.Size > k.SizeSlabMin && j.Size < k.SizeSlabMax {
					Found = true

					if Amt == 0 {
						Pid = k.PartnerId

						Amt = getMax(k.CostGB*j.Size, k.MinCost)
					} else {
						if Amt > getMax(k.CostGB*j.Size, k.MinCost) {

							Pid = k.PartnerId
							Amt = getMax(k.CostGB*j.Size, k.MinCost)
						}

					}

				}

			}

		}
		res := []string{}
		res = append(res, Did)
		res = append(res, fmt.Sprintf("%v", Found))
		if Found {
			res = append(res, Pid)
			res = append(res, fmt.Sprintf("%v", Amt))
		} else {
			res = append(res, "''")
			res = append(res, "''")
		}
		result = append(result, res)
	}
	f, err := os.Create("output.csv")
	defer f.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(result) // calls Flush internally
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func getMin(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Records struct {
	Theatre     string
	SizeSlabMin int
	SizeSlabMax int
	MinCost     int
	CostGB      int
	PartnerId   string
}
type Input struct {
	Did  string
	Size int
	Tid  string
}

var Arranged map[string][]Records
var Data []Input

func init() {
	var wg sync.WaitGroup
	wg.Add(2)
	go ReadFeeder(&wg)
	go ReadInput(&wg)
	wg.Wait()
}

func ReadFeeder(wg *sync.WaitGroup) {
	arranged := map[string][]Records{}

	f, err := os.Open("partners.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Read()
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var record Records
		record.Theatre = strings.Replace(rec[0], " ", "", -1)
		record.SizeSlabMin = Conversion(strings.Split(strings.Replace(rec[1], " ", "", -1), "-")[0])
		record.SizeSlabMax = Conversion(strings.Split(strings.Replace(rec[1], " ", "", -1), "-")[1])
		record.MinCost = Conversion(strings.Replace(rec[2], " ", "", -1))
		record.CostGB = Conversion(strings.Replace(rec[3], " ", "", -1))
		record.PartnerId = strings.Replace(rec[4], " ", "", -1)
		arranged[strings.Replace(rec[0], " ", "", -1)] = append(arranged[strings.Replace(rec[0], " ", "", -1)], record)
	}
	Arranged = arranged
	wg.Done()
}
func ReadInput(wg *sync.WaitGroup) {
	data := []Input{}
	f, err := os.Open("input.csv")
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var input Input
		input.Did = strings.Replace(rec[0], " ", "", -1)
		input.Size = Conversion(strings.Replace(rec[1], " ", "", -1))
		input.Tid = strings.Replace(rec[2], " ", "", -1)
		data = append(data, input)
	}

	Data = data
	wg.Done()
}
func Conversion(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err)

	}
	return res
}
