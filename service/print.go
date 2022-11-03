package service

import (
	"bufio"
	"fmt"
	"os"
	"qube/constants"
	"qube/model"
)

func PrintProblemStatement1(result []model.Result) {
	f, err := os.Create(constants.OUTPUT1_FILE)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	for _, elem := range result {
		if elem.Cost == -1 {
			fmt.Fprintf(w, "%v,false,\"\",\"\"\n", elem.Name)
		} else {
			fmt.Fprintf(w, "%v,true ,%v,%v\n", elem.Name, elem.Partner, elem.Cost)
		}
	}

	w.Flush()
}

func PrintProblemStatement2(result []model.Result2) {
	f, err := os.Create(constants.OUTPUT2_FILE)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	for _, elem := range result {
		if elem.Cost == -1 {
			fmt.Fprintf(w, "%v,false,\"\",\"\"\n", elem.Name)
		} else {
			fmt.Fprintf(w, "%v,true ,%v,%v\n", elem.Name, elem.Partner, elem.Cost)
		}
	}

	w.Flush()
}
