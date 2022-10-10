package solve

import (
	"bufio"
	"challenge2019/models"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// getPartners gets partner details
func getPartners(filename string) (models.TheatreMap, error) {
	theatreMap := make(models.TheatreMap)

	file, err := os.Open(filename)

	if err != nil {
		return theatreMap, fmt.Errorf("solve/getPartners() failed opening file:\n %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")

		var minimumCost int
		minimumCost, err = strconv.Atoi(strings.Trim(data[2], " "))
		if err != nil {
			return models.TheatreMap{}, fmt.Errorf("solve/getPartners() error reading minimum cost:\n %w", err)
		}

		var costPerGB int
		costPerGB, err = strconv.Atoi(strings.Trim(data[3], " "))
		if err != nil {
			return models.TheatreMap{}, fmt.Errorf("solve/getPartners() error reading cost per gb:\n %w", err)
		}
		minmax := strings.Split(data[1], "-")

		var min int
		min, err = strconv.Atoi(strings.Trim(minmax[0], " "))
		if err != nil {
			return models.TheatreMap{}, fmt.Errorf("solve/getPartners() error reading minimum slab:\n %w", err)
		}
		var max int
		max, err = strconv.Atoi(strings.Trim(minmax[1], " "))
		if err != nil {
			return models.TheatreMap{}, fmt.Errorf("solve/getPartners() error reading maximum slab:\n %w", err)
		}
		partner := models.PartnerDetails{
			TheatreID: strings.Trim(data[0], " "),
			SizeSlab: models.Slab{
				Min: min,
				Max: max,
			},
			MinimumCost: minimumCost,
			CostPerGB:   costPerGB,
			PartnerID:   strings.Trim(data[4], " "),
		}
		_, theatreOk := theatreMap[partner.TheatreID]
		if theatreOk {
			_, partnerOK := theatreMap[partner.TheatreID][partner.PartnerID]
			if partnerOK {
				theatreMap[partner.TheatreID][partner.PartnerID] = append(theatreMap[partner.TheatreID][partner.PartnerID], partner)
			} else {
				theatreMap[partner.TheatreID][partner.PartnerID] = []models.PartnerDetails{}
				theatreMap[partner.TheatreID][partner.PartnerID] = append(theatreMap[partner.TheatreID][partner.PartnerID], partner)
			}

		} else {
			theatreMap[partner.TheatreID] = models.PartnerMap{}
			theatreMap[partner.TheatreID][partner.PartnerID] = append([]models.PartnerDetails{}, partner)
		}
	}
	return theatreMap, nil
}
