package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//Partner ...
type Partner map[string][]*config

var data = make(Partner)

type config struct {
	TID         string
	MinSlabSize int
	MaxSlabSize int
	MinCost     int
	CperGB      int
}

var input [][]string
var cap = make(map[string]int)

var distribution [][]int
var partnerids []string

func main() {
	for key := range data {
		partnerids = append(partnerids, key)
	}

	distribution := make([][]int, len(input))
	for i, d := range input {
		distribution[i] = make([]int, len(partnerids))
		for j, p := range partnerids {
			distribution[i][j] = findCost(d, p)
		}
	}

	for _, value := range distribution {
		fmt.Println(value, "\n")
	}

	cummilative(distribution)
	//fmt.Println(hungarianAlgorithm.Solve(distribution))
}

type group struct {
	sum   int
	order string
	check bool
}

func cummilative(distribution [][]int) {
	var res []group
	var length = 0
	for _, v := range distribution {
		res = add(res, v, length)
		length = len(res)
	}

	for index, value := range res {

		if !checkOrder(value.order) {
			res[index].check = false
		}
	}

	var mincost int = 99999999
	var final string
	final += ""
	for _, value := range res {
		if value.check && mincost > value.sum {
			mincost = value.sum
			final = value.order
		}
	}

	fmt.Println(final)

	var partnerRes []string
	for i := 2; i <= len(final); i = i + 2 {
		partnerRes = append(partnerRes, final[i-2:i])
	}

	var output1 [][]string

	for index, inputrow := range input {
		//var resultrow []string
		resultrow := []string{inputrow[0]}
		if partnerRes[index] == "  " {
			resultrow = append(resultrow, "false", ` `, ` `)
		} else {
			resultrow = append(resultrow, "true", partnerRes[index], strconv.Itoa(findCost(inputrow, partnerRes[index])))
		}
		output1 = append(output1, resultrow)
	}
	fmt.Println(output1)
	OutputWriter("output2.csv", output1)

}

func checkOrder(seq string) bool {
	defer loadcap()
	for i := 2; i <= len(seq); i = i + 2 {
		if seq[i-2:i] == "  " {
			return true
		}
		value := cap[seq[i-2:i]] - toInt(input[(i-2)/2][1])
		if value < 0 {
			return false
		}
		cap[seq[i-2:i]] = value
	}
	return true
}

func add(result []group, input []int, length int) []group {
	if len(result) != 0 {
		//final := make([]int, len(result)*len(input))
		presult := result
		result = result[length:]
		for _, p1 := range presult {
			for index, p2 := range input {
				if p2 == -1 {
					if check(input) {
						result = append(result, group{p1.sum + 0, p1.order + "  ", true})
					} else {
						result = append(result, group{p1.sum + p2, p1.order + partnerids[index], false})

					}
					continue
				}
				result = append(result, group{p1.sum + p2, p1.order + partnerids[index], true})
			}
		}
	} else {
		//final := make([]int, len(input))
		for index, p2 := range input {
			result = append(result, group{p2, partnerids[index], true})
		}
	}
	//fmt.Println(result)
	return result
}

func check(input []int) bool {
	for _, v := range input {
		if v != -1 {
			return false
		}
	}
	return true
}

func findCost(delivery []string, pid string) int {
	configarr := data[pid]
	deliveryContent := toInt(delivery[1])
	for _, value := range configarr {
		if deliveryContent >= value.MinSlabSize && deliveryContent <= value.MaxSlabSize {
			c := value.CperGB * deliveryContent
			if c <= value.MinCost {
				return value.MinCost
			}
			return c
		}
	}
	return -1
}

func init() {
	input = InputReader("input.csv")
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			input[i][j] = strings.TrimSpace(input[i][j])
		}
	}

	partners := InputReader("partners.csv")
	for i := 0; i < len(partners); i++ {
		for j := 0; j < len(partners[i]); j++ {
			partners[i][j] = strings.TrimSpace(partners[i][j])
		}
	}

	for _, row := range partners[1:] {
		slabArr := strings.Split(row[1], "-")
		data[row[4]] = append(data[row[4]], &config{row[0], toInt(slabArr[0]), toInt(slabArr[1]), toInt(row[2]), toInt(row[3])})
	}
	loadcap()
}

func loadcap() {
	capacities := InputReader("capacities.csv")
	for _, capacitiesRow := range capacities[1:] {
		var err error
		cap[strings.TrimSpace(capacitiesRow[0])], err = strconv.Atoi(capacitiesRow[1])
		if err != nil {
			log.Fatal(err)
		}
	}
}

//InputReader ...
func InputReader(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

func toInt(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Fatal(err)
	}
	return n
}

//OutputWriter ...
func OutputWriter(filename string, output [][]string) {
	file, _ := os.Create(filename)
	defer file.Close()

	r := csv.NewWriter(file)
	err := r.WriteAll(output)
	if err != nil {
		log.Fatal(err)
	}
}
