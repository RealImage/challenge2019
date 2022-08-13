package main

import (
	solve "challenge2019"
	"github.com/zeebo/errs"
	"log"
)

var (
	// Error is an error class that indicates error in main func.
	Error = errs.Class("cmd/main error")
)

func main() {
	log.Print(solve.SolveFirstTask())
}
