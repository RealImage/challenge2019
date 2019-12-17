package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		partners   = flag.String("partners", "partners.csv", "path to the partners file [required]")
		input      = flag.String("input", "input.csv", "path to the input file [required]")
		output     = flag.String("ouput", "output.csv", "path to the output file [required]")
		capacities = flag.String("capacities", "", "path to the capacities file. If set will also consider maximum capacity of partners")
	)

	flag.Parse()

	if *partners == "" || *input == "" || *output == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err := run(*partners, *input, *output, *capacities); err != nil {
		fmt.Println(err)
	}
}

func run(partners, input, output, capacities string) error {
	rates, err := ReadRates(partners)
	if err != nil {
		return fmt.Errorf("unable to read rates from partners file %q, err: %w", partners, err)
	}

	req, err := ReadDeliveryRequests(input)
	if err != nil {
		return fmt.Errorf("unable to read delivery requests from input file %q, err: %w", input, err)
	}

	var cap map[string]int
	if capacities != "" {
		cap, err = ReadCapacities(capacities)
		if err != nil {
			return fmt.Errorf("unable to read partner capacities from capacities file %q, err: %w", capacities, err)
		}
	}

	res := Process(rates, req, cap)

	if err := WriteResponses(res, output); err != nil {
		return fmt.Errorf("unable to write delivery responses to output file %q, err: %w", output, err)
	}

	return nil
}

func Process(rates []*Rate, requests []*DeliveryRequest, capacities map[string]int) []*DeliveryResponse {
	res := make([]*DeliveryResponse, 0, len(requests))

	sizes := map[string]int{}
	prs := map[string]DeliveryResponses{}
	for _, req := range requests {
		r := &DeliveryResponse{
			ID:      req.ID,
			Cost:    -1,
			Request: req,
		}
		for _, rate := range rates { // Can be optimised using a map of theatres to rates depending on input
			updateResponse(rate, req, r)
		}

		if r.Accepted {
			sizes[r.Partner] = sizes[r.Partner] + req.Size
			prs[r.Partner] = append(prs[r.Partner], r)
		}
		res = append(res, r)
	}

	if capacities != nil {
		for p, s := range sizes {
			if capacities[p] < s {
				sort.Sort(prs[p])

				for i, r := range prs[p] {
					req := r.Request
					r.Cost = -1
					r.Accepted = false
					r.Partner = ""
					for _, rate := range rates {
						if rate.Partner == p {
							continue
						}
						if sizes[rate.Partner]+req.Size > capacities[rate.Partner] {
							continue
						}
						updateResponse(rate, req, r)
					}

					if r.Accepted {
						sizes[r.Partner] = sizes[r.Partner] + req.Size
						prs[r.Partner] = append(prs[r.Partner], r)
					}

					sizes[p] -= req.Size
					if capacities[p] >= sizes[p] {
						prs[p] = prs[p][i+1:]
						break
					}
				}
			}
		}
	}

	return res
}

func updateResponse(rate *Rate, req *DeliveryRequest, r *DeliveryResponse) {
	if rate.Theatre == req.Theatre && req.Size >= rate.Lower && req.Size <= rate.Upper { // Can be optimised by sorting based on bounds. Also depends on input size
		cost := req.Size * rate.CostPerGB
		if cost < rate.MinimumCost {
			cost = rate.MinimumCost
		}
		if r.Cost > cost || !r.Accepted {
			r.Cost = cost
			r.Accepted = true
			r.Partner = rate.Partner
		}
	}
}

type Rate struct {
	Theatre string
	Partner string

	CostPerGB    int
	MinimumCost  int
	Lower, Upper int
}

func ReadRates(file string) ([]*Rate, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("unable to open partners file: %w", err)
	}
	defer f.Close()

	pr := csv.NewReader(f)
	pr.ReuseRecord = true

	pr.Read() // The header, remove this if there is no header

	rr := make([]*Rate, 0, 20) // Optimisation based on example partners file.
	for {
		rec, err := pr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("unknown error while reading from partners file: %w", err)
		}

		if len(rec) != 5 {
			return nil, fmt.Errorf("malformed partners file, record contains %d entries", len(rec))
		}

		for i := range rec {
			rec[i] = strings.TrimSpace(rec[i])
		}

		r := Rate{
			Theatre: rec[0],
			Partner: rec[4],
		}

		r.MinimumCost, err = strconv.Atoi(rec[2])
		if err != nil {
			return nil, fmt.Errorf("malformed minimum cost in partners file, minimum cost: %q", rec[2])
		}

		r.CostPerGB, err = strconv.Atoi(rec[3])
		if err != nil {
			return nil, fmt.Errorf("malformed cost for gb in partners file, cost for gb: %q", rec[2])
		}

		slab := strings.Split(rec[1], "-")
		if len(slab) != 2 {
			return nil, fmt.Errorf("malformed slab in partners file, slab: %q", rec[1])
		}

		r.Lower, err = strconv.Atoi(slab[0])
		if err != nil {
			return nil, fmt.Errorf("malformed slab in partners file, slab: %q", rec[1])
		}

		r.Upper, err = strconv.Atoi(slab[1])
		if err != nil {
			return nil, fmt.Errorf("malformed slab in partners file, slab: %q", rec[1])
		}

		rr = append(rr, &r)
	}

	return rr, nil
}

type DeliveryRequest struct {
	ID      string
	Size    int
	Theatre string
}

func ReadDeliveryRequests(file string) ([]*DeliveryRequest, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("unable to open input file: %w", err)
	}
	defer f.Close()

	dr := csv.NewReader(f)
	dr.ReuseRecord = true
	dd := make([]*DeliveryRequest, 0, 5) // Optimisation based on example delivery file.
	for {
		rec, err := dr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("unknown error while reading from input file: %w", err)
		}

		if len(rec) != 3 {
			return nil, fmt.Errorf("malformed input file, record contains %d entries", len(rec))
		}

		for i := range rec {
			rec[i] = strings.TrimSpace(rec[i])
		}

		d := DeliveryRequest{
			ID:      rec[0],
			Theatre: rec[2],
		}

		d.Size, err = strconv.Atoi(rec[1])
		if err != nil {
			return nil, fmt.Errorf("malformed size in input file, size: %q", rec[2])
		}

		dd = append(dd, &d)
	}

	return dd, nil
}

type DeliveryResponse struct {
	ID string

	Accepted bool

	Partner string
	Cost    int

	Request *DeliveryRequest
}

func WriteResponses(rr []*DeliveryResponse, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("unable to create/truncate output file: %w", err)
	}
	defer f.Close()

	rw := csv.NewWriter(f)

	rec := make([]string, 4) // Can be resused instead of creating a new one for each record
	for _, r := range rr {
		rec[0] = r.ID
		rec[1] = strconv.FormatBool(r.Accepted)
		rec[2] = r.Partner
		if r.Cost < 0 {
			rec[3] = ""
		} else {
			rec[3] = strconv.Itoa(r.Cost)
		}

		if err := rw.Write(rec); err != nil {
			return err
		}
	}

	rw.Flush()
	if err := rw.Error(); err != nil {
		return err
	}

	return nil
}

func ReadCapacities(capacities string) (map[string]int, error) {
	f, err := os.Open(capacities)
	if err != nil {
		return nil, fmt.Errorf("unable to open capacities file: %w", err)
	}
	defer f.Close()

	cr := csv.NewReader(f)
	cr.ReuseRecord = true

	cr.Read() // header
	caps := map[string]int{}
	for {
		rec, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("unknown error while reading from capacities file: %w", err)
		}

		if len(rec) != 2 {
			return nil, fmt.Errorf("malformed capacities file, record contains %d entries", len(rec))
		}

		for i := range rec {
			rec[i] = strings.TrimSpace(rec[i])
		}

		cap, err := strconv.Atoi(rec[1])
		if err != nil {
			return nil, fmt.Errorf("malformed size in input file, size: %q", rec[1])
		}

		caps[rec[0]] = cap
	}

	return caps, nil
}

type DeliveryResponses []*DeliveryResponse

func (h DeliveryResponses) Len() int           { return len(h) }
func (h DeliveryResponses) Less(i, j int) bool { return h[i].Request.Size < h[j].Request.Size }
func (h DeliveryResponses) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
