package main

import (
	"challenge2019/src/p1"
	"challenge2019/src/p2"
)

const (
	PARTNERSPATH   = "./resource/partners.csv"
	INPUTPATH      = "./resource/input.csv"
	OUTPUTPATHP1   = "./resource/output-p1.csv"
	CAPACITIESPATH = "./resource/capacities.csv"
	OUTPUTPATHP2   = "./resource/output-p2.csv"
)

func main() {
	p1.Soultion(PARTNERSPATH, INPUTPATH, OUTPUTPATHP1)
	p2.Soultion(PARTNERSPATH, INPUTPATH, CAPACITIESPATH, OUTPUTPATHP2)
}
