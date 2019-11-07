package Prob

import (
	"strconv"
	"strings"
)

func SplitString(str string) []string {
	res := strings.Split(str, "-")
	return res
}

func ConvertToInt(str string) int {
	ss := strings.TrimSpace(str)
	s, _ := strconv.Atoi(ss)
	return s
}

func ConvertToFloat(str string) float64 {
	ss := strings.TrimSpace(str)
	s, _ := strconv.ParseFloat(ss, 64)
	return s
}
