package db

import (
	"log"

	"github.com/shreeyashnaik/challenge2019/common/schemas"
	"github.com/shreeyashnaik/challenge2019/common/utils"
)

func LoadPartnersCsv(path string) error {
	TheatrePartner = make(map[string][]schemas.PartnerDetail)

	partnerRows, err := utils.ReadCsv(path)
	if err != nil {
		return err
	}

	for num, row := range partnerRows {
		if num == 0 {
			continue
		}
		theatreID := utils.Trim(row[0])
		partnerDetail := schemas.PartnerDetail{
			PartnerID:   utils.Trim(row[4]),
			SizeSlab:    utils.Trim(row[1]),
			MinimumCost: utils.ToInt(row[2]),
			CostPerGB:   utils.ToInt(row[3]),
		}

		TheatrePartner[theatreID] = append(TheatrePartner[theatreID], partnerDetail)
	}

	log.Println(TheatrePartner)
	return nil
}

func LoadCapacitiesCsv(path string) error {
	Capacities = make(map[string]int)

	capacities, err := utils.ReadCsv(path)
	if err != nil {
		return err
	}

	for num, row := range capacities {
		if num == 0 {
			continue
		}
		partnerID := utils.Trim(row[0])
		capacity := utils.ToInt(row[1])

		Capacities[partnerID] = capacity
	}

	log.Println(Capacities)
	return nil
}
