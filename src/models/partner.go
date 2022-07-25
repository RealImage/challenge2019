package models

import (
	"strconv"
	"strings"
)

type PartnerRecord struct {
	TheatreID string
	Min       int
	Max       int
	MinCost   int
	CostPerGB int
	PartnerID string
}

func NewPartnerRecord(csvRecord []string) PartnerRecord {
	slab := strings.Split(csvRecord[1], "-")
	min, _ := strconv.Atoi(strings.TrimSpace(slab[0]))
	max, _ := strconv.Atoi(strings.TrimSpace(slab[1]))
	minCost, _ := strconv.Atoi(strings.TrimSpace(csvRecord[2]))
	costPerGB, _ := strconv.Atoi(strings.TrimSpace(csvRecord[3]))

	record := PartnerRecord{
		TheatreID: strings.TrimSpace(csvRecord[0]),
		Min:       min,
		Max:       max,
		MinCost:   minCost,
		CostPerGB: costPerGB,
		PartnerID: strings.TrimSpace(csvRecord[4]),
	}
	return record
}
