package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Partner holds detailed information about a Partner's delivery capability.
type Partner struct {
	ID          string
	TID         string
	MinimumCost int
	CostPerGB   int
	MinGB       int
	MaxGB       int
}

// Capacity is a map of PartnerID and Capacity in GB for that partner.
type Capacity map[string]int

// Delivery contains a unique delivery ID with delivery Size in GB and a Theatre ID.
type Delivery struct {
	ID   string
	Size int
	TID  string
}

// Schedule is a record of the output.
type Schedule struct {
	DID      string
	Possible bool
	PID      string
	Cost     int
}

var problemNumber int
var partnersFile, capacitiesFile, inputFile string

func init() {
	const (
		defaultProblemNumber  = 1
		defaultPartnersFile   = "partners.csv"
		defaultCapacitiesFile = "capacities.csv"
		defaultInputFile      = "input.csv"
	)
	flag.IntVar(&problemNumber, "problem", defaultProblemNumber, "Choose between the two problems")
	flag.StringVar(&partnersFile, "partners", defaultPartnersFile, "Specify a partners list")
	flag.StringVar(&capacitiesFile, "capacities", defaultCapacitiesFile, "Specify capacities")
	flag.StringVar(&inputFile, "input", defaultInputFile, "Input")
}

func readRecords(reader *csv.Reader, o chan []string) {
	for {
		r, err := reader.Read()
		if err == io.EOF {
			close(o)
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		o <- r
	}
}

func readPartners() []Partner {
	partnersCSV, err := os.Open(partnersFile)
	defer partnersCSV.Close()
	if err != nil {
		log.Fatal(err)
	}

	pr := csv.NewReader(partnersCSV)
	var partners []Partner
	pr.Read() // Ignore header
	records := make(chan []string)
	go readRecords(pr, records)
	for r := range records {
		minCost, err := strconv.Atoi(strings.TrimSpace(r[2]))
		if err != nil {
			log.Fatal(err)
		}
		costGB, err := strconv.Atoi(strings.TrimSpace(r[3]))
		if err != nil {
			log.Fatal(err)
		}
		minMax := strings.Split(r[1], "-")
		minGB, err := strconv.Atoi(strings.TrimSpace(minMax[0]))
		if err != nil {
			log.Fatal(err)
		}
		maxGB, err := strconv.Atoi(strings.TrimSpace(minMax[1]))
		if err != nil {
			log.Fatal(err)
		}
		p := Partner{
			strings.TrimSpace(r[4]),
			strings.TrimSpace(r[0]),
			minCost,
			costGB,
			minGB,
			maxGB,
		}
		partners = append(partners, p)

	}
	return partners
}

func readCapacities() Capacity {
	capCSV, err := os.Open(capacitiesFile)
	defer capCSV.Close()
	if err != nil {
		log.Fatal(err)
	}

	cr := csv.NewReader(capCSV)
	capacities := make(Capacity)
	cr.Read() // dump header
	c := make(chan []string)
	go readRecords(cr, c)
	for r := range c {
		cap, err := strconv.Atoi(strings.TrimSpace(r[1]))
		if err != nil {
			log.Fatal(err)
		}
		capacities[strings.TrimSpace(r[0])] = cap
	}
	return capacities
}

func readInput() []Delivery {
	inputCSV, err := os.Open(inputFile)
	defer inputCSV.Close()
	if err != nil {
		log.Fatal(err)
	}

	ir := csv.NewReader(inputCSV)
	var deliveries []Delivery
	c := make(chan []string)
	go readRecords(ir, c)
	for r := range c {
		size, err := strconv.Atoi(strings.TrimSpace(r[1]))
		if err != nil {
			log.Fatal(err)
		}
		d := Delivery{
			strings.TrimSpace(r[0]),
			size,
			strings.TrimSpace(r[2]),
		}
		deliveries = append(deliveries, d)
	}
	return deliveries
}

func solveProblemOne(partners []Partner, deliveries []Delivery) []Schedule {
	var schedules []Schedule
	for _, d := range deliveries {
		s := Schedule{Possible: false, Cost: 0, DID: d.ID}
		for _, p := range partners {
			if d.TID == p.TID {
				if d.Size < p.MinGB || d.Size > p.MaxGB {
					if s.Possible == true {
						// We already know that there is atleast one solution
						continue
					}
					s.Possible = false
				} else {
					s.Possible = true
					cost := d.Size * p.CostPerGB
					if cost < p.MinimumCost {
						cost = p.MinimumCost
					}
					if s.Cost == 0 {
						s.Cost = cost
						s.PID = p.ID
					}
					if cost < s.Cost {
						s.Cost = cost
						s.PID = p.ID
					}
				}
			}
		}
		schedules = append(schedules, s)
	}
	return schedules
}

func solveProblemTwo(partners []Partner, capacities Capacity, deliveries []Delivery) []Schedule {
	var schedules []Schedule
	for _, d := range deliveries {
		deliveryPossible := false
		for _, p := range partners {
			s := Schedule{Possible: false, Cost: 0, DID: d.ID}
			if d.TID == p.TID {
				if d.Size < p.MinGB || d.Size > p.MaxGB {
					if deliveryPossible == true {
						// We already know that there is atleast one solution
						continue
					}
				} else {
					deliveryPossible = true
					s.Possible = true
					cost := d.Size * p.CostPerGB
					if cost < p.MinimumCost {
						cost = p.MinimumCost
					}
					if s.Cost == 0 {
						s.Cost = cost
						s.PID = p.ID
					}
					if cost < s.Cost {
						s.Cost = cost
						s.PID = p.ID
					}
					schedules = append(schedules, s)
				}
			}
		}
	}

	sort.Slice(deliveries, func(i, j int) bool { return deliveries[i].Size > deliveries[j].Size })

	var output []Schedule
	for _, d := range deliveries {
		found := false
		var sIdx int
		minCost := math.MaxInt64
		for i, s := range schedules {
			if s.DID == d.ID {
				if capacities[s.PID] >= d.Size {
					if s.Cost < minCost {
						found = true
						minCost = s.Cost
						capacities[s.PID] -= d.Size
						sIdx = i
					}
				}
			}
		}
		if found == false {
			output = append(output, Schedule{d.ID, false, "", 0})
		} else {
			output = append(output, schedules[sIdx])
		}
	}
	return output
}

func writeSchedules(pNum int, schedules []Schedule) {
	sort.Slice(schedules, func(i, j int) bool { return schedules[i].DID < schedules[j].DID })
	outFile := "problem-output" + strconv.Itoa(pNum) + ".csv"
	out, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer out.Close()
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(out)
	for _, s := range schedules {
		var cost string
		if s.Cost == 0 {
			cost = ""
		} else {
			cost = strconv.Itoa(s.Cost)
		}
		possible := strconv.FormatBool(s.Possible)
		w.Write([]string{s.DID, possible, s.PID, cost})
	}
	w.Flush()
	log.Printf("Wrote %d schedules to %s", len(schedules), outFile)
}

func main() {
	flag.Parse()
	partners := readPartners()
	log.Printf("Read %d partner records\n", len(partners))
	capacities := readCapacities()
	log.Printf("Read %d capacity records\n", len(capacities))
	deliveries := readInput()
	log.Printf("Read %d deliveries from the input file\n", len(deliveries))

	if problemNumber == 1 {
		writeSchedules(problemNumber, solveProblemOne(partners, deliveries))
	} else if problemNumber == 2 {
		writeSchedules(problemNumber, solveProblemTwo(partners, capacities, deliveries))
	} else {
		log.Fatal("That problem does not exist.")
	}
}
