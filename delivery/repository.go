package delivery

import "github.com/challenge2019/file"

type Repository struct{}

func (r *Repository) FetchDeliveries(c chan<- error, fileName string, deliveries *[]Input) {
	file.ReadAsync(c, fileName, deliveries)
}

func (r *Repository) FetchPartners(c chan<- error, fileName string, partners *[]Partner) {
	file.ReadAsync(c, fileName, partners)
}

func (r *Repository) SaveDeliveriesOutput(fileName string, out *[]Output) error {
	err := file.Write(fileName, out)
	if err != nil {
		return err
	}
	return nil
}
