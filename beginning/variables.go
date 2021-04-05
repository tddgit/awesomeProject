package main

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
)

const boilingF float64 = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Boiling Temperature = %gF or %gC\n", f, c)

	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, fToC(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, fToC(boilingF))

	var s string
	fmt.Printf("%sS", s)
	fmt.Println()

	var i, j, k int
	var b, d, e = true, 1., 4

	fmt.Println(reflect.TypeOf(d), e, d, i, j, k, b)

	//var file, err = os.Open("first.go")
	//fmt.Println(file, err, reflect.TypeOf(file))

	freq := rand.Float64() * 3.0
	t := 0.0
	fmt.Println(freq, t)

	z := 100
	var boiling float64 = 100
	var names []string
	var err1 error
	var someBool bool
	fmt.Println(z, boiling, reflect.TypeOf(names), reflect.TypeOf(err1), someBool)

	i, j = 1, 2
	i, j = j, i
	fmt.Println(i, j)

	file, err := os.Open("first.go")
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
	
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
