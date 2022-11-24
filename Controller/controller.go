package Controller

import (
	"fmt"

	"../Utils"
)

func Controller() bool {
	status, PartnerDetails := Utils.ReadPartnersDetails()
	if status {
		fmt.Println("ReadPartnersDetails Success")
		status, InputDetails := Utils.ReadInput()
		if status {
			fmt.Println("ReadInput Success")
			if Utils.DeliveryIsPossibleCheck(PartnerDetails, InputDetails) {
				fmt.Println("DeliveryISPossibleCheck Success")
				return true
			} else {
				fmt.Println("DeliveryISPossibleCheck Failed")
			}
		} else {
			fmt.Println("ReadInput Failed")
		}
	} else {
		fmt.Println("ReadPartnersDetails Failed")
	}
	return false
}
