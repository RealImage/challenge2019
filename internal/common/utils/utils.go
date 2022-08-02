package utils

import (
	"log"
	"strconv"
	"strings"
)

// Parse string to int
func ToInt(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Fatal(err)
	}
	return n
}

// Selects correct slab measures
func CheckSlab(slab string, value int) bool {
	slabArr := strings.Split(slab, "-")
	if value >= ToInt(slabArr[0]) && value <= ToInt(slabArr[1]) {
		return true
	}
	return false
}
