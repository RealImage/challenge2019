package file

import (
	"log"
	"os"

	"github.com/c1pca/challenge2019/models"
	"github.com/gocarina/gocsv"
)

func Read(path string, callback func(*os.File) interface{}) interface{} {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return callback(file)
}

// Write writes the data to CSV file
func Write(filename string, output []models.DeliveryResponse) {
	file, _ := os.Create(filename)
	defer file.Close()

	err := gocsv.MarshalFile(&output, file)
	if err != nil {
		log.Fatal(err)
	}
}
