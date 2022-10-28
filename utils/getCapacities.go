package utils

import (
	"bufio"
	"challenge2019/models"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// getCapacities gets capacity details
func GetCapacities(filename string) (models.CapacityMap, error) {
	capacityMap := make(models.CapacityMap)
	file, err := os.Open(filename)

	if err != nil {
		return models.CapacityMap{}, fmt.Errorf("utils/getCapacities() failed opening file:\n %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	for scanner.Scan() {

		data := strings.Split(scanner.Text(), ",")
		var capacity int

		if len(data) < 2 {
			break
		}
		capacity, err = strconv.Atoi(strings.Trim(data[1], " "))
		if err != nil {
			return models.CapacityMap{}, fmt.Errorf("utils/getCapacities() error reading size:\n %w", err)
		}

		capacityMap[strings.Trim(data[0], " ")] = capacity

	}

	return capacityMap, nil
}
