package main

import "github.com/challenge2019/delivery"

func main() {
	d := delivery.NewDeliveryService(&delivery.Repository{})
	_, err := d.Assign("./data/input.csv")
	if err != nil {
		panic(err)
	}
}
