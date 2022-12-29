package utils

import (
	"log"
	"strconv"
	"strings"
)

// Converts valid string int to int
func ToInt(x string) int {
	x = strings.Trim(x, " ")
	val, err := strconv.ParseInt(x, 10, 64)
	if err != nil {
		log.Println("Unable to parse to int:", err)
	}

	return int(val)
}

// Trims all the spaces
func Trim(x string) string {
	return strings.Trim(x, " ")
}
