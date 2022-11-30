package main

import (
	"challenge2019/models"
	"github.com/gocarina/gocsv"
	"os"
)

func main() {
	options := models.Options{}
	if err := ReadCsv("partners.csv", &options); err != nil {
		panic(err)
	}
	deliveriesInput := []*models.DeliveryInput{}
	if err := ReadCsvWithoutHeaders("input.csv", &deliveriesInput); err != nil {
		panic(err)
	}
	output1 := Task1(deliveriesInput, options)
	if err := WriteCsvWithoutHeaders("myoutput1.csv", output1); err != nil {
		panic(err)
	}

}

func Task1(deliveriesInput []*models.DeliveryInput, options models.Options) []*models.DeliveryOutput {
	output := []*models.DeliveryOutput{}
	for _, delivery := range deliveriesInput {
		optionsByTheater, err := options.GetOptionsByTheater(delivery.TheatreID)
		if err != nil {
			output = append(output, &models.DeliveryOutput{
				Delivery:   delivery.Delivery,
				IsPossible: false,
				PartnerID:  "",
				Cost:       0,
			})
			continue
		}
		optionsBySize, err := optionsByTheater.GetOptionsBySize(delivery.Size)
		if err != nil {
			output = append(output, &models.DeliveryOutput{
				Delivery:   delivery.Delivery,
				IsPossible: false,
				PartnerID:  "",
				Cost:       0,
			})
			continue
		}
		bestOption := optionsBySize.GetOptionWithBestPrice(delivery.Size)
		output = append(output, &models.DeliveryOutput{
			Delivery:   delivery.Delivery,
			IsPossible: true,
			PartnerID:  models.PartnerID(bestOption.PartnerID),
			Cost:       models.Cost(bestOption.GetPrice(delivery.Size)),
		})
	}
	return output
}

func ReadCsv(fileName string, model interface{}) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	if err = gocsv.UnmarshalFile(f, model); err != nil {
		return err
	}
	return nil
}

func ReadCsvWithoutHeaders(fileName string, model interface{}) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	if err = gocsv.UnmarshalWithoutHeaders(f, model); err != nil {
		return err
	}
	return nil
}

func WriteCsvWithoutHeaders(fileName string, model interface{}) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	if err = gocsv.MarshalWithoutHeaders(model, f); err != nil {
		return err
	}
	return nil
}
