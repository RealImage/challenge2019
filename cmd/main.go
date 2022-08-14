package main

import (
	"log"

	"github.com/zeebo/errs"

	solve "challenge2019"
)

var (
	// Error is an error class that indicates error in main func.
	Error = errs.Class("cmd/main error")
)

func main() {
	if err := solve.TaskOne(); err != nil {
		log.Fatal(Error.Wrap(err))
	}

	log.Println("Solved successful: result saved in output.csv")
}
