package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

type PartnersStruct struct {
	Theatre string `csv:"theatre"`
	MinSize int    `csv:"minsize"`
	MaxSize int    `csv:"maxsize"`
	MinCost int    `csv:"mincost"`
	CostPer int    `csv:"costper"`
	Partner string `csv:"partner"`
}

type InputStruct struct {
	Delivery string `csv:"delivery"`
	Size     int    `csv:"size"`
	Theatres string `csv:"theatre"`
}

func main() {
	partnersvalues := partners()
	inputsvalues := inputs()
	for i := 0; i < len(partners()); i++ {
		for j := 0; j < len(inputs()); j++ {
			if partnersvalues[i].MinSize < inputsvalues[j].Size && partnersvalues[i].MaxSize > inputsvalues[j].Size && strings.TrimSpace(partnersvalues[i].Theatre) == strings.TrimSpace(inputsvalues[j].Theatres) {
				fmt.Println(inputsvalues[j].Delivery, true, partnersvalues[i].Partner, partnersvalues[i].CostPer*inputsvalues[j].Size)
			}

			if partnersvalues[i].MinSize <= inputsvalues[j].Size && partnersvalues[i].MaxSize <= inputsvalues[j].Size && strings.TrimSpace(partnersvalues[i].Theatre) == strings.TrimSpace(inputsvalues[j].Theatres) {
				fmt.Println(inputsvalues[j].Delivery, false)
			}

		}

	}

}

func partners() []PartnersStruct {

	partnersValue, partnerserr := os.OpenFile("partners.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if partnerserr != nil {
		log.Fatal("Unable to read input file ", partnerserr)
	}
	defer partnersValue.Close()

	partnersclients := []PartnersStruct{}

	if err := gocsv.UnmarshalFile(partnersValue, &partnersclients); err != nil {
		panic(err)
	}
	return partnersclients
}

func inputs() []InputStruct {

	inputsValue, inputserr := os.OpenFile("input.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if inputserr != nil {
		log.Fatal("Unable to read input file ", inputserr)
	}
	defer inputsValue.Close()

	inputsclients := []InputStruct{}

	if err := gocsv.UnmarshalFile(inputsValue, &inputsclients); err != nil {
		panic(err)
	}
	return inputsclients
}
