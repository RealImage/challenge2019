package file

import (
	"github.com/gocarina/gocsv"
	"os"
)

func ReadAsync(c chan<- error, fileName string, out interface{}) {
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
	return
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
