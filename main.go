package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type PartnerID string
type TheatreID string
type DeliveryID string
type SizeGB int
type Cost int

type RawDataCost struct {
	Theatre     TheatreID
	SizeSlab    string
	MinimumCost string
	CostPerGB   string
	Partner     PartnerID
}

type PartnerRange struct {
	SizeSlabMin SizeGB
	SizeSlabMax SizeGB
	MinimumCost Cost
	CostPerGB   Cost
}

type PartnerParam struct {
	Ranges []PartnerRange
}

type PartnerData struct {
	Partners map[PartnerID]PartnerParam
}

type TheatreData struct {
	Theatres map[TheatreID]PartnerData
}

type PartnersLimit struct {
	Pl map[PartnerID]SizeGB
}

type Result struct {
	Delivery DeliveryID
	Size     SizeGB
	Partner  map[PartnerID]PartnerChar
}
type PartnerChar struct {
	Sum     Cost
	Posible bool
}

type Results struct {
	R []Result
}

type InputData struct {
	Delivery     DeliveryID
	SizeDelivery SizeGB
	Theatre      TheatreID
}

type OutputData1 struct {
	Delivery DeliveryID
	Posible  bool
	Partner  PartnerID
	Sum      Cost
}

type OutputData struct {
	Data []OutputData1
}

const (
	csvExtention = ".csv"
)

func removeSpaces(str string) string {
	return strings.Replace(str, " ", "", -1)
}

func toSizeGB(num string) SizeGB {
	n, _ := strconv.ParseInt(num, 10, 0)
	return SizeGB(n)
}

func toCost(num string) Cost {
	n, _ := strconv.ParseInt(num, 10, 0)
	return Cost(n)
}

func copyResultWithoutPartner(a Result, p PartnerID) (b Result) {
	b.Delivery = a.Delivery
	b.Size = a.Size
	b.Partner = make(map[PartnerID]PartnerChar)
	for key, value := range a.Partner {
		if key != p {
			b.Partner[key] = value
		}
	}
	return
}

// ReadPartnersData read data from file with partner table
func (td *TheatreData) ReadPartnersData(filename string) (err error) {
	if !strings.HasSuffix(filename, csvExtention) {
		filename += csvExtention
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
	}

	var (
		rdc  RawDataCost
		rdcs []RawDataCost
	)
	// if len(records[0]) > 4 {
	// 	rdc.Theatre = TheaterID(records[0][0])
	// 	rdc.SizeSlab = records[0][1]
	// 	rdc.MinimumCost = records[0][2]
	// 	rdc.CostPerGB = records[0][3]
	// 	rdc.Partner = PartnerID(records[0][4])
	// }
	for col := 1; col < len(records); col++ {
		if len(records[col]) > 4 {
			rdc.Theatre = TheatreID(removeSpaces(records[col][0]))
			rdc.SizeSlab = removeSpaces(records[col][1])
			rdc.MinimumCost = removeSpaces(records[col][2])
			rdc.CostPerGB = removeSpaces(records[col][3])
			rdc.Partner = PartnerID(removeSpaces(records[col][4]))
			rdcs = append(rdcs, rdc)
		} else {
			return err
		}
	}
	for col := range rdcs {

		var pr PartnerRange
		pr.CostPerGB = toCost(rdcs[col].CostPerGB)
		pr.MinimumCost = toCost(rdcs[col].MinimumCost)
		SizeSlab := strings.Split(rdcs[col].SizeSlab, "-")
		pr.SizeSlabMin = toSizeGB(SizeSlab[0])
		pr.SizeSlabMax = toSizeGB(SizeSlab[1])

		v1, b1 := td.Theatres[rdcs[col].Theatre]
		if b1 {
			v2, b2 := v1.Partners[rdcs[col].Partner]
			if b2 {
				v2.Ranges = append(v2.Ranges, pr)
				td.Theatres[rdcs[col].Theatre].Partners[rdcs[col].Partner] = v2
			} else {
				var pp PartnerParam
				pp.Ranges = append(pp.Ranges, pr)
				if v1.Partners == nil {
					v1.Partners = make(map[PartnerID]PartnerParam)
				}
				v1.Partners[rdcs[col].Partner] = pp
			}
		} else {
			if td.Theatres == nil {
				td.Theatres = make(map[TheatreID]PartnerData)
			}
			var pd PartnerData
			var pp PartnerParam
			pp.Ranges = append(pp.Ranges, pr)
			if pd.Partners == nil {
				pd.Partners = make(map[PartnerID]PartnerParam)
			}
			pd.Partners[rdcs[col].Partner] = pp
			td.Theatres[rdcs[col].Theatre] = pd
		}
	}
	return
}

// ReadInputData gets inpuut data
func ReadInputData(filename string) (in []InputData, err error) {
	if !strings.HasSuffix(filename, csvExtention) {
		filename += csvExtention
	}
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	for i := range records {
		delivery := DeliveryID(records[i][0])
		sizeDelivery := toSizeGB(records[i][1])
		theatre := TheatreID(records[i][2])
		in = append(in, InputData{delivery, sizeDelivery, theatre})
	}
	return
}

// ReadMaxLimitData read file with limits
func (pms *PartnersLimit) ReadMaxLimitData(filename string) {
	if !strings.HasSuffix(filename, csvExtention) {
		filename += csvExtention
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
	}
	// if len(records[0]) > 2 {
	// 	pm.PartnerID = records[0][0]
	// 	pm.Capacity = records[0][1]
	// }
	pms.Pl = make(map[PartnerID]SizeGB)
	for col := 1; col < len(records); col++ {
		if len(records[col]) > 1 {
			partnerID := PartnerID(removeSpaces(records[col][0]))
			capacity := toSizeGB(removeSpaces(records[col][1]))
			if pms.Pl == nil {
				pms.Pl = make(map[PartnerID]SizeGB)
			}
			pms.Pl[partnerID] = capacity
		}
	}
}

// AddResults concatenate all results together
func (rs *Results) AddResults(r ...Result) {
	rs.R = append(rs.R, r...)
	// for _, v := range r {
	// 	if len(v.Partner) > 0 {
	// 		rs.R = append(rs.R, v)
	// 	}
	// }
}

// CalculateVariants find all available variants
func (td *TheatreData) CalculateVariants(in InputData) (result Result) {
	theater := td.Theatres[in.Theatre]
	for key, value := range theater.Partners {
		for _, v := range value.Ranges {
			if in.SizeDelivery > v.SizeSlabMin && in.SizeDelivery <= v.SizeSlabMax {
				result.Size = in.SizeDelivery
				num := Cost(result.Size) * v.CostPerGB
				if num < v.MinimumCost {
					num = v.MinimumCost
				}
				if result.Partner == nil {
					result.Partner = make(map[PartnerID]PartnerChar)
				}
				result.Partner[key] = PartnerChar{num, true}
				break
			}
		}
	}
	result.Delivery = in.Delivery
	return
}

// FindResults give you final result
func (rs *Results) FindResults(pms ...*PartnersLimit) (out OutputData) {
	if len(pms) > 1 {
		var pls PartnersLimit
		pls.Pl = make(map[PartnerID]SizeGB)
		for _, v := range pms {
			if len(v.Pl) > 0 {
				for key, value := range v.Pl {
					if pls.Pl[key] > value {
						pls.Pl[key] = value
					}
				}
			}
		}
		return rs.FindResults(&pls)
	} else if len(pms) == 1 {
		pls := *pms[0]
		out = rs.FindResults()
		sum := make(map[PartnerID]SizeGB)
		for i, v := range out.Data {
			if v.Posible {
				sum[v.Partner] += rs.R[i].Size
			}
		}
		for key, val := range sum {
			if v, b := pls.Pl[key]; b && val > v {
				// find some variant with the limitations
				outs := []OutputData{}
				for key1, val1 := range out.Data {
					if key == val1.Partner {
						nrs := Results{R: []Result{}}
						nrs.R = append(nrs.R, rs.R...)
						nrs.R[key1] = copyResultWithoutPartner(nrs.R[key1], key)
						out1 := nrs.FindResults(&pls)
						outs = append(outs, out1)
					}
				}
				iterMin, possMax, sumMin := 0, 0, Cost(-1)
				for i, v := range outs {
					sum, possibility := Cost(0), 0
					for _, v1 := range v.Data {
						if v1.Posible {
							possibility++
						}
						sum += v1.Sum
					}
					if (possibility >= possMax && sumMin > sum) || sumMin == -1 {
						iterMin = i
						sumMin = sum
						possMax = possibility
					}
				}
				out = outs[iterMin]
			}
		}
	} else {
		for _, v := range rs.R {
			var out1 OutputData1
			partnerID, b := v.CalculateLowerVariant()
			out1.Delivery = v.Delivery
			out1.Posible = b
			if b {
				out1.Partner = partnerID
				out1.Sum = v.Partner[partnerID].Sum
			}
			out.Data = append(out.Data, out1)
		}
	}
	return
}

// CalculateLowerVariant find the best variant for this result
func (r *Result) CalculateLowerVariant() (PartnerID, bool) {
	partnerID, min := PartnerID(""), Cost(-1)
	if len(r.Partner) == 0 {
		return partnerID, false
	}
	for key, value := range r.Partner {
		if value.Sum < min || min == -1 {
			min = value.Sum
			partnerID = key
		}
	}
	return partnerID, true
}

// WriteOutputData writes in file outpuut data
func (out *OutputData) WriteOutputData(filename string) {
	if !strings.HasSuffix(filename, csvExtention) {
		filename += csvExtention
	}
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		log.Println(err)
	}

	records := [][]string{}
	for _, v := range out.Data {
		record := []string{}
		record = append(record, string(v.Delivery), fmt.Sprint(v.Posible))
		if v.Posible {
			record = append(record, string(v.Partner), fmt.Sprint(v.Sum))
		} else {
			// empty := ""
			empty := `""`
			record = append(record, empty, empty)
		}
		records = append(records, record)
	}

	if err := WriteAlignedCsv(f, records); err != nil {
		fmt.Println(err)
	}

	// w := csv.NewWriter(f) // csv writes data in not aligned format
	// if err := w.WriteAll(records); err != nil {
	// 	fmt.Println(err)
	// }

}

type Align int

const (
	Left Align = iota
	Right
)

// GetAlignedData add spaces to strings
func GetAlignedData(records [][]string, rule Align) (data [][]string) {
	max := []int{}
	for _, v := range records {
		for i, v2 := range v {
			if i >= len(max) {
				max = append(max, 0)
			}
			if l := len(v2); l > max[i] {
				max[i] = l
			}
		}
	}
	for i := range records {
		record := []string{}
		for j := range records[i] {
			str := records[i][j]
			if str == "" {
				// str = `""`
			}
			spaces := ""
			if l := max[j] - len(str); l > 0 {
				spaces = strings.Repeat(" ", l)
			}
			if rule == Left {
				str = str + spaces
			}
			if rule == Right {
				str = spaces + str
			}
			record = append(record, str)
		}
		data = append(data, record)
	}
	return
}

// WriteAlignedCsv write aligned data in csv file
func WriteAlignedCsv(f *os.File, records [][]string) (err error) {

	records = GetAlignedData(records, Left)

	w := bufio.NewWriter(f)
	defer w.Flush()
	bytes := []byte{}
	for _, record := range records {
		for i, v := range record {
			if i != 0 {
				bytes = append(bytes, ',')
			}
			bytes = append(bytes, []byte(v)...)
		}
		bytes = append(bytes, '\n')
	}
	_, err = w.Write(bytes)
	return err
}

var (
	partners   = flag.String("p", "partners", "Path to the CSV partners file")
	input      = flag.String("i", "input", "Path to the CSV input file")
	output     = flag.String("o", "output", "Path to the CSV output file")
	capacities = flag.String("c", "", "Path to the CSV capasities file")
	help       = flag.Bool("h", false, "")
)

func usage() {
	fmt.Printf(
		"%s: -p=<CSV Partners File> -i=<CSV Input File> -o=<CSV Output File> -c=<CSV Capacities File>\nDefaults are partner, input, output",
		os.Args[0],
	)
}

func main() {
	flag.Parse()
	if *help {
		usage()
		return
	}
	err := Start()
	if err != nil {
		fmt.Print(err.Error())
		return
	}

}

func Start() error {
	var td TheatreData
	err := td.ReadPartnersData(*partners)
	if err != nil {
		return err
	}

	in, err := ReadInputData(*input)
	if err != nil {
		return err
	}

	var results Results
	for _, v := range in {
		res := td.CalculateVariants(v)
		results.AddResults(res)
	}

	if *capacities == "" {
		out := results.FindResults()
		out.WriteOutputData(*output)
	} else {
		var pms PartnersLimit
		pms.ReadMaxLimitData(*capacities)

		out := results.FindResults(&pms)
		out.WriteOutputData(*output)
	}
	return nil
}
