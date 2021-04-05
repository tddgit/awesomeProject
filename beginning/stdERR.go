package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please provide more arguments"

	} else {
		myString = arguments[1]

	}
	io.WriteString(os.Stdout, "This is a standart output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")

}
