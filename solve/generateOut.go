package solve

import (
	"challenge2019/models"
	"fmt"
	"os"
	"strconv"
)

func generateOut(f string, outArray []models.OutputDetails) error {
	file, err := os.Create(f)
	if err != nil {
		return err
	}

	defer file.Close()
	for _, out := range outArray {
		line := out.DeliveryID
		if out.Feasibility {
			line = fmt.Sprintf("%s,%s,%s,%s\n", line, "true ", out.PartnerID, strconv.Itoa(out.Cost))
		} else {
			line = fmt.Sprintf("%s,%s,%s,%s\n", line, "false", "\"\"", "\"\"")
		}
		_, err = file.WriteString(line)

		if err != nil {
			return err
		}
	}
	return nil
}
