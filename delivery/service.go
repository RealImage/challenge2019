package delivery

import (
	"sync"
)

type Service struct {
	Repo *Repository
}

func NewDeliveryService(r *Repository) *Service {
	return &Service{Repo: r}
}

func fanOut(partners []Partner, out ...chan Partner) {
	go func() {
		for _, p := range partners {
			for _, c := range out {
				c <- p
			}
		}

		for _, c := range out {
			close(c)
		}
	}()
}

func merge(in ...chan Output) chan Output {
	wg := sync.WaitGroup{}
	out := make(chan Output, len(in))
	wg.Add(len(in))

	for _, c := range in {
		go func(ch chan Output) {
			defer wg.Done()
			for t := range ch {
				out <- t
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func (s *Service) checkPartner(in <-chan Partner, d Input) chan Output {
	out := make(chan Output)
	go func() {
		for p := range in {
			out <- func(p Partner) Output {
				if d.TheatreID == p.TheatreID && d.Amount >= p.Slab.MinSlab && d.Amount <= p.Slab.MaxSlab {
					cost := d.Amount * p.CostPerGB
					if cost < p.MinCost {
						cost = p.MinCost
					}
					return Output{DeliveryID: d.DeliveryID, IsPossible: true, Cost: cost, PartnerID: p.PartnerID}
				}
				return Output{DeliveryID: d.DeliveryID, IsPossible: false}
			}(p)
		}
		close(out)
	}()

	return out
}

func (s *Service) FindMinCostPartners(deliveryFile, partnerFile string) ([]Output, error) {
	var (
		deliveries []Input
		partners   []Partner
	)

	filesCount := 2
	errors := make(chan error, filesCount)

	go s.Repo.FetchDeliveries(errors, deliveryFile, &deliveries)
	go s.Repo.FetchPartners(errors, partnerFile, &partners)

	for i := 0; i < filesCount; i++ {
		if err := <-errors; err != nil {
			return nil, err
		}
	}

	size := len(deliveries)
	out := make([]chan Output, 0, size)
	outMap := make(map[TrimString]Output)
	result := make([]Output, 0, size)
	deliveryChannels := make([]chan Partner, size, size)

	for i := range deliveries {
		c := make(chan Partner)
		deliveryChannels[i] = c
		out = append(out, s.checkPartner(c, deliveries[i]))
	}

	fanOut(partners, deliveryChannels...)

	for o := range merge(out...) {
		if v, ok := outMap[o.DeliveryID]; !ok || (o.IsPossible && !v.IsPossible) || (o.IsPossible && o.Cost < v.Cost) {
			outMap[o.DeliveryID] = o
		}
	}

	for _, o := range outMap {
		result = append(result, o)
	}

	err := s.Repo.SaveDeliveriesOutput("./data/output.csv", &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
