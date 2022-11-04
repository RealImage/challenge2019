package file

import (
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"os"
)

func ReadAsync(fileName string, out interface{}) chan error {
	c := make(chan error)

	go func() {
		defer close(c)

		file, err := os.Open(fileName)
		if err != nil {
			c <- err
			return
		}

		defer file.Close()

		if err := gocsv.UnmarshalFile(file, out); err != nil {
			c <- err
			return
		}

		c <- nil
	}()

	return c
}

func ReadToMapAsync(fileName string, out interface{}) chan error {
	c := make(chan error)

	go func() {
		defer close(c)

		file, err := os.Open(fileName)
		if err != nil {
			c <- err
			return
		}

		r := csv.NewReader(file)

		defer file.Close()

		if err := gocsv.UnmarshalCSVToMap(r, out); err != nil {
			c <- err
			return
		}

		c <- nil
	}()

	return c
}

func Write(fileName string, data interface{}) (err error) {
	err = os.Remove(fileName)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	file, _ := os.Create(fileName)
	defer file.Close()

	err = gocsv.MarshalFile(data, file)
	if err != nil {
		return
	}
	return
}
