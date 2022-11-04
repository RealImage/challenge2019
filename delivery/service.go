package delivery

import (
	"github.com/challenge2019/util"
)

const (
	partnerFileName  = "./data/partners.csv"
	capacityFileName = "./data/capacities.csv"
)

type Service struct {
	Repo *Repository
}

func NewDeliveryService(r *Repository) *Service {
	return &Service{Repo: r}
}

func (s *Service) checkPartner(in <-chan Partner, d Input) chan Output {
	out := make(chan Output)
	min := Output{DeliveryID: d.DeliveryID, IsPossible: false}
	go func() {
		for p := range in {
			if d.TheatreID == p.TheatreID && d.Amount >= p.Slab.MinSlab && d.Amount <= p.Slab.MaxSlab {
				cost := d.Amount * p.CostPerGB
				if cost < p.MinCost {
					cost = p.MinCost
				}
				min = Output{DeliveryID: d.DeliveryID, IsPossible: true, Cost: cost, PartnerID: p.PartnerID}
			}
		}
		out <- min
		close(out)
	}()

	return out
}

func (s *Service) FindMinCostPartners(input string) ([]Output, error) {
	var (
		deliveries []Input
		partners   []Partner
	)

	errors := util.Merge(
		s.Repo.FetchDeliveries(input, &deliveries),
		s.Repo.FetchPartners(partnerFileName, &partners),
	)

	for err := range errors {
		if err != nil {
			return nil, err
		}
	}

	size := len(deliveries)
	out := make([]chan Output, 0, size)
	result := make([]Output, 0, size)
	deliveryChannels := make([]chan Partner, size, size)

	for i := range deliveries {
		c := make(chan Partner)
		deliveryChannels[i] = c
		out = append(out, s.checkPartner(c, deliveries[i]))
	}

	util.FanOut(partners, deliveryChannels...)

	for o := range util.Merge(out...) {
		result = append(result, o)
	}

	err := s.Repo.SaveDeliveriesOutput("./data/output_problem1.csv", &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func toMap(pp []Partner) map[TrimString][]Partner {
	m := make(map[TrimString][]Partner)

	for _, p := range pp {
		m[p.TheatreID] = append(m[p.TheatreID], p)
	}

	return m
}

type container struct {
	value        int
	okCount      int
	deliveries   []Output
	capacityLeft map[TrimString]int
}

func (c *container) copyAndAdd(d Input, p Partner) container {
	cl := copyMap(c.capacityLeft)
	out := Output{DeliveryID: d.DeliveryID, IsPossible: false}
	ok := c.okCount
	value := c.value
	if cl[p.PartnerID] >= d.Amount {
		ok++
		cost := d.Amount * p.CostPerGB
		if cost < p.MinCost {
			cost = p.MinCost
		}
		value += cost
		out = Output{DeliveryID: d.DeliveryID, IsPossible: true, Cost: cost, PartnerID: p.PartnerID}
	}
	cl[p.PartnerID] = cl[p.PartnerID] - d.Amount

	dd := make([]Output, len(c.deliveries), len(c.deliveries)+1)
	copy(dd, c.deliveries)

	return container{
		value:        value,
		capacityLeft: cl,
		okCount:      ok,
		deliveries:   append(dd, out),
	}
}

func copyMap(m map[TrimString]int) map[TrimString]int {
	cm := make(map[TrimString]int)
	for k, v := range m {
		cm[k] = v
	}
	return cm
}

func newContainer(d Input, p Partner, cc map[TrimString]int) container {
	var value int
	cl := copyMap(cc)
	out := Output{DeliveryID: d.DeliveryID, IsPossible: false}
	ok := 0

	if cl[p.PartnerID] >= d.Amount {
		ok++
		value = d.Amount * p.CostPerGB
		if value < p.MinCost {
			value = p.MinCost
		}
		out = Output{DeliveryID: d.DeliveryID, IsPossible: true, Cost: value, PartnerID: p.PartnerID}
	}
	cl[p.PartnerID] = cl[p.PartnerID] - d.Amount

	return container{
		value:        value,
		capacityLeft: cl,
		okCount:      ok,
		deliveries:   []Output{out},
	}
}

func (s *Service) Assign(input string) ([]Output, error) {
	var (
		deliveries []Input
		partners   []Partner
		capacities map[TrimString]int
	)

	errors := util.Merge(
		s.Repo.FetchDeliveries(input, &deliveries),
		s.Repo.FetchPartners(partnerFileName, &partners),
		s.Repo.FetchCapacities(capacityFileName, &capacities),
	)

	for err := range errors {
		if err != nil {
			return nil, err
		}
	}

	partnersMap := toMap(partners)

	var (
		containerSet []container
		opt          container
	)

	for _, d := range deliveries {
		pp := partnersMap[d.TheatreID]
		var cc []container
		for _, p := range pp {
			if d.Amount < p.Slab.MinSlab || d.Amount > p.Slab.MaxSlab {
				continue
			}

			cc = append(cc, newContainer(d, p, capacities))

			for i := 0; i < len(containerSet); i++ {
				newC := containerSet[i].copyAndAdd(d, p)
				cc = append(cc, newC)

				if newC.okCount > opt.okCount || (newC.okCount == opt.okCount && newC.value <= opt.value) {
					opt = newC
				}
			}
		}
		containerSet = append(containerSet, cc...)
	}

	err := s.Repo.SaveDeliveriesOutput("./data/output_problem2.csv", &opt.deliveries)
	if err != nil {
		return nil, err
	}

	return opt.deliveries, nil
}
