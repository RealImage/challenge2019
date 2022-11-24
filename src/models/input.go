package models

import (
	"strconv"
	"strings"
)

type Input struct {
	DeliveryID string
	Volume     int
	TheatreID  string
}

func NewInput(input []string) Input {
	volume, _ := strconv.Atoi(strings.TrimSpace(input[1]))
	return Input{
		DeliveryID: strings.TrimSpace(input[0]),
		Volume:     volume,
		TheatreID:  strings.TrimSpace(input[2]),
	}
}
