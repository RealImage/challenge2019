package util

import "fmt"

func GetInput(input string) (output string) {
	output = ""
	fmt.Println("Enter the input for", input)
	fmt.Scanln("%s", output)
	return output
}
