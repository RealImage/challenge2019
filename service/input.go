package service

import (
	"bufio"
	"os"
	"qube/constants"
	"qube/model"
	"strconv"
	"strings"
	"unicode"
)

func ReadCapacity() map[string]int {
	var capacities map[string]int = make(map[string]int)
	readFile, err := os.Open(constants.CAPACITY_FILE)

	if err != nil {
		panic(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		data := strings.Split(SpaceStringsBuilder(fileScanner.Text()), ",")

		amount, err := parseInt(data[1])

		if err != nil {
			continue
		}

		capacities[data[0]] = amount
	}

	return capacities
}

func ReadDelivery() []model.Delivery {
	var delivery []model.Delivery
	readFile, err := os.Open(constants.INPUT_FILE)

	if err != nil {
		panic(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		data := strings.Split(SpaceStringsBuilder(fileScanner.Text()), ",")

		newDelivery, err := parseDelivery(data)

		if err != nil {
			continue
		}

		delivery = append(delivery, *newDelivery)
	}

	return delivery
}

func ReadPartners() map[string][]model.Partner {
	var partners map[string][]model.Partner = make(map[string][]model.Partner)

	readFile, err := os.Open(constants.PARTNERS_FILE)

	if err != nil {
		panic(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		data := strings.Split(SpaceStringsBuilder(fileScanner.Text()), ",")

		_, ok := partners[data[0]]

		if !ok {
			partners[data[0]] = make([]model.Partner, 0)
		}

		newPartner, err := parsePartner(data)

		if err != nil {
			continue
		}

		partners[data[0]] = append(partners[data[0]], *newPartner)
	}

	return partners
}

func parseDelivery(data []string) (*model.Delivery, error) {
	amount, err := parseInt(data[1])

	if err != nil {
		return nil, err
	}

	return &model.Delivery{
		Name:    data[0],
		Amount:  amount,
		Theatre: data[2],
	}, nil
}

func parsePartner(data []string) (*model.Partner, error) {
	minCost, err := parseInt(data[2])

	if err != nil {
		return nil, err
	}

	perGB, err := parseInt(data[3])

	if err != nil {
		return nil, err
	}

	minGB, err := parseInt(strings.Split(data[1], "-")[0])

	if err != nil {
		return nil, err
	}

	maxGB, err := parseInt(strings.Split(data[1], "-")[1])

	if err != nil {
		return nil, err
	}

	return &model.Partner{
		MinCost: minCost,
		PerGB:   perGB,
		Partner: data[4],
		MinGB:   minGB,
		MaxGB:   maxGB,
	}, nil
}

func parseInt(number string) (int, error) {
	num, err := strconv.Atoi(number)

	if err != nil {
		return 0, err
	}

	return num, nil
}

func SpaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
