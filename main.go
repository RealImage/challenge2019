package main

import (
	"log"

	"github.com/shreeyashnaik/challenge2019/common/db"
	"github.com/shreeyashnaik/challenge2019/config"
	"github.com/shreeyashnaik/challenge2019/src"
)

func main() {
	// Load current working directory to config.SrcPath
	config.LoadSrcPath()
	log.Println(config.SrcPath + "/__data__")

	// Load partners.csv
	if err := db.LoadPartnersCsv(config.SrcPath + "/__data__/partners.csv"); err != nil {
		log.Println("Unable to load Partners csv to DB:", err)
	}

	// Load capacities.csv
	if err := db.LoadCapacitiesCsv(config.SrcPath + "/__data__/capacities.csv"); err != nil {
		log.Println("Unable to load Capacities csv to DB:", err)
	}

	// Problem Statement 1 solution
	if err := src.ProblemStatementOne(config.SrcPath+"/__data__/input.csv", config.SrcPath+"/__data__/my_output1.csv"); err != nil {
		log.Println("Error:", err)
	}

	// Problem Statement 2 solution
	if err := src.ProblemStatementTwo(config.SrcPath+"/__data__/input.csv", config.SrcPath+"/__data__/my_output2.csv"); err != nil {
		log.Println("Error:", err)
	}
}
