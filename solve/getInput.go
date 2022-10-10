package solve

import (
	"bufio"
	"challenge2019/models"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// getInput gets input details
func getInput(filename string) (*[]models.InputDetails, error) {
	var inputList []models.InputDetails

	file, err := os.Open(filename)

	if err != nil {
		return &[]models.InputDetails{}, fmt.Errorf("solve/getInput() failed opening file:\n %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		var size int
		size, err = strconv.Atoi(strings.Trim(data[1], " "))
		if err != nil {
			return &[]models.InputDetails{}, fmt.Errorf("solve/getInput() error reading size:\n %w", err)
		}
		inputList = append(inputList, models.InputDetails{
			DeliveryID: strings.Trim(data[0], " "),
			Size:       size,
			TheatreID:  strings.Trim(data[2], " "),
		})
	}

	return &inputList, nil
}
