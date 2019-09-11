/*
@author: Vinay Kumar
*/
package preprocess

import(
	"strings"
	"strconv"
 	"io"
 	"log"
	"github.com/funcoding/challenge2019/helpers"
	"github.com/funcoding/challenge2019/structs"
	"path/filepath"
	"go/build"
)

type Capacity map[string] int


type Partner map[string] map[string] [] structs.PartnerStruct

// To export
var Capacities Capacity = createCapacityHashTable()
var Partners Partner = createPartnerHashTable()

func createPartnerHashTable() Partner {
	csvPath, csvPathErr := filepath.Abs(build.Default.GOPATH+"/src/github.com/funcoding/challenge2019/static/partners.csv")
	if csvPathErr != nil{
		log.Fatal(csvPathErr)
	}
	r := helpers.ParseContentsFromCsv(csvPath)

	partner := Partner{}
	line := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		
		if line == 0{
			line++
			continue
		}
		theatreId := strings.TrimSpace(record[0])
		partnerId := strings.TrimSpace(record[4])
		minCost, minCostConversionErr := strconv.Atoi(strings.TrimSpace(record[2]))
		costPerGB, costPerGBConversionErr := strconv.Atoi(strings.TrimSpace(record[3]))
		usageSlab := strings.Split(strings.TrimSpace(record[1]), "-")
		minSlab, _ := strconv.Atoi(usageSlab[0])
		maxSlab, _ := strconv.Atoi(usageSlab[1])
		if minCostConversionErr != nil || costPerGBConversionErr != nil {
			log.Fatal("Something went wrong", minCostConversionErr, costPerGBConversionErr)
		}
		partnerData := structs.PartnerStruct{MinSlab: minSlab, MaxSlab: maxSlab, MinCost: minCost, CostPerGB: costPerGB}
		if partner[theatreId] == nil {
			partner[theatreId] = map [string] [] structs.PartnerStruct{}
		}
		partner[theatreId][partnerId] = append(partner[theatreId][partnerId], partnerData)
		line++
	}
	return partner
}

func createCapacityHashTable() Capacity{
	csvPath, csvPathErr := filepath.Abs(build.Default.GOPATH+"/src/github.com/funcoding/challenge2019/static/capacities.csv")
	if csvPathErr != nil{
		log.Fatal(csvPathErr)
	}
	r := helpers.ParseContentsFromCsv(csvPath)

	capacity := Capacity{}
	line := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		
		if line == 0{
			line++
			continue
		}
		number, err := strconv.Atoi(record[1])
		if err != nil{
			log.Fatal(err)
		}
		capacity[strings.TrimSpace(record[0])] = number
		line++
	}
	return capacity
}