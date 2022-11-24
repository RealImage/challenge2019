package main

import (
	"fmt"

	customType "github.com/praveenpkg8/challenge2019/customtype"
	"github.com/praveenpkg8/challenge2019/fileparser"
)

func main() {
	const inputFilePath string = "./input.csv"
	const capcitiesFilePath string = "./capacities.csv"
	const parternsFilePath string = "./partners.csv"
	var theaterDetials customType.Theatre = fileparser.LoadTheatreDetials(parternsFilePath)
	fileparser.ParseCapacitiesDetials(capcitiesFilePath)
	fileparser.GenerateOutput(inputFilePath, theaterDetials)
	fileparser.GenerateOutputV1(inputFilePath, theaterDetials)
	fmt.Println("Output generated successfully")

}
