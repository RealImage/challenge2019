package main

import (
	"fmt"

	"./Controller"
)

func init() {
	fmt.Println("Code Challenge Started")
}

func main() {
	status := Controller.Controller()
	if status {
		fmt.Println("Success")
	} else {
		fmt.Println("Failed")
	}
}
