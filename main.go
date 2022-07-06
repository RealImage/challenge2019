package main

import (
	"flag"
	// ""
)

type ProblemOps struct {
	inputFile      string
	outputFile     string
	capacitiesFile string
	partnersFile   string
}

var Ops = ProblemOps{}

func init() {
	flag.StringVar(&Ops.inputFile, "inputFile", "", "specify the input Filepath(Absolute or Relavtive path)")
	flag.StringVar(&Ops.partnersFile, "partnersFile", "", "specify the partners Filepath(Absolute or Relavtive path)")
	flag.StringVar(&Ops.capacitiesFile, "capacitiesFile", "", "specify the capacities Filepath(Absolute or Relavtive path)")
	flag.StringVar(&Ops.outputFile, "outputFile", "", "specify the output Filepath(Absolute or Relavtive path)")
}

func main() {
	flag.Parse()

	// if Ops.partnersFile == "" {
	// Ops.partnersFile = util.GetInput("partnersFile")
	// }
}
