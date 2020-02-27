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
	var theatreDetials customType.Theatre = fileparser.UpdateTheatreDetials(parternsFilePath)
	fileparser.ParseCapacitiesDetials(capcitiesFilePath)
	fileparser.GenerateOutput(inputFilePath, theatreDetials)
	fileparser.GenerateOutputV1(inputFilePath, theatreDetials)
	fmt.Println("Output generated successfully")

}
