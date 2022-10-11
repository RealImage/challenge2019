package main

import (
	"challenge2019/models"
	"challenge2019/solve"
	"log"
)

var filenames = models.FileDetails{
	Partners:   "partners.csv",
	Capacities: "capacities.csv",
	Input:      "input.csv",
	Solution1:  "output1.csv",
	Solution2:  "output2.csv",
}

func main() {
	//..To "catch" panic and exit gracefully
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("panic occurred:", err)
		}
	}()
	log.Println("Solving....")

	if err := solve.Solution(&filenames); err != nil {
		log.Fatal(err)
	}

}
