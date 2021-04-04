package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument"
	} else {
		myString = arguments[1]
	}
	io.WriteString(os.Stdout, myString+"\n")
	io.WriteString(os.Stderr, "Testing Error output")
	io.WriteString(os.Stdout, "\n")

}
