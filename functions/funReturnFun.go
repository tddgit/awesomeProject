package main

import "fmt"

func funReturnFun() func() int {
	i := 0
	return func() int {
		i++
		return i * i
	}

}

func main() {
	i := funReturnFun()
	j := funReturnFun()

	fmt.Println("i1: ", i())
	fmt.Println("i2: ", i())
	fmt.Println("j1: ", j())
	fmt.Println("j2: ", j())

}
