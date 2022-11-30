package models

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type Option struct {
	Theatre     string `csv:"Theatre"`
	Size        Size   `csv:"Size Slab (in GB)"`
	MinimumCost int    `csv:"Minimum cost"`
	CostPerGB   int    `csv:"Cost Per GB"`
	PartnerID   string `csv:"Partner ID"`
}

func (o Option) GetPrice(size int) int {
	sizeCost := o.CostPerGB * size
	if sizeCost < o.MinimumCost {
		return o.MinimumCost
	}
	return sizeCost
}

type Size struct {
	SizeSlab string
	MinSize  int
	MaxSize  int
}

func (s *Size) UnmarshalCSV(csv string) (err error) {
	s.SizeSlab = csv
	sizeRange := strings.Split(csv, "-")
	s.MinSize, err = strconv.Atoi(sizeRange[0])
	if err != nil {
		return err
	}
	s.MaxSize, err = strconv.Atoi(strings.TrimSpace(sizeRange[1]))
	if err != nil {
		return err
	}
	return nil
}

func (s *Size) MarshalCSV() (string, error) {
	return s.SizeSlab, nil
}

type Options []*Option

func (o Options) GetOptionsByTheater(theaterID string) (Options, error) {
	optionsByTheater := []*Option{}
	for _, option := range o {
		if strings.TrimSpace(option.Theatre) == theaterID {
			optionsByTheater = append(optionsByTheater, option)
		}
	}
	if len(optionsByTheater) == 0 {
		return nil, errors.New("there is no appropriate data")
	}
	return optionsByTheater, nil
}

func (o Options) GetOptionsBySize(size int) (Options, error) {
	optionsBySize := []*Option{}
	for _, option := range o {
		if option.Size.MinSize < size && option.Size.MaxSize > size {
			optionsBySize = append(optionsBySize, option)
		}
	}
	if len(optionsBySize) == 0 {
		return nil, errors.New("there is no appropriate data")
	}
	return optionsBySize, nil
}

func (o Options) GetOptionWithBestPrice(size int) *Option {
	if len(o) == 1 {
		return o[0]
	}
	sort.SliceStable(o, func(i, j int) bool {
		return o[i].GetPrice(size) < o[j].GetPrice(size)
	})

	return o[0]
}
