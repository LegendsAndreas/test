package main

import (
	"fmt"
	"log"
)

func main() {
	var input, binaryTxt string

	fmt.Print("Enter a string you wish to translate to binary> ")
	_, err := fmt.Scan(&input)

	if err != nil {
		log.Fatal(err)
	}

	binaryTxt = stringToBinary(input)

	fmt.Printf("%s in binary is: %s", input, binaryTxt)
}

func stringToBinary(text string) string {
	var binaryTxt string

	return binaryTxt
}
