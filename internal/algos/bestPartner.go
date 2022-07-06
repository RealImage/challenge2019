package algos

import (
	"github.com/purush7/challenge2019/v1/internal/populate"
	"github.com/purush7/challenge2019/v1/types"
	"github.com/purush7/challenge2019/v1/util"
)

func BestPartner(ops types.ProblemOps) {
	//populate data

	data := populate.PopulateData(ops.PartnersFile)
	inputRecords := util.ReadCsv(ops.InputFile)

	var outputRecords [][4]string

	for _, record := range inputRecords {
		outputRecord := [4]string{}
		outputRecord[0] = record[0]

		outputRecords = append(outputRecords, outputRecord)
	}
}
