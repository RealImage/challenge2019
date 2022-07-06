package main

import (
	"flag"

	"github.com/purush7/challenge2019/v1/definitions"
	"github.com/purush7/challenge2019/v1/internal/algos"
	"github.com/purush7/challenge2019/v1/util"
)

var Ops = definitions.ProblemOps{}

func init() {
	flag.StringVar(&Ops.InputFile, "inputFile", "", "specify the input Filepath(Absolute or Relavtive path)")
	flag.StringVar(&Ops.PartnersFile, "partnersFile", "", "specify the partners Filepath(Absolute or Relavtive path)")
	flag.StringVar(&Ops.CapacitiesFile, "capacitiesFile", "", "specify the capacities Filepath(Absolute or Relavtive path)")
	flag.StringVar(&Ops.OutputFile, "outputFile", "", "specify the output Filepath(Absolute or Relavtive path)")
}

func main() {
	flag.Parse()

	// check for inputs and get the inputs
	if Ops.PartnersFile == "" {
		Ops.PartnersFile = util.GetInput("partnersFile")
	}

	if Ops.InputFile == "" {
		Ops.InputFile = util.GetInput("inputFile")
	}

	if Ops.OutputFile == "" {
		Ops.OutputFile = util.GetInput("outputFile")
	}

	//get the abs paths
	Ops.InputFile = util.GetAbsPaths(Ops.InputFile)
	Ops.OutputFile = util.GetAbsPaths(Ops.OutputFile)
	Ops.PartnersFile = util.GetAbsPaths(Ops.PartnersFile)
	Ops.CapacitiesFile = util.GetAbsPaths(Ops.CapacitiesFile)

	// Prob1 or Prob2
	if Ops.CapacitiesFile == "" {
		algos.BestPartner(Ops)
	} else {
		algos.MinimizeCost(Ops)
	}

}
