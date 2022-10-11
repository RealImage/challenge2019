package solve

import (
	"challenge2019/models"
	"challenge2019/utils"
	"fmt"
)

func Solution(f *models.FileDetails) error {
	//Getting input partners and capacities from csv files
	input, err := utils.GetInput(f.Input)
	if err != nil {
		return fmt.Errorf("solve/Solution(): error reading input: \n %w", err)
	}
	partners, err := utils.GetPartners(f.Partners)
	if err != nil {
		return fmt.Errorf("solve/Solution(): error reading partners: \n %w", err)
	}
	capacityMap, err := utils.GetCapacities(f.Capacities)
	if err != nil {
		return fmt.Errorf("solve/Solution(): error reading capacities: \n %w", err)
	}

	output1Map, totalDataPerPartner := solution1(input, partners)
	if err := utils.GenerateOut(f.Solution1, output1Map); err != nil {
		return fmt.Errorf("solve/Solution(): error generating output files:%s \n %w", f.Solution1, err)
	}

	output2Map := solution2(input, capacityMap, totalDataPerPartner, output1Map, partners)
	if err := utils.GenerateOut(f.Solution2, output2Map); err != nil {
		return fmt.Errorf("solve/Solution(): error generating output files:%s \n %w", f.Solution2, err)
	}
	return nil
}
