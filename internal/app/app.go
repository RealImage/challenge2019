package app

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/c1pca/challenge2019/internal/common/storage/file"
	"github.com/c1pca/challenge2019/internal/common/utils"
	"github.com/c1pca/challenge2019/models"
	"github.com/gocarina/gocsv"
)

func FindPartnerWithMinCost(inputPath, partnersPath, outputPath string) {
	partners := file.Read(partnersPath, func(file *os.File) interface{} {
		var partners []*models.Partner
		if err := gocsv.UnmarshalFile(file, &partners); err != nil {
			panic(err)
		}
		return partners
	})

	deliveryRequest := file.Read(inputPath, func(file *os.File) interface{} {
		var deliveryRequest []*models.DeliveryRequest
		if err := gocsv.UnmarshalFile(file, &deliveryRequest); err != nil {
			panic(err)
		}
		return deliveryRequest
	})

	d, isValidRequest := deliveryRequest.([]*models.DeliveryRequest)
	p, isValidPartners := partners.([]*models.Partner)

	if isValidPartners && isValidRequest {
		wrappedMinDeliveryCostV1 := ElapsedTimeMiddleware(minDeliveryCostV1)
		result1 := wrappedMinDeliveryCostV1(d, p)

		wrappedMinDeliveryCostV2 := ElapsedTimeMiddleware(minDeliveryCostV2)
		result2 := wrappedMinDeliveryCostV2(d, p)

		file.Write(fmt.Sprintf("%soutput.csv", outputPath), result1)
		file.Write(fmt.Sprintf("%soutput_concurrently.csv", outputPath), result2)
	}
}

// Search available partners for delivery
func minDeliveryCostV1(req []*models.DeliveryRequest, partners []*models.Partner) []models.DeliveryResponse {
	fmt.Println("Best delivery partners:")
	var output []models.DeliveryResponse
	for _, delivery := range req {

		result := findBestPartner(delivery, partners)

		fmt.Println(result)
		output = append(output, result)
	}
	return output
}

// Concurrently search available partners for delivery
func minDeliveryCostV2(req []*models.DeliveryRequest, partners []*models.Partner) []models.DeliveryResponse {
	fmt.Println("Best delivery partners (found concurrently):")
	in := fanOut(req)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := worker(in, partners)
	c2 := worker(in, partners)
	c3 := worker(in, partners)
	c4 := worker(in, partners)

	var output []models.DeliveryResponse

	for n := range fanIn(c1, c2, c3, c4) {
		fmt.Println(n)
		output = append(output, n)
	}
	return output
}

func fanIn(cs ...<-chan models.DeliveryResponse) <-chan models.DeliveryResponse {
	var wg sync.WaitGroup
	out := make(chan models.DeliveryResponse)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan models.DeliveryResponse) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func worker(in <-chan *models.DeliveryRequest, partners []*models.Partner) <-chan models.DeliveryResponse {
	out := make(chan models.DeliveryResponse)
	go func() {
		for n := range in {
			out <- findBestPartner(n, partners)
		}
		close(out)
	}()
	return out
}

func findBestPartner(delivery *models.DeliveryRequest, partners []*models.Partner) models.DeliveryResponse {
	result := models.DeliveryResponse{
		DeliveryId: strings.TrimSpace(delivery.DeliveryId),
		PartnerId:  " ",
		Cost:       " ",
	}

	for _, partner := range partners {
		if strings.TrimSpace(delivery.TheatreId) == strings.TrimSpace(partner.TheatreId) && utils.CheckSlab(partner.Slab, int(delivery.Amount)) {
			computedCost := delivery.Amount * partner.CostPerGB
			if computedCost <= partner.CostMinimal {
				computedCost = partner.CostMinimal
			}

			if result.Cost == " " {
				result.Cost = strconv.Itoa(computedCost)
				result.PartnerId = strings.TrimSpace(partner.Id)
				result.IsPossible = true
			}
		}
	}

	return result
}

func fanOut(deliveries []*models.DeliveryRequest) <-chan *models.DeliveryRequest {
	out := make(chan *models.DeliveryRequest)
	go func() {
		for _, n := range deliveries {
			out <- n
		}
		close(out)
	}()
	return out
}

// Middleware that wraps minDeliveryCostV* functions and prints elapsed time
func ElapsedTimeMiddleware(next FinderFunc) FinderFunc {
	return func(req []*models.DeliveryRequest, partners []*models.Partner) []models.DeliveryResponse {
		start := time.Now()
		defer fmt.Printf("Processing succesfully finished!\r\nElapsed Time: %s \r\n\r\n", time.Since(start))
		return next(req, partners)
	}
}

type FinderFunc func(req []*models.DeliveryRequest, partners []*models.Partner) []models.DeliveryResponse
