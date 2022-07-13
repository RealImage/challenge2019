package main

import (
    "encoding/csv"
    "log"
    "os"
    "strconv"
    "strings"
)

//Partner ...
type Partner map[string][]*config

type config struct {
    TID         string
    MinSlabSize int
    MaxSlabSize int
    MinCost     int
    CperGB      int
}

var (
    data          = make(Partner)
    input         [][]string
    cap           = make(map[string]int)
    distributions [][]int
    partnerIds    []string
)

const MIN_COST = 99999999

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

    go loadCapacities()
}

func main() {
    for key := range data {
        partnerIds = append(partnerIds, key)
    }

    distributions = make([][]int, len(input))
    for idx, distribution := range input {
        distributions[idx] = make([]int, len(partnerIds))
        for j, p := range partnerIds {
            distributions[idx][j] = getCost(distribution, p)
        }
    }

    createResult(distributions)

}

type group struct {
    sum   int
    order string
    check bool
}

func createResult(distribution [][]int) {
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

    var minCost = MIN_COST
    var final string
    final += ""
    for _, value := range res {
        if value.check && minCost > value.sum {
            minCost = value.sum
            final = value.order
        }
    }

    var partnerRes []string
    for i := 2; i <= len(final); i = i + 2 {
        partnerRes = append(partnerRes, final[i-2:i])
    }

    var output1 [][]string

    for index, inputRow := range input {

        resultRow := []string{inputRow[0]}
        if partnerRes[index] == "  " {
            resultRow = append(resultRow, "false", " ", " ")
        } else {
            resultRow = append(resultRow, "true", partnerRes[index], strconv.Itoa(getCost(inputRow, partnerRes[index])))
        }
        output1 = append(output1, resultRow)
    }

    OutputWriter("output2.csv", output1)
}

func checkOrder(seq string) bool {
    defer loadCapacities()
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
        preResult := result
        result = result[length:]
        for _, p1 := range preResult {
            for index, p2 := range input {
                if p2 == -1 {
                    if validateInput(input) {
                        result = append(result, group{p1.sum + 0, p1.order + "  ", true})
                    } else {
                        result = append(result, group{p1.sum + p2, p1.order + partnerIds[index], false})
                    }
                    continue
                }
                result = append(result, group{p1.sum + p2, p1.order + partnerIds[index], true})
            }
        }
    } else {
        for index, p2 := range input {
            result = append(result, group{p2, partnerIds[index], true})
        }
    }
    return result
}

func validateInput(input []int) bool {
    for _, v := range input {
        if v != -1 {
            return false
        }
    }
    return true
}

func getCost(delivery []string, pid string) int {
    configArr := data[pid]
    deliveryContent := toInt(delivery[1])
    for _, value := range configArr {
        if deliveryContent >= value.MinSlabSize && deliveryContent <= value.MaxSlabSize {
            cost := value.CperGB * deliveryContent
            if cost <= value.MinCost {
                return value.MinCost
            }
            return cost
        }
    }
    return -1
}

func loadCapacities() {
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
