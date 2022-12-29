package db

import (
	"github.com/shreeyashnaik/challenge2019/common/schemas"
)

var (
	// Stores details from partner.csv as a mapping of theatre_id -> all partner details
	TheatrePartner map[string][]schemas.PartnerDetail

	// Stores details of partner capacities as partner_id -> capacity (in GB)
	Capacities map[string]int
)
