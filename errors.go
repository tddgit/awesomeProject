package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide one or more arguments")
	}
	arguments := os.Args
	var err error = errors.New("An error")
	k := 1
	var n float64

	for err != nil {
		if k > len(arguments) {
			fmt.Println("None of ")
		}
	}

}
