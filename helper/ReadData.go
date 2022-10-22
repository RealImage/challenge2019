package helper

// Contains two helper function

// 1.ReadCsvFile - reads the csv file and return the data in [][]string format
// 2.Sscan - takes the string value , it converts it into integer and return that value

import (
	"encoding/csv"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type HelperInterface interface {
	ReadCsvFile(FileName string, skip bool) ([][]string, error)
	Sscan(strVal string) int
}

type HelperStruct struct {
}

func NewHelper() HelperInterface {
	return &HelperStruct{}
}

func (H *HelperStruct) ReadCsvFile(FileName string, skip bool) ([][]string, error) {
	File, err := os.Open(FileName)
	if err != nil {
		log.Error("Failed to open file")
		return nil, err
	}
	defer File.Close()

	csvRead := csv.NewReader(File)

	if skip {
		record, err := csvRead.Read()
		if err != nil {
			log.Error("Unable to read record")
		}
		log.Info(record)
	}

	records, err := csvRead.ReadAll()
	if err != nil {
		log.Error("Not able to retrieve records from csv file", err.Error())
		return nil, err
	}

	return records, nil
}

func (H *HelperStruct) Sscan(strVal string) int {
	var intVal int
	_, _ = fmt.Sscan(strVal, &intVal)
	return intVal
}
