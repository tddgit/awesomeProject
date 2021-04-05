package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	fmt.Println(min)

	fmt.Println(retThree(min))

	n1, n2, _ := retThree(min)

	fmt.Println(n1, n2)

	n1, n2 = n2, n1
	fmt.Println(n1, n2)

	x1, x2, x3 := n1, n1*n1, -n1
	fmt.Println(x1, x2, x3)

}

func retThree(x float64) (float64, float64, float64) {
	return 2 * x, x * x, -x
}
