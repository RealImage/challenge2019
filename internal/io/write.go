package io

import (
	"github.com/gocarina/gocsv"
	"github.com/sureshn04/challenge2019/models"
	"os"
)

//Write writes the data to CSV file
func Write(filename string, output []models.Output) {
	file, _ := os.Create(filename)
	defer file.Close()

	err := gocsv.MarshalFile(&output, file) // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}
}
