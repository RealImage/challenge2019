package main

import (
	"encoding/csv"
	"io"
	"reflect"
	"strconv"
	"os"
	"strings"
	"errors"
	"sort"
	"github.com/olekukonko/tablewriter"
)

type Partner struct {
	TheatreID	string
	SizeSlabInGB	string
	MinimumCost	int
	CostPerGB	int
	PartnerID	string
}

type Input struct {
	DeliveryID	string
	ContentSize	int
	TheatreID	string
}

type Output struct {
	DeliveryID		string
	DeliveryPossible	bool
	PartnerID		string
	CostOfDelivery		string
}

type Capacity struct {
	PartnerID	string
	Cap		int
}


func convertToRange(s string) (a,b int, err error) {
	s_arr := strings.Split(s, "-")
	a, err = strconv.Atoi(s_arr[0])
	if err != nil {
		return
	}
	b, err = strconv.Atoi(s_arr[1])
	return
}

func (p Partner) CostOfDelivery(content_size, capacity int) (result int, err error) {
	if content_size > capacity {
		err = errors.New("Content size is greater than capacity")
		return
	}
	var a,b int
	err = nil
	a,b,err = convertToRange(p.SizeSlabInGB)
	if err != nil {
		return
	}
	if a <= content_size && b >= content_size {
		result = content_size * p.CostPerGB
		if result < p.MinimumCost {
			result = p.MinimumCost
		}
	} else {
		err = errors.New("Invalid content size")
	}
	return
}

func Unmarshal(reader *csv.Reader, v interface{}) error {
	record, err := reader.Read()
	if err != nil {
		return err
	}
	s := reflect.ValueOf(v).Elem()
	if s.NumField() != len(record) {
		return &FieldMismatch{s.NumField(), len(record)}
	}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Type().String() {
		case "string":
			f.SetString(strings.Trim(record[i], " "))
		case "int":
			ival, err := strconv.ParseInt(strings.Trim(record[i], " "), 10, 0)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		default:
			return &UnsupportedType{f.Type().String()}
		}
	}
	return nil
}



func GetOutput(input Input, filename string, cap_map map[string]int) (Output,error) {
	output := Output{}
	partners, err := ReadPartnersCSV(filename, input.TheatreID)
	if err != nil {
		return output, err
	}
	min_cost := 0
	p_res := ""
	if len(partners) > 0 {
		for _, p := range partners {
			cost, e := p.CostOfDelivery(input.ContentSize, cap_map[p.PartnerID])
			if e == nil && min_cost == 0 {
				min_cost = cost
				p_res = p.PartnerID
			} else if e == nil && min_cost > cost {
				min_cost = cost
				p_res = p.PartnerID
			}
		}
	}

	if p_res == "" {
		output.PartnerID = "\"\""
		output.CostOfDelivery = "\"\""
	} else {
		output.PartnerID = p_res
		output.CostOfDelivery = strconv.Itoa(min_cost)
		output.DeliveryPossible = true
	}
	output.DeliveryID = input.DeliveryID
	cap_map[p_res] -= input.ContentSize

	return output, nil

}

func WriteOutputCSV(filename string, outputs []Output) error {
	f, err := os.OpenFile(filename, os.O_RDWR | os.O_CREATE, 0755)
	defer f.Close()
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Delivery ID", "Is Delivery Possible", "Partner ID", "Cost Of Delivery"})

	var writer = csv.NewWriter(f)
	writer.Comma = ','
	for _, o := range outputs {
		var record []string
		record = append(record, o.DeliveryID)
		record = append(record, strconv.FormatBool(o.DeliveryPossible))
		record = append(record, o.PartnerID)
		record = append(record, o.CostOfDelivery)
		err = writer.Write(record)
		if err != nil {
			return err
		}
		table.Append(record)
	}
	writer.Flush()
	table.Render()
	return nil
}


func ReadInputCSV(filename string) ([]Input, error) {
	//var source = "John;Smith;42\nPiter;Abel;50"
	f, err := os.Open(filename)
	defer f.Close()
	//var reader = csv.NewReader(strings.NewReader(source))
	if err != nil {
		return nil, err
	}

	var reader = csv.NewReader(f)
	reader.Comma = ','
	var inputs []Input
	for {
		var i Input
		err := Unmarshal(reader, &i)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		inputs = append(inputs, i)
	}
	return inputs,nil
}


func ReadCapacityCSV(filename string) (map[string]int, error) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	var reader = csv.NewReader(f)
	reader.Comma = ','
	m := make(map[string]int)
	reader.Read()
	for {
		var c Capacity
		err := Unmarshal(reader, &c)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		m[c.PartnerID] = c.Cap
	}
	return m,nil
}


func ReadPartnersCSV(filename string, tid string) ([]Partner, error) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	var reader = csv.NewReader(f)
	reader.Comma = ','
	var partners []Partner
	reader.Read()
	for {
		var p Partner
		err := Unmarshal(reader, &p)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if p.TheatreID == tid {
			partners = append(partners, p)
		}
	}
	return partners,nil
}

type FieldMismatch struct {
	expected, found int
}

func (e *FieldMismatch) Error() string {
	return "CSV line fields mismatch. Expected " + strconv.Itoa(e.expected) + " found " + strconv.Itoa(e.found)
}

type UnsupportedType struct {
	Type string
}

func (e *UnsupportedType) Error() string {
	return "Unsupported type: " + e.Type
}


func main() {
	inputs, err := ReadInputCSV("input.csv")
	if err != nil {
		panic(err)
	}
	sort.Slice(inputs, func(i,j int) bool {
		return inputs[i].ContentSize > inputs[j].ContentSize
	})

	cap_map, err := ReadCapacityCSV("capacities.csv")
	if err != nil {
		panic(err)
	}


	var outputs []Output
	for _, input := range inputs {
		output, err := GetOutput(input, "partners.csv", cap_map)
		if err != nil {
			panic(err)
		}
		outputs = append(outputs, output)
	}
	sort.Slice(outputs, func(i,j int) bool {
		return outputs[i].DeliveryID < outputs[j].DeliveryID
	})

	err = WriteOutputCSV("problem2_output.csv", outputs)
	if err != nil {
		panic(err)
	}

}


