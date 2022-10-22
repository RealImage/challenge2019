package controller

import "C"
import (
	"challenge2019/db/entities"
	"challenge2019/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type ControllerStrcut struct {
	Service service.ServiceInterface
}

func NewController(serviceInterface service.ServiceInterface) *ControllerStrcut {
	return &ControllerStrcut{Service: serviceInterface}
}

var (
	InputData           []entities.Input
	ok                  bool
	PartnerDetails      []entities.Partner
	PartnerCapacityData []entities.PartnerCapacity
	output              []entities.Output
)

// Routing

func (c *ControllerStrcut) InstallRoute(g *gin.Engine) {
	apiRouter := g.Group("/QubeCinema")
	// Problem 1 handlers
	apiRouter.GET("/PartnerDetails", c.partnerDetails)
	apiRouter.GET("/InputDetails", c.inputDetails)
	apiRouter.GET("/Isdeliverable", c.Deliverable)
	// Problem 2 handlers
	apiRouter.GET("/PartnerCapacityDetails", c.PartnerCapacityDetails)
	apiRouter.GET("/CapcacityCheck", c.CapacityCheck)
}

// Problem 1

// Reads partner details from partner.csv file , store it as array of struct and returns it

func (c *ControllerStrcut) partnerDetails(g *gin.Context) {
	PartnerDetails, ok = c.Service.PartnerDetails()
	if !ok {
		log.Error("Failed to read PartnerDetails")
		g.JSON(http.StatusBadGateway, "Failed to read PartnerDetails")
		return
	}
	g.JSON(http.StatusOK, PartnerDetails)
}

// Reads Input details from partner.csv file , store it as array of struct and returns it

func (c *ControllerStrcut) inputDetails(g *gin.Context) {
	InputData, ok = c.Service.InputDetails()
	if !ok {
		log.Error("Failed to read PartnerDetails")
		g.JSON(http.StatusBadGateway, "Failed to read PartnerDetails")
		return
	}
	g.JSON(http.StatusOK, InputData)
}

// checks and returns if it's possible to deliver, which partner will deliver it and with what cost

func (c *ControllerStrcut) Deliverable(g *gin.Context) {
	PartnerDetails, ok = c.Service.PartnerDetails()
	if !ok {
		log.Error("Failed to read PartnerDetails")
		g.JSON(http.StatusBadGateway, "Failed to read PartnerDetails")
		return
	}

	InputData, ok = c.Service.InputDetails()
	if !ok {
		log.Error("Failed to read PartnerDetails")
		g.JSON(http.StatusBadGateway, "Failed to read PartnerDetails")
		return
	}
	PartnerCapacityData, ok = c.Service.PartnerCapacityDetails()
	if !ok {
		log.Error("Failed to read PartnerCapacityDetails")
		g.JSON(http.StatusBadGateway, "Failed to read PartnerCapacityDetails")
		return
	}

	_, output, ok = c.Service.Deliverable(PartnerDetails, InputData, PartnerCapacityData)
	if ok {
		g.JSON(http.StatusOK, "Deliverable")
		g.JSON(http.StatusOK, output)
		// If wants to download csv
		//g.Header("Content-type", "text/csv")
		//g.Header("Content-Disposition", "attachment; filename=\"output.csv\"")
		return
	} else {
		g.JSON(http.StatusBadGateway, "Not deliverable")
		return
	}
}

// Problem 2
func (c *ControllerStrcut) PartnerCapacityDetails(g *gin.Context) {
	PartnerCapacityData, ok := c.Service.PartnerCapacityDetails()
	var pIDcapacity = make(map[string]int)
	for _, j := range PartnerCapacityData {
		temp, _ := strconv.Atoi(j.Capacity)
		pIDcapacity[j.PartnerID] = temp
	}
	pIDcapacity["P1"] = pIDcapacity["P1"] - 21
	pIDcapacity["P1"] = pIDcapacity["P1"] - 22
	if !ok {
		log.Error("Failed to read PartnerCapacityDetails")
		g.JSON(http.StatusBadGateway, "Failed to read PartnerCapacityDetails")
		return
	}
	g.JSON(http.StatusOK, pIDcapacity)
	// g.JSON(http.StatusOK, PartnerCapacityData)
}

func (c *ControllerStrcut) CapacityCheck(g *gin.Context) {
	PartnerCapacityData, ok = c.Service.PartnerCapacityDetails()
	if !ok {
		log.Error("Failed to read PartnerCapacityDetails")
		g.JSON(http.StatusBadGateway, "Failed to read PartnerCapacityDetails")
		return
	}
	// g.JSON(http.StatusOK, PartnerCapacityData)
	var pIdCapacity = make(map[string]int)
	for _, j := range PartnerCapacityData {
		temp, _ := strconv.Atoi(j.Capacity)
		pIdCapacity[j.PartnerID] = temp
	}

	// Delivery check
	_, output, ok = c.Service.Deliverable(PartnerDetails, InputData, PartnerCapacityData)
	if ok {
		g.JSON(http.StatusOK, "Deliverable")
		g.JSON(http.StatusOK, output)
		return
	} else {
		g.JSON(http.StatusBadGateway, "Not deliverable")
		return
	}
}
