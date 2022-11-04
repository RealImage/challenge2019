package delivery

import "github.com/challenge2019/file"

type Repository struct{}

func (r *Repository) FetchDeliveries(fileName string, deliveries *[]Input) chan error {
	return file.ReadAsync(fileName, deliveries)
}

func (r *Repository) FetchPartners(fileName string, partners *[]Partner) chan error {
	return file.ReadAsync(fileName, partners)
}

func (r *Repository) FetchCapacities(fileName string, capacities *map[TrimString]int) chan error {
	return file.ReadToMapAsync(fileName, capacities)
}

func (r *Repository) SaveDeliveriesOutput(fileName string, out *[]Output) error {
	err := file.Write(fileName, out)
	if err != nil {
		return err
	}
	return nil
}
